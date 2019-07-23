# Ververica Platform Controller

Makes Ververica Platform resources Kubernetes-Native! Defines CustomResourceDefinitions
for the 

## Supported Resources

Since the resources names of K8s and the Ververica Platform somewhat class, the 
custom VP Resources will all be prefixed with `VP`.

* `DeploymentTarget` -> `VPDeploymentTarget`
* `Deployment` -> `VPDeployment`
* `Namespace` -> `VPNamespace`

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


### `VPNamespace`

Schema:
```yaml
apiVersion: ververicaplatform.fintechstudios.com/v1beta1
kind: VPNamespace
metadata:
  name: String # Required
  id: UUID String # Dynamic
  createdAt: Timestamp # Dynamic
  modifiedAt: Timestamp  # Dynamic
  resourceVersion: Integer # Dynamic
status:
  state: String # Dynamic
```

## Development

Built using [`kubebuilder`](https://github.com/kubernetes-sigs/kubebuilder),
which requires `kustomize`.

Also built as a Go 1.11 module - no vendor files here.

System Pre-requisites:
- `go` >= 1.12
- `kubebuilder` >= 2.0.0-beta.0
- `docker`

### Ververica Platform API

The API Client is auto-generated using the 

The original Swagger file was taken from their live API documentation (available at `${VP_URL}/api/swagger`),
but they are still using Swagger 2.0 and the docs don't exactly match their API, which
makes the generated client incorrect.

Main changes necessary:
* Timestamps are returned as ISO8601 strings, not numbers
* `DeploymentTarget.deploymentPatchSet` is a JSON Patch Array, not a JsonNode

Another annoying thing about the swagger-codegen is that the `optional` package is missing
from many of the imports in the generated code, as must be added manually.

Is this all better than creating an API Client from scratch? Yes. Can I still complain about it? TBD. 

