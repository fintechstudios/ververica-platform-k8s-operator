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