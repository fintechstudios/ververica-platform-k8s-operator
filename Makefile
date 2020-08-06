# Image URL to use all building/pushing image targets
VERSION?=v0.8.1
IMG?=fintechstudios/ververica-platform-k8s-operator
PKG=github.com/fintechstudios.com/ververica-platform-k8s-operator
VERSION_PKG=main
BUILD=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
# Produce CRDs that work back to Kubernetes 1.11 (no version conversion)
CRD_OPTIONS?="crd"

DOCKER_REPO=fintechstudios/ververica-platform-k8s-operator

LD_FLAGS="-X $(VERSION_PKG).operatorVersion='$(VERSION)' -X $(VERSION_PKG).gitCommit='$(GIT_COMMIT)' -X $(VERSION_PKG).buildDate='$(BUILD)'"

TEST_CLUSTER_NAME=ververica-platform-k8s-operator-cluster

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

all: manager

# find or download controller-gen
.PHONY: controller-gen
controller-gen:
ifeq (, $(shell which controller-gen))
	@{ \
	set -e ;\
	CONTROLLER_GEN_TMP_DIR=$$(mktemp -d) ;\
	cd $$CONTROLLER_GEN_TMP_DIR ;\
	go mod init tmp ;\
	go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.2.4 ;\
	rm -rf $$CONTROLLER_GEN_TMP_DIR ;\
	}
CONTROLLER_GEN=$(GOBIN)/controller-gen
else
CONTROLLER_GEN=$(shell which controller-gen)
endif

# find or download kustomize
.PHONY: kustomize
kustomize:
ifeq (, $(shell which kustomize))
	@{ \
	set -e ;\
	KUSTOMIZE_TMP_DIR=$$(mktemp -d) ;\
	cd $$KUSTOMIZE_TMP_DIR ;\
	go mod init tmp ;\
	go get sigs.k8s.io/kustomize/kustomize/v3@v3.3.0 ;\
	rm -rf $$KUSTOMIZE_TMP_DIR ;\
	}
KUSTOMIZE=$(GOBIN)/kustomize
else
KUSTOMIZE=$(shell which kustomize)
endif

# find or download kind
.PHONY: kind
kind:
ifeq (, $(shell which kind))
	@{ \
	set -e ;\
	KIND_TMP_DIR=$$(mktemp -d) ;\
	cd $$KIND_TMP_DIR ;\
	go mod init tmp ;\
	go get sigs.k8s.io/kind@v0.7.0 ;\
	rm -rf $$KIND_TMP_DIR ;\
	}
KIND=$(GOBIN)/kind
else
KIND=$(shell which kind)
endif

# find or download mocker
.PHONY: mockery
mockery:
ifeq (, $(shell which mockery))
	@{ \
	set -e ;\
	MOCKERY_TMP_DIR=$$(mktemp -d) ;\
	cd $$MOCKERY_TMP_DIR ;\
	go mod init tmp ;\
	go get github.com/vektra/mockery/.../ ;\
	rm -rf $$MOCKERY_TMP_DIR ;\
	}
MOCKERY=$(GOBIN)/mockery
else
MOCKERY=$(shell which mockery)
endif


# Run tests
.PHONY: test
test: generate manifests
	go test -ldflags $(LD_FLAGS) ./pkg/... ./api/... ./controllers/... ./ -coverprofile cover.out

# Build manager binary
.PHONY: manager
manager: generate
	go build $(ARGS) -ldflags $(LD_FLAGS) -o bin/manager main.go

mocks: mockery
	rm -rf mocks \
		&& $(MOCKERY) -dir pkg/vvp/appmanager -all -output ./mocks/vvp/appmanager -case=underscore \
		&& $(MOCKERY) -dir pkg/vvp/platform -all -output ./mocks/vvp/platform -case=underscore

# Run against the configured Kubernetes cluster in ~/.kube/config
.PHONY: run
run: generate
	go run -ldflags $(LD_FLAGS) ./main.go

# Install CRDs into a cluster
.PHONY: install
install: manifests
	$(KUSTOMIZE) build config/crd | kubectl apply -f -

# Uninstall CRDs from a cluster
uninstall: manifests
	$(KUSTOMIZE) build config/crd | kubectl delete -f -

# Deploy controller in the configured Kubernetes cluster in ~/.kube/config
.PHONY: deploy
deploy: manifests
	cd config/manager && $(KUSTOMIZE) edit set image controller=${IMG}
	$(KUSTOMIZE) build config/default | kubectl apply -f -

# Generate manifests e.g. CRD, RBAC etc.
.PHONY: manifests
manifests: controller-gen
	$(CONTROLLER_GEN) $(CRD_OPTIONS) rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases

# Run gofmt against non-generated code
.PHONY: fmt
fmt:
	gofmt -s -w ./api ./controllers ./pkg *.go

.PHONY: lint
lint:
	golangci-lint run --timeout=120s --verbose

