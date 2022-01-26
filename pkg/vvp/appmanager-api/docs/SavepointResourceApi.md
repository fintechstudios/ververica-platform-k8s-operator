# \SavepointResourceApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSavepointUsingPOST**](SavepointResourceApi.md#CreateSavepointUsingPOST) | **Post** /api/v1/namespaces/{namespace}/savepoints | Create a new savepoint
[**GetSavepointUsingGET**](SavepointResourceApi.md#GetSavepointUsingGET) | **Get** /api/v1/namespaces/{namespace}/savepoints/{savepointId} | Get a savepoint by id
[**GetSavepointsUsingGET**](SavepointResourceApi.md#GetSavepointsUsingGET) | **Get** /api/v1/namespaces/{namespace}/savepoints | List all savepoints. Can be filtered by DeploymentId


# **CreateSavepointUsingPOST**
> Savepoint CreateSavepointUsingPOST(ctx, namespace, savepointChange)
Create a new savepoint

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| namespace | 
  **savepointChange** | [**Savepoint**](Savepoint.md)| savepointChange | 

### Return type

[**Savepoint**](Savepoint.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSavepointUsingGET**
> Savepoint GetSavepointUsingGET(ctx, namespace, savepointId)
Get a savepoint by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| namespace | 
  **savepointId** | [**string**](.md)| savepointId | 

### Return type

[**Savepoint**](Savepoint.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSavepointsUsingGET**
> ResourceListOfSavepoint GetSavepointsUsingGET(ctx, namespace, optional)
List all savepoints. Can be filtered by DeploymentId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| namespace | 
 **optional** | ***SavepointResourceApiGetSavepointsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SavepointResourceApiGetSavepointsUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deploymentId** | [**optional.Interface of string**](.md)| deploymentId | 
 **jobId** | [**optional.Interface of string**](.md)| jobId | 

### Return type

[**ResourceListOfSavepoint**](ResourceListOfSavepoint.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

