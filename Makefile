.PHONY: clean build dev deploy doc

clean:
	go clean
	rm -rf ./bin

build: clean doc
	go mod tidy
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/create-user cmd/create-user/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/get-users cmd/get-users/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/serve-docs cmd/serve-docs/main.go

dev:  clean build doc
	sls offline --useDocker start --host 0.0.0.0

deploy-dev: clean build
	sls deploy --verbose --stage dev

deploy-prod: clean build
	sls deploy --verbose --stage production

doc:
	npx @redocly/cli build-docs --output ./docs/index.html docs/openapi/openapi.json