# Ververica Platform Controller

Makes Ververica Platform resources Kubernetes-Native! Defines CustomResourceDefinitions
for mapping resources to K8s!

This is built as a CRD / custom controller. Though an Aggregated API might be a better choice,
the tooling (currently) does not have as good support and I just learned how to `go`.

Built for Ververica Platform `1.4.1`, API Version 1.

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
- `make`
- `kubebuilder` >= 2.0.0-beta.0
- `kustomize` >= v3.0.1
- `docker`
- [`minikube`](https://github.com/kubernetes/minikube) or similar


### `make` Scripts

- `make`
- `make install`
- `make docker-build`

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

Effected files:
- `api_api_tokens.go`
- `api_events.go`
- `api_jobs.go`
- `api_namespaces.go`
- `api_savepoints.go`



There is also a bug that cannot handle an empty Swagger type to represent any type, so
you must manually change [`model_any.go`](./ververica-platform-api/model_any.go) to:

```go
package ververicaplatformapi

type Any interface {}
```

You'll also have to change any usages of this type in `structs` to be embedded, instead of by pointer ref.

Is this all better than creating an API Client from scratch? Yes. Can I still complain about it? TBD. 


## Future Work

This is a MVP for Flink deployments at FinTech Studios. We would love to see this
improved (ha)! 

Some known issues + places to improve:
* Mapping of more VP resources!
* `DeploymentTarget.deploymentPatchSet` values can only be `strings`.
* The nesting of `metadata` and `spec` is a little wonky.
* Polling the VP API for updates to Deployments, Jobs, etc would be excellent.
* Adding more `status` subresources to link everything together would also be most excellent.
* It might make sense to have a 1-1 mapping between K8s namespaces and names and VP namespaces and names, but 
will there ever be more than on VP running in a cluster?
* Improvements on the Swagger API generator / moving that to OpenAPI V3.
* Memory management / over-allocation / embed-by-value vs embed-by-pointer could probably be improved.
* Various `TODO`s should give us a place to start!

## Acknowledgements

Other OSS that influenced this project:
* [Kong Ingress Controller](https://github.com/Kong/kubernetes-ingress-controller)