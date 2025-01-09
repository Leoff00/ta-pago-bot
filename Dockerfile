# BUILD STAGE
FROM golang:1.22.0 AS base
WORKDIR /app

# Set timezone
ENV TZ="America/Sao_Paulo"

# Create directories for database migrations
RUN mkdir -p ./db/migrations

# Copy source code and Go modules
COPY . .
RUN go mod download 

# Install SQLite3 dependencies
RUN apt-get update && apt-get install -y sqlite3 libsqlite3-dev

# Install the migrate CLI with SQLite3 support
RUN go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.1

# Run migrations
RUN migrate -path ./db/migrations -database sqlite3://./db/ta_pago.db up

# Build the Go binary with CGO enabled
RUN CGO_ENABLED=1 GOOS=linux go build -o tapagobot .

# DEPLOY STAGE
FROM debian:bookworm-slim
WORKDIR /app

# Set timezone
ENV TZ="America/Sao_Paulo"

# Install runtime dependencies (including SQLite and glibc)
RUN apt-get update && apt-get install -y \
    ca-certificates \
    sqlite3 \
    libsqlite3-0 && \
    rm -rf /var/lib/apt/lists/*

# Copy the binary and migrations directory
COPY --from=base /app/tapagobot .
COPY --from=base /app/db/migrations ./db/migrations

# Expose the application port
EXPOSE 4000

# Run the application
ENTRYPOINT ["/app/tapagobot"]
