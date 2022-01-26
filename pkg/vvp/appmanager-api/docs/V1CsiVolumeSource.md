# V1CsiVolumeSource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Driver** | **string** | Driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster. | [default to null]
**FsType** | **string** | Filesystem type to mount. Ex. \&quot;ext4\&quot;, \&quot;xfs\&quot;, \&quot;ntfs\&quot;. If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply. | [optional] [default to null]
**NodePublishSecretRef** | [***V1LocalObjectReference**](V1LocalObjectReference.md) |  | [optional] [default to null]
**ReadOnly** | **bool** | Specifies a read-only configuration for the volume. Defaults to false (read/write). | [optional] [default to null]
**VolumeAttributes** | **map[string]string** | VolumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver&#39;s documentation for supported values. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


