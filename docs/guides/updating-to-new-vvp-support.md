# Updating to Support a new VVP version

The main task is mirroring the new VVP API to the CRDs and generating clients for it.

## Getting new VVP API swagger spec

The swagger spec is only available by running the platform. The easiest way to do that is via the 
[playground](https://github.com/ververica/ververica-platform-playground), which has branches for each
version (e.g., `release-2.4`).

Then, port-forward the VVP service and download the Application Manager and Platform swagger.json files.

```shell
kubectl port-forward -n vvp svc/vvp-ververica-platform 8081:80
curl -o appmanager-api-swagger.json http://localhost:8081/api/swagger.json
curl -o platform-api-swagger.json http://localhost:8081/swagger.json
```

After, pretty the files (json via editor, `jq`, whatever) :)

```shell
make lint
```

## Generating + updating the clients

Next, we have to update the go clients.

First, run the make target.
```shell
make swagger-gen
```

Then, follow the **Post-Generation Changes** notes in the main readme for type changes.

## Testing

Once the clients are generated and fixed up, manual testing is the best (read: only) way to check compatibility.

Locally running the operator with a debugger is useful. CRDs without webhook configuration can be found and applied in the
`config/crd/bases` directory.

```shell
kubectl apply -f config/crd/bases/
```

Alternatively, you can built the image and run it via the Helm chart.

`config/samples` contains some resources that can be applied for testing.

```shell
kubectl  apply -f config/samples/ververicaplatform_v1beta2_vpdeploymenttarget.yaml
kubectl  apply -f config/samples/ververicaplatform_v1beta2_vpdeployment.yaml
kubectl  apply -f config/samples/ververicaplatform_v1beta1_vpsavepoint.yaml
```
