# \DevicesV2PassApi

All URIs are relative to *http://api-dev.arduino.cc*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevicesV2PassCheck**](DevicesV2PassApi.md#DevicesV2PassCheck) | **Post** /iot/v2/devices/{id}/pass | check devices_v2_pass
[**DevicesV2PassDelete**](DevicesV2PassApi.md#DevicesV2PassDelete) | **Delete** /iot/v2/devices/{id}/pass | delete devices_v2_pass
[**DevicesV2PassGet**](DevicesV2PassApi.md#DevicesV2PassGet) | **Get** /iot/v2/devices/{id}/pass | get devices_v2_pass
[**DevicesV2PassSet**](DevicesV2PassApi.md#DevicesV2PassSet) | **Put** /iot/v2/devices/{id}/pass | set devices_v2_pass



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

- **Content-Type**: application/json
- **Accept**: Not defined

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
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2PassGet

> ArduinoDevicev2Pass DevicesV2PassGet(ctx, id)
get devices_v2_pass

Returns whether the password for this device is set or not. It doesn't return the password.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the device | 

### Return type

[**ArduinoDevicev2Pass**](ArduinoDevicev2Pass.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.devicev2.pass+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2PassSet

> DevicesV2PassSet(ctx, id, setDevicesV2PassPayload)
set devices_v2_pass

Sets the password for the device. It can never be read back.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the device | 
**setDevicesV2PassPayload** | [**SetDevicesV2PassPayload**](SetDevicesV2PassPayload.md)|  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.goa.error+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

