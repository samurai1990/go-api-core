.PHONY: build run product createsuperuser


all:build

build:
	go build -o web_api
run:
	go run main.go
product:
	GIN_MODE=release go run main.go
createsuperuser:
	./web_api createsuperuser
