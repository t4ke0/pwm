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
	generator "github.com/t4ke0/pwm/pwm_manager/pw_generator"
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

	claims, err := getTokenInfo(request.JwtToken)
	if err != nil {
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
				Username:          p.Username,
				Category:          p.Category,
				Site:              p.Site,
			},
		})

	}
	return passwds, nil
}

// StorePassword store user password
func (ms *managerServer) StorePassword(ctx context.Context, req *pb.ManagerRequest) (*pb.Empty, error) {
	claims, err := getTokenInfo(req.JwtToken)
	if err != nil {
		return nil, err
	}

	userKey := claims.SymmetricKey
	userKeyByte, err := hex.DecodeString(userKey)
	if err != nil {
		return nil, err
	}
	encryptedPassword, err := passwords.EncryptPassword(userKeyByte, []byte(req.Password.ClearTextPassword))
	if err != nil {
		return nil, err
	}

	conn, err := db.New(postgresURL)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	if err := conn.StoreUserPassword(claims.UserID, db.Passwords{
		EncryptedPassword: hex.EncodeToString(encryptedPassword),
		Username:          req.Password.Username,
		Category:          req.Password.Category,
		Site:              req.Password.Site,
	}); err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

// UpdatePassword update user credentials `username, password, etc ....`
func (ms *managerServer) UpdatePassword(ctx context.Context, req *pb.ManagerUpdateRequest) (*pb.Empty, error) {
	if len(req.Mode) != len(req.Value) {
		return nil, fmt.Errorf("amount of values is not the same as the amount of the items to update")
	}

	conn, err := db.New(postgresURL)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	claims, err := getTokenInfo(req.JwtToken)
	if err != nil {
		return nil, err
	}

	var itemToUpdate = map[db.ElementToUpdate]string{}

	for index, v := range req.Mode {
		item := pb.ItemToUpdate_name[int32(v)]
		switch item {
		case "Password":
			userEncryptionKey, err := hex.DecodeString(claims.SymmetricKey)
			if err != nil {
				return nil, err
			}
			encryptedNewPassword, err := passwords.EncryptPassword(userEncryptionKey, []byte(req.Value[index]))
			if err != nil {
				return nil, err
			}
			itemToUpdate[db.Password] = hex.EncodeToString(encryptedNewPassword)
		case "Category":
			itemToUpdate[db.Category] = req.Value[index]
		case "Site":
			itemToUpdate[db.Site] = req.Value[index]
		case "Username":
			itemToUpdate[db.Username] = req.Value[index]
		}
	}

	if err := conn.UpdateUserPassword(claims.UserID, int(req.PasswordID), itemToUpdate); err != nil {
		return nil, err
	}
	return nil, nil
}

func getTokenInfo(jwtToken string) (api.TokenClaims, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/info", authenticatorAddress), nil)
	if err != nil {
		return api.TokenClaims{}, err
	}
	req.Header.Set("token", jwtToken)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return api.TokenClaims{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return api.TokenClaims{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return api.TokenClaims{}, fmt.Errorf("failed to get token info [%v] [%v]", resp.StatusCode, string(data))
	}

	var claims api.TokenClaims
	err = json.Unmarshal(data, &claims)

	return claims, err
}

// GeneratedPassword generates passwords
func (ms *managerServer) GeneratePassword(ctx context.Context, req *pb.GeneratePasswordRequest) (*pb.GeneratedPassword, error) {
	pw, err := generator.Generate(int(req.Length), req.Mode)
	if err != nil {
		return nil, err
	}
	return &pb.GeneratedPassword{
		Password: pw,
	}, nil
}

// DeletePassword delete a particular password.
func (ms *managerServer) DeletePassword(ctx context.Context, req *pb.DeletePasswordRequest) (*pb.Empty, error) {
	conn, err := db.New(postgresURL)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	claims, err := getTokenInfo(req.JwtToken)
	if err != nil {
		return nil, err
	}
	if err := conn.DeletePassword(claims.UserID, int(req.PasswordID)); err != nil {
		return nil, err
	}
	return nil, nil
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
