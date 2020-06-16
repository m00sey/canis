.PHONY := clean test tools agency

CANIS_ROOT=$(abspath .)

all: clean tools wire test build

commit: cover build

# Cleanup files (used in Jenkinsfile)
clean:
	rm -f bin/*

wire: wire-steward

wire-steward:
	@. ./canis.sh; cd pkg/steward && wire

tools:
	go get bou.ke/staticfiles
	go get github.com/google/wire/cmd/wire
	go get github.com/vektra/mockery/.../
	go get golang.org/x/tools/cmd/cover

build: bin/steward bin/agency bin/router
build-steward: bin/steward

steward: bin/steward
bin/steward: wire-steward
	cd cmd/steward && go build -o $(CANIS_ROOT)/bin/steward

.PHONY: steward-docker agency-docker router-docker
package: steward-docker agency-docker router-docker

steward-docker: bin/steward
	@echo "Building steward agent docker image"
	@docker build -f ./docker/steward/Dockerfile -t scoir/steward:latest .

build-agency: bin/agency
build-router: bin/router

wire-agency:
	cd pkg/agency && wire

agency: bin/agency bin/router
bin/agency: wire-agency
	cd cmd/agency && go build -o $(CANIS_ROOT)/bin/agency

wire-router:
	cd pkg/router && wire

router: bin/router
bin/router: wire-router
	cd cmd/router && go build -o $(CANIS_ROOT)/bin/router

agency-docker: bin/agency
	@echo "Building agency docker image"
	@docker build -f ./docker/agency/Dockerfile --no-cache -t scoir/agency:latest .

router-docker: bin/router
	@echo "Building router docker image"
	@docker build -f ./docker/router/Dockerfile --no-cache -t scoir/router:latest .

canis-docker: build
	@echo "Building canis docker image"
	@docker build -f ./docker/canis/Dockerfile --no-cache -t scoir/canis:latest .

demo-web:
	cd demo && npm run build

# Development Local Run Shortcuts
urn: run
run: bin/steward
	@bin/scoir-agent

test:
	go test ./pkg/...

cover:
	go test -coverprofile cover.out ./pkg/...
	go tool cover -html=cover.out

dev-setup:
	@./scripts/dev-setup.sh

initialize:
	@minikube delete
	@minikube start --vm-driver virtualbox --insecure-registry registry.hyades.svc.cluster.local:5000
	@./scripts/minikube-setup.sh

install:
	@helm install canis ./canis-chart --values ./k8s/config/local/values.yaml --kubeconfig ./k8s/config/local/kubeconfig.yaml

uninstall:
	@helm uninstall canis && ([ $$? -eq 0 ] && echo "") || echo "nothing to uninstall!"

expose:
	minikube service -n hyades steward-loadbalancer --url

von-ip:
	@docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' von_webserver_1

cycle: clean build test
	@./scripts/cycle.sh

ledger:
	@helm upgrade --install ledger ./ledger-chart --values ./k8s/config/local/values.yaml --kubeconfig ./k8s/config/local/kubeconfig.yaml

prune:
	@echo
	@echo "These might be overly aggressive but they work and I just reclaimed 7gb of docker images sooooooooooo"
	@echo
	@echo "Deletes dangling data"
	@echo -e "\t \U0001F92F  #docker system prune"
	@echo
	@echo "Deletes any image not referenced by any container"
	@echo -e "\t \U0001F92F  #docker image prune -a"