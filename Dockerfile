FROM golang:1.16 AS build
LABEL org.opencontainers.image.source https://github.com/rssnyder/discord-stock-ticker

WORKDIR /go/src/app

COPY . .

RUN CGO_ENABLED=0 go build -o /bin/ticker

FROM scratch
COPY --from=build /bin/ticker /bin/ticker
EXPOSE 8080

ENTRYPOINT ["/bin/ticker"]
