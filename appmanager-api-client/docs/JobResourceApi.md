# \JobResourceApi

All URIs are relative to *https://localhost:8081*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetJobUsingGET**](JobResourceApi.md#GetJobUsingGET) | **Get** /api/v1/namespaces/{namespace}/jobs/{jobId} | Get a job by id
[**GetJobsUsingGET**](JobResourceApi.md#GetJobsUsingGET) | **Get** /api/v1/namespaces/{namespace}/jobs | List all jobs. Can be filtered by DeploymentId


# **GetJobUsingGET**
> Job GetJobUsingGET(ctx, jobId, namespace)
Get a job by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | [**string**](.md)| jobId | 
  **namespace** | **string**| namespace | 

### Return type

[**Job**](Job.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJobsUsingGET**
> ResourceListOfJob GetJobsUsingGET(ctx, namespace, optional)
List all jobs. Can be filtered by DeploymentId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| namespace | 
 **optional** | ***GetJobsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetJobsUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deploymentId** | [**optional.Interface of string**](.md)| deploymentId | 

### Return type

[**ResourceListOfJob**](ResourceListOfJob.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

