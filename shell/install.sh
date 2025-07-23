#!/bin/bash

set -e

TARGET="/usr/local/bin/passkey"
URL="https://raw.githubusercontent.com/ichinya/passkey/main/shell/passkey.sh"

echo "🔐 Установка passkey..."

# Загрузка скрипта
curl -fsSL "$URL" -o "$TARGET"

# Делаем исполняемым
chmod +x "$TARGET"

echo "✅ Установка завершена!"
echo "Теперь можно использовать команду:"
echo '  PASSCRYPT_KEY="mykey" passkey e "my-password"'