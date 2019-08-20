# \ThingsV1Api

All URIs are relative to *http://api-dev.arduino.cc*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ThingsV1Create**](ThingsV1Api.md#ThingsV1Create) | **Put** /iot/v1/things | create things_v1
[**ThingsV1CreateSketch**](ThingsV1Api.md#ThingsV1CreateSketch) | **Put** /iot/v1/things/{id}/sketch | createSketch things_v1
[**ThingsV1Delete**](ThingsV1Api.md#ThingsV1Delete) | **Delete** /iot/v1/things/{id} | delete things_v1
[**ThingsV1DeleteSketch**](ThingsV1Api.md#ThingsV1DeleteSketch) | **Delete** /iot/v1/things/{id}/sketch | deleteSketch things_v1
[**ThingsV1Layout**](ThingsV1Api.md#ThingsV1Layout) | **Get** /iot/v1/things/{id}/layout | layout things_v1
[**ThingsV1List**](ThingsV1Api.md#ThingsV1List) | **Get** /iot/v1/things | list things_v1
[**ThingsV1Shadow**](ThingsV1Api.md#ThingsV1Shadow) | **Post** /iot/v1/things/{id}/shadow/send | shadow things_v1
[**ThingsV1Show**](ThingsV1Api.md#ThingsV1Show) | **Get** /iot/v1/things/{id} | show things_v1
[**ThingsV1Update**](ThingsV1Api.md#ThingsV1Update) | **Post** /iot/v1/things/{id} | update things_v1
[**ThingsV1UpdateSketch**](ThingsV1Api.md#ThingsV1UpdateSketch) | **Put** /iot/v1/things/{id}/sketch/{sketchId} | updateSketch things_v1



## ThingsV1Create

> ArduinoThing ThingsV1Create(ctx, createThingsV1Payload, optional)
create things_v1

Creates a new thing associated to the user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createThingsV1Payload** | [**CreateThingsV1Payload**](CreateThingsV1Payload.md)| ThingPayload describes a thing | 
 **optional** | ***ThingsV1CreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsV1CreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| If true, detach device from the other thing, and attach to this thing | [default to false]

### Return type

[**ArduinoThing**](ArduinoThing.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.arduino.thing+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsV1CreateSketch

> ArduinoThing ThingsV1CreateSketch(ctx, id, thingSketch)
createSketch things_v1

Creates a new sketch thing associated to the thing

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the thing | 
**thingSketch** | [**ThingSketch**](ThingSketch.md)| ThingSketchPayload describes a sketch of a thing | 

### Return type

[**ArduinoThing**](ArduinoThing.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.arduino.thing+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsV1Delete

> ThingsV1Delete(ctx, id, optional)
delete things_v1

Removes a thing associated to the user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the thing | 
 **optional** | ***ThingsV1DeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsV1DeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| If true, hard delete the thing | [default to false]

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


## ThingsV1DeleteSketch

> ArduinoThing ThingsV1DeleteSketch(ctx, id)
deleteSketch things_v1

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the thing | 

### Return type

[**ArduinoThing**](ArduinoThing.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.thing+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsV1Layout

> ArduinoThing ThingsV1Layout(ctx, id, optional)
layout things_v1

Returns the thing requested by the user, without last values data

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the thing | 
 **optional** | ***ThingsV1LayoutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsV1LayoutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **showDeleted** | **optional.Bool**| If true, shows the soft deleted thing | [default to false]

### Return type

[**ArduinoThing**](ArduinoThing.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.thing+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsV1List

> []ArduinoThing ThingsV1List(ctx, optional)
list things_v1

Returns the list of things associated to the user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ThingsV1ListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsV1ListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acrossUserIds** | **optional.Bool**| If true, returns all the things | [default to false]
 **deviceId** | **optional.String**| The id of the device you want to filter | 
 **showDeleted** | **optional.Bool**| If true, shows the soft deleted things | [default to false]

### Return type

[**[]ArduinoThing**](ArduinoThing.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.thing+json; type=collection

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsV1Shadow

> ThingsV1Shadow(ctx, id, thingPropertiesValues)
shadow things_v1

Publish the last property values to MQTT

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the thing | 
**thingPropertiesValues** | [**ThingPropertiesValues**](ThingPropertiesValues.md)|  | 

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


## ThingsV1Show

> ArduinoThing ThingsV1Show(ctx, id, optional)
show things_v1

Returns the thing requested by the user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the thing | 
 **optional** | ***ThingsV1ShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsV1ShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **showDeleted** | **optional.Bool**| If true, shows the soft deleted thing | [default to false]

### Return type

[**ArduinoThing**](ArduinoThing.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.thing+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsV1Update

> ArduinoThing ThingsV1Update(ctx, id, thing, optional)
update things_v1

Updates a thing associated to the user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the thing | 
**thing** | [**Thing**](Thing.md)| ThingPayload describes a thing | 
 **optional** | ***ThingsV1UpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsV1UpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **optional.Bool**| If true, detach device from the other thing, and attach to this thing | [default to false]

### Return type

[**ArduinoThing**](ArduinoThing.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.arduino.thing+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsV1UpdateSketch

> ArduinoThing ThingsV1UpdateSketch(ctx, id, sketchId)
updateSketch things_v1

Update an existing thing sketch

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the thing | 
**sketchId** | **string**| The id of the sketch | 

### Return type

[**ArduinoThing**](ArduinoThing.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.thing+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

