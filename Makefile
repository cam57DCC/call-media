.PHONY:
.SILENT:
.DEFAULT_GOAL := run

build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/apiserver/main.go
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/consumer ./cmd/consumer/main.go

run: build
	docker-compose up --remove-orphans app supervisord