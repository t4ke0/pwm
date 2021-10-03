package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/grpc"

	keys_manager_pb "github.com/t4ke0/pwm/keys_manager/proto"

	"github.com/t4ke0/pwm/pwm_authenticator/api"
	"github.com/t4ke0/pwm/pwm_authenticator/passwords"
	db "github.com/t4ke0/pwm/pwm_db_api"
)

var (
	keysManagerHost = os.Getenv("KEYS_MANAGER_HOST")
	keysManagerPort = os.Getenv("KEYS_MANAGER_PORT")
	//
	postgresPW   = os.Getenv("POSTGRES_PASSWORD")
	postgresDB   = os.Getenv("POSTGRES_DATABASE")
	postgresUser = os.Getenv("POSTGRES_USER")
	postgresHost = os.Getenv("POSTGRES_HOST")

	postgresLink = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		postgresUser,
		postgresPW,
		postgresHost,
		postgresDB)
)

func dialKeysManager() (*grpc.ClientConn, error) {
	log.Printf("DEBUG keys manger hostname %v:%v", keysManagerHost, keysManagerPort)
	opts := []grpc.DialOption{grpc.WithInsecure()}
	return grpc.Dial(fmt.Sprintf("%v:%v", keysManagerHost, keysManagerPort), opts...)
}

func init() {
	if os.Getenv("LOCAL_TEST") == "true" {
		testpostgresURL, err := db.CreateTestingDatabase(postgresLink)
		if err != nil {
			log.Fatalf("couldn't setup testing database %v", err)
		}
		conn, err := db.New(testpostgresURL)
		if err != nil {
			log.Fatalf("couldn't dial postgres database %v", err)
		}
		if err := conn.ClearTestTables(); err != nil {
			log.Fatalf("couldn't clear test tables %v", err)
		}
		postgresLink = testpostgresURL
	}

	conn, err := db.New(postgresLink)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	k, err := conn.GetAuthServerKey()
	if err != nil {
		log.Fatal(err)
	}
	if k == "" {
		keyManagerConn, err := dialKeysManager()
		if err != nil {
			log.Fatal(err)
		}
		defer keyManagerConn.Close()

		// Generating auth key.
		client := keys_manager_pb.NewKeyManagerClient(keyManagerConn)
		key, err := client.GenKey(
			context.TODO(), &keys_manager_pb.KeyGenRequest{
				Mode: keys_manager_pb.Mode_ServerAuth,
			})
		if err != nil {
			log.Fatalf("Couldn't generate server auth key [%v]", err)
		}
		log.Printf("generated server auth key %v", key.Key)
	}
}

func handleError(c *gin.Context) {
	if r := recover(); r != nil {
		c.JSON(http.StatusInternalServerError, api.ErrResponse{
			ErrorMessage: r.(error).Error(),
		})
	}
}

var engine *gin.Engine

