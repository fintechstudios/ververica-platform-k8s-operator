# \ClusterRoleBindingsApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClusterRoleBinding**](ClusterRoleBindingsApi.md#CreateClusterRoleBinding) | **Post** /v1/cluster-role-bindings | Create a cluster role binding
[**DeleteClusterRoleBinding**](ClusterRoleBindingsApi.md#DeleteClusterRoleBinding) | **Delete** /v1/cluster-role-bindings/{name} | Delete a cluster role binding
[**GetClusterRoleBinding**](ClusterRoleBindingsApi.md#GetClusterRoleBinding) | **Get** /v1/cluster-role-bindings/{name} | Get a cluster role binding by name
[**GetClusterRoleBindings**](ClusterRoleBindingsApi.md#GetClusterRoleBindings) | **Get** /v1/cluster-role-bindings | List all cluster role bindings
[**UpdateClusterRoleBinding**](ClusterRoleBindingsApi.md#UpdateClusterRoleBinding) | **Patch** /v1/cluster-role-bindings/{name} | Update a cluster role binding


# **CreateClusterRoleBinding**
> CreateClusterRoleBinding(ctx, body)
Create a cluster role binding



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ClusterRoleBinding**](ClusterRoleBinding.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteClusterRoleBinding**
> ClusterRoleBinding DeleteClusterRoleBinding(ctx, name)
Delete a cluster role binding



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**ClusterRoleBinding**](ClusterRoleBinding.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterRoleBinding**
> ClusterRoleBinding GetClusterRoleBinding(ctx, name)
Get a cluster role binding by name



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**ClusterRoleBinding**](ClusterRoleBinding.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterRoleBindings**
> ResourceListClusterRoleBinding GetClusterRoleBindings(ctx, optional)
List all cluster role bindings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetClusterRoleBindingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetClusterRoleBindingsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **labelSelector** | **optional.String**|  | 

### Return type

[**ResourceListClusterRoleBinding**](ResourceListClusterRoleBinding.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateClusterRoleBinding**
> ClusterRoleBinding UpdateClusterRoleBinding(ctx, name, body)
Update a cluster role binding



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **body** | [**ClusterRoleBinding**](ClusterRoleBinding.md)|  | 

### Return type

[**ClusterRoleBinding**](ClusterRoleBinding.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

