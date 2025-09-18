run: 
	@go run cmd/app/*.go

seed: 
	@go run cmd/seed/*.go

migration: 
	@migrate create -ext sql -dir cmd/migrate/migrations

migrate-up: 
	@go run cmd/migrate/main.go up

migrate-down: 
	@go run cmd/migrate/main.go down