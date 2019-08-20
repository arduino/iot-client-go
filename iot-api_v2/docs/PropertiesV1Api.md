# \PropertiesV1Api

All URIs are relative to *http://api-dev.arduino.cc*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PropertiesV1Create**](PropertiesV1Api.md#PropertiesV1Create) | **Put** /iot/v1/things/{id}/properties | create properties_v1
[**PropertiesV1Delete**](PropertiesV1Api.md#PropertiesV1Delete) | **Delete** /iot/v1/things/{id}/properties/{pid} | delete properties_v1
[**PropertiesV1List**](PropertiesV1Api.md#PropertiesV1List) | **Get** /iot/v1/things/{id}/properties | list properties_v1
[**PropertiesV1Publish**](PropertiesV1Api.md#PropertiesV1Publish) | **Put** /iot/v1/things/{id}/properties/{pid}/publish | publish properties_v1
[**PropertiesV1Show**](PropertiesV1Api.md#PropertiesV1Show) | **Get** /iot/v1/things/{id}/properties/{pid} | show properties_v1
[**PropertiesV1Update**](PropertiesV1Api.md#PropertiesV1Update) | **Post** /iot/v1/things/{id}/properties/{pid} | update properties_v1



## PropertiesV1Create

> ArduinoProperty PropertiesV1Create(ctx, id, arduinoProperty)
create properties_v1

Creates a new property associated to a thing

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the thing | 
**arduinoProperty** | [**ArduinoProperty**](ArduinoProperty.md)| PropertyPayload describes a property of a thing. No field is mandatory | 

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


## PropertiesV1Delete

> PropertiesV1Delete(ctx, id, pid, optional)
delete properties_v1

Removes a property associated to a thing

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the thing | 
**pid** | **string**| The id of the property | 
 **optional** | ***PropertiesV1DeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PropertiesV1DeleteOpts struct


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


## PropertiesV1List

> []ArduinoProperty PropertiesV1List(ctx, id, optional)
list properties_v1

Returns the list of properties associated to the thing

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the thing | 
 **optional** | ***PropertiesV1ListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PropertiesV1ListOpts struct


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


## PropertiesV1Publish

> PropertiesV1Publish(ctx, id, pid, propertyValue)
publish properties_v1

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


## PropertiesV1Show

> ArduinoProperty PropertiesV1Show(ctx, id, pid, optional)
show properties_v1

Returns the property requested by the user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the thing | 
**pid** | **string**| The id of the property | 
 **optional** | ***PropertiesV1ShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PropertiesV1ShowOpts struct


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


## PropertiesV1Update

> ArduinoProperty PropertiesV1Update(ctx, id, pid, arduinoProperty)
update properties_v1

Updates a property associated to a thing

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the thing | 
**pid** | **string**| The id of the property | 
**arduinoProperty** | [**ArduinoProperty**](ArduinoProperty.md)| PropertyPayload describes a property of a thing. No field is mandatory | 

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

