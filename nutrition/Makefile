LOCAL_BIN:=$(CURDIR)/bin
MIGRATION_DIR:=$(CURDIR)/migrations

include .env

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install -mod=mod github.com/gojuno/minimock/v3/cmd/minimock@v3.3.12
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.20.0

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc



generate:
	make generate-nutrition-api



generate-nutrition-api:
	mkdir -p pkg/nutrition_v1
	protoc --proto_path api/nutrition_v1 \
	--go_out=pkg/nutrition_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/nutrition_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/nutrition_v1/nutrition.proto

.PHONY: test
test:
	go clean -testcache
	go test ./... -covermode count -coverpkg=github.com/gerathewitcher/healthstyle/nutrition/internal/service/...,github.com/gerathewitcher/healthstyle/nutrition/internal/api/... -count 5

.PHONY: test-coverage
test-coverage:
	go clean -testcache
	go test ./... -coverprofile-coverage.tmp.out -covermode count -coverpkg=github.com/gerathewitcher/healthstyle/nutrition/internal/service/...,github.com/gerathewitcher/healthstyle/nutrition/internal/api/... -count 5
	grep -v 'mocks\|config' coverage.tmp.out > coverage.out
	rm coverage.tmp.out
	go tool cover -html=coverage.out;
	go tool cover -func=./coverage.out | grep "total";
	grep -sqFx "/coverage.out" .gitignore || echo "/coverage.out" >> .gitignore


local-migrations-status:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${PG_DSN} status -v

local-migration-create:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} create ${MIGRATION_NAME} sql

local-migration-up:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${PG_DSN} up -v

local-migration-down:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${PG_DSN} down -v