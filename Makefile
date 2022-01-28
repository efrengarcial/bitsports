# Migrate scheme in ent to database
migrate_schema:
	APP_ENV=mig go run ./cmd/migration/main.go

docker:
	docker build -t product-service .

run:
	docker-compose  up --build -d

stop:
	docker-compose down

test:
	go test -coverpkg=./... -coverprofile=coverage.out ./...
	go tool cover -html coverage.out -o coverage.html

unittest:
	go test -short  ./...


.PHONY: test run
