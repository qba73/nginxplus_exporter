VERSION 0.6
FROM golang:1.19
WORKDIR /nginx_exporter

deps:
    COPY go.mod go.sum ./
    COPY collector/*.go ./collector/
    COPY client/*.go ./client/
    RUN go mod download
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum

build:
    COPY exporter.go .
    RUN go build -o build/nginx_exporter exporter.go
    SAVE ARTIFACT build/nginx_exporter /nginx_exporter AS LOCAL build/nginx_exporter

unit-test:
    FROM +deps
    COPY exporter.go exporter_internal_test.go ./
    RUN go test -v ./...