# Generate code
.PHONY: generate
generate: controller-gen
	$(CONTROLLER_GEN) object:headerFile=./hack/boilerplate.go.txt paths=./api/...

# Patch the latest image version into the default kustomize image patch
.PHONY: patch-image
patch-image:
	sed -i'' -e 's@image: .*@image: '"$(IMG):$(VERSION)"'@' ./config/default/manager_image_patch.yaml

# Build the k8s resources for deployment
.PHONY: kustomize-build
kustomize-build: patch-image manifests
	rm -f resources.yaml && $(KUSTOMIZE) build config/default > resources.yaml

# Update the Swagger Client API
.PHONY: swagger-gen
swagger-gen:
	./hack/update-app-manager-swagger-codegen.sh \
	 && ./hack/update-platform-swagger-codegen.sh

.PHONY: docker-build-image
docker-build-image:
	docker build -f build.Dockerfile \
		--cache-from $(DOCKER_REPO)-builder:$(VERSION) \
		--tag $(DOCKER_REPO)-builder:$(VERSION) \
		.

.PHONY: docker-build
docker-build: docker-build-image
	docker build \
		--build-arg BUILD_IMG=$(DOCKER_REPO)-builder:$(VERSION) \
		--build-arg GIT_COMMIT=$(GIT_COMMIT) \
		--build-arg VERSION=$(VERSION) \
		--cache-from $(DOCKER_REPO):$(VERSION) \
		-f Dockerfile \
		--tag $(DOCKER_REPO):$(VERSION) \
		.

.PHONY: test-cluster-load-image
test-cluster-load-image: kind
	$(KIND) load docker-image --name $(TEST_CLUSTER_NAME) $(DOCKER_REPO):$(VERSION)

# Create the test cluster using kind
# install local path storage as defult storage class (see: https://github.com/kubernetes-sigs/kind/issues/118#issuecomment-475134086)
.PHONY: test-cluster-create kind
test-cluster-create:
	$(KIND) create cluster --name $(TEST_CLUSTER_NAME)

# Delete the test cluster using kind
.PHONY: test-cluster-delete kind
test-cluster-delete:
	$(KIND) delete cluster --name $(TEST_CLUSTER_NAME)

.PHONY: test-cluster-install-vvp-enterprise
test-cluster-install-vvp-enterprise:
	. ./hack/helm-init.sh \
	&& helm upgrade --install \
		--version 4.0.0 \
		--namespace vvp \
		--set vvp.persistence.type=local \
		vvp \
		ververica/ververica-platform \
		$(HELM_EXTRA_ARGS)

.PHONY: test-cluster-install-vvp-community
test-cluster-install-vvp-community:
	. ./hack/helm-init.sh \
	&& helm upgrade --install \
		--version 4.0.0 \
		--namespace vvp \
		--set vvp.persistence.type=local \
		--set acceptCommunityEditionLicense=true \
		vvp \
		ververica/ververica-platform \
		$(HELM_EXTRA_ARGS)

.PHONY: test-cluster-install-cert-manager
test-cluster-install-cert-manager:
	kubectl apply \
		--validate=false \
		-f https://github.com/jetstack/cert-manager/releases/download/v0.14.1/cert-manager.crds.yaml \
	&& . ./hack/helm-init.sh \
	&& helm upgrade --install \
		--version v0.14.1 \
		--namespace cert-manager \
		cert-manager \
		jetstack/cert-manager

.PHONY: test-cluster-install-chart
test-cluster-install-chart: docker-build test-cluster-load-image
	. ./hack/helm-init.sh \
	&& helm upgrade --install \
		--namespace vvp \
		vp-k8s-operator \
		./charts/vp-k8s-operator \
		--set imageRepository=$(DOCKER_REPO) \
		--set imageTag=$(VERSION) \
		--set vvpEdition=community \
		--set vvpUrl=http://vvp-ververica-platform \
		$(HELM_EXTRA_ARGS)

.PHONY: test-cluster-install-crds
test-cluster-install-crds:
	. ./hack/helm-init.sh \
	&& helm upgrade --install \
		--namespace vvp \
		vp-k8s-operator-crds \
		./charts/vp-k8s-operator-crds \
		--set webhookCert.name=vp-k8s-operator-serving-cert \
		--set webhookService.name=vp-k8s-operator-webhook-service \
		$(HELM_EXTRA_ARGS)

.PHONY: test-cluster-wait-for-cert-manager
test-cluster-wait-for-cert-manager:
	kubectl -n cert-manager wait --for=condition=available deployments --all

.PHONY: test-cluster-wait-for-vvp
test-cluster-wait-for-vvp:
	kubectl -n vvp wait --for=condition=available deployments --all

# Requires tiller to be running
.PHONY: test-cluster-setup
test-cluster-setup: test-cluster-install-cert-manager \
					test-cluster-install-vvp-community \
					test-cluster-wait-for-cert-manager \
					test-cluster-wait-for-vvp \
					test-cluster-install-chart \
					test-cluster-install-crds
