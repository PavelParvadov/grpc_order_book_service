FROM golang:1.24.4-alpine AS builder

WORKDIR /app


COPY . .

WORKDIR /app/order-service
RUN go mod download

RUN go build -o order-service ./cmd/order-service/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/order-service/order-service .
COPY --from=builder /app/order-service/config/config.yaml ./config.yaml

EXPOSE 5554
CMD ["./order-service"]
