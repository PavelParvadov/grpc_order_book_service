FROM golang:1.24.4-alpine AS builder

WORKDIR /app
COPY . .

RUN go build -o order-service ./cmd/order-service/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/order-service .
COPY config/config.yaml ./config.yaml

EXPOSE 5554

CMD ["./order-service"]
