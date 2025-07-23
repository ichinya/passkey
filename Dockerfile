# Используем минимальный образ с Go
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY . .
WORKDIR /app/go
RUN ls -l /app/go
RUN go mod download
RUN go build -o ../passkey

# Финальный образ: только бинарник
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/passkey /usr/local/bin/passkey

# Установка переменной по умолчанию
ENV PASSCRYPT_KEY=""

ENTRYPOINT ["/usr/local/bin/passkey"]
