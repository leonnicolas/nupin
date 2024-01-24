.PHONY: generate build

build:
	CGO_ENABLED=0 go build .

generate: api/nuki/swagger.json
	go generate ./...

generate-all: generate fe/src/client/index.ts

api/nuki/swagger.json:
	curl -o - https://api.nuki.io/static/swagger/swagger.json | sed 's/"int"/"integer"/' > $@

fe/src/client/index.ts:
	docker run --rm -v "$$(pwd)":/src --user $$(id -u):$$(id -g) openapitools/openapi-generator-cli:v7.3.0 generate -i /src/api/v0/v0.yaml -g typescript-fetch -o /src/fe/src/client

fe/build:
	cd fe && yarn build
