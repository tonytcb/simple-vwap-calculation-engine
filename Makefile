#!sh

APP_IMAGE_NAME:=tonycb-simple-vwap

docker-build:
	- docker build -t ${APP_IMAGE_NAME} .

docker-run: docker-build
	- docker run -it ${APP_IMAGE_NAME} go run . ${MAX_TRADINGS}

docker-tests: docker-build
	- docker run -it ${APP_IMAGE_NAME} go test -race -v ./...

tests:
	go test -race -v ./...

run:
	go run .

lint: ## Lint with golangci-lint
	@docker run --rm -it -v $(PWD):/app -w /app golangci/golangci-lint:v1.39-alpine \
	golangci-lint run \
	--timeout 5m0s \
	--exclude-use-default=false \
	--enable=prealloc