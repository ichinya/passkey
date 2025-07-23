# 🔐 passkey

`passkey` — минималистичная утилита для шифрования, расшифровки и генерации паролей. Бинарник написан на Go, есть версия на bash и Docker-образ. Все операции выполняются локально через ключ из переменной `PASSCRYPT_KEY`.

## Возможности

- `passkey e <строка>` — зашифровать строку
- `passkey d <cipher>` — расшифровать строку
- `passkey g` — генерация паролей
  - `--length`/`-l` — длина
  - `--level`/`-L` — уровень сложности (`low`, `medium`, `strong`, `paranoid`)
  - `--batch`/`-b` — количество паролей
  - `--encrypt`/`-e` — зашифровать результат
- `--mode shell|safe` — выбор режима шифрования (совместимый с OpenSSL `shell` или безопасный `safe`)

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
docker run --rm -e PASSCRYPT_KEY=mykey ghcr.io/ichinya/passkey e "secret"
```

Через bash без установки:

```bash
curl -sSL https://raw.githubusercontent.com/Ichinya/passkey/main/shell/passkey.sh | PASSCRYPT_KEY=mykey bash -s e "secret"
```

## Примеры

```bash
# Шифрование
PASSCRYPT_KEY=mykey ./passkey e "mypassword"

# Расшифровка
PASSCRYPT_KEY=mykey ./passkey d "cipher"

# Генерация 10 паролей и их шифрование
PASSCRYPT_KEY=mykey ./passkey g -b 10 -e
```

## Лицензия

MIT

