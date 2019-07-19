# \RolesApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRole**](RolesApi.md#CreateRole) | **Post** /v1/namespaces/{namespace}/roles | Create a role
[**DeleteRole**](RolesApi.md#DeleteRole) | **Delete** /v1/namespaces/{namespace}/roles/{name} | Delete a role
[**GetRole**](RolesApi.md#GetRole) | **Get** /v1/namespaces/{namespace}/roles/{name} | Get a role by name
[**GetRoles**](RolesApi.md#GetRoles) | **Get** /v1/namespaces/{namespace}/roles | List all roles
[**UpdateRole**](RolesApi.md#UpdateRole) | **Patch** /v1/namespaces/{namespace}/roles/{name} | Update a role


# **CreateRole**
> CreateRole(ctx, namespace, body)
Create a role



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**|  | 
  **body** | [**Role**](Role.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRole**
> Role DeleteRole(ctx, namespace, name)
Delete a role



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**|  | 
  **name** | **string**|  | 

### Return type

[**Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRole**
> Role GetRole(ctx, namespace, name)
Get a role by name



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**|  | 
  **name** | **string**|  | 

### Return type

[**Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRoles**
> ResourceListRole GetRoles(ctx, namespace, optional)
List all roles



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**|  | 
 **optional** | ***GetRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetRolesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **labelSelector** | **optional.String**|  | 

### Return type

[**ResourceListRole**](ResourceListRole.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRole**
> Role UpdateRole(ctx, namespace, name, body)
Update a role



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**|  | 
  **name** | **string**|  | 
  **body** | [**Role**](Role.md)|  | 

### Return type

[**Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

