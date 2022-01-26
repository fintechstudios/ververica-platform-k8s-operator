# \ConnectorApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnectorUsingPOST**](ConnectorApi.md#CreateConnectorUsingPOST) | **Post** /sql/v1beta1/namespaces/{ns}/connectors | createConnector
[**CreateFormatUsingPOST**](ConnectorApi.md#CreateFormatUsingPOST) | **Post** /sql/v1beta1/namespaces/{ns}/formats | createFormat
[**DeleteConnectorUsingDELETE**](ConnectorApi.md#DeleteConnectorUsingDELETE) | **Delete** /sql/v1beta1/namespaces/{ns}/connectors/{name} | deleteConnector
[**DeleteFormatUsingDELETE**](ConnectorApi.md#DeleteFormatUsingDELETE) | **Delete** /sql/v1beta1/namespaces/{ns}/formats/{name} | deleteFormat
[**GetConnectorUsingGET**](ConnectorApi.md#GetConnectorUsingGET) | **Get** /sql/v1beta1/namespaces/{ns}/connectors/{name} | getConnector
[**GetFormatUsingGET**](ConnectorApi.md#GetFormatUsingGET) | **Get** /sql/v1beta1/namespaces/{ns}/formats/{name} | getFormat
[**ListConnectorsUsingGET**](ConnectorApi.md#ListConnectorsUsingGET) | **Get** /sql/v1beta1/namespaces/{ns}/connectors | listConnectors
[**ListFormatsUsingGET**](ConnectorApi.md#ListFormatsUsingGET) | **Get** /sql/v1beta1/namespaces/{ns}/formats | listFormats
[**ListTablesReferencingConnectorUsingGET**](ConnectorApi.md#ListTablesReferencingConnectorUsingGET) | **Get** /sql/v1beta1/namespaces/{ns}/connectors/{name}:list-tables | listTablesReferencingConnector
[**ListTablesReferencingFormatUsingGET**](ConnectorApi.md#ListTablesReferencingFormatUsingGET) | **Get** /sql/v1beta1/namespaces/{ns}/formats/{name}:list-tables | listTablesReferencingFormat
[**UpdateConnectorUsingPUT**](ConnectorApi.md#UpdateConnectorUsingPUT) | **Put** /sql/v1beta1/namespaces/{ns}/connectors/{name} | updateConnector
[**UpdateFormatUsingPUT**](ConnectorApi.md#UpdateFormatUsingPUT) | **Put** /sql/v1beta1/namespaces/{ns}/formats/{name} | updateFormat


# **CreateConnectorUsingPOST**
> CreateConnectorResponse CreateConnectorUsingPOST(ctx, connector, ns)
createConnector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connector** | [**Connector**](Connector.md)| connector | 
  **ns** | **string**| ns | 

### Return type

[**CreateConnectorResponse**](CreateConnectorResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFormatUsingPOST**
> CreateFormatResponse CreateFormatUsingPOST(ctx, format, ns)
createFormat

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **format** | [**Format**](Format.md)| format | 
  **ns** | **string**| ns | 

### Return type

[**CreateFormatResponse**](CreateFormatResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteConnectorUsingDELETE**
> DeleteConnectorResponse DeleteConnectorUsingDELETE(ctx, name, ns)
deleteConnector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name | 
  **ns** | **string**| ns | 

### Return type

[**DeleteConnectorResponse**](DeleteConnectorResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFormatUsingDELETE**
> DeleteFormatResponse DeleteFormatUsingDELETE(ctx, name, ns)
deleteFormat

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name | 
  **ns** | **string**| ns | 

### Return type

[**DeleteFormatResponse**](DeleteFormatResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectorUsingGET**
> GetConnectorResponse GetConnectorUsingGET(ctx, connectorResourceId, ns)
getConnector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectorResourceId** | **string**| connectorResourceId | 
  **ns** | **string**| ns | 

### Return type

[**GetConnectorResponse**](GetConnectorResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFormatUsingGET**
> GetFormatResponse GetFormatUsingGET(ctx, formatResourceId, ns)
getFormat

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **formatResourceId** | **string**| formatResourceId | 
  **ns** | **string**| ns | 

### Return type

[**GetFormatResponse**](GetFormatResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListConnectorsUsingGET**
> ListConnectorsResponse ListConnectorsUsingGET(ctx, ns)
listConnectors

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ns** | **string**| ns | 

### Return type

[**ListConnectorsResponse**](ListConnectorsResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFormatsUsingGET**
> ListFormatsResponse ListFormatsUsingGET(ctx, ns)
listFormats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ns** | **string**| ns | 

### Return type

[**ListFormatsResponse**](ListFormatsResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTablesReferencingConnectorUsingGET**
> ListTablesReferencingConnectorResponse ListTablesReferencingConnectorUsingGET(ctx, name, ns)
listTablesReferencingConnector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name | 
  **ns** | **string**| ns | 

### Return type

[**ListTablesReferencingConnectorResponse**](ListTablesReferencingConnectorResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTablesReferencingFormatUsingGET**
> ListTablesReferencingFormatResponse ListTablesReferencingFormatUsingGET(ctx, name, ns)
listTablesReferencingFormat

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name | 
  **ns** | **string**| ns | 

### Return type

[**ListTablesReferencingFormatResponse**](ListTablesReferencingFormatResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateConnectorUsingPUT**
> UpdateConnectorResponse UpdateConnectorUsingPUT(ctx, connector, name, ns)
updateConnector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connector** | [**Connector**](Connector.md)| connector | 
  **name** | **string**| name | 
  **ns** | **string**| ns | 

### Return type

[**UpdateConnectorResponse**](UpdateConnectorResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFormatUsingPUT**
> UpdateFormatResponse UpdateFormatUsingPUT(ctx, format, name, ns)
updateFormat

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **format** | [**Format**](Format.md)| format | 
  **name** | **string**| name | 
  **ns** | **string**| ns | 

### Return type

[**UpdateFormatResponse**](UpdateFormatResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

