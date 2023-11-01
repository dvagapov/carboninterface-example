BINARY_NAME=carboninterface-example

build:
	go build -o ${BINARY_NAME} main.go

clean:
	go clean