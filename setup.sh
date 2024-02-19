#!/bin/bash

## THIS SCRIPT SETUPS THE MACHINE DEPLOY ENVIRONMENT FOR THE PROJECT (NOT DOCKER CONTAINER)
## FUNCTIONS ---------------------------
function is_database_empty {
    tables=$(sqlite3 ./db/ta_pago.db ".tables")
    if [ -z "$tables" ]; then
        return 0 # Database is empty
    else
        return 1 # Database is not empty
    fi
}

# Function to check if seed files exist
function seed_files_exist {
    if [ -f ./db/seed.sql ] || [ -f ./db/seed_update.sql ]; then
        return 0 # Seed files exist
    else
        return 1 # Seed files do not exist
    fi
}
## END FUNCTIONS -------------------------

# Check if Go is installed  (necessary to run migrate)
if command -v go &> /dev/null; then
    echo "Go is installed."
else
    echo "Go is not installed. Installing..."
    curl -L https://git.io/vQhTU | bash -s -- --version 1.21.5
    echo "Go is installed."
fi

# Check if make is installed
if command -v make &> /dev/null; then
    echo "make is already installed."
else
    echo "make is not installed. Installing..."
    sudo apt update
    sudo apt install -y make
fi

# Check if golang-migrate/migrate is installed
if command -v migrate &> /dev/null; then
    echo "golang-migrate/migrate is installed."
else
    echo "golang-migrate/migrate is not installed. Installing..."
    curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash
    sudo apt update
    sudo apt install -y migrate
fi

# Check if migrate has SQLite3 support
if migrate -help | grep -q 'sqlite3'; then
    echo "golang-migrate/migrate already installed SQLite3 driver."
else
    echo "golang-migrate/migrate does not have SQLite3 driver. Installing..."
    sudo apt update
    sudo apt install -y sqlite3 libsqlite3-dev
    go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
fi

# do migrations
sqlite3 ./db/ta_pago.db ""
chmod 666 ./db/ta_pago.db
make migration_exec

# Check if the database is empty to seed
if is_database_empty; then
    echo "Database is empty. running seeds..."
    # Seed the database
    echo "Seeding the database..."
    sqlite3 ./db/ta_pago.db < ./db/seed.sql
    sqlite3 ./db/ta_pago.db < ./db/seed_update.sql
else
    echo "Database is not empty. Skipping migrations and seeds."
fi


echo "-------------- Setup complete. ----------------------------"
