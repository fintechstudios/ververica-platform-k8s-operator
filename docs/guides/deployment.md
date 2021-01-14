# Deployment

The only currently supported method of deploying the operator
is through Helm. This guide also assumes that you have [Cert-Manager](https://cert-manager.io/)
running in the cluster to provision certificates for the CRD webhooks.

## Helm

### Installing the Ververica Platform

For the most up-to-date guide, refer to [the official getting started guide](https://www.ververica.com/getting-started).

```shell
# 0. Add the Ververica chart repository
helm repo add ververica https://charts.ververica.com

# 1. Deploy the Ververica Platform

# For deploying jobs into namespaces outside of the VVP deployment namespace,
# specify the `rbac.additionalNamespaces` value with a set of namespaces.
# In this case, add the `top-speed` namespace associated with the sample manifests in `/config/samples`.

## For the Enterprise Edition
helm install --namespace vvp \
            --name vvp \
            ververica/ververica-platform \
            -f /path/to/values-with-license.yaml  \ # must specify a values file with the enterprise license
            --set rbac.additionalNamespaces={top-speed}

## Or, for the Community Edition
For helm 2:

```bash
helm install --namespace vvp \
            --name vvp \
            ververica/ververica-platform \
            --set acceptCommunityEditionLicense=true \
            --set rbac.additionalNamespaces={top-speed}
```

For helm 3:

```bash
helm install vvp ververica/ververica-platform -n vvp --set acceptCommunityEditionLicense=true
```

## 2.Wait for the deployment to come up

```bash
kubectl --namespace vvp wait --for=condition=available deployments --all
```

## 3. Access the platform UI locally at port 8080

```bash
kubectl port-forward --namespace vvp service/vvp-ververica-platform 8080:80
```

NOTE: If you don't have cert-manager you may get following error:

```bash
Error: Internal error occurred: failed calling webhook "webhook.cert-manager.io": Post https://cert-manager-webhook.cert-manager.svc:443/mutate?timeout=30s: no endpoints available for service "cert-manager-webhook"
```

To deploy cert-manager[`check out this guide`](./cert-manager.md)

### Installing the Operator

Until there is a Helm chart repository set up for this operator, copies of the charts
in the base [`charts`](../../charts) directory are needed to install with helm.
First, the operator should be deployed, followed by the CRDs. They should both be deployed
in the same namespace as the Ververica Platform.

This guide assumes you are operating in the base `ververica-platform-k8s-operator` directory.

```shell
# 4. Install the Operator
# Pointed at the service deployed with the Ververica Platform.
# NOTE: the pods might crash on startup and enter a restart loop until the CRDs
#       are present, but this should be fine.

## Enterprise
helm install --namespace vvp \
    --name vp-k8s-operator \
    ./charts/vp-k8s-operator \
    --set vvpEdition=enterprise \
    --set vvpUrl=http://vvp-ververica-platform

## Community

### For helm 2:
helm install --namespace vvp \
    --name vp-k8s-operator \
    ./charts/vp-k8s-operator \
    --set vvpEdition=community \
    --set vvpUrl=http://vvp-ververica-platform

### For helm 3:
# Try dryrun to check the rendered templates
$ helm install -n vvp vp-k8s-operator ./charts/vp-k8s-operator \
    --set vvpEdition=community \
    --set vvpUrl=http://vvp-ververica-platform --dry-run

# Deploy the chart
$ helm install -n vvp vp-k8s-operator ./charts/vp-k8s-operator \
    --set vvpEdition=community \
    --set vvpUrl=http://vvp-ververica-platform
```


## 5.Install the CRDs

### Using the Cert created by the operator chart for serving webhooks

Pointed at the webhook conversion service of the operator.

With helm2:

```bash
helm install --namespace vvp \
    --name vp-k8s-operator-crds \
    ./charts/vp-k8s-operator-crds \
    --set webhookCert.name=vp-k8s-operator-serving-cert \
    --set webhookService.name=vp-k8s-operator-webhook-service
```

```bash
# Try dryrun to check the rendered templates
$ helm install -n vvp vp-k8s-operator-crds ./charts/vp-k8s-operator-crds \
    --set webhookCert.name=vp-k8s-operator-serving-cert \
    --set webhookService.name=vp-k8s-operator-webhook-service --dry-run

# Deploy
$ helm install -n vvp vp-k8s-operator-crds ./charts/vp-k8s-operator-crds \
    --set webhookCert.name=vp-k8s-operator-serving-cert \
    --set webhookService.name=vp-k8s-operator-webhook-service
```

## 6. Wait for the deployment to come up

```bash
kubectl --namespace vvp wait --for=condition=available deployments --all
```

### Deploying the samples

The samples can all be deployed through `kubectl`.

```shell
# 7. Install the samples in the top-speed namespace.

# Create namespace if it doesn't exist
kubectl create namespace top-speed || true

# Create the VpDeploymentTarget
kubectl apply -f config/samples/ververicaplatform_v1beta2_vpdeploymenttarget.yaml

# Create the VpDeployment
kubectl apply -f config/samples/ververicaplatform_v1beta2_vpdeployment.yaml

# 8. Watch the Top Speed V2 deployment come online

# Visit http://localhost:8080/app/#/namespaces/default/deployments to see the UI

# View through K8s Events
kubectl -n top-speed get events --sort-by='lastTimestamp'
# Or live
kubectl -n top-speed get events --watch

# Once started, editing the resource vpdeployment/top-speed-v2 in the top-speed namespace
# will trigger upgrades

# Ex: cancell the deployment with a JSON Patch
kubectl patch -n top-speed vpdeployment/top-speed-v2 --type merge --patch '{ "spec": { "spec": { "state": "CANCELLED" } } }'
```
