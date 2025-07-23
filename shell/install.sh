#!/bin/bash

set -e

echo "⬇️ Установка passkey..."

TARGET="/usr/local/bin/passkey"

curl -sSL https://raw.githubusercontent.com/ichinya/passkey/main/shell/passkey.sh -o "$TARGET"
chmod +x "$TARGET"

echo "✅ Установлено в $TARGET"
echo "Пример использования:"
echo 'PASSCRYPT_KEY="mykey" passkey e "password"'
