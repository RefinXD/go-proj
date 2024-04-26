up:
	migrate -path database/migrations -database 'postgres://postgres:password@localhost:5432/postgres?sslmode=disable' -verbose up

down:
	migrate -path database/migrations -database 'postgres://postgres:password@localhost:5432/postgres?sslmode=disable' -verbose down

run:
	go run main.go
test:
	go test ./...
sqlc:
	sqlc generate -f database/sqlc.yaml