# \EventResourceApi

All URIs are relative to *https://localhost:8081*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEventsUsingGET**](EventResourceApi.md#GetEventsUsingGET) | **Get** /api/v1/namespaces/{namespace}/events | Filter all events for a deployment or job


# **GetEventsUsingGET**
> ResourceListOfEvent GetEventsUsingGET(ctx, namespace, optional)
Filter all events for a deployment or job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| namespace | 
 **optional** | ***GetEventsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetEventsUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deploymentId** | [**optional.Interface of string**](.md)| deploymentId | 
 **jobId** | [**optional.Interface of string**](.md)| jobId | 

### Return type

[**ResourceListOfEvent**](ResourceListOfEvent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

