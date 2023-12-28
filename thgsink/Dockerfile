# build with 'docker buildx build -t thgsink:latest.'
FROM golang:bookworm AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
RUN GOOS=linux go build -tags=musl ./cmd/sink 

RUN apt update
# RUN apt install strace

CMD ["/bin/sh"]
