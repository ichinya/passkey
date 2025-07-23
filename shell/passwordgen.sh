#!/bin/bash

LENGTH=16
LEVEL="medium"

# Парсинг аргументов
while [[ "$#" -gt 0 ]]; do
  case $1 in
    -l|--length) LENGTH="$2"; shift ;;
    -L|--level) LEVEL="$2"; shift ;;
    *) echo "Неизвестный аргумент: $1" && exit 1 ;;
  esac
  shift
done

# Символьные наборы
LETTERS="abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
NUMBERS="0123456789"
SYMBOLS_BASIC="@#-_+="
SYMBOLS_ADVANCED="!@#\$%^&*()[]{}<>?~"
UNICODE_PARANOID="🔐💡🧠🌐🎲🚀你好†±λ@#€"

# Выбор алфавита
case "$LEVEL" in
  low)
    CHARS="${LETTERS}${NUMBERS}"
    ;;
  medium)
    CHARS="${LETTERS}${NUMBERS}${SYMBOLS_BASIC}"
    ;;
  strong)
    CHARS="${LETTERS}${NUMBERS}${SYMBOLS_BASIC}${SYMBOLS_ADVANCED}"
    ;;
  paranoid)
    CHARS="${LETTERS}${NUMBERS}${SYMBOLS_BASIC}${SYMBOLS_ADVANCED}${UNICODE_PARANOID}"
    ;;
  *)
    echo "❌ Неизвестный уровень сложности: $LEVEL (доступны: low, medium, strong, paranoid)"
    exit 1
    ;;
esac

# Генерация пароля
PASSWORD=""
for i in $(seq 1 $LENGTH); do
  INDEX=$(( RANDOM % ${#CHARS} ))
  CHAR="${CHARS:INDEX:1}"
  PASSWORD+="$CHAR"
done

echo "$PASSWORD"
