# \NamespacesApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNamespaceUsingPOST**](NamespacesApi.md#CreateNamespaceUsingPOST) | **Post** /namespaces/v1/namespaces | createNamespace
[**DeleteNamespaceUsingDELETE**](NamespacesApi.md#DeleteNamespaceUsingDELETE) | **Delete** /namespaces/v1/namespaces/{ns} | deleteNamespace
[**GetNamespaceUsingGET**](NamespacesApi.md#GetNamespaceUsingGET) | **Get** /namespaces/v1/namespaces/{ns} | getNamespace
[**ListNamespacesUsingGET**](NamespacesApi.md#ListNamespacesUsingGET) | **Get** /namespaces/v1/namespaces | listNamespaces
[**PatchNamespaceUsingPATCH**](NamespacesApi.md#PatchNamespaceUsingPATCH) | **Patch** /namespaces/v1/namespaces/{ns} | patchNamespace
[**SetPreviewSessionClusterUsingPOST**](NamespacesApi.md#SetPreviewSessionClusterUsingPOST) | **Post** /namespaces/v1/namespaces/{ns}:setPreviewSessionCluster | setPreviewSessionCluster
[**UpdateNamespaceUsingPUT**](NamespacesApi.md#UpdateNamespaceUsingPUT) | **Put** /namespaces/v1/namespaces/{ns} | updateNamespace


# **CreateNamespaceUsingPOST**
> CreateNamespaceResponse CreateNamespaceUsingPOST(ctx, namespace)
createNamespace

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | [**Namespace**](Namespace.md)| namespace | 

### Return type

[**CreateNamespaceResponse**](CreateNamespaceResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNamespaceUsingDELETE**
> DeleteNamespaceResponse DeleteNamespaceUsingDELETE(ctx, ns)
deleteNamespace

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ns** | **string**| ns | 

### Return type

[**DeleteNamespaceResponse**](DeleteNamespaceResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNamespaceUsingGET**
> GetNamespaceResponse GetNamespaceUsingGET(ctx, ns)
getNamespace

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ns** | **string**| ns | 

### Return type

[**GetNamespaceResponse**](GetNamespaceResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNamespacesUsingGET**
> ListNamespacesResponse ListNamespacesUsingGET(ctx, )
listNamespaces

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListNamespacesResponse**](ListNamespacesResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchNamespaceUsingPATCH**
> UpdateNamespaceResponse PatchNamespaceUsingPATCH(ctx, namespacePatch, ns)
patchNamespace

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespacePatch** | [**Namespace**](Namespace.md)| namespacePatch | 
  **ns** | **string**| ns | 

### Return type

[**UpdateNamespaceResponse**](UpdateNamespaceResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetPreviewSessionClusterUsingPOST**
> SetPreviewSessionClusterResponse SetPreviewSessionClusterUsingPOST(ctx, ns, request)
setPreviewSessionCluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ns** | **string**| ns | 
  **request** | [**SetPreviewSessionClusterRequest**](SetPreviewSessionClusterRequest.md)| request | 

### Return type

[**SetPreviewSessionClusterResponse**](SetPreviewSessionClusterResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNamespaceUsingPUT**
> UpdateNamespaceResponse UpdateNamespaceUsingPUT(ctx, namespace, ns)
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

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

