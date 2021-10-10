package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"

	"google.golang.org/grpc"

	pb "github.com/t4ke0/pwm/pwm_manager/proto"

	db "github.com/t4ke0/pwm/pwm_db_api"
	"github.com/t4ke0/pwm/pwm_manager/passwords"
	"github.com/t4ke0/pwm/pwm_manager/server/api"
)

var (
	authenticatorAddress = os.Getenv("AUTHENTICATOR_ADDRESS")
	//
	postgresPW   = os.Getenv("POSTGRES_PASSWORD")
	postgresHost = os.Getenv("POSTGRES_HOST")
	postgresUser = os.Getenv("POSTGRES_USER")
	postgresDB   = os.Getenv("POSTGRES_DB")

	postgresURL = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		postgresUser,
		postgresPW,
		postgresHost,
		postgresDB)
)

var (
	getJWTinfoError = errors.New("Couldn't get JWT token info from authenticator")
)

type managerServer struct {
	pb.UnimplementedManagerServer

	userKey []byte
}

// GetPasswords method that accept context and GetPasswordsRequest proto buff
// type and returns UserPasswords.
func (ms *managerServer) GetPasswords(ctx context.Context,
	request *pb.GetPasswordsRequest) (*pb.UserPasswords, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/info", authenticatorAddress), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("token", request.JwtToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, getJWTinfoError
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var claims api.TokenClaims
	if err := json.Unmarshal(data, &claims); err != nil {
		return nil, err
	}

	userKeyBytes, err := hex.DecodeString(claims.SymmetricKey)
	if err != nil {
		return nil, err
	}
	ms.userKey = userKeyBytes

	conn, err := db.New(postgresURL)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	outChan, err := conn.GetUserPasswords(claims.UserID)
	if err != nil {
		return nil, err
	}

	userPwds, err := ms.decryptUserPassword(outChan)
	if err != nil {
		return nil, err
	}

	return userPwds, nil
}

func (ms *managerServer) decryptUserPassword(passwdChan <-chan db.Passwords) (*pb.UserPasswords, error) {
	passwds := new(pb.UserPasswords)
	for p := range passwdChan {
		pwByte, err := hex.DecodeString(p.EncryptedPassword)
		if err != nil {
			return nil, err
		}
		pw, err := passwords.DecryptPassword(ms.userKey, pwByte)
		if err != nil {
			return nil, err
		}

		passwds.Passwords = append(passwds.Passwords, &pb.PasswordItem{
			PasswordID: int64(p.ID),
			Data: &pb.PasswordData{
				ClearTextPassword: string(pw),
				Category:          p.Category,
				Site:              p.Site,
			},
		})

	}
	return passwds, nil
}

const serviceAddress = ":8989"

func main() {
	listener, err := net.Listen("tcp", serviceAddress)
	if err != nil {
		log.Fatalf("Cannot listen on %s [%v]", serviceAddress, err)
	}
	server := grpc.NewServer()
	pb.RegisterManagerServer(server, &managerServer{})
	log.Printf("listening on 127.0.0.0%v", serviceAddress)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Couldn't Serve on %s [%v]", serviceAddress, err)
	}
}
