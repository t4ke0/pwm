BINARY = launcher


.PHONY: build
build:
	go mod download
	go build -o ${BINARY} 
	

.PHONY: clean
clean:
	rm ${BINARY}
