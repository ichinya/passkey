# Passcrypt

## Утилита `passcrypt` на Go

Использование утилиты `passcrypt` для шифрования и дешифрования паролей.

```bash
# Шифрование (e = encrypt)
PASSCRYPT_KEY="ключ" ./passkey e "my-password"

# Расшифровка (d = decrypt)
PASSCRYPT_KEY="ключ" ./passkey d "base64-cipher"

# Без переменных
go run main.go encrypt
```

### Сборка

```bash
go build -o passkey go/main.go
```

Сделать исполняемый файл можно с помощью команды `go build`, которая создаст файл `passkey` в текущей директории.

```shell
chmod +x passkey.sh
sudo mv passkey.sh /usr/local/bin/passkey

```

## Использование через скрипт

```bash
export PASSCRYPT_KEY="superkey"

# Шифрование
./passkey.sh e "my-password"
# => U2FsdGVkX1+...

# Расшифровка
./passkey.sh d "U2FsdGVkX1+..."
# => my-password
```

### Запуск с curl/wget

```shell
PASSCRYPT_KEY="yourkey" bash <(wget -qO- https://raw.githubusercontent.com/ichinya/passkey/main/shell/passkey.sh) d "ciphertext"
```

### Установка в систему

```bash
curl -sSL https://raw.githubusercontent.com/ichinya/passkey/main/shell/install.sh | bash
```