#!/usr/bin/env bash
set -e
cd $(dirname "$0")/..

export PASSCRYPT_KEY="testkey"

out=$(bash shell/passkey.sh e "secret")
dec=$(bash shell/passkey.sh d "$out")
[[ "$dec" == "secret" ]]

pw=$(bash shell/passkey.sh g -l 12)
test ${#pw} -eq 12

echo "shell tests passed"

