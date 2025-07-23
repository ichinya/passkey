#!/usr/bin/env bash

set -e

MODE="$1"
shift || true

KEY_REQUIRED() {
  if [[ -z "$PASSCRYPT_KEY" ]]; then
    echo "PASSCRYPT_KEY not set" >&2
    exit 1
  fi
}

case "$MODE" in
  e)
    KEY_REQUIRED
    if [[ -z "$1" ]]; then
      echo "usage: passkey.sh e <string>" >&2
      exit 1
    fi
    echo -n "$1" | openssl enc -aes-256-cbc -a -salt -pbkdf2 -pass pass:"$PASSCRYPT_KEY"
    ;;
  d)
    KEY_REQUIRED
    if [[ -z "$1" ]]; then
      echo "usage: passkey.sh d <cipher>" >&2
      exit 1
    fi
    echo "$1" | openssl enc -aes-256-cbc -a -d -salt -pbkdf2 -pass pass:"$PASSCRYPT_KEY"
    ;;
  g)
    LENGTH=16
    LEVEL="medium"
    BATCH=1
    ENCRYPT=0
    while [[ $# -gt 0 ]]; do
      case "$1" in
        -l|--length) LENGTH="$2"; shift 2;;
        -L|--level) LEVEL="$2"; shift 2;;
        -b|--batch) BATCH="$2"; shift 2;;
        -e|--encrypt) ENCRYPT=1; shift;;
        *) shift;;
      esac
    done
    LETTERS="abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    NUMBERS="0123456789"
    BASIC="@#-_+="
    ADV="!@#$%^&*()[]{}<>?~"
    EMOJI="🔐💡🧠🌐🎲🚀你好†±λ@#€"
    case "$LEVEL" in
      low) CHARS="$LETTERS$NUMBERS";;
      medium) CHARS="$LETTERS$NUMBERS$BASIC";;
      strong) CHARS="$LETTERS$NUMBERS$BASIC$ADV";;
      paranoid) CHARS="$LETTERS$NUMBERS$BASIC$ADV$EMOJI";;
      *) echo "unknown level" >&2; exit 1;;
    esac
    for ((i=0;i<BATCH;i++)); do
      PASS=$(tr -dc "$CHARS" < /dev/urandom | head -c "$LENGTH")
      if [[ $ENCRYPT -eq 1 ]]; then
        KEY_REQUIRED
        PASS=$(echo -n "$PASS" | openssl enc -aes-256-cbc -a -salt -pbkdf2 -pass pass:"$PASSCRYPT_KEY")
      fi
      echo "$PASS"
    done
    ;;
  *)
    echo "usage: passkey.sh [e|d|g]" >&2
    exit 1
    ;;
esac
