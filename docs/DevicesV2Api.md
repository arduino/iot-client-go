# \DevicesV2Api

All URIs are relative to *https://api2.arduino.cc/iot*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevicesV2Create**](DevicesV2Api.md#DevicesV2Create) | **Put** /v2/devices | create devices_v2
[**DevicesV2Delete**](DevicesV2Api.md#DevicesV2Delete) | **Delete** /v2/devices/{id} | delete devices_v2
[**DevicesV2GetEvents**](DevicesV2Api.md#DevicesV2GetEvents) | **Get** /v2/devices/{id}/events | getEvents devices_v2
[**DevicesV2GetProperties**](DevicesV2Api.md#DevicesV2GetProperties) | **Get** /v2/devices/{id}/properties | getProperties devices_v2
[**DevicesV2List**](DevicesV2Api.md#DevicesV2List) | **Get** /v2/devices | list devices_v2
[**DevicesV2Show**](DevicesV2Api.md#DevicesV2Show) | **Get** /v2/devices/{id} | show devices_v2
[**DevicesV2Timeseries**](DevicesV2Api.md#DevicesV2Timeseries) | **Get** /v2/devices/{id}/properties/{pid} | timeseries devices_v2
[**DevicesV2Update**](DevicesV2Api.md#DevicesV2Update) | **Post** /v2/devices/{id} | update devices_v2
[**DevicesV2UpdateProperties**](DevicesV2Api.md#DevicesV2UpdateProperties) | **Put** /v2/devices/{id}/properties | updateProperties devices_v2



## DevicesV2Create

> ArduinoDevicev2 DevicesV2Create(ctx, createDevicesV2Payload)

create devices_v2

Creates a new device associated to the user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createDevicesV2Payload** | [**CreateDevicesV2Payload**](CreateDevicesV2Payload.md)| DeviceV2 describes a device. | 

### Return type

[**ArduinoDevicev2**](ArduinoDevicev2.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2Delete

> DevicesV2Delete(ctx, id)

delete devices_v2

Removes a device associated to the user

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


## DevicesV2GetEvents

> ArduinoDevicev2EventProperties DevicesV2GetEvents(ctx, idoptional)

getEvents devices_v2

GET device events

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the device | 
 **optional** | ***DevicesV2GetEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DevicesV2GetEventsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.**| The number of events to select | 
 **start** | **optional.**| The time at which to start selecting events | 

### Return type

[**ArduinoDevicev2EventProperties**](ArduinoDevicev2EventProperties.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2GetProperties

> ArduinoDevicev2properties DevicesV2GetProperties(ctx, idoptional)

getProperties devices_v2

GET device properties

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the device | 
 **optional** | ***DevicesV2GetPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DevicesV2GetPropertiesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **showDeleted** | **optional.**| If true, shows the soft deleted properties | 

### Return type

[**ArduinoDevicev2properties**](ArduinoDevicev2properties.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2List

> []ArduinoDevicev2 DevicesV2List(ctx, optional)

list devices_v2

Returns the list of devices associated to the user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DevicesV2ListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DevicesV2ListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acrossUserIds** | **optional.**| If true, returns all the devices | 

### Return type

[**[]ArduinoDevicev2**](ArduinoDevicev2.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2Show

> ArduinoDevicev2 DevicesV2Show(ctx, id)

show devices_v2

Returns the device requested by the user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the device | 

### Return type

[**ArduinoDevicev2**](ArduinoDevicev2.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2Timeseries

> ArduinoDevicev2propertyvalues DevicesV2Timeseries(ctx, idpidoptional)

timeseries devices_v2

GET device properties values in a range of time

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the device | 
**pid** | **string**| The id of the property | 
 **optional** | ***DevicesV2TimeseriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DevicesV2TimeseriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.**| The number of properties to select | 
 **start** | **optional.**| The time at which to start selecting properties | 

### Return type

[**ArduinoDevicev2propertyvalues**](ArduinoDevicev2propertyvalues.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2Update

> ArduinoDevicev2 DevicesV2Update(ctx, iddevicev2)

update devices_v2

Updates a device associated to the user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the device | 
**devicev2** | [**Devicev2**](Devicev2.md)| DeviceV2 describes a device. | 

### Return type

[**ArduinoDevicev2**](ArduinoDevicev2.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2UpdateProperties

> DevicesV2UpdateProperties(ctx, idpropertiesValues)

updateProperties devices_v2

Update device properties last values

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the device | 
**propertiesValues** | [**PropertiesValues**](PropertiesValues.md)|  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

