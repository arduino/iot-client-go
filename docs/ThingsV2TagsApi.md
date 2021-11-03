# \ThingsV2TagsApi

All URIs are relative to *https://api2.arduino.cc/iot*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ThingsV2TagsDelete**](ThingsV2TagsApi.md#ThingsV2TagsDelete) | **Delete** /v2/things/{id}/tags/{key} | delete things_v2_tags
[**ThingsV2TagsList**](ThingsV2TagsApi.md#ThingsV2TagsList) | **Get** /v2/things/{id}/tags | list things_v2_tags
[**ThingsV2TagsUpsert**](ThingsV2TagsApi.md#ThingsV2TagsUpsert) | **Put** /v2/things/{id}/tags | upsert things_v2_tags



## ThingsV2TagsDelete

> ThingsV2TagsDelete(ctx, id, key)

delete things_v2_tags

Delete a tag associated to the thing given its key.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the thing | 
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


## ThingsV2TagsList

> ArduinoTags ThingsV2TagsList(ctx, id)

list things_v2_tags

List tags associated to the thing.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the thing | 

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


## ThingsV2TagsUpsert

> ThingsV2TagsUpsert(ctx, id, tag)

upsert things_v2_tags

Creates or updates a tag associated to the thing.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the thing | 
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

