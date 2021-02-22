FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn


WORKDIR /build

COPY . .

RUN go build -o app .

WORKDIR /dist

RUN cp /build/app .

FROM scrach

COPY --from=builder /build/app /

ENTRYPOINT ["/app"]

