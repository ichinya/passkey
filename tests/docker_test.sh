#!/usr/bin/env bash
set -e
cd $(dirname "$0")/..

docker build -t passkey-test . >/dev/null
export PASSCRYPT_KEY="testkey"

enc=$(docker run --rm -e PASSCRYPT_KEY=testkey passkey-test e secret)
dec=$(docker run --rm -e PASSCRYPT_KEY=testkey passkey-test d "$enc")
[[ "$dec" == "secret" ]]

echo "docker tests passed"

