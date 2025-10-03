# 1. Билдим Go бинарь
FROM golang:1.24.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o task-service ./cmd/server/main.go

# 2. Минимальный рантайм
FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/task-service .

CMD ["./task-service"]
