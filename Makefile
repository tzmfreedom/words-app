PORT := 8080

.PHONY: run
run: format
	PORT=$(PORT) go run .

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
db/migrate: sql-migrate
	sql-migrate up

.PHONY: sql-migrate
sql-migrate:
	go get -v github.com/rubenv/sql-migrate/...

.PHONY: install
install: sql-migrate
