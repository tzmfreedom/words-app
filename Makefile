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
