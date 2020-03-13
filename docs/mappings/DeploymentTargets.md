# Deployment Targets

Somewhat more difficult than, namespaces, as arbitrary JSON doesn't 
seem to play nice with the K8s code generators.

[Official Ververica Docs](https://docs.ververica.com/administration/deployment_targets.html)

## Ververica Platform Definition

```yaml
apiVersion: v1
kind: DeploymentTarget
metadata:
    name: String
    namespace: String # defaults to "default"
    id: UUID String
    createdAt: ISO8601 String
    modifiedAt: ISO8601 String
    resourceVersion: Integer
spec:
  kubernetes:
    namespace: String
```

## K8s Definition

```yaml
apiVersion: ververicaplatform.fintechstudios.com/v1beta1
kind: VpDeploymentTarget
metadata:
  name: String # Required
spec:
  metadata:
    namespace: String # defaults to "default"
    id: UUID String # Dynamic
    createdAt: Time # Dynamic
    modifiedAt: Time  # Dynamic
    resourceVersion: Integer # Dynamic
  spec:
    kubernetes: # Required
      namespace: String # Optional
```

You can find an example in [config/samples/ververicaplatform_v1beta1_vpdeploymenttarget.yaml](../../config/samples/ververicaplatform_v1beta1_vpdeploymenttarget.yaml).
