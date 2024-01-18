TARGET_DIR=bin
GOBUILD=go build
MAIN_DIR=cmd
BINARY_NAME=go_ta_pago_bot

install:
	go get ./...

local:
	go run $(MAIN_DIR)/main.go

build:
	$(GOBUILD) -o $(TARGET_DIR)/$(BINARY_NAME) 

run:
	./$(TARGET_DIR)/$(BINARY_NAME)

clean:
	go clean
	rm -f $(TARGET_DIR)/$(BINARY_NAME)


release: build run

.PHONY: local build run release install