# \ThingsV2Api

All URIs are relative to *https://api2.arduino.cc/iot*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ThingsV2Create**](ThingsV2Api.md#ThingsV2Create) | **Put** /v2/things | create things_v2
[**ThingsV2CreateSketch**](ThingsV2Api.md#ThingsV2CreateSketch) | **Put** /v2/things/{id}/sketch | createSketch things_v2
[**ThingsV2Delete**](ThingsV2Api.md#ThingsV2Delete) | **Delete** /v2/things/{id} | delete things_v2
[**ThingsV2DeleteSketch**](ThingsV2Api.md#ThingsV2DeleteSketch) | **Delete** /v2/things/{id}/sketch | deleteSketch things_v2
[**ThingsV2List**](ThingsV2Api.md#ThingsV2List) | **Get** /v2/things | list things_v2
[**ThingsV2Show**](ThingsV2Api.md#ThingsV2Show) | **Get** /v2/things/{id} | show things_v2
[**ThingsV2Update**](ThingsV2Api.md#ThingsV2Update) | **Post** /v2/things/{id} | update things_v2
[**ThingsV2UpdateSketch**](ThingsV2Api.md#ThingsV2UpdateSketch) | **Put** /v2/things/{id}/sketch/{sketchId} | updateSketch things_v2



## ThingsV2Create

> ArduinoThing ThingsV2Create(ctx, createThingsV2Payload, optional)

create things_v2

Creates a new thing associated to the user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createThingsV2Payload** | [**CreateThingsV2Payload**](CreateThingsV2Payload.md)| ThingPayload describes a thing | 
 **optional** | ***ThingsV2CreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsV2CreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| If true, detach device from the other thing, and attach to this thing | 

### Return type

[**ArduinoThing**](ArduinoThing.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsV2CreateSketch

> ArduinoThing ThingsV2CreateSketch(ctx, id, thingSketch)

createSketch things_v2

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsV2Delete

> ThingsV2Delete(ctx, id, optional)

delete things_v2

Removes a thing associated to the user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the thing | 
 **optional** | ***ThingsV2DeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsV2DeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| If true, hard delete the thing | 

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


## ThingsV2DeleteSketch

> ArduinoThing ThingsV2DeleteSketch(ctx, id)

deleteSketch things_v2

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsV2List

> []ArduinoThing ThingsV2List(ctx, optional)

list things_v2

Returns the list of things associated to the user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ThingsV2ListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsV2ListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acrossUserIds** | **optional.Bool**| If true, returns all the things | 
 **deviceId** | **optional.String**| The id of the device you want to filter | 
 **ids** | [**optional.Interface of []string**](string.md)| Filter only the desired things | 
 **showDeleted** | **optional.Bool**| If true, shows the soft deleted things | 
 **showProperties** | **optional.Bool**| If true, returns things with their properties, and last values | 

### Return type

[**[]ArduinoThing**](ArduinoThing.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsV2Show

> ArduinoThing ThingsV2Show(ctx, id, optional)

show things_v2

Returns the thing requested by the user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the thing | 
 **optional** | ***ThingsV2ShowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsV2ShowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **showDeleted** | **optional.Bool**| If true, shows the soft deleted thing | 

### Return type

[**ArduinoThing**](ArduinoThing.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsV2Update

> ArduinoThing ThingsV2Update(ctx, id, thing, optional)

update things_v2

Updates a thing associated to the user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the thing | 
**thing** | [**Thing**](Thing.md)| ThingPayload describes a thing | 
 **optional** | ***ThingsV2UpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsV2UpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **optional.Bool**| If true, detach device from the other thing, and attach to this thing | 

### Return type

[**ArduinoThing**](ArduinoThing.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsV2UpdateSketch

> ArduinoThing ThingsV2UpdateSketch(ctx, id, sketchId, optional)

updateSketch things_v2

Update an existing thing sketch

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the thing | 
**sketchId** | **string**| The id of the sketch | 
 **optional** | ***ThingsV2UpdateSketchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsV2UpdateSketchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateSketch** | [**optional.Interface of UpdateSketch**](UpdateSketch.md)|  | 

### Return type

[**ArduinoThing**](ArduinoThing.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

