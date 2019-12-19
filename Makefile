# Image URL to use all building/pushing image targets
VERSION?=v0.4.0
IMG?=fintechstudios/ververica-platform-k8s-operator
PKG=github.com/fintechstudios.com/ververica-platform-k8s-operator
VERSION_PKG=main
BUILD=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
# Produce CRDs that work back to Kubernetes 1.11 (no version conversion)
CRD_OPTIONS?="crd:trivialVersions=true"

LD_FLAGS="-X $(VERSION_PKG).controllerVersion='$(VERSION)' -X $(VERSION_PKG).gitCommit='$(GIT_COMMIT)' -X $(VERSION_PKG).buildDate='$(BUILD)'"

TEST_CLUSTER_NAME=ververica-platform-k8s-operator-cluster

all: manager

# find or download controller-gen
.PHONY: controller-gen
controller-gen:
ifeq (, $(shell which controller-gen))
	go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.2.4
CONTROLLER_GEN=$(shell go env GOPATH)/bin/controller-gen
else
CONTROLLER_GEN=$(shell which controller-gen)
endif

# find or download kustomize
.PHONY: kustomize
kustomize:
ifeq (, $(shell which kustomize))
	go get sigs.k8s.io/kustomize/kustomize/v3@v3.3.0
KUSTOMIZE=$(shell go env GOPATH)/bin/kustomize
else
KUSTOMIZE=$(shell which kustomize)
endif

# Run tests
.PHONY: test
test: generate manifests
	go test -ldflags $(LD_FLAGS) ./api/... ./controllers/... -coverprofile cover.out

# Build manager binary
.PHONY: manager
manager: generate
	go build $(ARGS) -ldflags $(LD_FLAGS) -o bin/manager main.go

# Run against the configured Kubernetes cluster in ~/.kube/config
.PHONY: run
run: generate
	go run -ldflags $(LD_FLAGS) ./main.go

# Install CRDs into a cluster
.PHONY: install
install: manifests
	kubectl apply -f config/crd/bases

# Deploy controller in the configured Kubernetes cluster in ~/.kube/config
.PHONY: deploy
deploy: install
	kustomize build config/default | kubectl apply -f -

# Generate manifests e.g. CRD, RBAC etc.
.PHONY: manifests
manifests: controller-gen
	$(CONTROLLER_GEN) $(CRD_OPTIONS) rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases

# Run go fmt against code
.PHONY: fmt
fmt:
	go fmt ./...

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
kustomize-build: patch-image
	$(KUSTOMIZE) build config/default > resources.yaml

# Update the Swagger Client API
.PHONY: swagger-gen
swagger-gen:
	./hack/update-app-manager-swagger-codegen.sh && \
	./hack/update-platform-swagger-codegen.sh

# Create the test cluster using kind
.PHONY: test-cluster-create
test-cluster-create:
	kind create cluster --name $(TEST_CLUSTER_NAME) && $(MAKE) install

# Delete the test cluster using kind
.PHONY: test-cluster-delete
test-cluster-delete:
	kind delete cluster --name $(TEST_CLUSTER_NAME)

