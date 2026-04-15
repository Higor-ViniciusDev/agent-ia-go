.PHONY: up down test

up:
	docker compose -f ./docker/docker-compose.yml --env-file .env up --build -d

down:
	docker compose -f ./docker/docker-compose.yml down -v

test:
	go test ./...

lint:
	docker run --rm -v $(CURDIR):/app -w /app golangci/golangci-lint:v2.11.4 golangci-lint run

build-grpc:
	docker run --rm -v $(CURDIR):/app -w /app protoc-gen protoc -I . -I /googleapis --go_out=. --go-grpc_out=. --grpc-gateway_out=. internal/infra/grpc/proto/protofiles/hello.proto

migrate-up:
	docker run --rm -v $(CURDIR):/app -w /app migrate/migrate -path ./sql/migrations -database "postgres://postgres:postgres@host.docker.internal:5432/work_agent?sslmode=disable" up