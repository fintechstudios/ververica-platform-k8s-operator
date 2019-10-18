# Savepoints

Like K8s, the Ververica Platform also has the concept of namespaces.
These were the easiest to map over.

[Official Ververica Platform Docs](https://docs.ververica.com/user_guide/deployments/savepoints.html)

## Ververica Platform Definition

```yaml
apiVersion: v1
kind: Namespace
metadata:
    name: String # Required
    id: UUID String # Dynamic
    createdAt: Timestamp # Dynamic
    modifiedAt: Timestamp  # Dynamic
    resourceVersion: Integer # Dynamic
status:
  state: String # Dynamic
```

## K8s Definition

```yaml
apiVersion: ververicaplatform.fintechstudios.com/v1beta1
kind: VpNamespace
metadata:
  name: String # Required
spec:
  metadata:
    name: String # Dynamic
    id: UUID String # Dynamic
    createdAt: Timestamp # Dynamic
    modifiedAt: Timestamp  # Dynamic
    resourceVersion: Integer # Dynamic
status:
  state: String # Dynamic
```

You can find an example in [config/samples/ververicaplatform_v1beta1_vpnamespace.yaml](../../config/samples/ververicaplatform_v1beta1_vpnamespace.yaml).

