.PHONY: generate
generate:
	go generate ./...
	openapi-generator generate -g typescript-fetch -i api/swagger.json -o web/src/lib/models/api

.PHONY: run
run:
	go run cmd/main.go