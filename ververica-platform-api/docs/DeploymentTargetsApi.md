# \DeploymentTargetsApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeploymentTarget**](DeploymentTargetsApi.md#CreateDeploymentTarget) | **Post** /v1/namespaces/{namespace}/deployment-targets | Create a deployment target
[**DeleteDeploymentTarget**](DeploymentTargetsApi.md#DeleteDeploymentTarget) | **Delete** /v1/namespaces/{namespace}/deployment-targets/{name} | Delete a deployment target
[**GetDeploymentTarget**](DeploymentTargetsApi.md#GetDeploymentTarget) | **Get** /v1/namespaces/{namespace}/deployment-targets/{name} | Get a deployment target by name
[**GetDeploymentTargets**](DeploymentTargetsApi.md#GetDeploymentTargets) | **Get** /v1/namespaces/{namespace}/deployment-targets | List all deployment targets


# **CreateDeploymentTarget**
> DeploymentTarget CreateDeploymentTarget(ctx, namespace, body)
Create a deployment target



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**|  | 
  **body** | [**DeploymentTarget**](DeploymentTarget.md)|  | 

### Return type

[**DeploymentTarget**](DeploymentTarget.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDeploymentTarget**
> DeploymentTarget DeleteDeploymentTarget(ctx, namespace, name)
Delete a deployment target



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**|  | 
  **name** | **string**|  | 

### Return type

[**DeploymentTarget**](DeploymentTarget.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeploymentTarget**
> DeploymentTarget GetDeploymentTarget(ctx, namespace, name)
Get a deployment target by name



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**|  | 
  **name** | **string**|  | 

### Return type

[**DeploymentTarget**](DeploymentTarget.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeploymentTargets**
> ResourceListDeploymentTarget GetDeploymentTargets(ctx, namespace)
List all deployment targets



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**|  | 

### Return type

[**ResourceListDeploymentTarget**](ResourceListDeploymentTarget.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

