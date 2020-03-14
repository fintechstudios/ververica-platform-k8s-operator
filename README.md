# Ververica Platform K8s Operator

[![go reportcard](https://goreportcard.com/badge/github.com/fintechstudios/ververica-platform-k8s-operator)](https://goreportcard.com/report/github.com/fintechstudios/ververica-platform-k8s-operator)
[![FOSSA Status](https://app.fossa.io/api/projects/custom%2B12442%2Fgit%40github.com%3Afintechstudios%2Fververica-platform-k8s-operator.git.svg?type=shield)](https://app.fossa.io/projects/custom%2B12442%2Fgit%40github.com%3Afintechstudios%2Fververica-platform-k8s-operator.git?ref=badge_shield)[![pipeline status](https://gitlab.com/fintechstudios/ververica-platform-k8s-operator/badges/master/pipeline.svg)](https://gitlab.com/fintechstudios/ververica-platform-k8s-operator/commits/master)
[![coverage report](https://gitlab.com/fintechstudios/ververica-platform-k8s-operator/badges/master/coverage.svg)](https://gitlab.com/fintechstudios/ververica-platform-k8s-operator/commits/master)

Makes Ververica Platform resources Kubernetes-Native! Defines CustomResourceDefinitions
for mapping resources to K8s!

Built for Ververica Platform version `2.x` with go version `1.13`.

[More about the Ververica Platform](https://www.ververica.com/platform-overview)  
[Ververica Platform Docs](https://docs.ververica.com/)

## Supported Resources

Since the resources names of K8s and the Ververica Platform somewhat clash, the 
custom VP Resources will all be prefixed with `Vp`.

* `DeploymentTarget` -> `VpDeploymentTarget`
* `Deployment` -> `VpDeployment`
* `Namespace` -> `VpNamespace`
* `Savepoint` -> `VpSavepoint`
* `Event` -> native K8s `Event`

## Unsupported

* `Job`
* `Secret Value`
* `Status`

To avoid naming conflicts, and for simplicity, and VP `metadata` and `spec` fields
are nested under the top-level `spec` field of the K8s resource.

Look in [docs/mappings](./docs/mappings) for information on each supported resource.

## Running

To run the binary directly, after building run `./bin/manager`.

**Flags:**
* `--help` prints usage
* `--app-manager-api-url=http://localhost:8081/api` the url, without trailing slash, for the Ververica Platform's AppManager API
* `--platform-api-url=http://localhost:8081` the url, without trailing slash, for the Ververica Platform's AppManager API
* `--debug` debug mode for logging
* `--enable-leader-election` to ensure only one manager is active with a multi-replica deployment
* `--metrics-addr=:8080` address to bind metrics to 
* `--watch-namespace=all-namespaces` the namespace to watch resources on
* `[--env-file]` the path to an environment (`.env`) file to be loaded

For authorization with the AppManager's API, a token is needed. This can be provided in the environment on either a
per-namespace or one-token-to-rule-them-all basis. If it is not provided in the environment, an "owner" token will be created
for each namespace that resources are managed in.

Specifying in the environment is a good way to integrate with namespaces that aren't defined in Kubernetes.

**Environment:**
* `APPMANAGER_API_TOKEN_{NAMESPACE}` a token to use for resources in a specific Ververica Platform namespace, upper-cased
* `APPMANAGER_API_TOKEN` if no namespace-specific token can be found, this value will be used. 


## Docker

Images are published to [Docker Hub](https://hub.docker.com/r/fintechstudios/ververica-platform-k8s-operator).
*  The `latest` tag always refers to the current HEAD in the master branch.
* Each master commit hash is also tagged and published.
* Git tags are published with the same tag. 

## Helm

A base Helm chart is provided in [`./charts/vp-k8s-operator`](charts/vp-k8s-operator).

This sets up a deployment with a metrics server, RBAC policies, CRDs, and, optionally, an RBAC proxy for the metrics over HTTPS.

## Development

Built using [`kubebuilder`](https://github.com/kubernetes-sigs/kubebuilder).
[`kind`](https://github.com/kubernetes-sigs/kind) is used for running a local test cluster,
though something like `minikube` will also do.  

More on the design of the controller and its resources can be found
in [docs/design.md](./docs/design.md).

Also built as a Go module - no vendor files here.

System Pre-requisites:
- `go` >= `1.13.x`
- `make` >= `4`
- `kubebuilder` == [`v2.2.0`](https://github.com/kubernetes-sigs/kubebuilder/releases/tag/v2.2.0)
- `docker` >= `19`
- `kind` >= `0.6.0`

### `make` Scripts

- `make` alias for `manager`
- `make manager` builds the entire app binary
- `make run` runs the entire app
- `make manifests` builds the CRDs
- `make install` installs the CRDs on the cluster
- `make deploy` installs the entire app on the cluster
- `make docker-build` builds the docker image
- `make docker-push` pushes the built docker image
- `make generate` generates the controller code from the `./api` package
- `make controller-gen` loads the correct controller-gen binary
- `make swagger-gen` generates the swagger code
- `make lint` runs the golangci linter 
- `make fmt` runs `go fmt` on the package
- `make test` runs the test suites with coverage
- `make test-cluster-create` initializes a cluster for testing, using kind
- `make test-cluster-delete` deletes the testing cluster
- `make patch-image` sets the current version as the default deployment image tag
- `make kustomize-build` builds the default k8s resources for deployment

### Environment

To use the default test cluster, you'll need to store a `KUBECONFIG` env var pointed to it.

[`godotenv`](https://github.com/joho/godotenv) automatically loads this when running `main`.

### AppManager + Platform APIs

The API Clients are auto-generated using the [Swagger Codegen utility](https://github.com/swagger-api/swagger-codegen.git).

#### AppManager

The original Swagger file was taken from their live API documentation (available at `${VP_URL}/api/swagger`),
but the docs don't exactly match their API, which makes the generated client incorrect.

##### Post-Generation Changes

The `optional` package is missing from many of the imports in the generated code, as must be added manually.

```go
package ververicaplatformapi

import (
	// ...
    "github.com/antihax/optional"
    // ...
)
```

Affected files:
- `api_event_resource.go`
- `api_job_resource.go`
- `api_savepoint_resource.go`


Type Changes:
* `model_pods.go` needs to be updated with the proper Kubernetes types
* `model_volume_and_mount.go` needs to be updated with the proper Kubernetes types


There is also a bug that cannot handle an empty Swagger type to represent the `any` type, so
you must manually change [`model_any.go`](pkg/vvp/appmanager-api/model_any.go) to:

```go
package ververicaplatformapi

type Any interface {}
```

You'll also have to change any usages of this type in `structs` to be embedded, instead of by pointer ref, namely in:
- `model_json_patch_generic.go`


### Building Images

The images are built in two steps:
1. The [`Dockerfile_build`](./Dockerfile_build) image is a full development environment for running tests, linting,
and building the source with the correct tooling. This can also be used for development if you so like,
just override the entrypoint.
2. The build image is then passed as a build arg to the main [`Dockerfile`](./Dockerfile), which builds
the manager binary and copies it over into an image for distribution.


## Future Work

This is a MVP for Flink deployments at FinTech Studios. We would love to see this
improved! 

Some known issues + places to improve:
* Make the Platform APIClient an `interface` so that something like [`mockery`](https://github.com/vektra/mockery) can mock it for tests.
* Splitting tests into unit, integration, e2e tests against a live cluster, etc.

## Acknowledgements

Other OSS that influenced this project:
* [Kong Ingress Controller](https://github.com/Kong/kubernetes-ingress-controller)


## License
[![FOSSA Status](https://app.fossa.io/api/projects/custom%2B12442%2Fgit%40github.com%3Afintechstudios%2Fververica-platform-k8s-operator.git.svg?type=large)](https://app.fossa.io/projects/custom%2B12442%2Fgit%40github.com%3Afintechstudios%2Fververica-platform-k8s-operator.git?ref=badge_large)
