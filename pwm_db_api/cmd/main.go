package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"

	"github.com/t4ke0/pwm/keys_manager/common"
	keys_manager_pb "github.com/t4ke0/pwm/keys_manager/proto"

	db "github.com/t4ke0/pwm/pwm_db_api"
)

func main() {
	var (
		postgresPW   = os.Getenv("POSTGRES_PASSWORD")
		postgresDB   = os.Getenv("POSTGRES_DATABASE")
		postgresUser = os.Getenv("POSTGRES_USER")
		postgresHost = os.Getenv("POSTGRES_HOST")

		postgresLink = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
			postgresUser,
			postgresPW,
			postgresHost,
			postgresDB)
		//
		keysManagerHost = os.Getenv("KEYS_MANAGER_HOST")
	)

	log.Printf("DEBUG postgres Link [%v]", postgresLink)

	// When starting inital the database. support [prod & test] env
	conn, err := db.New(postgresLink)
	if err != nil {
		log.Fatalf("Couldn't connect to the Database %v", err)
	}
	defer conn.Close()
	db.SchemaFile = os.Getenv("SCHEMA_FILE_PATH")
	if err := conn.InitDB(); err != nil {
		log.Fatalf("Couldn't init the Database %v", err)
	}
	log.Printf("DEBUG initialized DB successfully!")

	if os.Getenv("TEST") == "true" {
		//		pqLink, err := db.CreateTestingDatabase(postgresLink)
		//		if err != nil {
		//			log.Fatal(err)
		//		}
		//		postgresLink = pqLink
		wordsFilePath := "../../keys_manager/common/words.txt"
		key, err := common.GenerateEncryptionKey(wordsFilePath, 0)
		if err != nil {
			log.Fatalf("Error generating server key %v", err)
		}
		if err := conn.StoreServerKey(key.String()); err != nil {
			log.Fatalf("Error storing server key into database %v", err)
		}

	}

	serverKey, err := conn.GetServerEncryptionKey()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DEBUG %v", serverKey == "")
	if serverKey == "" {
		grpcConn, err := grpc.Dial(keysManagerHost, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("GRPC ERROR: %v", err)
		}
		defer grpcConn.Close()
		client := keys_manager_pb.NewKeyManagerClient(grpcConn)
		key, err := client.GenKey(context.TODO(), &keys_manager_pb.KeyGenRequest{
			Mode: keys_manager_pb.Mode_Server,
		})
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("DEBUG gen server key %v", key.Key)
		return
	}
	log.Printf("DEBUG: server key %v", serverKey)
}
