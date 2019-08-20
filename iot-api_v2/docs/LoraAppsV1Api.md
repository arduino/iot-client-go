# \LoraAppsV1Api

All URIs are relative to *http://api-dev.arduino.cc*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LoraAppsV1List**](LoraAppsV1Api.md#LoraAppsV1List) | **Get** /iot/v1/lora-apps/ | list lora_apps_v1
[**LoraAppsV1Update**](LoraAppsV1Api.md#LoraAppsV1Update) | **Post** /iot/v1/lora-apps/ | update lora_apps_v1



## LoraAppsV1List

> LoraAppsV1List(ctx, userId)
list lora_apps_v1

List the available apps for the given user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| The user id. Could be a uuid, or the string &#39;me&#39; | 

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


## LoraAppsV1Update

> LoraAppsV1Update(ctx, userId, updateLoraAppsV1Payload)
update lora_apps_v1

Update the available apps for the given user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| The user id. Could be a uuid, or the string &#39;me&#39; | 
**updateLoraAppsV1Payload** | [**UpdateLoraAppsV1Payload**](UpdateLoraAppsV1Payload.md)|  | 

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

