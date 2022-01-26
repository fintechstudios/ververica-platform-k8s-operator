# \CatalogApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCatalogUsingGET**](CatalogApi.md#GetCatalogUsingGET) | **Get** /catalog/v1beta2/namespaces/{ns}/catalogs/{cat} | getCatalog
[**GetDatabaseUsingGET**](CatalogApi.md#GetDatabaseUsingGET) | **Get** /catalog/v1beta2/namespaces/{ns}/catalogs/{cat}:getDatabase | getDatabase
[**GetFunctionUsingGET**](CatalogApi.md#GetFunctionUsingGET) | **Get** /catalog/v1beta2/namespaces/{ns}/catalogs/{cat}:getFunction | getFunction
[**GetTableUsingGET**](CatalogApi.md#GetTableUsingGET) | **Get** /catalog/v1beta2/namespaces/{ns}/catalogs/{cat}:getTable | getTable
[**ListCatalogsUsingGET**](CatalogApi.md#ListCatalogsUsingGET) | **Get** /catalog/v1beta2/namespaces/{ns}/catalogs | listCatalogs
[**ListDatabasesUsingGET**](CatalogApi.md#ListDatabasesUsingGET) | **Get** /catalog/v1beta2/namespaces/{ns}/catalogs/{cat}:listDatabases | listDatabases
[**ListFunctionsUsingGET**](CatalogApi.md#ListFunctionsUsingGET) | **Get** /catalog/v1beta2/namespaces/{ns}/catalogs/{cat}:listFunctions | listFunctions
[**ListTablesUsingGET**](CatalogApi.md#ListTablesUsingGET) | **Get** /catalog/v1beta2/namespaces/{ns}/catalogs/{cat}:listTables | listTables


# **GetCatalogUsingGET**
> GetCatalogResponse GetCatalogUsingGET(ctx, cat, ns)
getCatalog

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cat** | **string**| cat | 
  **ns** | **string**| ns | 

### Return type

[**GetCatalogResponse**](GetCatalogResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDatabaseUsingGET**
> GetDatabaseResponse GetDatabaseUsingGET(ctx, cat, database, ns)
getDatabase

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cat** | **string**| cat | 
  **database** | **string**| database | 
  **ns** | **string**| ns | 

### Return type

[**GetDatabaseResponse**](GetDatabaseResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFunctionUsingGET**
> GetFunctionResponse GetFunctionUsingGET(ctx, cat, function, ns, optional)
getFunction

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cat** | **string**| cat | 
  **function** | **string**| function | 
  **ns** | **string**| ns | 
 **optional** | ***CatalogApiGetFunctionUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CatalogApiGetFunctionUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **database** | **optional.String**| database | 

### Return type

[**GetFunctionResponse**](GetFunctionResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTableUsingGET**
> GetTableResponse GetTableUsingGET(ctx, cat, ns, table, optional)
getTable

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cat** | **string**| cat | 
  **ns** | **string**| ns | 
  **table** | **string**| table | 
 **optional** | ***CatalogApiGetTableUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CatalogApiGetTableUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **database** | **optional.String**| database | 

### Return type

[**GetTableResponse**](GetTableResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCatalogsUsingGET**
> ListCatalogsResponse ListCatalogsUsingGET(ctx, ns)
listCatalogs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ns** | **string**| ns | 

### Return type

[**ListCatalogsResponse**](ListCatalogsResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDatabasesUsingGET**
> ListDatabasesResponse ListDatabasesUsingGET(ctx, cat, ns)
listDatabases

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cat** | **string**| cat | 
  **ns** | **string**| ns | 

### Return type

[**ListDatabasesResponse**](ListDatabasesResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFunctionsUsingGET**
> ListFunctionsResponse ListFunctionsUsingGET(ctx, cat, ns, optional)
listFunctions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cat** | **string**| cat | 
  **ns** | **string**| ns | 
 **optional** | ***CatalogApiListFunctionsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CatalogApiListFunctionsUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **database** | **optional.String**| database | 

### Return type

[**ListFunctionsResponse**](ListFunctionsResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTablesUsingGET**
> ListTablesResponse ListTablesUsingGET(ctx, cat, ns, optional)
listTables

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cat** | **string**| cat | 
  **ns** | **string**| ns | 
 **optional** | ***CatalogApiListTablesUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CatalogApiListTablesUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **database** | **optional.String**| database | 

### Return type

[**ListTablesResponse**](ListTablesResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

