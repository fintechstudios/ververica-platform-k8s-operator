# Deployment

The only currently supported method of deploying the operator
is through Helm. This guide also assumes that you have Cert-Manager
running in the cluster to provision certificates for the CRD webhooks.

## Helm

### Installing the Ververica Platform

For the most up-to-date guide, refer to [the official getting started guide](https://www.ververica.com/getting-started).

```shell
helm repo add ververica https://charts.ververica.com

# Only for deploying jobs into namespaces outside of the VVP deployment namespace
# In this case, the namespace associated with the sample configs
VVP_ARGS="--set vvp.appmanager.rbac.additionalNamespaces[0]=top-speed"

# For the Enterprise Edition
helm install --namespace vvp \
            vvp ververica/ververica-platform -f values-with-license.yaml \
            ${VVP_ARGS}

# For the Community Edition
helm install --namespace vvp \
            vvp ververica/ververica-platform --set acceptCommunityEditionLicense=true \
             ${VVP_ARGS}

# Wait for the deployment to come up
kubectl --namespace vvp wait --for=condition=available deployments --all

# Access the platform UI locally at port 8080
kubectl port-forward --namespaces service/vvp-ververica-platform 8080:80
```

### Installing the Operator

Until there is a Helm chart repository set up for this operator, copies of the charts
in the base [`charts`](../../charts) directory are needed to install with helm.
First, the operator should be deployed, followed by the CRDs. They should both be deployed
in the same namespace as the Ververica Platform.

This guide assumes you are operating in the base `ververica-platform-k8s-operator` directory.

```shell
# Install the Operator 
helm install --namespace vvp vp-k8s-operator ./charts/vp-k8s-operator \
    --set vvpEdition=community \ # Or enterprise
    --set vvpUrl=http://vvp-ververica-platform # pointed at the service deployed with the Ververica Platform 

# Install the CRDs
helm install --namespace vvp vp-k8s-operator-crds ./charts/vp-k8s-operator-crds \
    --set webhookCert.name=vp-k8s-operator-serving-cert \ # the Cert created by the operator chart for serving
    --set webhookCert.name=vp-k8s-operator-webhook-service # the webhook conversion service of the operator

# Wait for the deployment to come up
kubectl --namespace vvp wait --for=condition=available deployments --all
```

### Deploying the samples

The samples can all be deployed through `kubectl`.

```shell
# Now install a DeploymentTarget in the top-speed namespace
kubectl create namespace top-speed
kubectl apply -f config/samples/ververicaplatform_v1beta2_vpdeploymenttarget.yaml
kubectl apply -f config/samples/ververicaplatform_v1beta2_vpdeployment.yaml
```
