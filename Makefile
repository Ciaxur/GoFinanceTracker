BIN=./bin
SRC=./src
APP=app


# Compile Project
build:
	go build -o $(BIN)/$(APP) $(SRC)


# Initialize all required dependencies
init:
	go get -v -t -d ./...
