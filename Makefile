.PHONY: clean build dev deploy doc

clean:
	go clean
	rm -rf ./bin

build: clean
	go mod tidy
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/create-user cmd/create-user/main.go

dev:  clean build
	sls offline --useDocker start --host 0.0.0.0

deploy: clean build
	sls deploy --verbose

doc:
	npx @redocly/cli build-docs --output ./docs/index.html docs/openapi/openapi.json