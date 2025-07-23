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
  echo "  $0 g [--length <длина>] [--level <уровень>]"
  echo "Пример:"
  echo "  PASSCRYPT_KEY=key123 $0 e \"mypassword\""
  exit 1
fi

if [[ "$MODE" == "e" ]]; then
  echo -n "$INPUT" | openssl enc -aes-256-cbc -a -salt -pbkdf2 -pass pass:"$PASSCRYPT_KEY"
elif [[ "$MODE" == "d" ]]; then
  echo "$INPUT" | openssl enc -aes-256-cbc -a -d -salt -pbkdf2 -pass pass:"$PASSCRYPT_KEY"
if [[ "$MODE" == "g" ]]; then
  LENGTH=16
  LEVEL="medium"
  while [[ "$#" -gt 0 ]]; do
    case $1 in
      -l|--length) LENGTH="$2"; shift ;;
      -L|--level) LEVEL="$2"; shift ;;
      *) ;;
    esac
    shift
  done

  LETTERS="abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
  NUMBERS="0123456789"
  SYMBOLS_BASIC="@#-_+="
  SYMBOLS_ADVANCED="!@#\$%^&*()[]{}<>?~"
  UNICODE_PARANOID="🔐💡🧠🌐🎲🚀你好†±λ@#€"

  case "$LEVEL" in
    low) CHARS="${LETTERS}${NUMBERS}" ;;
    medium) CHARS="${LETTERS}${NUMBERS}${SYMBOLS_BASIC}" ;;
    strong) CHARS="${LETTERS}${NUMBERS}${SYMBOLS_BASIC}${SYMBOLS_ADVANCED}" ;;
    paranoid) CHARS="${LETTERS}${NUMBERS}${SYMBOLS_BASIC}${SYMBOLS_ADVANCED}${UNICODE_PARANOID}" ;;
    *) echo "❌ Неизвестный уровень сложности: $LEVEL"; exit 1 ;;
  esac

  PASSWORD=""
  for i in $(seq 1 $LENGTH); do
    INDEX=$(( RANDOM % ${#CHARS} ))
    CHAR="${CHARS:INDEX:1}"
    PASSWORD+="$CHAR"
  done

  echo "$PASSWORD"
  exit 0
fi
else
  echo "❌ Неизвестный режим: $MODE (допустимы e или d)"
  exit 1
fi