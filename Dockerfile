FROM golang:latest AS base

WORKDIR /go/src/discord-stock-ticker

COPY . .

ARG TARGETOS=linux
ARG TARGETARCH=amd64

RUN GOOS=$TARGETOS GOARCH=$TARGETARCH CGO_ENABLED=1 go build -o discord-stock-ticker

ENTRYPOINT ["/go/src/discord-stock-ticker/discord-stock-ticker"]
