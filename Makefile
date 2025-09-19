run: 
	@go run cmd/app/*.go

seed: scrape 
	@go run cmd/seed/*.go

scrape:  
	@python3 web-scraper/scraper.py 

migration: 
	@migrate create -ext sql -dir cmd/migrate/migrations

migrate-up: 
	@go run cmd/migrate/main.go up

migrate-down: 
	@go run cmd/migrate/main.go down