# \PropertiesV2Api

All URIs are relative to *http://api2.arduino.cc/iot*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PropertiesV2Create**](PropertiesV2Api.md#PropertiesV2Create) | **Put** /v2/things/{id}/properties | create properties_v2
[**PropertiesV2Delete**](PropertiesV2Api.md#PropertiesV2Delete) | **Delete** /v2/things/{id}/properties/{pid} | delete properties_v2
[**PropertiesV2List**](PropertiesV2Api.md#PropertiesV2List) | **Get** /v2/things/{id}/properties | list properties_v2
[**PropertiesV2Publish**](PropertiesV2Api.md#PropertiesV2Publish) | **Put** /v2/things/{id}/properties/{pid}/publish | publish properties_v2
[**PropertiesV2Show**](PropertiesV2Api.md#PropertiesV2Show) | **Get** /v2/things/{id}/properties/{pid} | show properties_v2
[**PropertiesV2Update**](PropertiesV2Api.md#PropertiesV2Update) | **Post** /v2/things/{id}/properties/{pid} | update properties_v2



## PropertiesV2Create

> ArduinoProperty PropertiesV2Create(ctx, id, property)

create properties_v2

Creates a new property associated to a thing

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the thing | 
**property** | [**Property**](Property.md)| PropertyPayload describes a property of a thing. No field is mandatory | 

### Return type

[**ArduinoProperty**](ArduinoProperty.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.arduino.property+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PropertiesV2Delete

> PropertiesV2Delete(ctx, id, pid, optional)

delete properties_v2

Removes a property associated to a thing

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the thing | 
**pid** | **string**| The id of the property | 
 **optional** | ***PropertiesV2DeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PropertiesV2DeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **optional.Bool**| If true, hard delete the property | [default to false]

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.goa.error+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PropertiesV2List

> []ArduinoProperty PropertiesV2List(ctx, id, optional)

list properties_v2

Returns the list of properties associated to the thing

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the thing | 
 **optional** | ***PropertiesV2ListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PropertiesV2ListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **showDeleted** | **optional.Bool**| If true, shows the soft deleted properties | [default to false]

### Return type

[**[]ArduinoProperty**](ArduinoProperty.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.property+json; type=collection, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PropertiesV2Publish

> PropertiesV2Publish(ctx, id, pid, propertyValue)

publish properties_v2

Publish a property value to MQTT

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the thing | 
**pid** | **string**| The id of the property | 
**propertyValue** | [**PropertyValue**](PropertyValue.md)| PropertyValuePayload describes a property value | 

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


## PropertiesV2Show

> ArduinoProperty PropertiesV2Show(ctx, id, pid, optional)

show properties_v2

Returns the property requested by the user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the thing | 
**pid** | **string**| The id of the property | 
 **optional** | ***PropertiesV2ShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PropertiesV2ShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **showDeleted** | **optional.Bool**| If true, shows the soft deleted properties | [default to false]

### Return type

[**ArduinoProperty**](ArduinoProperty.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.property+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PropertiesV2Update

> ArduinoProperty PropertiesV2Update(ctx, id, pid, property)

update properties_v2

Updates a property associated to a thing

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the thing | 
**pid** | **string**| The id of the property | 
**property** | [**Property**](Property.md)| PropertyPayload describes a property of a thing. No field is mandatory | 

### Return type

[**ArduinoProperty**](ArduinoProperty.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.arduino.property+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

