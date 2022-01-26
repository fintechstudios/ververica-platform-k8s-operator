# V1HttpGetAction

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | Host name to connect to, defaults to the pod IP. You probably want to set \&quot;Host\&quot; in httpHeaders instead. | [optional] [default to null]
**HttpHeaders** | [**[]V1HttpHeader**](V1HTTPHeader.md) | Custom headers to set in the request. HTTP allows repeated headers. | [optional] [default to null]
**Path** | **string** | Path to access on the HTTP server. | [optional] [default to null]
**Port** | [***IntOrString**](IntOrString.md) | IntOrString is a type that can hold an int32 or a string.  When used in JSON or YAML marshalling and unmarshalling, it produces or consumes the inner type.  This allows you to have, for example, a JSON field that can accept a name or number. | [default to null]
**Scheme** | **string** | Scheme to use for connecting to the host. Defaults to HTTP. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


