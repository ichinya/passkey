# Documentation

Использование утилиты `passcrypt` для шифрования и дешифрования паролей.

```bash
# Используя переменные окружения
PASSCRYPT_PASS="my-secret" PASSCRYPT_KEY="key123" go run main.go encrypt

# Расшифровка
PASSCRYPT_CIPHER="..." PASSCRYPT_KEY="key123" go run main.go decrypt

# Без переменных
go run main.go encrypt
```