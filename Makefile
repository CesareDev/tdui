SRC_DIR = ./cmd/tdui
SRC = main.go

BIN_DIR = ./bin
BIN = tdui 

all: main

main: $(SRC_DIR)/$(SRC)
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BIN) $(SRC_DIR)/$(SRC)

run:
	$(BIN_DIR)/$(BIN)
