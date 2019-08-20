# \DevicesV1Api

All URIs are relative to *http://api-dev.arduino.cc*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevicesV1Connect**](DevicesV1Api.md#DevicesV1Connect) | **Post** /iot/v1/devices/connect | connect devices_v1
[**DevicesV1Create**](DevicesV1Api.md#DevicesV1Create) | **Put** /iot/v1/devices | create devices_v1
[**DevicesV1Delete**](DevicesV1Api.md#DevicesV1Delete) | **Delete** /iot/v1/devices/{id} | delete devices_v1
[**DevicesV1Image**](DevicesV1Api.md#DevicesV1Image) | **Post** /iot/v1/devices/{id}/image | image devices_v1
[**DevicesV1List**](DevicesV1Api.md#DevicesV1List) | **Get** /iot/v1/devices | list devices_v1
[**DevicesV1Show**](DevicesV1Api.md#DevicesV1Show) | **Get** /iot/v1/devices/{id} | show devices_v1
[**DevicesV1Update**](DevicesV1Api.md#DevicesV1Update) | **Post** /iot/v1/devices/{id} | update devices_v1



## DevicesV1Connect

> ArduinoMqttv1 DevicesV1Connect(ctx, )
connect devices_v1

Returns a signed websocket url to use for a connection

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ArduinoMqttv1**](ArduinoMqttv1.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.mqttv1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV1Create

> ArduinoDevicev1 DevicesV1Create(ctx, arduinoDevicev1)
create devices_v1

Creates a new device associated to the user. By default it creates a device with no name and no certificates attached. If you provide a name or a certificate request it will set them up.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**arduinoDevicev1** | [**ArduinoDevicev1**](ArduinoDevicev1.md)| DevicePayload describes a device. No field is mandatory | 

### Return type

[**ArduinoDevicev1**](ArduinoDevicev1.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.arduino.devicev1+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV1Delete

> DevicesV1Delete(ctx, id)
delete devices_v1

Removes a device associated to the user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the resource | 

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


## DevicesV1Image

> DevicesV1Image(ctx, id, configv1)
image devices_v1

Returns an iso image to be burned into a usb key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the resource | 
**configv1** | [**[]Configv1**](configv1.md)|  | 

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


## DevicesV1List

> []ArduinoDevicev1 DevicesV1List(ctx, )
list devices_v1

Returns the list of devices associated to the user

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]ArduinoDevicev1**](ArduinoDevicev1.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.devicev1+json; type=collection

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV1Show

> ArduinoDevicev1 DevicesV1Show(ctx, id)
show devices_v1

Returns the device requested by the user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the resource | 

### Return type

[**ArduinoDevicev1**](ArduinoDevicev1.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.devicev1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV1Update

> ArduinoDevicev1 DevicesV1Update(ctx, id, arduinoDevicev1)
update devices_v1

Updates a device associated to the user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the resource | 
**arduinoDevicev1** | [**ArduinoDevicev1**](ArduinoDevicev1.md)| DevicePayload describes a device. No field is mandatory | 

### Return type

[**ArduinoDevicev1**](ArduinoDevicev1.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.arduino.devicev1+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

