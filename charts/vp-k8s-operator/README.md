# VP K8s Operator

A Helm chart for deploying the Ververica Platform Kubernetes Operator.

## Installing the Chart

| Parameter                   | Description                                                           | Default                                                                         |
| --------------------------- | --------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `rbac.enabled`              | Whether or not to create RBAC resources.                              | `true`                                                                          |
| `rbacProxy.enabled`         | Whether or not to create an RBAC proxy over `https`.                  | `true`                                                                          |
| `rbacProxy.logLevel`        | Log level for the proxy.                                              | `10`                                                                            |
| `rbacProxy.imageRepository` |                                                                       | `gcr.io/kubebuilder/kube-rbac-proxy`                                            |
| `rbacProxy.imageTag`        |                                                                       | `v0.4.0`                                                                        |
| `rbacProxy.imagePullPolicy` |                                                                       | `IfNotPresent`                                                                  |
| `rbacProxy.port`            |                                                                       | `8443`                                                                          |
| `imageRepository`           | Image repository for the Manager                                      | `fintechstudios/ververica-platform-k8s-operator`                                |
| `imageTag`                  |                                                                       | `latest`                                                                        |
| `imagePullPolicy`           |                                                                       | `IfNotPresent`                                                                  |
| `metricsHost`               | Host for the metrics reporter.                                        | `127.0.0.1`                                                                     |
| `metricsPort`               | Port for the metrics reporter.                                        | `8080`                                                                          |
| `metricsMonitorEnabled`     | Whether or not to create a Prometheus ServiceMonitor.                 | `false`                                                                         |
| `certs.enabled`             | Whether or not to create CertManager certs for webhook serving.       | `true`                                                                          |
| `certs.existingSecret`      | If not creating certs, must specify a secret with pre-existing certs. | `nil`                                                                           |
| `vvpUrl`                    | URL for the Ververica Platform.                                       | `http://ververica-platform`                                                     |
| `vvpEdition`                | Ververica Platform Edition. Either `community` or `enterprise`.       | `enterprise`                                                                    |
| `extraArgs`                 | Extra CLI args to pass to the controller manager.                     | `[]`                                                                            |
| `resources`                 | Resource specs for the manager deployment.                            | `{ limits: { cpu: 100m, memory: 30Mi }, requests: { cpu: 100m, memory 20Mi } }` |
| `livenessProbe`             | Liveness Probe configuration for manager container.                   | `{}`                                                                            |
| `podAnnotations`            | Annotations to be added to the pods (inside deployment).              | `{}`                                                                            |
