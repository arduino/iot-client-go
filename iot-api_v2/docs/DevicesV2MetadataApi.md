# \DevicesV2MetadataApi

All URIs are relative to *http://api-dev.arduino.cc*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevicesV2MetadataDelete**](DevicesV2MetadataApi.md#DevicesV2MetadataDelete) | **Delete** /iot/v2/devices/{id}/metadata/{key} | delete devices_v2_metadata
[**DevicesV2MetadataList**](DevicesV2MetadataApi.md#DevicesV2MetadataList) | **Get** /iot/v2/devices/{id}/metadata | list devices_v2_metadata
[**DevicesV2MetadataSet**](DevicesV2MetadataApi.md#DevicesV2MetadataSet) | **Put** /iot/v2/devices/{id}/metadata | set devices_v2_metadata



## DevicesV2MetadataDelete

> DevicesV2MetadataDelete(ctx, id, key)
delete devices_v2_metadata

Removes the a metadata key of the device.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the device | 
**key** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2MetadataList

> DevicesV2MetadataList(ctx, id)
list devices_v2_metadata

Returns metadata of a device

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the device | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2MetadataSet

> DevicesV2MetadataSet(ctx, id, requestBody)
set devices_v2_metadata

Sets metadata of a device

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the device | 
**requestBody** | [**map[string]map[string]interface{}**](map[string]interface{}.md)|  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

