PORT := 8080
DATABASE_URL := "postgres://postgres:@localhost:5432/myapp?sslmode=disable"

.PHONY: run
run: format
	DATABASE_URL=$(DATABASE_URL) PORT=$(PORT) go run .

.PHONY: format
format:
	gofmt -w .

.PHONY: import
import:
	goimports -w .

.PHONY: dep
dep:
	dep ensure

.PHONY: deploy
deploy:
	git push heroku master

.PHONY: db/migrate
db/migrate:
	migrate -database $(DATABASE_URL) -path db/migrations up

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
