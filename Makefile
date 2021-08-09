#!sh

APP_IMAGE_NAME:=tonycb-simple-vwap

docker-build:
	- docker build -t ${APP_IMAGE_NAME} .

docker-run: docker-build
	- docker run -it ${APP_IMAGE_NAME} go run .

docker-tests: docker-build
	- docker run -it ${APP_IMAGE_NAME} go test -race -v ./...

tests:
	go test -race -v ./...

run:
	go run .
