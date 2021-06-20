package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/t4ke0/pwm/keys_manager/common"
	pb "github.com/t4ke0/pwm/keys_manager/proto"
)

const port = 9090

var (
	wordListFilePath = os.Getenv("WORD_LIST_PATH")
)

type KeyManagerServer struct {
	pb.UnimplementedKeyManagerServer
}

// NOTE: will be replaced with the server key that we will get from the DB.
var tempServerKey common.Key

func (s *KeyManagerServer) GenKey(ctx context.Context,
	genRequest *pb.KeyGenRequest) (*pb.KeyResponse, error) {

	if wordListFilePath == "" {
		return nil, fmt.Errorf("word list path not set")
	}

	switch genRequest.Mode {
	case pb.Mode_Server:
		// TODO: check if we already generated a server key from the postgres db.
		// for now we are just generating new key each time.
		// for the sake of testing.
		serverKey, err := common.GenerateEncryptionKey(wordListFilePath,
			int(genRequest.Size))
		if err != nil {
			return nil, err
		}
		tempServerKey = serverKey
		return &pb.KeyResponse{
			Key: serverKey.String(),
		}, nil
	case pb.Mode_User:
		// TODO load server key from db so we can encrypt the user key with it
		// then we serve it .
		userKey, err := common.GenerateEncryptionKey(wordListFilePath,
			int(genRequest.Size))
		if err != nil {
			return nil, err
		}
		if tempServerKey == nil {
			return nil, fmt.Errorf("server's key is not yet generated")
		}
		key, err := tempServerKey.Encrypt(userKey)
		if err != nil {
			return nil, err
		}
		return &pb.KeyResponse{
			Key: common.Key(key).String(),
		}, nil
	}
	return nil, fmt.Errorf("no mode has been set")
}

func (s *KeyManagerServer) GetUserKey(ctx context.Context,
	fetchMsg *pb.KeyFetchRequest) (*pb.KeyResponse, error) {

	// TODO: get user key from database.
	return nil, nil
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
