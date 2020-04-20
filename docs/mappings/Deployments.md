# Deployments

Somewhat more difficult than, namespaces, as arbitrary JSON doesn't 
seem to play nicely with the K8s code generators.

[Official Ververica Docs](https://docs.ververica.com/user_guide/deployments/index.html)

## Ververica Platform Definition

```yaml
apiVersion: v1
kind: Deployment
metadata:
    namespace: String
spec:
    state: STATE String
    deploymentTargetId: String
    upgradeStrategy:
      kind: UPGRADE_STRATEGY String
    restoreStrategy:
      kind: RESTORE_STRATEGY String
    maxSavepointCreationAttempts: Integer
    maxJobCreationAttempts: Integer
    template:
        metadata:
        artifact:
          kind: String
          jarUri: String
          mainArgs: String
          entryClass: String
          flinkVersion: String
          flinkImageRegistry: String
          flinkImageRepository: String
          flinkImageTag: String
        parallelism: Integer
        numberOfTaskManagers: Integer
        resources:
          jobmanager:
            cpu: Number
            memory: String # with memory unit, ie 1g
          taskmanager:
            cpu: Number
            memory: String # with memory unit, ie 1g
        flinkConfiguration: map[string]string 
        logging:
          log4jLoggers: map[string]string
        kubernetes:
          pods:
            annotations: map[string]string
status:
    state: STATE string
```

## K8s Definition

```yaml
apiVersion: ververicaplatform.fintechstudios.com/v1beta2
kind: VpDeployment
metadata:
  name: String
spec:
  deploymentTargetName: String # Addition, will dynamically look up target id
  metadata:
    namespace: String
  spec:
    state: STATE String
    upgradeStrategy:
      kind: UPGRADE_STRATEGY String
    restoreStrategy:
      kind: RESTORE_STRATEGY String
    maxSavepointCreationAttempts: Integer
    maxJobCreationAttempts: Integer
    template:
      metadata:
        annotations: map[string]string
      spec:
        artifact:
          kind: String
          jarUri: String
          mainArgs: String
          entryClass: String
          flinkVersion: String
          flinkImageRegistry: String
          flinkImageRepository: String
          flinkImageTag: String
        parallelism: Integer
        numberOfTaskManagers: Integer
        resources:
          jobmanager:
            cpu: Quantity String # the k8s float 
            memory: String
          taskmanager:
            cpu: Quantity String # the k8s float
            memory: String
        flinkConfiguration: map[string]string
        logging: map[string]string
        kubernetes:
          pods:
            annotations: map[string]string
            labels: map[string]string
            # ...
status:
    state: STATE string
```

The main changes from the Ververica Platform are:
- resource CPU must be a `Quantity` string, though this should normally just mean you must quote your value
- You can specify a `spec.deploymentTargetName` instead of an ID 

You can find an example in [config/samples/ververicaplatform_v1beta1_vpdeployment.yaml](../../config/samples/ververicaplatform_v1beta1_vpdeployment.yaml).

