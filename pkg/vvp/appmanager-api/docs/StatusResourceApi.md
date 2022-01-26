# \StatusResourceApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStatusUsingGET**](StatusResourceApi.md#GetStatusUsingGET) | **Get** /api/v1/status | Check that the server is running
[**GetSystemInfoUsingGET**](StatusResourceApi.md#GetSystemInfoUsingGET) | **Get** /api/v1/status/system-info | Get system&#39;s information


# **GetStatusUsingGET**
> interface{} GetStatusUsingGET(ctx, )
Check that the server is running

### Required Parameters
This endpoint does not need any parameter.

### Return type

**interface{}**

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSystemInfoUsingGET**
> SystemInformation GetSystemInfoUsingGET(ctx, )
Get system's information

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SystemInformation**](SystemInformation.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

