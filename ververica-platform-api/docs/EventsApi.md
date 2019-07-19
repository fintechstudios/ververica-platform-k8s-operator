# \EventsApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEvents**](EventsApi.md#GetEvents) | **Get** /v1/namespaces/{namespace}/events | Filter all events for a deployment or job


# **GetEvents**
> ResourceListEvent GetEvents(ctx, namespace, optional)
Filter all events for a deployment or job



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**|  | 
 **optional** | ***GetEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetEventsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deploymentId** | [**optional.Interface of string**](.md)|  | 
 **jobId** | [**optional.Interface of string**](.md)|  | 

### Return type

[**ResourceListEvent**](ResourceListEvent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

