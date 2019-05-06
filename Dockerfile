FROM golang:latest

MAINTAINER "whywaita <https://github.com/whywaita>"

WORKDIR /go/src/github.com/whywaita/google-drive-checker

COPY go.mod go.mod
ENV GO111MODULE=on
RUN go mod download

COPY . .

ENTRYPOINT go test -v ./...