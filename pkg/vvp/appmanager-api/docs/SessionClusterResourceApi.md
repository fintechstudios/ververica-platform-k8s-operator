# \SessionClusterResourceApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSessionClusterUsingPOST**](SessionClusterResourceApi.md#CreateSessionClusterUsingPOST) | **Post** /api/v1/namespaces/{namespace}/sessionclusters | Create a session cluster
[**DeleteSessionClusterUsingDELETE**](SessionClusterResourceApi.md#DeleteSessionClusterUsingDELETE) | **Delete** /api/v1/namespaces/{namespace}/sessionclusters/{name} | Delete a session cluster
[**GetSessionClusterUsingGET**](SessionClusterResourceApi.md#GetSessionClusterUsingGET) | **Get** /api/v1/namespaces/{namespace}/sessionclusters/{name} | Get a session cluster by name
[**GetSessionClustersUsingGET**](SessionClusterResourceApi.md#GetSessionClustersUsingGET) | **Get** /api/v1/namespaces/{namespace}/sessionclusters | List all session clusters
[**UpdateSessionClusterUsingPATCH**](SessionClusterResourceApi.md#UpdateSessionClusterUsingPATCH) | **Patch** /api/v1/namespaces/{namespace}/sessionclusters/{name} | Update a session cluster
[**UpsertSessionClusterUsingPUT**](SessionClusterResourceApi.md#UpsertSessionClusterUsingPUT) | **Put** /api/v1/namespaces/{namespace}/sessionclusters/{name} | Create or replace the session cluster


# **CreateSessionClusterUsingPOST**
> SessionCluster CreateSessionClusterUsingPOST(ctx, namespace, sessionCluster)
Create a session cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| namespace | 
  **sessionCluster** | [**SessionCluster**](SessionCluster.md)| sessionCluster | 

### Return type

[**SessionCluster**](SessionCluster.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSessionClusterUsingDELETE**
> SessionCluster DeleteSessionClusterUsingDELETE(ctx, name, namespace)
Delete a session cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name | 
  **namespace** | **string**| namespace | 

### Return type

[**SessionCluster**](SessionCluster.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSessionClusterUsingGET**
> SessionCluster GetSessionClusterUsingGET(ctx, name, namespace)
Get a session cluster by name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name | 
  **namespace** | **string**| namespace | 

### Return type

[**SessionCluster**](SessionCluster.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSessionClustersUsingGET**
> ResourceListOfSessionCluster GetSessionClustersUsingGET(ctx, namespace)
List all session clusters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| namespace | 

### Return type

[**ResourceListOfSessionCluster**](ResourceListOfSessionCluster.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSessionClusterUsingPATCH**
> SessionCluster UpdateSessionClusterUsingPATCH(ctx, body, name, namespace)
Update a session cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SessionCluster**](SessionCluster.md)|  | 
  **name** | **string**| name | 
  **namespace** | **string**| namespace | 

### Return type

[**SessionCluster**](SessionCluster.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpsertSessionClusterUsingPUT**
> SessionCluster UpsertSessionClusterUsingPUT(ctx, name, namespace, sessionCluster)
Create or replace the session cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name | 
  **namespace** | **string**| namespace | 
  **sessionCluster** | [**SessionCluster**](SessionCluster.md)| sessionCluster | 

### Return type

[**SessionCluster**](SessionCluster.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

