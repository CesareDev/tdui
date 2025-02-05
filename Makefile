SRC_DIR = ./cmd
BIN_DIR = ./bin
BIN = tdui

all: build 

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR) $(SRC_DIR)/...

run:
	$(BIN_DIR)/$(BIN)
