# Deployment

The only currently supported method of deploying the operator
is through Helm. This guide also assumes that you have [Cert-Manager](https://cert-manager.io/)
running in the cluster to provision certificates for the CRD webhooks.

## Helm

### Installing the cert-manager
According to [cert-manager](https://cert-manager.io/docs/installation/helm/#steps) site:

```bash
helm repo add jetstack https://charts.jetstack.io
helm repo update
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.8.0/cert-manager.crds.yaml

helm install \
  cert-manager jetstack/cert-manager \
  --namespace cert-manager \
  --create-namespace \
  --version v1.8.0
```

### Installing the Ververica Platform

For the most up-to-date guide, refer to [the official getting started guide](https://www.ververica.com/getting-started).

#### Add the Ververica chart repository
```bash
helm repo add ververica https://charts.ververica.com
```

#### Deploy the Ververica Platform
For deploying jobs into namespaces outside of the VVP deployment namespace, specify the `rbac.additionalNamespaces` value with a set of namespaces.

In this case, add the `top-speed` namespace associated with the sample manifests in `/config/samples`. Create the `top-speed` namespace if it doesn't exist:
```bash
kubectl create namespace vvp || true
kubectl create namespace top-speed || true
```

Deploy the platform:
```bash
# For the Enterprise Edition
helm install vvp ververica/ververica-platform \
            -n vvp \
            -f /path/to/values-with-license.yaml  \ # must specify a values file with the enterprise license
            --set rbac.additionalNamespaces={top-speed}

# Or, for the Community Edition
helm install vvp ververica/ververica-platform \
            -n vvp \
            --set acceptCommunityEditionLicense=true \
            --set rbac.additionalNamespaces={top-speed}
```

Wait for the deployment to come up:
```bash
kubectl --namespace vvp wait --for=condition=available deployments --all
```

And make Ververica available:
```bash
kubectl port-forward --namespace vvp service/vvp-ververica-platform 8080:80
```

### Installing the Operator

Until there is a Helm chart repository set up for this operator, copies of the charts
in the base [`charts`](../../charts) directory are needed to install with helm.
First, the operator should be deployed, followed by the CRDs. They should both be deployed
in the same namespace as the Ververica Platform.

This guide assumes you are operating in the base `ververica-platform-k8s-operator` directory.

#### Install the Operator

First, check that everything is okay with the help of the `--dry-run`, then run the command without this flag:
```bash
# NOTE: the pods might crash on startup and enter a restart loop until the CRDs
#       are present, but this should be fine. 

# Enterprise
helm install --namespace vvp \
    vp-k8s-operator \
    ./charts/vp-k8s-operator \
    --set vvpEdition=enterprise \
    --set vvpUrl=http://vvp-ververica-platform

# Community
helm install -n vvp \
    vp-k8s-operator \
    ./charts/vp-k8s-operator \
    --set vvpEdition=community \
    --set vvpUrl=http://vvp-ververica-platform \
    --dry-run
```

#### Install the CRDs
```bash
# Using the Cert created by the operator chart for serving webhooks.
# Pointed at the webhook conversion service of the operator.
helm install --namespace vvp \
    vp-k8s-operator-crds \
    ./charts/vp-k8s-operator-crds \
    --set webhookCert.name=vp-k8s-operator-serving-cert \
    --set webhookService.name=vp-k8s-operator-webhook-service
```

And, finally, wait for the deployment to come up:
```bash
kubectl --namespace vvp wait --for=condition=available deployments --all
```

### Deploying the samples

The samples can all be deployed through `kubectl`.

```bash
# Create the VpDeploymentTarget
kubectl apply -f config/samples/ververicaplatform_v1beta2_vpdeploymenttarget.yaml

# Create the VpDeployment
kubectl apply -f config/samples/ververicaplatform_v1beta2_vpdeployment.yaml
```

Watch the Top Speed V2 deployment come online. Either visit http://localhost:8080/app/#/namespaces/default/deployments to see the UI, or view through the K8s Events:
```bash
# View through K8s Events
kubectl -n top-speed get events --sort-by='lastTimestamp'
# Or live
kubectl -n top-speed get events --watch
```

```bash
# Ex: cancell the deployment with a JSON Patch
kubectl patch -n top-speed vpdeployment/top-speed-v2 --type merge --patch '{ "spec": { "spec": { "state": "CANCELLED" } } }'
```
