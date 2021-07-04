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
		_, err := conn.GetStoredServerKey()
		if err != nil && err != db.ErrNoRows {
			return nil, fmt.Errorf("server key is already exists in the database")
		}
		if err != nil {
			return nil, err
		}
		serverKey, err := common.GenerateEncryptionKey(wordListFilePath,
			int(genRequest.Size))
		if err != nil {
			return nil, err
		}

		if err := conn.StoreServerKey(serverKey.String()); err != nil {
			return nil, err
		}

		return &pb.KeyResponse{
			Key: serverKey.String(),
		}, nil

	case pb.Mode_User:
		encodedServerKey, err := conn.GetStoredServerKey()
		if err != nil && err == db.ErrNoRows {
			// we need server key to encrypt user key.
			return nil, fmt.Errorf("server key is not yet generated.")
		}
		if err != nil {
			return nil, err
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

	userkey, err := conn.LoadUserKey(fetchMsg.Username)
	if err != nil && err == db.ErrNoRows {
		return nil, fmt.Errorf("user's key not found")
	} else if err != nil {
		return nil, err
	}
	return &pb.KeyResponse{
		Key: userkey,
	}, nil
}

func init() {
	// Verify env vars
	for _, arg := range []string{
		"WORD_LIST_PATH",
		"POSTGRES_HOST",
		"POSTGRES_DB",
		"POSTGRES_USER",
		"POSTGRES_PASSWORD",
	} {
		if os.Getenv(arg) == "" {
			panic(fmt.Sprintf("%v env variable is not set", arg))
		}
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
