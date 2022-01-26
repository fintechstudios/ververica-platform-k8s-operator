# V1PodSecurityContext

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FsGroup** | **int64** | A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod:  1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR&#39;d with rw-rw----  If unset, the Kubelet will not modify the ownership and permissions of any volume. | [optional] [default to null]
**FsGroupChangePolicy** | **string** | fsGroupChangePolicy defines behavior of changing ownership and permission of the volume before being exposed inside Pod. This field will only apply to volume types which support fsGroup based ownership(and permissions). It will have no effect on ephemeral volume types such as: secret, configmaps and emptydir. Valid values are \&quot;OnRootMismatch\&quot; and \&quot;Always\&quot;. If not specified defaults to \&quot;Always\&quot;. | [optional] [default to null]
**RunAsGroup** | **int64** | The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. | [optional] [default to null]
**RunAsNonRoot** | **bool** | Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. | [optional] [default to null]
**RunAsUser** | **int64** | The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. | [optional] [default to null]
**SeLinuxOptions** | [***V1SeLinuxOptions**](V1SELinuxOptions.md) |  | [optional] [default to null]
**SeccompProfile** | [***V1SeccompProfile**](V1SeccompProfile.md) |  | [optional] [default to null]
**SupplementalGroups** | **[]int64** | A list of groups applied to the first process run in each container, in addition to the container&#39;s primary GID.  If unspecified, no groups will be added to any container. | [optional] [default to null]
**Sysctls** | [**[]V1Sysctl**](V1Sysctl.md) | Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported sysctls (by the container runtime) might fail to launch. | [optional] [default to null]
**WindowsOptions** | [***V1WindowsSecurityContextOptions**](V1WindowsSecurityContextOptions.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


