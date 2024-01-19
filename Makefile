TARGET_DIR=bin
GOBUILD=go build
MAIN_DIR=cmd
BINARY_NAME=go_ta_pago_bot
MIGRATIONS_DIR=database/migrations

install:
	go get ./...

local: db_up
	go run $(MAIN_DIR)/main.go

local_no_db:
	go run $(MAIN_DIR)/main.go

db_up:
	docker compose up db -d

db_down:
	docker compose down db 

migrate_create:
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(migration_name)

migrate_up:
	docker compose up db_migrate

build:
	$(GOBUILD) -o $(TARGET_DIR)/$(BINARY_NAME) 

run: db_up
	./$(TARGET_DIR)/$(BINARY_NAME)

clean:
	go clean
	rm -f $(TARGET_DIR)/$(BINARY_NAME)

release: build run

.PHONY: local build run release install db_up db_down