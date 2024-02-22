# BUILD STAGE
FROM golang:1.21.5
ENV TZ="America/Sao_Paulo"
WORKDIR /app

# Go dependencies download
COPY go.mod go.sum ./
RUN go mod download && go mod verify
# Copy everything from PWD build to /app container
COPY . .
# Remove unwanted acidentally files
RUN find . -type f \( -name '*.db' -o -name '*.sqlite' -o -name '*.sqlite3' -o -name '*.sql' -o -name '*.env' -o -name 'ta_pago_bot' -o -name 'tenant.json' \) -delete
# Remove db directory if it exists
RUN if [ -d "/app/db" ]; then rm -rf /app/db; fi
# Create directories for database migrations
RUN mkdir -p ./db/migrations
# SQLite3 dependencies
RUN apt-get update && apt-get install -y sqlite3 libsqlite3-dev
# Build the go app to /app/ta_pago_bot
RUN go build -o ./ta_pago_bot ./cmd/main.go

EXPOSE 4000

CMD ["./ta_pago_bot"]
