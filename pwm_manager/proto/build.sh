#!/bin/sh

set -xe

protoc --go_out=. --go_opt=paths=source_relative \
	   --go-grpc_out=. --go-grpc_opt=paths=source_relative \
		./pwm_manager.proto
