.PHONY := clean test tools agency

CANIS_ROOT=$(abspath .)

all: clean tools wire steward agent

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
	go get -u github.com/golang/protobuf/protoc-gen-go

swagger_pack: pkg/static/steward_agent_swagger.go
pkg/static/steward_agent_swagger.go: steward-pb pkg/steward/api/spec/steward_agent.swagger.json
	staticfiles -o pkg/static/steward_agent_swagger.go --package static pkg/steward/api/spec

build: bin/steward bin/agency bin/router
build-steward: bin/steward

steward: bin/steward
bin/steward: steward-pb swagger_pack wire-steward
	cd cmd/steward && go build -o $(CANIS_ROOT)/bin/steward

.PHONY: steward-docker agency-docker router-docker
package: steward-docker agency-docker router-docker

steward-docker: bin/steward
	@echo "Building steward agent docker image"
	@docker build -f ./docker/steward/Dockerfile -t scoir/steward:latest .

build-agent: bin/agent
build-agency: bin/agency
build-router: bin/router

wire-agent:
	cd pkg/agent && wire

agent: bin/agent
bin/agent: wire-agent
	cd cmd/agent && go build -o $(CANIS_ROOT)/bin/agent

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

steward-pb: pkg/steward/api/steward_agent.pb.go
pkg/steward/api/steward_agent.pb.go:pkg/steward/api/steward_agent.proto
	cd pkg && protoc -I/home/pfeairheller/opt/protoc-3.6.1/include -I . -I steward/api/ steward/api/steward_agent.proto --go_out=plugins=grpc:.
	cd pkg && protoc -I/home/pfeairheller/opt/protoc-3.6.1/include -I . -I steward/api/ steward/api/steward_agent.proto --grpc-gateway_out=logtostderr=true:.
	cd pkg && protoc -I/home/pfeairheller/opt/protoc-3.6.1/include -I . -I steward/api/ steward/api/steward_agent.proto --swagger_out=logtostderr=true:.
	mv pkg/steward/api/steward_agent.swagger.json pkg/steward/api/spec
demo-web:
	cd demo && npm run build

# Development Local Run Shortcuts
urn: run
run: bin/steward
	@bin/scoir-agent

test: clean tools
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