# \DevicesV2WebhooksApi

All URIs are relative to *http://api-dev.arduino.cc*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevicesV2WebhooksCreate**](DevicesV2WebhooksApi.md#DevicesV2WebhooksCreate) | **Put** /iot/v2/devices/{id}/webhooks | create devices_v2_webhooks
[**DevicesV2WebhooksDelete**](DevicesV2WebhooksApi.md#DevicesV2WebhooksDelete) | **Delete** /iot/v2/devices/{id}/webhooks/{wid} | delete devices_v2_webhooks
[**DevicesV2WebhooksTest**](DevicesV2WebhooksApi.md#DevicesV2WebhooksTest) | **Post** /iot/v2/devices/{id}/webhooks/test | test devices_v2_webhooks



## DevicesV2WebhooksCreate

> ArduinoDevicev2Webhook DevicesV2WebhooksCreate(ctx, id, arduinoDevicev2Webhook)
create devices_v2_webhooks

Creates a new webhook associated to a device

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the device | 
**arduinoDevicev2Webhook** | [**ArduinoDevicev2Webhook**](ArduinoDevicev2Webhook.md)|  | 

### Return type

[**ArduinoDevicev2Webhook**](ArduinoDevicev2Webhook.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.arduino.devicev2.webhook+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2WebhooksDelete

> DevicesV2WebhooksDelete(ctx, id, wid)
delete devices_v2_webhooks

Removes a webhook associated to a device

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the device | 
**wid** | **string**| The id of the webhook | 

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


## DevicesV2WebhooksTest

> DevicesV2WebhooksTest(ctx, id, arduinoDevicev2Webhook, optional)
test devices_v2_webhooks

Tests a webhook

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the device | 
**arduinoDevicev2Webhook** | [**ArduinoDevicev2Webhook**](ArduinoDevicev2Webhook.md)|  | 
 **optional** | ***DevicesV2WebhooksTestOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DevicesV2WebhooksTestOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wid** | **optional.String**| The id of the webhook | 

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

