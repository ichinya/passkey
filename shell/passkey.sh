#!/bin/bash

# Проверка ключа
if [[ -z "$PASSCRYPT_KEY" ]]; then
  echo "Ошибка: переменная PASSCRYPT_KEY не задана"
  exit 1
fi

MODE=$1
DATA=$2

if [[ -z "$MODE" || -z "$DATA" ]]; then
  echo "Использование:"
  echo "  $0 e \"пароль для шифрования\""
  echo "  $0 d \"зашифрованный base64\""
  echo "  (ключ должен быть в переменной PASSCRYPT_KEY)"
  exit 1
fi

if [[ "$MODE" == "e" ]]; then
  echo -n "$DATA" | openssl enc -aes-256-cbc -a -salt -pbkdf2 -pass pass:"$PASSCRYPT_KEY"
elif [[ "$MODE" == "d" ]]; then
  echo "$DATA" | openssl enc -aes-256-cbc -a -d -salt -pbkdf2 -pass pass:"$PASSCRYPT_KEY"
else
  echo "Неизвестный режим: $MODE (используй e или d)"
  exit 1
fi
