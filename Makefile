VERSION := $(shell git rev-parse HEAD | cut -c1-8)
PROJECT_NAME := places
BACKEND_IMAGE := ${PROJECT_NAME}:${VERSION}

build:
	GOOS=linux GOARCH=amd64 go build -o program main.go

docker: build
	docker build -t ${PROJECT_NAME}:${VERSION} .

run:  docker
	API_KEY=${API_KEY} BACKEND_IMAGE=${BACKEND_IMAGE} docker-compose up
test:
	go test ./... -v -short -p 1 -cover
lint:
	golangci-lint run
