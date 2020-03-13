# \SecretValueResourceApi

All URIs are relative to *https://localhost:8081*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecretValueUsingPOST**](SecretValueResourceApi.md#CreateSecretValueUsingPOST) | **Post** /api/v1/namespaces/{namespace}/secret-values | Create a secret value
[**DeleteSecretValueUsingDELETE**](SecretValueResourceApi.md#DeleteSecretValueUsingDELETE) | **Delete** /api/v1/namespaces/{namespace}/secret-values/{name} | Delete a secret value
[**GetSecretValueUsingGET**](SecretValueResourceApi.md#GetSecretValueUsingGET) | **Get** /api/v1/namespaces/{namespace}/secret-values/{name} | Get a secret value by name
[**GetSecretValuesUsingGET**](SecretValueResourceApi.md#GetSecretValuesUsingGET) | **Get** /api/v1/namespaces/{namespace}/secret-values | List all secrets values
[**UpdateSecretValueUsingPATCH**](SecretValueResourceApi.md#UpdateSecretValueUsingPATCH) | **Patch** /api/v1/namespaces/{namespace}/secret-values/{name} | Update a secret value


# **CreateSecretValueUsingPOST**
> SecretValue CreateSecretValueUsingPOST(ctx, namespace, secretValue)
Create a secret value

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| namespace | 
  **secretValue** | [**SecretValue**](SecretValue.md)| secretValue | 

### Return type

[**SecretValue**](SecretValue.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSecretValueUsingDELETE**
> SecretValue DeleteSecretValueUsingDELETE(ctx, name, namespace)
Delete a secret value

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name | 
  **namespace** | **string**| namespace | 

### Return type

[**SecretValue**](SecretValue.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSecretValueUsingGET**
> SecretValue GetSecretValueUsingGET(ctx, name, namespace)
Get a secret value by name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name | 
  **namespace** | **string**| namespace | 

### Return type

[**SecretValue**](SecretValue.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSecretValuesUsingGET**
> ResourceListOfSecretValue GetSecretValuesUsingGET(ctx, namespace)
List all secrets values

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| namespace | 

### Return type

[**ResourceListOfSecretValue**](ResourceListOfSecretValue.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSecretValueUsingPATCH**
> SecretValue UpdateSecretValueUsingPATCH(ctx, body, name, namespace)
Update a secret value

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SecretValue**](SecretValue.md)|  | 
  **name** | **string**| name | 
  **namespace** | **string**| namespace | 

### Return type

[**SecretValue**](SecretValue.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

