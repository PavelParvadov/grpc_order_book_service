
FROM golang:1.24.4-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o api_gateway cmd/gateway/main.go


FROM alpine
WORKDIR /app
COPY --from=builder /app/api_gateway .
COPY config/config.yaml ./config.yaml
EXPOSE 8080
CMD ["./api_gateway"]