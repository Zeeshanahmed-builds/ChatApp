DB_URL=postgres://postgres:123456@localhost:5433/chatapp?sslmode=disable
MIGRATION_DIR=./db/migrations

up:
	migrate -path $(MIGRATION_DIR) -database "$(DB_URL)" up

down:
	migrate -path $(MIGRATION_DIR) -database "$(DB_URL)" down

migrate-force:
	migrate -path $(MIGRATION_DIR) -database "$(DB_URL)" force 1

run:
	golangci-lint run ./...
	go run main.go