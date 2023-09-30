.DEFAULT_GOAL := build
.PHONY:fmt
.PHONY:lint
.PHONY:vet
.PHONY:build

fmt:
	go fmt ./...

lint: fmt
	golint ./...

vet: fmt
	go vet ./...

build: vet
	go build ./cmd/sink 

docker-build:
	sudo docker buildx build -t thgsink:latest .
