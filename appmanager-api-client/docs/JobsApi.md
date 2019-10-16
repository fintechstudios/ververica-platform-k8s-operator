# \JobsApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetJob**](JobsApi.md#GetJob) | **Get** /v1/namespaces/{namespace}/jobs/{jobId} | Get a job by id
[**GetJobs**](JobsApi.md#GetJobs) | **Get** /v1/namespaces/{namespace}/jobs | List all jobs. Can be filtered by DeploymentId


# **GetJob**
> Job GetJob(ctx, namespace, jobId)
Get a job by id



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**|  | 
  **jobId** | [**string**](.md)|  | 

### Return type

[**Job**](Job.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJobs**
> ResourceListJob GetJobs(ctx, namespace, optional)
List all jobs. Can be filtered by DeploymentId



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**|  | 
 **optional** | ***GetJobsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetJobsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deploymentId** | [**optional.Interface of string**](.md)|  | 

### Return type

[**ResourceListJob**](ResourceListJob.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

