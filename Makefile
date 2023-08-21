.PHONY: run product


run:
	go run main.go
product:
	GIN_MODE=release go run main.go