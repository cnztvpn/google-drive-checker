FROM golang:latest as builder

MAINTAINER "whywaita <https://github.com/whywaita>"

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/github.com/whywaita/google-drive-checker

COPY go.mod go.mod
ENV GO111MODULE=on
RUN go mod download

COPY . .

ENTRYPOINT go test -v ./...