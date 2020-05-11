# \DevicesV2CertsApi

All URIs are relative to *https://api2.arduino.cc/iot*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevicesV2CertsCreate**](DevicesV2CertsApi.md#DevicesV2CertsCreate) | **Put** /v2/devices/{id}/certs | create devices_v2_certs
[**DevicesV2CertsDelete**](DevicesV2CertsApi.md#DevicesV2CertsDelete) | **Delete** /v2/devices/{id}/certs/{cid} | delete devices_v2_certs
[**DevicesV2CertsList**](DevicesV2CertsApi.md#DevicesV2CertsList) | **Get** /v2/devices/{id}/certs | list devices_v2_certs
[**DevicesV2CertsShow**](DevicesV2CertsApi.md#DevicesV2CertsShow) | **Get** /v2/devices/{id}/certs/{cid} | show devices_v2_certs
[**DevicesV2CertsUpdate**](DevicesV2CertsApi.md#DevicesV2CertsUpdate) | **Post** /v2/devices/{id}/certs/{cid} | update devices_v2_certs



## DevicesV2CertsCreate

> ArduinoDevicev2Cert DevicesV2CertsCreate(ctx, id, createDevicesV2CertsPayload)

create devices_v2_certs

Creates a new cert associated to a device. The csr is signed and saved in database. The CommonName will be replaced with the device id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the device | 
**createDevicesV2CertsPayload** | [**CreateDevicesV2CertsPayload**](CreateDevicesV2CertsPayload.md)|  | 

### Return type

[**ArduinoDevicev2Cert**](ArduinoDevicev2Cert.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2CertsDelete

> DevicesV2CertsDelete(ctx, cid, id)

delete devices_v2_certs

Removes a cert associated to a device

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string**| The id of the cert | 
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


## DevicesV2CertsList

> []ArduinoDevicev2Cert DevicesV2CertsList(ctx, id)

list devices_v2_certs

Returns the list of certs associated to the device

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the device | 

### Return type

[**[]ArduinoDevicev2Cert**](ArduinoDevicev2Cert.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2CertsShow

> ArduinoDevicev2Cert DevicesV2CertsShow(ctx, cid, id)

show devices_v2_certs

Returns the cert requested by the user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string**| The id of the cert | 
**id** | **string**| The id of the device | 

### Return type

[**ArduinoDevicev2Cert**](ArduinoDevicev2Cert.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2CertsUpdate

> ArduinoDevicev2Cert DevicesV2CertsUpdate(ctx, cid, id, devicev2Cert)

update devices_v2_certs

Updates a cert associated to a device. The csr is signed and saved in database. The CommonName will be replaced with the device id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string**| The id of the cert | 
**id** | **string**| The id of the device | 
**devicev2Cert** | [**Devicev2Cert**](Devicev2Cert.md)|  | 

### Return type

[**ArduinoDevicev2Cert**](ArduinoDevicev2Cert.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

