# Используем минимальный образ с Go
FROM golang:1.21-alpine as builder

WORKDIR /app
COPY . .
RUN go build -o passkey go/main.go encrypt.go

# Финальный образ: только бинарник
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/passkey /usr/local/bin/passkey

# Установка переменной по умолчанию
ENV PASSCRYPT_KEY=""

ENTRYPOINT ["/usr/local/bin/passkey"]
