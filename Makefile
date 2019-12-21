PORT ?= 8080
DATABASE_URL ?= "postgres://postgres:@localhost:5432/myapp?sslmode=disable"
TEST_DATABASE_URL ?= "postgres://postgres:@localhost:5432/test?sslmode=disable"
AUTH_USER ?= user
AUTH_PASS ?= pass

.PHONY: run
run: format
	USER=$(AUTH_USER) PASS=$(AUTH_PASS) DATABASE_URL=$(DATABASE_URL) PORT=$(PORT) go run .

.PHONY: format
format:
	gofmt -w .

.PHONY: import
import:
	goimports -w .

.PHONY: deploy
deploy:
	git push heroku master

.PHONY: db/migrate
db/migrate:
	migrate -database $(DATABASE_URL) -path db/migrations up

.PHONY: db/migrate
db/migrate/test:
	migrate -database $(TEST_DATABASE_URL) -path db/migrations up

.PHONY: migrate
migrate:
	go get -u github.com/golang-migrate/migrate

.PHONY: install
install: sql-migrate

.PHONY: prod/db/migrate
prod/db/migrate:
	heroku run make db/migrate

.PHONY: up
up:
	docker-compose up -d

.PHONY: db/init
db/init:
	psql -U postgres -h localhost -p 5432 -c "CREATE DATABASE postgres;"

.PHONY: db/migrate/new
db/migrate/new:
	migrate create -ext sql -dir db/migrations -seq $(NAME)

.PHONY: db/console
db/console:
	psql -U postgres -h localhost -p 5432 myapp

.PHONY: prod/db/console
prod/db/console:
	heroku pg:psql

.PHONY: test
test:
	USER=$(AUTH_USER) PASS=$(AUTH_PASS) DATABASE_URL=$(TEST_DATABASE_URL) go test .
