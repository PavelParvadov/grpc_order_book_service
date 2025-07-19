FROM golang:1.24.4-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o book-service ./cmd/book-service/main.go && \
    go build -o migrator cmd/migrator/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/book-service .
COPY --from=builder /app/migrator .
COPY config/config.yaml ./config.yaml
COPY migrations ./migrations

EXPOSE 5555

CMD ["./book-service"]




