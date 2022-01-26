# \ApiTokensApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiTokenUsingPOST**](ApiTokensApi.md#CreateApiTokenUsingPOST) | **Post** /apitokens/v1/namespaces/{ns}/apitokens | createApiToken
[**DeleteApiTokenUsingDELETE**](ApiTokensApi.md#DeleteApiTokenUsingDELETE) | **Delete** /apitokens/v1/namespaces/{ns}/apitokens/{apiTokenName} | deleteApiToken
[**GetApiTokenUsingGET**](ApiTokensApi.md#GetApiTokenUsingGET) | **Get** /apitokens/v1/namespaces/{ns}/apitokens/{apiTokenName} | getApiToken
[**ListApiTokensUsingGET**](ApiTokensApi.md#ListApiTokensUsingGET) | **Get** /apitokens/v1/namespaces/{ns}/apitokens | listApiTokens


# **CreateApiTokenUsingPOST**
> CreateApiTokenResponse CreateApiTokenUsingPOST(ctx, apiToken, ns)
createApiToken

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiToken** | [**ApiToken**](ApiToken.md)| apiToken | 
  **ns** | **string**| ns | 

### Return type

[**CreateApiTokenResponse**](CreateApiTokenResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiTokenUsingDELETE**
> DeleteApiTokenResponse DeleteApiTokenUsingDELETE(ctx, apiTokenName, ns)
deleteApiToken

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiTokenName** | **string**| apiTokenName | 
  **ns** | **string**| ns | 

### Return type

[**DeleteApiTokenResponse**](DeleteApiTokenResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiTokenUsingGET**
> GetApiTokenResponse GetApiTokenUsingGET(ctx, apiTokenName, ns)
getApiToken

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiTokenName** | **string**| apiTokenName | 
  **ns** | **string**| ns | 

### Return type

[**GetApiTokenResponse**](GetApiTokenResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListApiTokensUsingGET**
> ListApiTokensResponse ListApiTokensUsingGET(ctx, ns)
listApiTokens

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ns** | **string**| ns | 

### Return type

[**ListApiTokensResponse**](ListApiTokensResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

