# \EventResourceApi

All URIs are relative to *https://localhost:8080*

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
 **optional** | ***EventResourceApiGetEventsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventResourceApiGetEventsUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deploymentId** | [**optional.Interface of string**](.md)| deploymentId | 
 **jobId** | [**optional.Interface of string**](.md)| jobId | 
 **sessionClusterId** | [**optional.Interface of string**](.md)| sessionClusterId | 

### Return type

[**ResourceListOfEvent**](ResourceListOfEvent.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