func setupGinEngine() {

	const headerTokenKey string = "token"

	engine = gin.Default()
	engine.Use(cors.Default())
	engine.POST("/login", func(c *gin.Context) {
		var req api.AuthRequest
		defer handleError(c)
		if err := c.BindJSON(&req); err != nil {
			panic(err)
		}
		if req.Username.IsEmpty() || req.Password.IsEmpty() {
			c.Status(http.StatusBadRequest)
			return
		}
		conn, err := db.New(postgresLink)
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		info, err := conn.GetUserAuthInfo(req.Username.String())
		if err != nil && err == db.ErrNoRows {
			c.Status(http.StatusUnauthorized)
			return
		}
		if err != nil {
			panic(err)
		}

		storedPassword := passwords.ToHashedPassword(info.Password)
		valid := storedPassword.IsCorrectPassword(req.Password.Byte())
		if !valid {
			c.Status(http.StatusUnauthorized)
			return
		}
		sessionID := uuid.New().String()

		grpcConn, err := dialKeysManager()
		if err != nil {
			panic(err)
		}
		defer grpcConn.Close()

		keysManagerClient := keys_manager_pb.NewKeyManagerClient(grpcConn)
		userKey, err := keysManagerClient.GetUserKey(context.TODO(), &keys_manager_pb.KeyFetchRequest{Username: req.Username.String()})
		if err != nil {
			panic(err)
		}

		authServerKey, err := conn.GetAuthServerKey()
		if err != nil {
			panic(err)
		}

		if authServerKey == "" {
			panic(fmt.Errorf("no auth server key in the database"))
		}

		jwtToken, err := getNewJWTtoken([]byte(authServerKey), tokenClaims{
			UserID:       info.ID,
			Username:     req.Username.String(),
			SessionID:    sessionID,
			SymmetricKey: userKey.Key,
		})
		if err != nil {
			panic(err)
		}

		if err := conn.InsertNewSession(sessionID, jwtToken, info.ID, time.Now()); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, api.AuthResponse{jwtToken})
	})

	engine.POST("/register", func(c *gin.Context) {
		var req api.RegisterRequest
		if err := c.BindJSON(&req); err != nil {
			log.Printf("debug error %v", err)
			c.String(http.StatusInternalServerError, "Error: [%v]", err)
			return
		}
		if req.Email.IsEmpty() || req.Username.IsEmpty() || req.Password.IsEmpty() {
			c.Status(http.StatusBadRequest)
			return
		}
		log.Printf("debug %v", postgresLink)
		conn, err := db.New(postgresLink)
		if err != nil {
			log.Printf("debug error %v", err)
			c.String(http.StatusInternalServerError, "PSQL conn: [%v]", err)
			return
		}
		defer conn.Close()
		ok, err := conn.UserExists(req.Username.String())
		if err != nil {
			log.Printf("debug error %v", err)
			c.String(http.StatusInternalServerError, "Check User existance: [%v]", err)
			return
		}
		if ok {
			c.String(http.StatusConflict, "username already taken")
			return
		}

		emailExists, err := conn.EmailExists(req.Email.String())
		if err != nil {
			log.Printf("debug error %v", err)
			c.String(http.StatusInternalServerError, "Check Email existance [%v]", err)
			return
		}
		if emailExists {
			c.String(http.StatusConflict, "email already registred")
			return
		}

		// Generate an encryption key for the user.
		clientConn, err := dialKeysManager()
		if err != nil {
			log.Printf("debug error %v", err)
			c.String(http.StatusInternalServerError, "GRPC [error] (%v)", err)
			return
		}
		defer clientConn.Close()

		keyManagerClient := keys_manager_pb.NewKeyManagerClient(clientConn)
		userKey, err := keyManagerClient.GenKey(
			context.TODO(), &keys_manager_pb.KeyGenRequest{
				Mode: keys_manager_pb.Mode_User,
			})
		if err != nil {
			log.Printf("debug error %v", err)
			c.String(http.StatusInternalServerError, "GRPC [error] (%v)", err)
			return
		}
		hashedPassword, err := passwords.Hash([]byte(req.Password))
		if err != nil {
			c.String(http.StatusInternalServerError, "hash user pw (%v)", err)
			return
		}
		regConf := db.RegistrationConfig{
			Username: req.Username.String(),
			Password: hashedPassword.String(),
			Email:    req.Email.String(),
			Key:      userKey.Key,
		}
		if err := conn.InsertNewUser(regConf); err != nil {
			c.String(http.StatusInternalServerError, "Store User (%v)", err)
			return
		}
		c.Status(http.StatusCreated)
		// TODO: in the future we can introduce an email service that sends an
		// email verification to the users
		return
	})

	engine.GET("/info", func(c *gin.Context) {
		tokenString := c.GetHeader(headerTokenKey)
		if tokenString == "" {
			c.Status(http.StatusBadRequest)
			return
		}

		defer handleError(c)
		conn, err := db.New(postgresLink)
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		authKey, err := conn.GetAuthServerKey()
		if err != nil {
			panic(err)
		}
		if authKey == "" {
			panic(fmt.Errorf("auth key not present in the database"))
		}
		tokenclaims, err := parseJWTtoken(tokenString, []byte(authKey))
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, tokenclaims)
	})

	engine.POST("/logout", func(c *gin.Context) {
		tokenAsStr := c.GetHeader(headerTokenKey)
		conn, err := db.New(postgresLink)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		defer conn.Close()
		if err := conn.RevokeSession(tokenAsStr); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
	})

}

func main() {

	setupGinEngine()
	// Default set to port 8080
	engine.Run()
}
