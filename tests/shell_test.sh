#!/usr/bin/env bash
set -e
cd $(dirname "$0")/..

export PASSCRYPT_KEY="testkey"

out=$(bash shell/passkey.sh e "secret")
dec=$(bash shell/passkey.sh d "$out")
[[ "$dec" == "secret" ]]

pw=$(bash shell/passkey.sh g -l 12)
test ${#pw} -eq 12

# cross compatibility with Go implementation
go_enc=$(cd go && go run . e -mode shell "secret")
dec_shell=$(bash shell/passkey.sh d "$go_enc")
[[ "$dec_shell" == "secret" ]]

sh_enc=$(bash shell/passkey.sh e "secret")
dec_go=$(cd go && go run . d -mode shell "$sh_enc")
[[ "$dec_go" == "secret" ]]

echo "shell tests passed"

