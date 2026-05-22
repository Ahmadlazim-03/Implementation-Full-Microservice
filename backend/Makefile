.PHONY: tidy run-auth run-places run-review run-gateway up down

tidy:
	go mod tidy

run-auth:
	go run ./cmd/auth-service

run-places:
	go run ./cmd/places-service

run-review:
	go run ./cmd/review-service

run-gateway:
	go run ./cmd/api-gateway

up:
	docker compose up -d --build

down:
	docker compose down
