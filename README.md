# Ververica Platform Controller

Makes Ververica Platform resources Kubernetes-Native! Defines CustomResourceDefinitions
for mapping resources to K8s!

This is built as a CRD / custom controller. Though an Aggregated API might be a better choice,
the tooling (currently) does not have as good support and I just learned how to `go`.

Built for Ververica Platform version `1.4.x`.

[More about the Ververica Platform](https://www.ververica.com/platform-overview)  
[Ververica Platform Docs](https://docs.ververica.com/)

## Supported Resources

Since the resources names of K8s and the Ververica Platform somewhat class, the 
custom VP Resources will all be prefixed with `Vp`.

* `DeploymentTarget` -> `VpDeploymentTarget`
* `Deployment` -> `VpDeployment`
* `Namespace` -> `VpNamespace`

## Unsupported

* `Job`
* `Event`
* `Role Binding`
* `Role`
* `Cluster Role Binding`
* `Cluster Role`
* `Savepoint`
* `Secret Value`
* `Status`


To avoid naming conflicts, and for simplicity, and VP `metadata` and `spec` fields
are nested under the top-level `spec` field of the K8s resource.

Look in [docs/mappings](./docs/mappings) for information on each supported resource.

## Development

Built using [`kubebuilder`](https://github.com/kubernetes-sigs/kubebuilder).

More on the design of the controller and its resources can be found
in [docs/design.md](./docs/design.md).


Also built as a Go 1.11 module - no vendor files here.

System Pre-requisites:
- `go` >= 1.12
- `make` >= 4
- `kubebuilder` >= 2.0.0-beta.0
- `kustomize` >= v3.0.1
- `docker`
- [`kind`](https://github.com/kubernetes-sigs/kind) (or similar, like `minikube`)  
  `kind` is seamless with docker and the tools of this repo, which is why it is built into the `Makefile` scripts. 

### `make` Scripts

- `make`
- `make run` runs the entire app
- `make manager` builds the entire app binary
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
- `make test-cluster-create` initializes a cluster fro testing, using kind


### Ververica Platform API

The API Client is auto-generated using the [Swagger Codegen utility](https://github.com/swagger-api/swagger-codegen.git).

#### Pre-Generation Changes

The original Swagger file was taken from their live API documentation (available at `${VP_URL}/api/swagger`),
but the docs don't exactly match their API, which makes the generated client incorrect.

Main changes necessary:
* Timestamps are returned as ISO8601 strings, not numbers
* `DeploymentTarget.deploymentPatchSet` is a JSON Patch Array, not a JsonNode
* `Artifact` needs many other fields
* `DeploymentUpgradeStrategy` needs choices
* `DeploymentRestoreStrategy` needs choices and `allowNonRestoredState` option
* `DeploymentStartFromSavepoint` needs choices
* `POST /namespaces/{namespace}/deployments` needs a `201` response with a `Deployment` in the body
* `POST /namespaces/{namespace}/deployment-targets` needs a `201` response with a `DeploymenTarget` in the body

#### Post-Generation Changes

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
- `api_api_tokens.go`
- `api_events.go`
- `api_jobs.go`
- `api_namespaces.go`
- `api_savepoints.go`


There is also a bug that cannot handle an empty Swagger type to represent the `any` type, so
you must manually change [`model_any.go`](./ververica-platform-api/model_any.go) to:

```go
package ververicaplatformapi

type Any interface {}
```

You'll also have to change any usages of this type in `structs` to be embedded, instead of by pointer ref, namely in:
- `model_json_patch_generic.go`

Is this all better than creating an API Client from scratch? Yes. Can I still gripe about it? TBD. 


## Future Work

This is a MVP for Flink deployments at FinTech Studios. We would love to see this
improved! 

Some known issues + places to improve:
* Mapping of more VP resources!
* `DeploymentTarget.deploymentPatchSet` values can only be `strings`.
* The nesting of `metadata` and `spec` is a little wonky.
* Watching the VP API for updates to Deployments, Jobs, etc would be excellent.
* Adding more `status` subresources to link everything together would also be most excellent.
* It might make sense to have a 1-1 mapping between K8s namespaces and names and VP namespaces and names, but 
will there ever be more than one VP running in a cluster?
* Improvements on the Swagger API generator / moving that to OpenAPI V3.
* Memory management / over-allocation / embed-by-value vs embed-by-pointer could probably be improved.
* Various `TODO`s should give us a place to start!
* Better package structure for internal code.

## Acknowledgements

Other OSS that influenced this project:
* [Kong Ingress Controller](https://github.com/Kong/kubernetes-ingress-controller)
