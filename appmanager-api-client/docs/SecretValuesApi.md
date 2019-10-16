# \SecretValuesApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecretValue**](SecretValuesApi.md#CreateSecretValue) | **Post** /v1/namespaces/{namespace}/secret-values | Create a secret value
[**DeleteSecretValue**](SecretValuesApi.md#DeleteSecretValue) | **Delete** /v1/namespaces/{namespace}/secret-values/{name} | Delete a secret value
[**GetSecretValue**](SecretValuesApi.md#GetSecretValue) | **Get** /v1/namespaces/{namespace}/secret-values/{name} | Get a secret value by name
[**GetSecretValues**](SecretValuesApi.md#GetSecretValues) | **Get** /v1/namespaces/{namespace}/secret-values | List all secrets values
[**UpdateSecretValue**](SecretValuesApi.md#UpdateSecretValue) | **Patch** /v1/namespaces/{namespace}/secret-values/{name} | Update a secret value


# **CreateSecretValue**
> SecretValue CreateSecretValue(ctx, namespace, body)
Create a secret value



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**|  | 
  **body** | [**SecretValue**](SecretValue.md)|  | 

### Return type

[**SecretValue**](SecretValue.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSecretValue**
> SecretValue DeleteSecretValue(ctx, namespace, name)
Delete a secret value



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**|  | 
  **name** | **string**|  | 

### Return type

[**SecretValue**](SecretValue.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSecretValue**
> SecretValue GetSecretValue(ctx, namespace, name)
Get a secret value by name



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**|  | 
  **name** | **string**|  | 

### Return type

[**SecretValue**](SecretValue.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSecretValues**
> ResourceListSecretValue GetSecretValues(ctx, namespace)
List all secrets values



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**|  | 

### Return type

[**ResourceListSecretValue**](ResourceListSecretValue.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSecretValue**
> SecretValue UpdateSecretValue(ctx, namespace, name, body)
Update a secret value



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**|  | 
  **name** | **string**|  | 
  **body** | [**SecretValue**](SecretValue.md)|  | 

### Return type

[**SecretValue**](SecretValue.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

