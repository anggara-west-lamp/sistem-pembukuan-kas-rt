.PHONY: run dev up down seed

run:
	go run ./cmd/server

dev:
	APP_ENV=development go run ./cmd/server

up:
	docker compose up -d --build

down:
	docker compose down -v

swagger:
	go install github.com/swaggo/swag/cmd/swag@latest && swag init -g cmd/server/main.go -o internal/docs
