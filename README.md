# Ververica Platform K8s Operator

[![go reportcard](https://goreportcard.com/badge/github.com/fintechstudios/ververica-platform-k8s-operator)](https://goreportcard.com/report/github.com/fintechstudios/ververica-platform-k8s-operator)
[![FOSSA Status](https://app.fossa.io/api/projects/custom%2B12442%2Fgit%40github.com%3Afintechstudios%2Fververica-platform-k8s-operator.git.svg?type=shield)](https://app.fossa.io/projects/custom%2B12442%2Fgit%40github.com%3Afintechstudios%2Fververica-platform-k8s-operator.git?ref=badge_shield)[![pipeline status](https://gitlab.com/fintechstudios/ververica-platform-k8s-operator/badges/master/pipeline.svg)](https://gitlab.com/fintechstudios/ververica-platform-k8s-operator/commits/master)
[![coverage report](https://gitlab.com/fintechstudios/ververica-platform-k8s-operator/badges/master/coverage.svg)](https://gitlab.com/fintechstudios/ververica-platform-k8s-operator/commits/master)

Makes Ververica Platform resources Kubernetes-Native! Defines CustomResourceDefinitions
for mapping resources to K8s!

Built for Ververica Platform version `2.x`.

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
* `DeploymentDefaults`
* `Secret Value`
* `Status`

To avoid naming conflicts, and for simplicity, and VP `metadata` and `spec` fields
are nested under the top-level `spec` field of the K8s resource.

Look in [docs/mappings](docs/mappings) for information on each supported resource.

## Getting Started

Please have a look at the [`docs`](docs/README.md) for information on getting started using
the operator.

### Editions

This operator works with both the Community and Enterprise editions of the Ververica Platform, with the caveats:
* `VpNamespaces` are not supported by the Community Edition, so the manager will not register those resources
* The `spec.metadata.namespace` field must either be left unset or set explicitly to `default` for all `Vp` resources

Find out more about [the editions here](https://www.ververica.com/pricing-editions).

## Running

To run the binary directly, after building run `./bin/manager`.

**Flags:**
* `--help` prints usage
* `--vvp-url=http://localhost:8081` the url, without trailing slash, for the Ververica Platform
* `--vvp-edition=enterprise` the Ververica Platform Edition to support. See [Editions](#Editions) for more.
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

A Helm chart for the operator lives in [`./charts/vp-k8s-operator`](charts/vp-k8s-operator),
which sets up a deployment with a metrics server, RBAC policies, CRDs, and, optionally,
an RBAC proxy for the metrics over HTTPS.

The CRDs are managed in a separate chart ([`./charts/vp-k8s-operator-crds`](charts/vp-k8s-operator-crds)), which also
needs to be installed.

## Development

Built using [`kubebuilder`](https://github.com/kubernetes-sigs/kubebuilder).
[`kind`](https://github.com/kubernetes-sigs/kind) is used for running a local test cluster,
though something like `minikube` will also do.  

More on the design of the controller and its resources can be found
in [docs/design.md](docs/design.md).

Also built as a Go module - no vendor files here.

System Pre-requisites:
- `go` >= `1.14.x`
- `make` >= `4`
- `kubebuilder` == [`v2.2.0`](https://github.com/kubernetes-sigs/kubebuilder/releases/tag/v2.2.0)
- `docker` >= `19`
- `kind` >= `0.6.0`

### `make` Scripts

- `make` alias for `manager`
- `make manager` builds the entire app binary
- `make run` runs the entire app locally
- `make manifests` builds the CRDs from `./config/crd`
- `make install` installs the CRDs from `./config/crd` on the cluster
- `make deploy` installs the entire app on the cluster
- `make docker-build` builds the docker image
- `make docker-push` pushes the built docker image
- `make generate` generates the controller code from the `./api` package
- `make swagger-gen` generates the swagger code
- `make lint` runs linting on the source code
- `make fmt` runs `go fmt` on the package
- `make test` runs the test suites with coverage
- `make patch-image` sets the current version as the default deployment image tag
- `make kustomize-build` builds the default k8s resources for deployment

#### For working with a local kind cluster

- `make test-cluster-create` initializes a cluster for testing, using kind
- `make test-cluster-delete` deletes the testing cluster
- `make test-cluster-setup` installs cert-manager, the Community VVP, the vp-k8s-crds, and the vp-k8s-operator on the test cluster
- `make test-cluster-instal-chart` builds the operator and installs it on the test cluster from the local chart
- `make test-cluster-instal-crds` installs the vp-k8s-operator CRDs on the test cluster from the local chart

### Environment

To use the default test cluster, you'll need to store a `KUBECONFIG` env var pointed to it.

Setting `DISABLE_WEBHOOKS` to any value does what you think it will.

[`godotenv`](https://github.com/joho/godotenv) automatically loads this when running `main`.

### AppManager + Platform APIs

The API Clients are auto-generated using the [Swagger Codegen utility](https://github.com/swagger-api/swagger-codegen.git).

#### AppManager

The [`appmanager-api` Swagger file](appmanager-api-swagger.json) is from the live API documentation (available at `${VP_URL}/api/swagger`),
but the generated client needs a few updates to work correctly.

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
* `model_env_var.go` needs to be updated with the proper Kubernetes types


There is also a bug that cannot handle an empty Swagger type to represent the `any` type, so
you must manually change [`model_any.go`](pkg/vvp/appmanager-api/model_any.go) to:

```go
package appmanagerapi

type Any interface {}
```

You'll also have to change any usages of this type in `structs` to be embedded, instead of by pointer ref, namely in:
- `model_json_patch_generic.go`


### Building Images

The images are built in two steps:
1. The [`build.Dockerfile`](build.Dockerfile) image is a full development environment for running tests, linting,
and building the source with the correct tooling. This can also be used for development if you so like,
just override the entrypoint.
2. The build image is then passed as a build arg to the main [`Dockerfile`](Dockerfile), which builds
the manager binary and copies it over into an image for distribution.


## Acknowledgements

Other OSS that influenced this project:
* [Kong Ingress Controller](https://github.com/Kong/kubernetes-ingress-controller)

## License

[Licensed under Apache 2.0](LICENSE)

[![FOSSA Status](https://app.fossa.io/api/projects/custom%2B12442%2Fgit%40github.com%3Afintechstudios%2Fververica-platform-k8s-operator.git.svg?type=large)](https://app.fossa.io/projects/custom%2B12442%2Fgit%40github.com%3Afintechstudios%2Fververica-platform-k8s-operator.git?ref=badge_large)
