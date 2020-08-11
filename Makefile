BINARY = launcher
ROUTER = ./router

.PHONY: build
build:
	go mod download
	go build -o ${BINARY} 
	
.PHONY: router
router:
	go mod download
	go build -o ${ROUTER} ./server/main.go

.PHONY: clean
clean:
	rm ${BINARY}
