APP_NAME := api-server
APP_PATH := apps/$(APP_NAME)

.PHONY: build run test docker-build docker-run docker-clean air

## Go build
build:
	cd $(APP_PATH) && go build -o bin/$(APP_NAME) ./cmd/server

## Run the built binary
run:
	cd $(APP_PATH) && ./bin/$(APP_NAME)

## Run tests
test:
	cd $(APP_PATH) && go test -v ./...

## Build Docker image
docker-build:
	docker build -t $(APP_NAME) $(APP_PATH)

## Run Docker container
docker-run:
	docker run --rm -p 8080:8080 $(APP_NAME)

## Remove built binary
clean:
	rm -rf $(APP_PATH)/bin

## Remove Docker image
docker-clean:
	docker rmi $(APP_NAME) || true

## Run development server
air:
	cd $(APP_PATH) && air