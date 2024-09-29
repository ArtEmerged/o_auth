include .env

LOCAL_BIN:=$(CURDIR)/bin

LOCAL_MIGRATION_DIR = $(MIGRATION_DIR)
LOCAL_MIGRATION_DSN = "host=localhost port=$(DB_PORT) dbname=$(DB_NAME) user=$(DB_USER) password=$(DB_PASSWORD) sslmode=disable"

SERVER_HOST = $(SERVER_HOST)
SSH_USERNAME = $(SSH_USERNAME)

REGISTRY_URL = $(REGISTRY_URL)
REGISTRY_USER = $(REGISTRY_USER)
REGISTRY_PASSWORD = $(REGISTRY_PASSWORD)

lint:
	GOBIN=$(LOCAL_BIN) golangci-lint run ./... --config .golangci.pipeline.yaml

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.0
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc

install-goose:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate-api: generate-auth-api

generate-auth-api:
	mkdir -p pkg/auth_v1

	protoc --proto_path api/auth_v1 \
	--go_out=pkg/auth_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/auth_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/auth_v1/auth.proto

local-migration-status:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

local-migration-up:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

local-migration-down:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v

build:
	GOOS=linux GOARCH=amd64 go build -o auth_server cmd/main.go

copy-to-server:
	scp auth_server $(SSH_USERNAME)@$(SERVER_HOST):

docker-build:
	docker buildx build --no-cache --platform linux/amd64 -t $(REGISTRY_URL)/auth-server:v0.0.1 .
	docker login -u $(REGISTRY_USER) -p $(REGISTRY_PASSWORD) $(REGISTRY_URL)

test :
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
