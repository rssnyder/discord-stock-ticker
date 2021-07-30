FROM golang:1.16-alpine AS build
LABEL org.opencontainers.image.source https://github.com/rssnyder/discord-stock-ticker

RUN apk --no-cache add ca-certificates

WORKDIR /go/src/app

COPY . .

RUN CGO_ENABLED=0 go build -o /bin/ticker

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /bin/ticker /bin/ticker
EXPOSE 8080

ENTRYPOINT ["/bin/ticker", "-address", "0.0.0.0:8080"]
