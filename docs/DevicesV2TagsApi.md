# \DevicesV2TagsApi

All URIs are relative to *https://api2.arduino.cc/iot*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevicesV2TagsDelete**](DevicesV2TagsApi.md#DevicesV2TagsDelete) | **Delete** /v2/devices/{id}/tags/{key} | delete devices_v2_tags
[**DevicesV2TagsList**](DevicesV2TagsApi.md#DevicesV2TagsList) | **Get** /v2/devices/{id}/tags | list devices_v2_tags
[**DevicesV2TagsUpsert**](DevicesV2TagsApi.md#DevicesV2TagsUpsert) | **Put** /v2/devices/{id}/tags | upsert devices_v2_tags



## DevicesV2TagsDelete

> DevicesV2TagsDelete(ctx, id, key)

delete devices_v2_tags

Delete a tag associated to the device given its key.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the device | 
**key** | **string**| The key of the tag | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2TagsList

> ArduinoTags DevicesV2TagsList(ctx, id)

list devices_v2_tags

List tags associated to the device.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the device | 

### Return type

[**ArduinoTags**](ArduinoTags.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2TagsUpsert

> DevicesV2TagsUpsert(ctx, id, tag)

upsert devices_v2_tags

Creates or updates a tag associated to the device.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the device | 
**tag** | [**Tag**](Tag.md)|  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

