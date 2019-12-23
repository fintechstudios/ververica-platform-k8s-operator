# Events

For each Ververica Platform deployment, the Event Log is watched and synchronized to native Kubernetes Events. Each
event is attached to the K8s VpDeployment object.

[Official Ververica Docs](https://docs.ververica.com/user_guide/deployments/event_log.html)

Some useful commands:
```shell
# Create a deployment
kubectl apply -f config/samples/ververicaplatform_v1beta1_{vpnamespace,vpdeploymenttarget_testing,vpdeployment}.yaml

# Watch the events in real time
kubectl get events --watch

# List by latest at the bottom
kubectl get events --sort-by='lastTimestamp'

# See the attached events
kubectl describe -f config/samples/ververicaplatform_v1beta1_vpdeployment.yaml
```
