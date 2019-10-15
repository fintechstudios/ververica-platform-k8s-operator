# \StatusApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStatus**](StatusApi.md#GetStatus) | **Get** /v1/status | Check that the server is running
[**GetSystemInfo**](StatusApi.md#GetSystemInfo) | **Get** /v1/status/system-info | Get system&#39;s information


# **GetStatus**
> GetStatus(ctx, )
Check that the server is running



### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSystemInfo**
> SystemInformation GetSystemInfo(ctx, )
Get system's information



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SystemInformation**](SystemInformation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

