package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	keys_manager_pb "github.com/t4ke0/pwm/keys_manager/proto"

	"github.com/t4ke0/pwm/pwm_authenticator/api"
	"github.com/t4ke0/pwm/pwm_authenticator/passwords"
	db "github.com/t4ke0/pwm/pwm_db_api"
)

// TODO: use JWT for authentication.

// TODO: make a server that authenticate the user by returning back a JWT token.

// TODO: make a simple JWT library that we can import from other services
//		 such as Authorizer in order to authorize the JWT token.

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

func main() {

	engine := gin.Default()
	engine.Use(cors.Default())

	engine.POST("/auth", func(c *gin.Context) {
	})

	engine.POST("/register", func(c *gin.Context) {
		var req api.RegisterRequest
		if err := c.BindJSON(&req); err != nil {
			c.String(http.StatusInternalServerError, "Error: [%v]", err)
			return
		}
		if req.Email.IsEmpty() || req.Username.IsEmpty() || req.Password.IsEmpty() {
			c.Status(http.StatusBadRequest)
			return
		}
		conn, err := db.New(postgresLink)
		if err != nil {
			c.String(http.StatusInternalServerError, "PSQL conn: [%v]", err)
			return
		}
		defer conn.Close()
		ok, err := conn.UserExists(req.Username.String())
		if err != nil {
			c.String(http.StatusInternalServerError, "Check User existance: [%v]", err)
			return
		}
		if ok {
			c.String(http.StatusConflict, "username already taken")
			return
		}
		// Generate an encryption key for the user.
		clientConn, err := dialKeysManager()
		if err != nil {
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

	// Default set to port 8080
	engine.Run()
}
