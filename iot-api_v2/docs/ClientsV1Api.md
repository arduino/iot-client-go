# \ClientsV1Api

All URIs are relative to *http://api-dev.arduino.cc*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClientsV1Create**](ClientsV1Api.md#ClientsV1Create) | **Put** /iot/v1/clients | create clients_v1
[**ClientsV1Delete**](ClientsV1Api.md#ClientsV1Delete) | **Delete** /iot/v1/clients/{id} | delete clients_v1
[**ClientsV1List**](ClientsV1Api.md#ClientsV1List) | **Get** /iot/v1/clients | list clients_v1
[**ClientsV1Show**](ClientsV1Api.md#ClientsV1Show) | **Get** /iot/v1/clients/{id} | show clients_v1



## ClientsV1Create

> ArduinoClient ClientsV1Create(ctx, client)
create clients_v1

Creates a new client associated to the user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**client** | [**Client**](Client.md)| ClientPayload describes an oauth client | 

### Return type

[**ArduinoClient**](ArduinoClient.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.arduino.client+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientsV1Delete

> ClientsV1Delete(ctx, id)
delete clients_v1

Removes a client associated to the user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the client | 

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


## ClientsV1List

> []ArduinoClient ClientsV1List(ctx, )
list clients_v1

Returns the list of oauth clients of the user

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]ArduinoClient**](ArduinoClient.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.client+json; type=collection, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientsV1Show

> ArduinoClient ClientsV1Show(ctx, id)
show clients_v1

Returns the client requested by the user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the thing | 

### Return type

[**ArduinoClient**](ArduinoClient.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.client+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

