FROM golang:1.26-trixie AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/thgsink ./cmd/sink

FROM gcr.io/distroless/static-debian13:nonroot

COPY --from=build /out/thgsink /thgsink


USER nonroot:nonroot
ENTRYPOINT ["/thgsink"]
