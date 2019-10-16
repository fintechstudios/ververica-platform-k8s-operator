# \SavepointsApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSavepoint**](SavepointsApi.md#CreateSavepoint) | **Post** /v1/namespaces/{namespace}/savepoints | Create a new savepoint
[**GetSavepoint**](SavepointsApi.md#GetSavepoint) | **Get** /v1/namespaces/{namespace}/savepoints/{savepointId} | Get a savepoint by id
[**GetSavepoints**](SavepointsApi.md#GetSavepoints) | **Get** /v1/namespaces/{namespace}/savepoints | List all savepoints. Can be filtered by DeploymentId


# **CreateSavepoint**
> CreateSavepoint(ctx, namespace, body)
Create a new savepoint



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**|  | 
  **body** | [**Savepoint**](Savepoint.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSavepoint**
> Savepoint GetSavepoint(ctx, namespace, savepointId)
Get a savepoint by id



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**|  | 
  **savepointId** | [**string**](.md)|  | 

### Return type

[**Savepoint**](Savepoint.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSavepoints**
> []Savepoint GetSavepoints(ctx, namespace, optional)
List all savepoints. Can be filtered by DeploymentId



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**|  | 
 **optional** | ***GetSavepointsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSavepointsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deploymentId** | [**optional.Interface of string**](.md)|  | 
 **jobId** | [**optional.Interface of string**](.md)|  | 

### Return type

[**[]Savepoint**](Savepoint.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

