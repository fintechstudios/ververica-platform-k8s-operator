# \DeploymentResourceApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeploymentUsingPOST**](DeploymentResourceApi.md#CreateDeploymentUsingPOST) | **Post** /api/v1/namespaces/{namespace}/deployments | Create a deployment
[**DeleteDeploymentUsingDELETE**](DeploymentResourceApi.md#DeleteDeploymentUsingDELETE) | **Delete** /api/v1/namespaces/{namespace}/deployments/{name} | Delete deployment
[**GetDeploymentUsingGET**](DeploymentResourceApi.md#GetDeploymentUsingGET) | **Get** /api/v1/namespaces/{namespace}/deployments/{name} | Get a deployment by name
[**GetDeploymentsUsingGET**](DeploymentResourceApi.md#GetDeploymentsUsingGET) | **Get** /api/v1/namespaces/{namespace}/deployments | List all deployments
[**UpdateDeploymentUsingPATCH**](DeploymentResourceApi.md#UpdateDeploymentUsingPATCH) | **Patch** /api/v1/namespaces/{namespace}/deployments/{name} | Update a deployment
[**UpsertDeploymentUsingPUT**](DeploymentResourceApi.md#UpsertDeploymentUsingPUT) | **Put** /api/v1/namespaces/{namespace}/deployments/{name} | Create or replace the deployment


# **CreateDeploymentUsingPOST**
> Deployment CreateDeploymentUsingPOST(ctx, body, namespace)
Create a deployment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Deployment**](Deployment.md)|  | 
  **namespace** | **string**| namespace | 

### Return type

[**Deployment**](Deployment.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDeploymentUsingDELETE**
> Deployment DeleteDeploymentUsingDELETE(ctx, name, namespace)
Delete deployment

This operation expects the deployment to be in desired state CANCELLED or SUSPENDED. Acting on deployments by ID is deprecated in favor of by name, though if the name is a valid ID this endpoint will first attempt to lookup by ID for backwards compatibility, and then fallback to a name lookup if none was found.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name | 
  **namespace** | **string**| namespace | 

### Return type

[**Deployment**](Deployment.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeploymentUsingGET**
> Deployment GetDeploymentUsingGET(ctx, name, namespace)
Get a deployment by name

Acting on deployments by ID is deprecated in favor of by name, though if the name is a valid ID this endpoint will first attempt to lookup by ID for backwards compatibility, and then fallback to a name lookup if none was found.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name | 
  **namespace** | **string**| namespace | 

### Return type

[**Deployment**](Deployment.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeploymentsUsingGET**
> ResourceListOfDeployment GetDeploymentsUsingGET(ctx, namespace, optional)
List all deployments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| namespace | 
 **optional** | ***DeploymentResourceApiGetDeploymentsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeploymentResourceApiGetDeploymentsUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **labelSelector** | **optional.String**| labelSelector | 

### Return type

[**ResourceListOfDeployment**](ResourceListOfDeployment.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDeploymentUsingPATCH**
> Deployment UpdateDeploymentUsingPATCH(ctx, body, name, namespace)
Update a deployment

Acting on deployments by ID is deprecated in favor of by name, though if the name is a valid ID this endpoint will first attempt to lookup by ID for backwards compatibility, and then fallback to a name lookup if none was found.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Deployment**](Deployment.md)|  | 
  **name** | **string**| name | 
  **namespace** | **string**| namespace | 

### Return type

[**Deployment**](Deployment.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpsertDeploymentUsingPUT**
> Deployment UpsertDeploymentUsingPUT(ctx, deployment, name, namespace)
Create or replace the deployment

Acting on deployments by ID is deprecated in favor of by name, though if the name is a valid ID this endpoint will first attempt to lookup by ID for backwards compatibility, and then fallback to a name lookup if none was found.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deployment** | [**ObjectNode**](ObjectNode.md)| deployment | 
  **name** | **string**| name | 
  **namespace** | **string**| namespace | 

### Return type

[**Deployment**](Deployment.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

