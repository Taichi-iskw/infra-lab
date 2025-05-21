APP_NAME := api-server
APP_PATH := apps/$(APP_NAME)
CLUSTER_NAME := infra-lab
NAMESPACE := monitoring

.PHONY: build run test docker-build docker-run clean docker-clean air \
        kind-create kind-delete kind-reload docker-kind-load \
        deploy-api deploy-prometheus deploy-grafana setup-local reset-local

## Build Go binary
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

## Create kind cluster
kind-create:
	kind create cluster --name $(CLUSTER_NAME) --config infra/k8s/kind-cluster.yaml

## Delete kind cluster
kind-delete:
	kind delete cluster --name $(CLUSTER_NAME)

## Load Docker image into kind
docker-kind-load:
	kind load docker-image infra-lab-api-server:latest --name $(CLUSTER_NAME)

## Reload image to kind and rollout restart
kind-reload: docker-build docker-kind-load
	kubectl rollout restart deployment/$(APP_NAME)

## Deploy API server
deploy-api:
	kubectl apply -f infra/k8s/base/$(APP_NAME)

## Deploy Prometheus with Helm
deploy-prometheus:
	helm upgrade --install prometheus prometheus-community/prometheus \
	  -f infra/k8s/monitoring/prometheus/prometheus-values.yaml \
	  -n $(NAMESPACE) --create-namespace

## Deploy Grafana with Helm
deploy-grafana:
	helm upgrade --install grafana grafana/grafana \
	  -f infra/k8s/monitoring/grafana/grafana-values.yaml \
	  -n $(NAMESPACE)

## Setup full local Kubernetes environment (kind + api-server + monitoring stack)
setup-local: kind-create docker-build docker-kind-load deploy-api deploy-prometheus deploy-grafana

## Tear down and fully recreate the local environment
reset-local: kind-delete setup-local
