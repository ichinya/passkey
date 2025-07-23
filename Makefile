build:
	go build -o passkey go/main.go encrypt.go

test:
	go test