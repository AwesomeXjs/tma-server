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


swagger:
	$(LOCAL_BIN)/swag init -g cmd/app/main.go