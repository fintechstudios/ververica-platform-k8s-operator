# Ververica Platform Controller

Makes Ververica Platform resources Kubernetes-Native! 

## Supported Resources

Since the namespaces of K8s and the Ververica Platform somewhat class, the 
custom VP Resources will all be prefixed with `VP`.

* `DeploymentTargets` -> `VPDeploymentTargets`
* `Deployments` -> `VPDeployments`
* `Namespaces` -> `VPNamespaces`

## Unsupported

* `Jobs`
* `Events`
* `Role Bindings`
* `Roles`
* `Cluster Role Bindings`
* `Cluster Roles`
* `Savepoints`
* `Secret Values`
* `Status`


## Development

Built using [`kubebuilder`](https://github.com/kubernetes-sigs/kubebuilder),
which requires `kustomize`.

Also built as a Go 1.11 module - no vendor files here.