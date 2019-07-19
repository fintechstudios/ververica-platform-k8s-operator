# \NamespacesApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteNamespace**](NamespacesApi.md#DeleteNamespace) | **Delete** /v1/namespaces/{name} | Delete a namespace by name
[**GetNamespace**](NamespacesApi.md#GetNamespace) | **Get** /v1/namespaces/{name} | Get a namespace by name
[**ListNamespaces**](NamespacesApi.md#ListNamespaces) | **Get** /v1/namespaces | List namespaces
[**PostNamespace**](NamespacesApi.md#PostNamespace) | **Post** /v1/namespaces | Create a namespace


# **DeleteNamespace**
> Namespace DeleteNamespace(ctx, name)
Delete a namespace by name



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**Namespace**](Namespace.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNamespace**
> Namespace GetNamespace(ctx, name)
Get a namespace by name



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**Namespace**](Namespace.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNamespaces**
> ResourceListNamespace ListNamespaces(ctx, )
List namespaces



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ResourceListNamespace**](ResourceListNamespace.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostNamespace**
> Namespace PostNamespace(ctx, optional)
Create a namespace



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PostNamespaceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostNamespaceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Namespace**](Namespace.md)|  | 

### Return type

[**Namespace**](Namespace.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

