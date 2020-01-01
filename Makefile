BIN=./bin
SRC=./src
APP=app
DEPS="github.com/fatih/color" "github.com/manifoldco/promptui"


# Compile Project
build:
	go build -o $(BIN)/$(APP) $(SRC)


# Initialize all required dependencies
init:
	go get $(DEPS)