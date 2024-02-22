#!/bin/bash

## THIS SCRIPT SETUPS THE MACHINE DEPLOY ENVIRONMENT FOR THE PROJECT (WARN: THIS IS NOT INTEND TO BE RUN INSIDE DOCKER CONTAINER)
## THIS IS A IDEMPOTENT SCRIPT, CAN BE RUN MULTIPLE TIMES WITHOUT ANY SIDE EFFECTS

# Check for docker
if command -v docker &> /dev/null; then
    echo "Docker is installed."
else
    echo "Docker is not installed. Installing..."
    curl -fsSL https://get.docker.com -o get-docker.sh
    sudo sh get-docker.sh
    sudo groupadd docker
    sudo usermod -aG docker $USER
    echo "Docker is installed."
fi

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


echo "-------------- Setup complete. ----------------------------"
