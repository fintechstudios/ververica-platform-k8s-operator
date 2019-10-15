# \ClusterRolesApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClusterRole**](ClusterRolesApi.md#CreateClusterRole) | **Post** /v1/cluster-roles | Create a cluster role
[**DeleteClusterRole**](ClusterRolesApi.md#DeleteClusterRole) | **Delete** /v1/cluster-roles/{name} | Delete a cluster role
[**GetClusterRole**](ClusterRolesApi.md#GetClusterRole) | **Get** /v1/cluster-roles/{name} | Get a cluster role by name
[**GetClusterRoles**](ClusterRolesApi.md#GetClusterRoles) | **Get** /v1/cluster-roles | List all cluster roles
[**UpdateClusterRole**](ClusterRolesApi.md#UpdateClusterRole) | **Patch** /v1/cluster-roles/{name} | Update a cluster role


# **CreateClusterRole**
> CreateClusterRole(ctx, body)
Create a cluster role



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ClusterRole**](ClusterRole.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteClusterRole**
> ClusterRole DeleteClusterRole(ctx, name)
Delete a cluster role



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**ClusterRole**](ClusterRole.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterRole**
> ClusterRole GetClusterRole(ctx, name)
Get a cluster role by name



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**ClusterRole**](ClusterRole.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterRoles**
> ResourceListClusterRole GetClusterRoles(ctx, optional)
List all cluster roles



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetClusterRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetClusterRolesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **labelSelector** | **optional.String**|  | 

### Return type

[**ResourceListClusterRole**](ResourceListClusterRole.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateClusterRole**
> ClusterRole UpdateClusterRole(ctx, name, body)
Update a cluster role



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **body** | [**ClusterRole**](ClusterRole.md)|  | 

### Return type

[**ClusterRole**](ClusterRole.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

