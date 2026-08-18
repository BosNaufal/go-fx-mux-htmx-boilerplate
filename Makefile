
run:
	go run cmd/app/main.go

migrate-up:
	go run cmd/migration/main.go up


migrate-down:
	go run cmd/migration/main.go down