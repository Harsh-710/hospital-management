build:
	@go build -o bin/hospital-management cmd/main.go

test:
	@go test -v ./...
	
run: build
	@./bin/hospital-management

migrate:
	@go run cmd/migrate/main.go

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down