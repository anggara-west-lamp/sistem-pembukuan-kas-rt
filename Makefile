.PHONY: run dev up down seed

run:
	go run ./cmd/server

dev:
	APP_ENV=development go run ./cmd/server

up:
	docker compose up -d --build

down:
	docker compose down -v

