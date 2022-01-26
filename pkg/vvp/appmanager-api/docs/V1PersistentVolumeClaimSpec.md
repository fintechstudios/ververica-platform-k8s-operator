# V1PersistentVolumeClaimSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessModes** | **[]string** | AccessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1 | [optional] [default to null]
**DataSource** | [***V1TypedLocalObjectReference**](V1TypedLocalObjectReference.md) |  | [optional] [default to null]
**Resources** | [***V1ResourceRequirements**](V1ResourceRequirements.md) |  | [optional] [default to null]
**Selector** | [***V1LabelSelector**](V1LabelSelector.md) |  | [optional] [default to null]
**StorageClassName** | **string** | Name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1 | [optional] [default to null]
**VolumeMode** | **string** | volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec. | [optional] [default to null]
**VolumeName** | **string** | VolumeName is the binding reference to the PersistentVolume backing this claim. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


