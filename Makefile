.PHONY: build

cli.build:
	go build -o eim cmd/eim-cli/*.go

api.build:
	go build -o eim-web web/*.go

api.up:
	go run api/*.go

