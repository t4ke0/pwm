package keys_manager_testing

import (
	"context"
	"testing"

	"google.golang.org/grpc"

	keys_manager_pb "github.com/t4ke0/pwm/keys_manager/proto"
)

const serverHost string = "localhost:9090"

func newGRPCconn() (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(serverHost, opts...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func queryKeyGen(mode keys_manager_pb.Mode) (*keys_manager_pb.KeyResponse, error) {
	conn, err := newGRPCconn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	keyManagerClient := keys_manager_pb.NewKeyManagerClient(conn)
	return keyManagerClient.GenKey(context.TODO(),
		&keys_manager_pb.KeyGenRequest{
			Mode: mode,
		})
}

func TestGenServerKey(t *testing.T) {
	key, err := queryKeyGen(keys_manager_pb.Mode_Server)
	if err != nil {
		t.Logf("generate server key [%v]", err)
		t.Fail()
		return
	}
	if key != nil && key.Key == "" {
		t.Logf("server key is empty")
		t.Fail()
	}
}

func TestGenAlreadyExistsServeKey(t *testing.T) {
	key, err := queryKeyGen(keys_manager_pb.Mode_Server)
	if err == nil || key != nil {
		t.Log("couldn't catch server key conflict")
		t.Fail()
		return
	}
}

func TestGenUserKey(t *testing.T) {
	key, err := queryKeyGen(keys_manager_pb.Mode_User)
	if err != nil {
		t.Logf("couldn't generate user key [%v]", err)
		t.Fail()
		return
	}

	if key == nil {
		t.Logf("couldn't generate user key [key is nil]")
		t.Fail()
		return
	}

	if key.Key == "" {
		t.Logf("couldn't generate user key [key is empty]")
		t.Fail()
		return
	}
	t.Logf("User key [%v]", key.Key)
}
