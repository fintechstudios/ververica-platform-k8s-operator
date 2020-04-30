# Install cert-manager

Install the CustomResourceDefinition resources separately.

```bash
# Kubernetes 1.15+
$ kubectl apply --validate=false -f https://github.com/jetstack/cert-manager/releases/download/v0.14.3/cert-manager.crds.yaml

# Kubernetes <1.15
$ kubectl apply --validate=false -f https://github.com/jetstack/cert-manager/releases/download/v0.14.3/cert-manager-legacy.crds.yaml
```

Create the namespace for cert-manager.

```bash
kubectl create namespace cert-manager
```

Add the Jetstack Helm repository.

```bash
helm repo add jetstack https://charts.jetstack.io
```

Update your local Helm chart repository cache.

```bash
helm repo update
```

Install the cert-manager Helm chart.

```bash
# For Helm v3+
$ helm install \
  cert-manager jetstack/cert-manager \
  --namespace cert-manager \
  --version v0.14.3

# For Helm v2
$ helm install \
  --name cert-manager \
  --namespace cert-manager \
  --version v0.14.3 \
  jetstack/cert-manager
```

Wait for the deployment to come up

```bash
$ kubectl --namespace cert-manager wait --for=condition=available deployments --all
deployment.extensions/cert-manager condition met
deployment.extensions/cert-manager-cainjector condition met
deployment.extensions/cert-manager-webhook condition met
```
