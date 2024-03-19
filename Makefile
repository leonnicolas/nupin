.PHONY: generate build build-all generate-all fe/build clean vet

ALL_ARCH := amd64 arm arm64

build:
	CGO_ENABLED=0 go build .

build-all:
	for arch in $(ALL_ARCH); do \
		GOOS=linux GOARCH=$$arch CGO_ENABLED=0 go build -o bin/linux/$$arch/nupin .; \
	done

generate: api/nuki/swagger.json
	go generate ./...

generate-all: generate fe/src/client/index.ts README.md

api/nuki/swagger.json: api/nuki/patch.json
	curl -o - https://api.nuki.io/static/swagger/swagger.json | sed 's/"int"/"integer"/' | go run github.com/evanphx/json-patch/cmd/json-patch -p ./api/nuki/patch.json | jq > $@

fe/src/client/index.ts:
	docker run --rm -v "$$(pwd)":/src --user $$(id -u):$$(id -g) openapitools/openapi-generator-cli:v7.3.0 generate -i /src/api/v0/v0.yaml -g typescript-fetch -o /src/fe/src/client

fe/build:
	cd fe && yarn install && yarn build

tmp/help.txt: build
	mkdir -p tmp
	./nupin --help 2> $@

README.md: tmp/help.txt
	go run github.com/campoy/embedmd -w $@
clean:
	rm -rf fe/dist
	rm -rf bin
