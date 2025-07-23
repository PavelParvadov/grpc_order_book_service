FROM golang:1.24.4-alpine AS builder

WORKDIR /app

COPY book-service/go.mod book-service/go.sum ./
RUN go mod download

COPY book-service/. .

RUN go build -o book-service ./cmd/book-service/main.go && \
    go build -o migrator ./cmd/migrator/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/book-service .
COPY --from=builder /app/migrator .
COPY book-service/config/config.yaml ./config.yaml
COPY book-service/migrations ./migrations

EXPOSE 5555
CMD ["./book-service"]