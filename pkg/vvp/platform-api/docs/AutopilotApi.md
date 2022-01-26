# \AutopilotApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAutopilotPolicyRecommendationUsingGET**](AutopilotApi.md#GetAutopilotPolicyRecommendationUsingGET) | **Get** /autopilot/v1alpha1/namespaces/{ns}/deployments/{deploymentId}/autopilotpolicy:recommendation | getAutopilotPolicyRecommendation
[**GetAutopilotPolicyStatusUsingGET**](AutopilotApi.md#GetAutopilotPolicyStatusUsingGET) | **Get** /autopilot/v1alpha1/namespaces/{ns}/deployments/{deploymentId}/autopilotpolicy:status | getAutopilotPolicyStatus
[**GetAutopilotPolicyUsingGET**](AutopilotApi.md#GetAutopilotPolicyUsingGET) | **Get** /autopilot/v1alpha1/namespaces/{ns}/deployments/{deploymentId}/autopilotpolicy | getAutopilotPolicy
[**ListAutopilotPoliciesUsingGET**](AutopilotApi.md#ListAutopilotPoliciesUsingGET) | **Get** /autopilot/v1alpha1/namespaces/{ns} | listAutopilotPolicies
[**UpdateAutopilotPolicyUsingPUT**](AutopilotApi.md#UpdateAutopilotPolicyUsingPUT) | **Put** /autopilot/v1alpha1/namespaces/{ns}/deployments/{deploymentId}/autopilotpolicy | updateAutopilotPolicy


# **GetAutopilotPolicyRecommendationUsingGET**
> GetAutopilotPolicyRecommendationResponse GetAutopilotPolicyRecommendationUsingGET(ctx, deploymentId, ns)
getAutopilotPolicyRecommendation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deploymentId** | **string**| deploymentId | 
  **ns** | **string**| ns | 

### Return type

[**GetAutopilotPolicyRecommendationResponse**](GetAutopilotPolicyRecommendationResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAutopilotPolicyStatusUsingGET**
> GetAutopilotPolicyStatusResponse GetAutopilotPolicyStatusUsingGET(ctx, deploymentId, ns)
getAutopilotPolicyStatus

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deploymentId** | **string**| deploymentId | 
  **ns** | **string**| ns | 

### Return type

[**GetAutopilotPolicyStatusResponse**](GetAutopilotPolicyStatusResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAutopilotPolicyUsingGET**
> GetAutopilotPolicyResponse GetAutopilotPolicyUsingGET(ctx, deploymentId, ns)
getAutopilotPolicy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deploymentId** | **string**| deploymentId | 
  **ns** | **string**| ns | 

### Return type

[**GetAutopilotPolicyResponse**](GetAutopilotPolicyResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAutopilotPoliciesUsingGET**
> ListAutopilotPoliciesResponse ListAutopilotPoliciesUsingGET(ctx, ns)
listAutopilotPolicies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ns** | **string**| ns | 

### Return type

[**ListAutopilotPoliciesResponse**](ListAutopilotPoliciesResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAutopilotPolicyUsingPUT**
> UpdateAutopilotPolicyResponse UpdateAutopilotPolicyUsingPUT(ctx, deploymentId, ns, policy)
updateAutopilotPolicy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deploymentId** | **string**| deploymentId | 
  **ns** | **string**| ns | 
  **policy** | [**AutopilotPolicy**](AutopilotPolicy.md)| policy | 

### Return type

[**UpdateAutopilotPolicyResponse**](UpdateAutopilotPolicyResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

