all: setup run

setup:
	docker compose -f ./deployments/docker-compose.yml up -d

run:
	go run ./cmd/security-service/main.go