# \NamespacesApi

All URIs are relative to *https://localhost:8081*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNamespace**](NamespacesApi.md#CreateNamespace) | **Post** /namespaces/v1/namespaces | createNamespace
[**DeleteNamespace**](NamespacesApi.md#DeleteNamespace) | **Delete** /namespaces/v1/namespaces/{ns} | deleteNamespace
[**GetNamespace**](NamespacesApi.md#GetNamespace) | **Get** /namespaces/v1/namespaces/{ns} | getNamespace
[**ListNamespaces**](NamespacesApi.md#ListNamespaces) | **Get** /namespaces/v1/namespaces | listNamespaces
[**UpdateNamespace**](NamespacesApi.md#UpdateNamespace) | **Put** /namespaces/v1/namespaces/{ns} | updateNamespace


# **CreateNamespace**
> CreateNamespaceResponse CreateNamespace(ctx, namespace)
createNamespace

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | [**Namespace**](Namespace.md)| namespace | 

### Return type

[**CreateNamespaceResponse**](CreateNamespaceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNamespace**
> DeleteNamespaceResponse DeleteNamespace(ctx, ns)
deleteNamespace

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ns** | **string**| ns | 

### Return type

[**DeleteNamespaceResponse**](DeleteNamespaceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNamespace**
> GetNamespaceResponse GetNamespace(ctx, ns)
getNamespace

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ns** | **string**| ns | 

### Return type

[**GetNamespaceResponse**](GetNamespaceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNamespaces**
> ListNamespacesResponse ListNamespaces(ctx, )
listNamespaces

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListNamespacesResponse**](ListNamespacesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNamespace**
> UpdateNamespaceResponse UpdateNamespace(ctx, namespace, ns)
updateNamespace

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | [**Namespace**](Namespace.md)| namespace | 
  **ns** | **string**| ns | 

### Return type

[**UpdateNamespaceResponse**](UpdateNamespaceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

