.DEFAULT_GOAL := build
.PHONY:fmt
.PHONY:lint
.PHONY:vet
.PHONY:build
.PHONY:docker-build

IMAGE_NAME ?= thgsink
IMAGE_TAG ?= latest

fmt:
	go fmt ./...

lint: fmt
	golint ./...

vet: fmt
	go vet ./...

build: vet
	go build ./cmd/sink 

docker-build:
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .

docker-build-dev:
	docker build -t $(IMAGE_NAME)-dev:$(IMAGE_TAG) .
