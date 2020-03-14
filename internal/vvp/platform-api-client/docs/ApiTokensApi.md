# \ApiTokensApi

All URIs are relative to *https://localhost:8081*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiToken**](ApiTokensApi.md#CreateApiToken) | **Post** /apitokens/v1/namespaces/{ns}/apitokens | createApiToken
[**DeleteApiToken**](ApiTokensApi.md#DeleteApiToken) | **Delete** /apitokens/v1/namespaces/{ns}/apitokens/{apiTokenName} | deleteApiToken
[**GetApiToken**](ApiTokensApi.md#GetApiToken) | **Get** /apitokens/v1/namespaces/{ns}/apitokens/{apiTokenName} | getApiToken
[**ListApiTokens**](ApiTokensApi.md#ListApiTokens) | **Get** /apitokens/v1/namespaces/{ns}/apitokens | listApiTokens


# **CreateApiToken**
> CreateApiTokenResponse CreateApiToken(ctx, apiToken, ns)
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

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiToken**
> DeleteApiTokenResponse DeleteApiToken(ctx, apiTokenName, ns)
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

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiToken**
> GetApiTokenResponse GetApiToken(ctx, apiTokenName, ns)
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

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListApiTokens**
> ListApiTokensResponse ListApiTokens(ctx, ns)
listApiTokens

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ns** | **string**| ns | 

### Return type

[**ListApiTokensResponse**](ListApiTokensResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

