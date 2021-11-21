package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	pwm_manager_pb "github.com/t4ke0/pwm/pwm_manager/proto"

	"github.com/t4ke0/pwm/gateway/pkg/http/api"
	http_toolkit "github.com/t4ke0/pwm/pkg/common/http"
)

var pwmManagerHostname string = os.Getenv("PWM_MANAGER_HOSTNAME")

type gRPCconnection struct {
	dialConn *grpc.ClientConn
	conn     interface{}
}

func dialPWMmanager() (*gRPCconnection, error) {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	clientConn, err := grpc.Dial(pwmManagerHostname, opts...)
	if err != nil {
		return nil, err
	}

	conn := pwm_manager_pb.NewManagerClient(clientConn)

	return &gRPCconnection{
		dialConn: clientConn,
		conn:     conn,
	}, nil
}

func handleEndpointError(c *gin.Context) {
	if r := recover(); r != nil {
		c.String(http.StatusInternalServerError, "Error: %v", r)
		return
	}
}

func authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Next()
	}
}

func main() {

	engine := gin.Default()
	engine.Use(http_toolkit.CorsMiddleware())

	r := http_toolkit.GetBasePATHroute(engine)
	r.Use(authorize())

	{
		r.GET("/list/passwords", func(c *gin.Context) {
			jwtToken := c.GetHeader("token")

			defer handleEndpointError(c)
			grpcConnection, err := dialPWMmanager()
			if err != nil {
				panic(err)
			}
			defer grpcConnection.dialConn.Close()

			client, ok := grpcConnection.conn.(pwm_manager_pb.ManagerClient)
			if !ok {
				panic("gRPCconnection.conn: type assertion error")
			}
			ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
			defer cancel()
			userPasswords, err := client.GetPasswords(ctx, &pwm_manager_pb.GetPasswordsRequest{
				JwtToken: jwtToken,
			})
			if err != nil {
				panic(err)
			}
			var resp []api.ListPasswordResponse
			for _, p := range userPasswords.Passwords {
				resp = append(resp, api.ListPasswordResponse{
					PasswordID:        int(p.PasswordID),
					Username:          p.Data.Username,
					PlainTextPassword: p.Data.ClearTextPassword,
					Category:          p.Data.Category,
					Site:              p.Data.Site,
				})
			}
			c.JSON(http.StatusOK, resp)
		})

		r.POST("/add/password", func(c *gin.Context) {
			token := c.GetHeader("token")
			var req api.StorePasswordRequest
			if err := c.BindJSON(&req); err != nil {
				c.Status(http.StatusBadRequest)
			}

			defer handleEndpointError(c)
			grpcConnection, err := dialPWMmanager()
			if err != nil {
				panic(err)
			}
			defer grpcConnection.dialConn.Close()

			client, ok := grpcConnection.conn.(pwm_manager_pb.ManagerClient)
			if !ok {
				panic("grpcConnection.conn: type assertion error")
			}
			ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
			defer cancel()
			_, err = client.StorePassword(ctx, &pwm_manager_pb.ManagerRequest{
				JwtToken: token,
				Password: &pwm_manager_pb.PasswordData{
					ClearTextPassword: req.PlainTextPassword,
					Username:          req.Username,
					Category:          req.Category,
					Site:              req.Site,
				},
			})
			if err != nil && err.Error() == "record already exists" {
				c.Status(http.StatusConflict)
				return
			}
			if err != nil {
				panic(err)
			}
			c.Status(http.StatusCreated)
		})

		r.PATCH("/update/password", func(c *gin.Context) {
			token := c.GetHeader("token")
			var req api.UpdateUserItemsRequest
			if err := c.BindJSON(&req); err != nil {
				c.Status(http.StatusBadRequest)
				return
			}

			defer handleEndpointError(c)

			grpcConn, err := dialPWMmanager()
			if err != nil {
				panic(err)
			}
			defer grpcConn.dialConn.Close()

			client, ok := grpcConn.conn.(pwm_manager_pb.ManagerClient)
			if !ok {
				panic("pwm_manager_pb.ManagerClient type assertion error")
			}

			ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
			defer cancel()

			var values []string
			var modes []pwm_manager_pb.ItemToUpdate
			for _, n := range req.Items {
				if !n.Item.Validate() {
					c.String(http.StatusBadRequest, "%v is not a valid cred item", n.Item)
				}
				modes = append(modes, pwm_manager_pb.ItemToUpdate(pwm_manager_pb.ItemToUpdate_value[n.Item.String()]))
				values = append(values, n.Value)
			}

			if _, err := client.UpdatePassword(ctx, &pwm_manager_pb.ManagerUpdateRequest{
				JwtToken:   token,
				PasswordID: int64(req.PasswordID),
				Mode:       modes,
				Value:      values,
			}); err != nil {
				panic(err)
			}

			c.Status(http.StatusOK)
		})

		r.DELETE("/delete/password", func(c *gin.Context) {
			token := c.GetHeader("token")
			var req api.DeleteUserCredRequest
			if err := c.BindJSON(&req); err != nil {
				c.Status(http.StatusBadRequest)
				return
			}

			defer handleEndpointError(c)

			grpcConn, err := dialPWMmanager()
			if err != nil {
				panic(err)
			}
			defer grpcConn.dialConn.Close()

			client, ok := grpcConn.conn.(pwm_manager_pb.ManagerClient)
			if !ok {
				panic("/delete/password : type assertion error pwm_manager_pb.ManagerClient")
			}

			ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
			defer cancel()

			_, err = client.DeletePasswords(ctx, &pwm_manager_pb.DeletePasswordRequest{
				JwtToken:   token,
				PasswordID: int64(req.PasswordID),
			})

			if err != nil {
				panic(err)
			}

			c.Status(http.StatusOK)
		})

		r.GET("/gen/password", func(c *gin.Context) {
			var req api.GeneratePasswordRequest
			if err := c.BindJSON(&req); err != nil {
				c.Status(http.StatusBadRequest)
				return
			}

			defer handleEndpointError(c)
			grpcConn, err := dialPWMmanager()
			if err != nil {
				panic(err)
			}
			defer grpcConn.dialConn.Close()

			conn, ok := grpcConn.conn.(pwm_manager_pb.ManagerClient)
			if !ok {
				panic("type assertion")
			}

			ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
			defer cancel()

			mode := pwm_manager_pb.PasswordMode(pwm_manager_pb.PasswordMode_value[req.PasswordComplexity.String()])

			generatedPassword, err := conn.GeneratePassword(ctx, &pwm_manager_pb.GeneratePasswordRequest{
				Length: req.PasswordLength,
				Mode:   mode,
			})

			if err != nil {
				panic(err)
			}

			c.JSON(http.StatusCreated, api.GeneratePasswordResponse{
				GeneratedPassword: generatedPassword.Password,
			})
		})

	}
	engine.Run(":6969")
}
