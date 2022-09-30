CURRENT_DIR=$(shell pwd)
APP=template
APP_CMD_DIR=./cmd

build:
	CGO_ENABLED=0 GOOS=darwin go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

lint: ## Run golangci-lint with printing to stdout
	golangci-lint -c .golangci.yaml run --build-tags "musl" ./...

swag-gen:
	echo ${REGISTRY}
	swag init -g api/router.go -o api/docs


create-migrations:
	migrate create -ext sql -dir migrations -seq addtokencolumn

migrate-up:
	migrate -path migrations/ -database postgres://postgres:1@localhost:5432/test?sslmode=disable up

migrate-down:
	migrate -path migrations/ -database postgres://postgres:1@localhost:5432/test?sslmode=disable down

migrate-forse:
	migrate -path migrations/ -database postgres://postgres:1@localhost:5432/test?sslmode=disable force 0
