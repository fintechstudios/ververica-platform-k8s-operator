# Design

The goal of the initial version is to provide native support for
Ververica Platform resources to Kubernetes with as little friction as 
possible. For this reason, the Ververica Platform resource definitions
are, whenever possible, directly embedded in the K8s resource's `spec`. The only exception here
is that the name is always taken from the K8s resource, and the API version and Kind
are automatically mapped.

The `status` of Ververica Platform resources is attempted to be mirrored back
to the K8s resource `status`, where a bit of extra information is also stored.