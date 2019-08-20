# \DevicesV2PubkeysApi

All URIs are relative to *http://api-dev.arduino.cc*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevicesV2PubkeysCreate**](DevicesV2PubkeysApi.md#DevicesV2PubkeysCreate) | **Put** /iot/v2/devices/{id}/pubkeys | create devices_v2_pubkeys
[**DevicesV2PubkeysDelete**](DevicesV2PubkeysApi.md#DevicesV2PubkeysDelete) | **Delete** /iot/v2/devices/{id}/pubkeys/{pid} | delete devices_v2_pubkeys
[**DevicesV2PubkeysList**](DevicesV2PubkeysApi.md#DevicesV2PubkeysList) | **Get** /iot/v2/devices/{id}/pubkeys | list devices_v2_pubkeys
[**DevicesV2PubkeysShow**](DevicesV2PubkeysApi.md#DevicesV2PubkeysShow) | **Get** /iot/v2/devices/{id}/pubkeys/{pid} | show devices_v2_pubkeys
[**DevicesV2PubkeysUpdate**](DevicesV2PubkeysApi.md#DevicesV2PubkeysUpdate) | **Post** /iot/v2/devices/{id}/pubkeys/{pid} | update devices_v2_pubkeys



## DevicesV2PubkeysCreate

> ArduinoDevicev2Pubkey DevicesV2PubkeysCreate(ctx, id, createDevicesV2PubkeysPayload)
create devices_v2_pubkeys

Creates a new pubkey associated to a device

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the device | 
**createDevicesV2PubkeysPayload** | [**CreateDevicesV2PubkeysPayload**](CreateDevicesV2PubkeysPayload.md)|  | 

### Return type

[**ArduinoDevicev2Pubkey**](ArduinoDevicev2Pubkey.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.arduino.devicev2.pubkey+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2PubkeysDelete

> DevicesV2PubkeysDelete(ctx, id, pid)
delete devices_v2_pubkeys

Removes a pubkey associated to a device

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the device | 
**pid** | **string**| The id of the pubkey | 

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


## DevicesV2PubkeysList

> []ArduinoDevicev2Pubkey DevicesV2PubkeysList(ctx, id)
list devices_v2_pubkeys

Returns the list of pubkeys associated to the device

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the device | 

### Return type

[**[]ArduinoDevicev2Pubkey**](ArduinoDevicev2Pubkey.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.devicev2.pubkey+json; type=collection

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2PubkeysShow

> ArduinoDevicev2Pubkey DevicesV2PubkeysShow(ctx, id, pid)
show devices_v2_pubkeys

Returns the pubkey requested by the user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the device | 
**pid** | **string**| The id of the pubkey | 

### Return type

[**ArduinoDevicev2Pubkey**](ArduinoDevicev2Pubkey.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.devicev2.pubkey+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2PubkeysUpdate

> ArduinoDevicev2Pubkey DevicesV2PubkeysUpdate(ctx, id, pid, devicev2Pubkey)
update devices_v2_pubkeys

Updates a pubkey associated to a device

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the device | 
**pid** | **string**| The id of the pubkey | 
**devicev2Pubkey** | [**Devicev2Pubkey**](Devicev2Pubkey.md)|  | 

### Return type

[**ArduinoDevicev2Pubkey**](ArduinoDevicev2Pubkey.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.arduino.devicev2.pubkey+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

