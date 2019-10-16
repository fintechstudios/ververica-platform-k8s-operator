# \RoleBindingsApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoleBinding**](RoleBindingsApi.md#CreateRoleBinding) | **Post** /v1/namespaces/{namespace}/role-bindings | Create a role binding
[**DeleteRoleBinding**](RoleBindingsApi.md#DeleteRoleBinding) | **Delete** /v1/namespaces/{namespace}/role-bindings/{name} | Delete a role binding
[**GetRoleBinding**](RoleBindingsApi.md#GetRoleBinding) | **Get** /v1/namespaces/{namespace}/role-bindings/{name} | Get a role binding by name
[**GetRoleBindings**](RoleBindingsApi.md#GetRoleBindings) | **Get** /v1/namespaces/{namespace}/role-bindings | List all role bindings
[**UpdateRoleBinding**](RoleBindingsApi.md#UpdateRoleBinding) | **Patch** /v1/namespaces/{namespace}/role-bindings/{name} | Update a role binding


# **CreateRoleBinding**
> CreateRoleBinding(ctx, namespace, body)
Create a role binding



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**|  | 
  **body** | [**RoleBinding**](RoleBinding.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRoleBinding**
> RoleBinding DeleteRoleBinding(ctx, namespace, name)
Delete a role binding



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**|  | 
  **name** | **string**|  | 

### Return type

[**RoleBinding**](RoleBinding.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRoleBinding**
> RoleBinding GetRoleBinding(ctx, namespace, name)
Get a role binding by name



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**|  | 
  **name** | **string**|  | 

### Return type

[**RoleBinding**](RoleBinding.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRoleBindings**
> ResourceListRoleBinding GetRoleBindings(ctx, namespace, optional)
List all role bindings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**|  | 
 **optional** | ***GetRoleBindingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetRoleBindingsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **labelSelector** | **optional.String**|  | 

### Return type

[**ResourceListRoleBinding**](ResourceListRoleBinding.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRoleBinding**
> RoleBinding UpdateRoleBinding(ctx, namespace, name, body)
Update a role binding



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**|  | 
  **name** | **string**|  | 
  **body** | [**RoleBinding**](RoleBinding.md)|  | 

### Return type

[**RoleBinding**](RoleBinding.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

