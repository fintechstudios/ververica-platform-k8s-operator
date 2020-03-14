# \DeploymentDefaultsResourceApi

All URIs are relative to *https://localhost:8081*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeploymentDefaultsUsingGET**](DeploymentDefaultsResourceApi.md#GetDeploymentDefaultsUsingGET) | **Get** /api/v1/namespaces/{namespace}/deployment-defaults | Get deployment defaults
[**UpdateDeploymentDefaultsUsingPATCH**](DeploymentDefaultsResourceApi.md#UpdateDeploymentDefaultsUsingPATCH) | **Patch** /api/v1/namespaces/{namespace}/deployment-defaults | Update a deployment defaults


# **GetDeploymentDefaultsUsingGET**
> DeploymentDefaults GetDeploymentDefaultsUsingGET(ctx, namespace)
Get deployment defaults

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| namespace | 

### Return type

[**DeploymentDefaults**](DeploymentDefaults.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDeploymentDefaultsUsingPATCH**
> DeploymentDefaults UpdateDeploymentDefaultsUsingPATCH(ctx, body, namespace)
Update a deployment defaults

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeploymentDefaults**](DeploymentDefaults.md)|  | 
  **namespace** | **string**| namespace | 

### Return type

[**DeploymentDefaults**](DeploymentDefaults.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

