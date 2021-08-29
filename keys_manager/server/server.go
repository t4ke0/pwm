package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	pb "github.com/t4ke0/pwm/keys_manager/proto"

	"github.com/t4ke0/pwm/keys_manager/common"
	"github.com/t4ke0/pwm/keys_manager/keys_manager_errors"
	db "github.com/t4ke0/pwm/pwm_db_api"
)

const port = 9090

var (
	wordListFilePath = os.Getenv("WORD_LIST_PATH")
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
	//
	test   = os.Getenv("TEST")
	isTest = (test == "true")
)

type KeyManagerServer struct {
	pb.UnimplementedKeyManagerServer
}

func (s *KeyManagerServer) GenKey(ctx context.Context,
	genRequest *pb.KeyGenRequest) (*pb.KeyResponse, error) {
	conn, err := db.New(postgresURL)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	switch genRequest.Mode {
	case pb.Mode_Server:
		key, err := conn.GetServerEncryptionKey()
		if err != nil {
			return nil, err
		}
		if key == "" {
			generatedServerKey, err := common.GenerateEncryptionKey(wordListFilePath,
				int(genRequest.Size))
			if err != nil {
				return nil, err
			}
			if err := conn.StoreServerKey(generatedServerKey.String()); err != nil {
				return nil, err
			}
			return &pb.KeyResponse{
				Key: generatedServerKey.String(),
			}, nil
		}

		return nil, keys_manager_errors.ErrKeyAlreadyExists

	case pb.Mode_ServerAuth:
		// try to load the auth server key.
		key, err := conn.GetAuthServerKey()
		// if the auth server key doesn't exists in the db
		// we generate it and store it into db.
		if err != nil {
			return nil, err
		}
		if key == "" {
			generatedAuthkey, err := common.GenerateEncryptionKey(
				wordListFilePath, int(genRequest.Size))
			if err != nil {
				return nil, err
			}
			if err := conn.StoreAuthServerKey(generatedAuthkey.String()); err != nil {
				return nil, err
			}
			return &pb.KeyResponse{
				Key: generatedAuthkey.String(),
			}, nil
		}
		return nil, keys_manager_errors.ErrKeyAlreadyExists

	case pb.Mode_User:
		encodedServerKey, err := conn.GetServerEncryptionKey()
		if err != nil {
			return nil, keys_manager_errors.ErrServerKeyNotExists
		}

		serverKey, err := common.DecodeStringKey(encodedServerKey)
		if err != nil {
			return nil, err
		}
		userKey, err := common.GenerateEncryptionKey(wordListFilePath,
			int(genRequest.Size))
		if err != nil {
			return nil, err
		}
		key, err := serverKey.Encrypt(userKey)
		if err != nil {
			return nil, err
		}
		return &pb.KeyResponse{
			Key: common.Key(key).String(),
		}, nil

	default:
		return nil, fmt.Errorf("no mode has been set")
	}
}

func (s *KeyManagerServer) GetUserKey(ctx context.Context,
	fetchMsg *pb.KeyFetchRequest) (*pb.KeyResponse, error) {

	conn, err := db.New(postgresURL)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	serverStrkey, err := conn.GetServerEncryptionKey()
	if err != nil {
		return nil, err
	}
	userStrkey, err := conn.LoadUserKey(fetchMsg.Username)
	if err != nil && err == db.ErrNoRows {
		return nil, fmt.Errorf("user's key not found")
	} else if err != nil {
		return nil, err
	}
	// decrypted user key to return
	serverKey, err := common.DecodeStringKey(serverStrkey)
	if err != nil {
		return nil, err
	}
	userKey, err := common.DecodeStringKey(userStrkey)
	if err != nil {
		return nil, err
	}
	key, err := serverKey.Decrypt(userKey)
	if err != nil {
		return nil, err
	}
	return &pb.KeyResponse{
		Key: key.String(),
	}, nil
}

func init() {
	// Verify env vars
	for _, arg := range []string{
		"WORD_LIST_PATH",
		"POSTGRES_HOST",
		"POSTGRES_DATABASE",
		"POSTGRES_USER",
		"POSTGRES_PASSWORD",
	} {
		if os.Getenv(arg) == "" {
			panic(fmt.Sprintf("%v env variable is not set", arg))
		}
	}
	if isTest {
		testPostgresPath, err := db.CreateTestingDatabase(postgresURL)
		if err != nil {
			panic(fmt.Sprintf("Failed To create test database [%v]", err))
		}
		postgresURL = testPostgresPath
		log.Printf("DEBUG POSTGRES_URL = %v", postgresURL)
	}
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal("Listener: %v", err)
	}
	log.Printf("Service Listening on %d ...", port)
	server := grpc.NewServer()
	pb.RegisterKeyManagerServer(server, &KeyManagerServer{})
	if err := server.Serve(listener); err != nil {
		log.Fatal("grpc serve: %v", err)
	}
}
