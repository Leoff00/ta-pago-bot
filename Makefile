TARGET_DIR=bin
GOBUILD=go build
MAIN_DIR=./cmd
BINARY_NAME=go_ta_pago_bot
MIGRATIONS_DIR=internal/db/migrations
MIGRATION_NAME=init_mg

install:
	go get ./... && go install ./...

local: migration_exec
	go run $(MAIN_DIR)/main.go

migration_create:
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(MIGRATION_NAME)

migration_exec:
	migrate -path $(MIGRATIONS_DIR) -database "sqlite3://./ta_pago.db" up

build:
	$(GOBUILD) -o $(TARGET_DIR)/$(BINARY_NAME) $(MAIN_DIR)

run:
	./$(TARGET_DIR)/$(BINARY_NAME)

clean:
	go clean
	rm -f $(TARGET_DIR)/$(BINARY_NAME)

release: build run

.PHONY: local build run release install migration_create migration_exec