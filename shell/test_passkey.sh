#!/bin/bash
set -e

export PASSCRYPT_KEY="test123"
TEST_PASS="supersecret"

# Зашифруем
ENC=$(bash shell/passkey.sh e "$TEST_PASS")

# Расшифруем
DEC=$(bash shell/passkey.sh d "$ENC")

if [[ "$DEC" == "$TEST_PASS" ]]; then
  echo "✅ Тест пройден: $DEC"
else
  echo "❌ Тест провален"
  exit 1
fi
