# \DevicesV2PassApi

All URIs are relative to *https://api2.arduino.cc/iot*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevicesV2PassCheck**](DevicesV2PassApi.md#DevicesV2PassCheck) | **Post** /v2/devices/{id}/pass | check devices_v2_pass
[**DevicesV2PassDelete**](DevicesV2PassApi.md#DevicesV2PassDelete) | **Delete** /v2/devices/{id}/pass | delete devices_v2_pass
[**DevicesV2PassGet**](DevicesV2PassApi.md#DevicesV2PassGet) | **Get** /v2/devices/{id}/pass | get devices_v2_pass
[**DevicesV2PassSet**](DevicesV2PassApi.md#DevicesV2PassSet) | **Put** /v2/devices/{id}/pass | set devices_v2_pass



## DevicesV2PassCheck

> DevicesV2PassCheck(ctx, id, checkDevicesV2PassPayload)

check devices_v2_pass

Check if the password matches.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the device | 
**checkDevicesV2PassPayload** | [**CheckDevicesV2PassPayload**](CheckDevicesV2PassPayload.md)|  | 

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


## DevicesV2PassDelete

> DevicesV2PassDelete(ctx, id)

delete devices_v2_pass

Removes the password for the device.

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2PassGet

> ArduinoDevicev2Pass DevicesV2PassGet(ctx, id, optional)

get devices_v2_pass

Returns whether the password for this device is set or not. It doesn't return the password.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the device | 
 **optional** | ***DevicesV2PassGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DevicesV2PassGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **suggestedPassword** | **optional.Bool**| If true, return a suggested password | [default to false]

### Return type

[**ArduinoDevicev2Pass**](ArduinoDevicev2Pass.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2PassSet

> ArduinoDevicev2Pass DevicesV2PassSet(ctx, id, devicev2Pass)

set devices_v2_pass

Sets the password for the device. It can never be read back.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the device | 
**devicev2Pass** | [**Devicev2Pass**](Devicev2Pass.md)|  | 

### Return type

[**ArduinoDevicev2Pass**](ArduinoDevicev2Pass.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

