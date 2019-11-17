# Namespaces

Like K8s, the Ververica Platform also has the concept of namespaces.
These were the easiest to map over.

[Official Ververica Platform Docs](https://docs.ververica.com/administration/namespaces.html)


## K8s Definition

```yaml
apiVersion: ververicaplatform.fintechstudios.com/v1beta1
kind: VpNamespace
metadata:
  name: String # Required
spec:
  roleBindings:
    - members: String[]
      role: String
status:
  lifecyclePhase: String
```

You can find an example in [config/samples/ververicaplatform_v1beta1_vpnamespace.yaml](../../config/samples/ververicaplatform_v1beta1_vpnamespace.yaml).

