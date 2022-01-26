# \JobGraphApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateJobGraphUsingPOST**](JobGraphApi.md#CreateJobGraphUsingPOST) | **Post** /sql/v1beta1/namespaces/{ns}/deployments/{deploymentId}:create-jobgraph | createJobGraph
[**DeleteJobGraphsUsingDELETE**](JobGraphApi.md#DeleteJobGraphsUsingDELETE) | **Delete** /sql/v1beta1/namespaces/{ns}/deployments/{deploymentName}:delete-jobgraphs | deleteJobGraphs
[**ValidateConfigChangeUsingPOST**](JobGraphApi.md#ValidateConfigChangeUsingPOST) | **Post** /sql/v1beta1/namespaces/{ns}/deployments/{deploymentName}:validate-config-change | validateConfigChange


# **CreateJobGraphUsingPOST**
> CreateJobGraphResponse CreateJobGraphUsingPOST(ctx, deploymentName, jobGraphSpec, ns)
createJobGraph

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deploymentName** | **string**| deploymentName | 
  **jobGraphSpec** | [**JobGraphSpec**](JobGraphSpec.md)| jobGraphSpec | 
  **ns** | **string**| ns | 

### Return type

[**CreateJobGraphResponse**](CreateJobGraphResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteJobGraphsUsingDELETE**
> DeleteJobGraphsResponse DeleteJobGraphsUsingDELETE(ctx, deploymentName, ns)
deleteJobGraphs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deploymentName** | **string**| deploymentName | 
  **ns** | **string**| ns | 

### Return type

[**DeleteJobGraphsResponse**](DeleteJobGraphsResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateConfigChangeUsingPOST**
> ValidateConfigChangeResponse ValidateConfigChangeUsingPOST(ctx, configSpec, deploymentName, ns)
validateConfigChange

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **configSpec** | [**ConfigSpec**](ConfigSpec.md)| configSpec | 
  **deploymentName** | **string**| deploymentName | 
  **ns** | **string**| ns | 

### Return type

[**ValidateConfigChangeResponse**](ValidateConfigChangeResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

