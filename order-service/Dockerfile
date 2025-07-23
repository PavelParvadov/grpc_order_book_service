FROM golang:1.24.4-alpine AS builder

WORKDIR /build


RUN go mod download

COPY . .

WORKDIR /build/order-service
RUN go build -o order-service ./cmd/order-service/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /build/order-service/order-service .
COPY --from=builder /build/order-service/config/config.yaml ./config.yaml

EXPOSE 5554
CMD ["./order-service"]
