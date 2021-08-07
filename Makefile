#!sh

APP_IMAGE_NAME:=tonycb-simple-vwap

docker-run:
	- docker build -t ${APP_IMAGE_NAME} .
	- docker run -it ${APP_IMAGE_NAME} go run .
#	docker run -it ${DOCKER_IMAGE} go run .