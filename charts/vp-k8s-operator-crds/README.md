# Vp K8s Operator CRDs

This chart manages the Custom Resource Definitions used by the Vp K8s Operator. Instead of
including the CRDs in the main `vp-k8s-operator` chart, they exist here to allow templating and
integrating with cert-manager for webhook certificates.

Find out more about [managing CRDs with Helm here](https://helm.sh/docs/chart_best_practices/custom_resource_definitions/). 

## Installing the Chart

All values default to the resources created in the default installation of the Operator chart with
the name `vp-k8s-operator`.

| Parameter                    | Description                                           | Default                                            |
|------------------------------|-------------------------------------------------------|----------------------------------------------------|
| `webhookCert.namespace`      | Namespace of the secret containing the TLS cert for the webhook.              | `{{ .Release.Namespace }}`                                             |
| `webhookCert.name`           | Name of the secret containing the TLS cert for the webhook.              | `vp-k8s-operator-serving-cert`                                             |
| `webhookCert.caBundle`       | PEM-encoded CA bundle for the webhook, if not using cert-manager.              | `Cg==`                                             |
| `webhookService.namespace`   | Namespace of the webhook service.              | `{{ .Release.Namespace }}`                                             |
| `webhookService.name`        | Name of the webhook service.             | `vp-k8s-operator-webhook-service`                                             |
