.PHONY: build run release createsuperuser test build_image run-debug build_ngix builddev down attach clean clean-all run-db


all:build

build:
	@go build -o web_api

run:
	@./web_api

release:
	@go build -ldflags "-s -w" -o web_api

createsuperuser:
	@./web_api createsuperuser

test:
	@go clean -testcache
	@go test --cover ./...

build_image:
	-docker builder prune
	@docker build -t api:1.0.0 docker-runner

run-debug:
	docker compose -f serivce.yml up

build_ngix:
	-docker builder prune
	-docker container prune -f
	-docker image rm api_nginx:1.0.0
	@docker build -t api_nginx:1.0.0 ./nginx

builddev: build_ngix
	-docker builder prune
	-docker container prune -f      
	-docker image rm api:1.0.0
	@docker build -f Dockerfile.develop -t api:1.0.0 .

down:
	@docker compose -f service.yml down

attach: ## Attach for debugging
	@docker exec -it $$(docker ps -a --filter name=api_web | awk '{ print $$1}'| tail -n+2) bash

clean:
	@rm -f web_api

clean-all: clean
	-docker container prune -f
	-docker image rm -f api:1.0.0
	-docker image rm -f api_nginx:1.0.0
	-docker system prune -f

run-db:
	@docker compose -f service.yml up -d db
	@echo IPAddress DB: $$(docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $$(docker-compose -f service.yml ps -q))