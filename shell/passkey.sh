#!/bin/bash

# passkey.sh — CLI для шифрования и расшифровки паролей
if [[ -z "$PASSCRYPT_KEY" ]]; then
  echo "❌ Ошибка: переменная окружения PASSCRYPT_KEY не задана"
  exit 1
fi

MODE=$1
INPUT=$2

if [[ -z "$MODE" || -z "$INPUT" ]]; then
  echo "Использование:"
  echo "  $0 e \"пароль для шифрования\""
  echo "  $0 d \"base64-зашифрованный текст\""
  echo "Пример:"
  echo "  PASSCRYPT_KEY=key123 $0 e \"mypassword\""
  exit 1
fi

if [[ "$MODE" == "e" ]]; then
  echo -n "$INPUT" | openssl enc -aes-256-cbc -a -salt -pbkdf2 -pass pass:"$PASSCRYPT_KEY"
elif [[ "$MODE" == "d" ]]; then
  echo "$INPUT" | openssl enc -aes-256-cbc -a -d -salt -pbkdf2 -pass pass:"$PASSCRYPT_KEY"
else
  echo "❌ Неизвестный режим: $MODE (допустимы e или d)"
  exit 1
fi
