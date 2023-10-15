.DEFAULT_GOAL := build
.PHONY:fmt
.PHONY:lint
.PHONY:vet
.PHONY:build
.PHONY:docker-build
.PHONY:docker-push

fmt:
	go fmt ./...

lint: fmt
	golint ./...

vet: fmt
	go vet ./...

build: vet
	go build ./cmd/sink 

test: build
	go test ./...

docker-build: test
	sudo docker buildx build -t thgsink:latest .

docker-push: docker-build
	sudo docker push rebay1982/thgsink:latest
