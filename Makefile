.PHONY: clean build dev deploy

clean:
	go clean
	rm -rf ./bin

build: clean
	go mod tidy
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/hello-world functions/hello-world/main.go

dev:
	sls offline --useDocker start --host 0.0.0.0

deploy: clean build
	sls deploy --verbose
