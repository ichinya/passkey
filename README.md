# 🔐 passkey

`passkey` — минималистичная утилита для шифрования, расшифровки и генерации паролей. Бинарник написан на Go, есть версия
на bash и Docker-образ. Все операции выполняются локально через ключ из переменной `PASSCRYPT_KEY`.

## Возможности

- `passkey e <строка>` — зашифровать строку
- `passkey d <cipher>` — расшифровать строку
`-mode shell|safe` — выбор режима шифрования (совместимый с OpenSSL `shell` или безопасный `safe`). По умолчанию используется режим `shell`.
- `passkey g` — генерация паролей
    - `--length`/`-l` — длина
    - `--level`/`-L` — уровень сложности (`low`, `medium`, `strong`, `paranoid`)
    - `--batch`/`-b` — количество паролей
    - `--encrypt`/`-e` — зашифровать результат
  **Важно:** для совместимости с bash-скриптом и `openssl` используйте `-mode shell` при шифровании и расшифровке. По
  умолчанию используется безопасный режим `safe`.

## Установка

```bash
git clone https://github.com/Ichinya/passkey
cd passkey
make build
export PASSCRYPT_KEY=mykey
./passkey e "secret"
```

Через Docker:

```bash
docker run --rm -e PASSCRYPT_KEY=mykey passkey e "secret"
```

Через bash без установки:

```bash
curl -sSL https://raw.githubusercontent.com/Ichinya/passkey/main/shell/passkey.sh | PASSCRYPT_KEY=mykey bash -s e "secret"
```

## Примеры

```bash
# Шифрование
PASSCRYPT_KEY=mykey ./passkey e -mode safe "mypassword"
PASSCRYPT_KEY=mykey ./passkey e "mypassword"

# Расшифровка
PASSCRYPT_KEY=mykey ./passkey d -mode safe "cipher_safe"
PASSCRYPT_KEY=mykey ./passkey d "cipher"

# Генерация 10 паролей и их шифрование
PASSCRYPT_KEY=mykey ./passkey g -b 10 -e
```

## Примеры через Docker

```bash
# Шифрование
docker run --rm -e PASSCRYPT_KEY=mykey ichinya/passkey e "mypassword
PASSCRYPT_KEY=sd make encrypt ARGS="ваш текст"
# Расшифровка
docker run --rm -e PASSCRYPT_KEY=mykey ichinya/passkey d "cipher"
# Генерация 10 паролей и их шифрование
docker run --rm -e PASSCRYPT_KEY=mykey passkey g  -batch 10
```

## Лицензия

MIT
