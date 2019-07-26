# Design

The goal of the initial version is to provide native support for
Ververica Platform resources to Kubernetes with as little friction as 
possible. For this reason, the Ververica Platform resource definitions
are, whenever possible, directly embedded in the K8s resource's `spec`. The only exception here
is that the name is always taken from the K8s resource, and the API version and Kind
are automatically mapped.

The `status` of Ververica Platform resources is attempted to be mirrored back
to the K8s resource `status`, where a bit of extra information is also stored.


## Reconciliation Loop

Each resource has a `Reconciler`, or a process that is trying to keep
all resources of that type in sync between Kubernetes and the Ververica Platform.

When an event happens on an object, it is up to the reconciler to determine
the desired state and make the current state match, if possible. When there are errors,
the reconcilliation attempt can be scheduled again and looped, helping to account for lags or race
conditions between resource-interdependency.  

Take a look at OpenShift's [Operator Best Practices](https://blog.openshift.com/kubernetes-operators-best-practices/)
section on _Resource Reconciliation Cycle_ for more of what we should be doing here.

You can see it in action by running the controller and then applying the entire samples directory:

```bash
kubectl apply -f ./config/samples
# Watch as some updates might fail the first time through
# as they wait for their dependencies to come online

kubectl delete -f ./config/samples
# Watch as some deletions might take more than one loop
# as their dependents wait to finish deletion
```

## Resource Deletion

All managed resources are appended with a finalizer. When the resource
is deleted in K8s, we wait until the resource is removed from the Ververica Platform
to remove the finalizer, allowing K8s to know something has not yet
finished cleaning up. 