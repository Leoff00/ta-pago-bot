TARGET_DIR=bin
GOBUILD=go build
MAIN_DIR=cmd
BINARY_NAME=go_ta_pago_bot

install:
	go get ./...

local: db_up
	go run $(MAIN_DIR)/main.go

db_up:
	docker compose up db -d

db_down:
	docker compose down db 

build:
	$(GOBUILD) -o $(TARGET_DIR)/$(BINARY_NAME) 

run: db_up
	./$(TARGET_DIR)/$(BINARY_NAME)

clean:
	go clean
	rm -f $(TARGET_DIR)/$(BINARY_NAME)


release: build run

.PHONY: local build run release install db_up db_down