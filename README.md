# Documentation

Использование утилиты `passcrypt` для шифрования и дешифрования паролей.

```bash
# Шифрование (e = encrypt)
PASSCRYPT_KEY="ключ" ./passkey e "my-password"

# Расшифровка (d = decrypt)
PASSCRYPT_KEY="ключ" ./passkey d "base64-cipher"

# Без переменных
go run main.go encrypt
```

## Сборка

```bash
go build -o passkey go/main.go
```