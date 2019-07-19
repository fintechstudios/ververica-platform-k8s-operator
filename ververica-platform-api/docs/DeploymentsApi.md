# \DeploymentsApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeployment**](DeploymentsApi.md#CreateDeployment) | **Post** /v1/namespaces/{namespace}/deployments | Create a deployment
[**DeleteDeployment**](DeploymentsApi.md#DeleteDeployment) | **Delete** /v1/namespaces/{namespace}/deployments/{deploymentId} | Delete deployment
[**GetDeployment**](DeploymentsApi.md#GetDeployment) | **Get** /v1/namespaces/{namespace}/deployments/{deploymentId} | Get a deployment by id
[**GetDeployments**](DeploymentsApi.md#GetDeployments) | **Get** /v1/namespaces/{namespace}/deployments | List all deployments
[**UpdateDeployment**](DeploymentsApi.md#UpdateDeployment) | **Patch** /v1/namespaces/{namespace}/deployments/{deploymentId} | Update a deployment


# **CreateDeployment**
> CreateDeployment(ctx, namespace, body)
Create a deployment



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**|  | 
  **body** | [**Deployment**](Deployment.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDeployment**
> Deployment DeleteDeployment(ctx, namespace, deploymentId)
Delete deployment

This operation expects the deployment to be in desired state CANCELLED

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**|  | 
  **deploymentId** | [**string**](.md)|  | 

### Return type

[**Deployment**](Deployment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeployment**
> Deployment GetDeployment(ctx, namespace, deploymentId)
Get a deployment by id



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**|  | 
  **deploymentId** | [**string**](.md)|  | 

### Return type

[**Deployment**](Deployment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeployments**
> ResourceListDeployment GetDeployments(ctx, namespace, optional)
List all deployments



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**|  | 
 **optional** | ***GetDeploymentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDeploymentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **labelSelector** | **optional.String**|  | 

### Return type

[**ResourceListDeployment**](ResourceListDeployment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDeployment**
> Deployment UpdateDeployment(ctx, namespace, deploymentId, body)
Update a deployment



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**|  | 
  **deploymentId** | [**string**](.md)|  | 
  **body** | [**Deployment**](Deployment.md)|  | 

### Return type

[**Deployment**](Deployment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

