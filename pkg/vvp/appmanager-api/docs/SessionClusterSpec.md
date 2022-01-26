# SessionClusterSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** |  | [optional] [default to null]
**DeploymentTargetName** | **string** |  | [optional] [default to null]
**FlinkVersion** | **string** |  | [optional] [default to null]
**FlinkImageRegistry** | **string** |  | [optional] [default to null]
**FlinkImageRepository** | **string** |  | [optional] [default to null]
**FlinkImageTag** | **string** |  | [optional] [default to null]
**NumberOfTaskManagers** | **int32** |  | [optional] [default to null]
**Resources** | [**map[string]ResourceSpec**](ResourceSpec.md) |  | [optional] [default to null]
**FlinkConfiguration** | **map[string]string** |  | [optional] [default to null]
**Logging** | [***Logging**](Logging.md) |  | [optional] [default to null]
**Kubernetes** | [***KubernetesOptions**](KubernetesOptions.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


