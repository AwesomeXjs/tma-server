include .env
LOCAL_BIN:=$(CURDIR)/bin

run:
	go run cmd/app/main.go

run-debug:
	go run cmd/app/main.go -l debug

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

install-deps:
	mkdir -p $(LOCAL_BIN)
	GOBIN=$(LOCAL_BIN) go install github.com/gojuno/minimock/v3/cmd/minimock@v3.4.0 # mocks
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.60.3 # lint
	GOBIN=$(LOCAL_BIN) go install github.com/swaggo/swag/cmd/swag@latest # swagger
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0 # migrations


swagger:
	$(LOCAL_BIN)/swag init -g cmd/app/main.go


LOCAL_MIGRATION_DIR=$(CURDIR)/internal/migrations
LOCAL_MIGRATION_DSN="host=localhost port=$(POSTGRES_PORT) dbname=$(POSTGRES_DB) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) sslmode=disable"

# migrations
local-migration-status:
	${LOCAL_BIN}/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

local-migration-up:
	${LOCAL_BIN}/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

local-migration-down:
	${LOCAL_BIN}/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v

init-migration:
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} create create_first_table sql