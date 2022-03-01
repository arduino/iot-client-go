# \DevicesV2OtaApi

All URIs are relative to *https://api2.arduino.cc/iot*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevicesV2OtaSend**](DevicesV2OtaApi.md#DevicesV2OtaSend) | **Put** /v2/devices/{id}/ota | send devices_v2_ota
[**DevicesV2OtaUpload**](DevicesV2OtaApi.md#DevicesV2OtaUpload) | **Post** /v2/devices/{id}/ota | upload devices_v2_ota



## DevicesV2OtaSend

> DevicesV2OtaSend(ctx, id, devicev2Otabinaryurl)

send devices_v2_ota

Send a binary url to a device

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the device | 
**devicev2Otabinaryurl** | [**Devicev2Otabinaryurl**](Devicev2Otabinaryurl.md)|  | 

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


## DevicesV2OtaUpload

> DevicesV2OtaUpload(ctx, id, otaFile, optional)

upload devices_v2_ota

Upload a binary and send it to a device

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the device | 
**otaFile** | ***os.File*****os.File**| OTA file | 
 **optional** | ***DevicesV2OtaUploadOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DevicesV2OtaUploadOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **async** | **optional.Bool**| If false, wait for the full OTA process, until it gets a result from the device | [default to true]
 **expireInMins** | **optional.Int32**| Binary expire time in minutes, default 10 mins | [default to 10]

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

