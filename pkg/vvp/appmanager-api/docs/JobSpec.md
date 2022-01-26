# JobSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowNonRestoredState** | **bool** |  | [optional] [default to null]
**Artifact** | [***Artifact**](Artifact.md) |  | [optional] [default to null]
**DeploymentTarget** | [***JobDeploymentTarget**](JobDeploymentTarget.md) |  | [optional] [default to null]
**FlinkConfiguration** | **map[string]string** |  | [optional] [default to null]
**Kubernetes** | [***KubernetesOptions**](KubernetesOptions.md) |  | [optional] [default to null]
**Logging** | [***Logging**](Logging.md) |  | [optional] [default to null]
**NumberOfTaskManagers** | **int32** |  | [optional] [default to null]
**Parallelism** | **int32** |  | [optional] [default to null]
**Resources** | [**map[string]ResourceSpec**](ResourceSpec.md) |  | [optional] [default to null]
**SavepointLocation** | **string** |  | [optional] [default to null]
**UserFlinkConfiguration** | **map[string]string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


