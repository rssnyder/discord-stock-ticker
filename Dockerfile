FROM golang:latest AS base
LABEL org.opencontainers.image.source https://github.com/rssnyder/discord-stock-ticker

RUN apt-get update && \
    apt-get -y --no-install-recommends install software-properties-common && \
    add-apt-repository "deb http://httpredir.debian.org/debian bullseye main" && \
    apt-get update && \
    apt-get -qq install -y libvips-dev && rm -rf /var/lib/apt/lists/*

WORKDIR /go/src/app

COPY . .

ARG TARGETOS=linux
ARG TARGETARCH=amd64

RUN GOOS=$TARGETOS GOARCH=$TARGETARCH CGO_ENABLED=1 go build -o /bin/ticker

ENTRYPOINT ["/bin/ticker"]
