FROM golang:alpine

ENV GO111MODULE=on
    GOOS=linux
    GOARCH=amd64

WORKDIR /build

COPY . .

RUN go build -o app .

WORKDIR /dist

RUN cp /build/app .

EXPOSE 9090

CMD ["/dist/app"]

