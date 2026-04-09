.PHONY: up down test

up:
	docker compose -f ./docker/docker-compose.yml --env-file .env up --build -d

down:
	docker compose -f ./docker/docker-compose.yml down -v

test:
	go test ./...