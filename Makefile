build:
	cd go && go build -o ../passkey

docker:
	docker build -t passkey .

encrypt:
	@if [ ! -f /run/secrets/PASSCRYPT_KEY ] && [ -z "$(PASSCRYPT_KEY)" ]; then \
		echo "Error: Secret PASSCRYPT_KEY not found and environment variable PASSCRYPT_KEY is not set."; \
		exit 1; \
	fi
	@if [ -f /run/secrets/PASSCRYPT_KEY ]; then \
		docker run --rm \
		--env PASSCRYPT_KEY=$(PASSCRYPT_KEY) \
		--mount type=bind,source=/run/secrets/PASSCRYPT_KEY,target=/run/secrets/PASSCRYPT_KEY \
		passkey e "$(ARGS)"; \
	else \
		docker run --rm \
		--env PASSCRYPT_KEY=$(PASSCRYPT_KEY) \
		passkey e "$(ARGS)"; \
	fi

decrypt:
	@if [ ! -f /run/secrets/PASSCRYPT_KEY ] && [ -z "$(PASSCRYPT_KEY)" ]; then \
		echo "Error: Secret PASSCRYPT_KEY not found and environment variable PASSCRYPT_KEY is not set."; \
		exit 1; \
	fi
	@if [ -f /run/secrets/PASSCRYPT_KEY ]; then \
		docker run --rm \
		--env PASSCRYPT_KEY=$(PASSCRYPT_KEY) \
		--mount type=bind,source=/run/secrets/PASSCRYPT_KEY,target=/run/secrets/PASSCRYPT_KEY \
		passkey d "$(ARGS)"; \
	else \
		docker run --rm \
		--env PASSCRYPT_KEY=$(PASSCRYPT_KEY) \
		passkey d "$(ARGS)"; \
	fi
