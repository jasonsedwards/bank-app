BINARY_NAME=get-balance

build:
	GOARCH=amd64 GOOS=darwin go build -o bin/${BINARY_NAME}-darwin ./src/main
	GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux ./src/main
	GOARCH=amd64 GOOS=windows go build -o bin/${BINARY_NAME}-windows ./src/main

run: build
	./bin/${BINARY_NAME}

clean:
	go clean
	rm ./bin/${BINARY_NAME}-darwin
	rm ./bin/${BINARY_NAME}-linux
	rm ./bin/${BINARY_NAME}-windows

dep:
	go mod download

lint:
	golangci-lint run --enable-all

docker:
	docker build --tag bank-app:local .

docker-run:
	docker run --env-file=env_vars --name bank-app bank-app:local

docker-stop:
	docker stop bank-app:local

docker-rm:
	docker rm bank-app

docker-all: docker docker-run

