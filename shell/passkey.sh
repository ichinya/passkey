#!/bin/bash
# Shell version of passkey

# Зашифровать и получить строку
passcrypt encrypt --password "my-secret" --key "superkey"

# Расшифровать строку
passcrypt decrypt --cipher "U2FsdGVkX1..." --key "superkey"

# Сохранить в хранилище
passcrypt save --name "gmail" --password "my-password" --key "masterkey"

# Получить из хранилища
passcrypt get --name "gmail" --key "masterkey"
