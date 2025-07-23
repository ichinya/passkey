build:
	go build -o passkey ./go

docker:
	docker build -t passkey .

encrypt:
	docker run --rm -e PASSCRYPT_KEY=$(key) passkey e "$(text)"

decrypt:
	docker run --rm -e PASSCRYPT_KEY=$(key) passkey d "$(text)"
