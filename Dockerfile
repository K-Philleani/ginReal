FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn

RUN go build -o go-app main.go

ADD ./go-app /bin/bash

RUN chmod +x /bin/bash

EXPOSE 9090

CMD main
