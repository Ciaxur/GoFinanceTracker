BIN=./bin
SRC=./src
APP=app


# Compile Project
build:
	go build -o $(BIN)/$(APP) $(SRC)


# Run Tests
run-test:
	go test ./test


# Initialize all required dependencies
init:
	go get -v -t -d ./...
