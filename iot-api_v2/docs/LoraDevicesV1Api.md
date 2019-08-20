# \LoraDevicesV1Api

All URIs are relative to *http://api-dev.arduino.cc*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LoraDevicesV1Create**](LoraDevicesV1Api.md#LoraDevicesV1Create) | **Put** /iot/v1/lora-devices/ | create lora_devices_v1



## LoraDevicesV1Create

> ArduinoLoradevicev1 LoraDevicesV1Create(ctx, createLoraDevicesV1Payload)
create lora_devices_v1

Create a new lora device. Its info are saved on our database, and on the a2a network. Creates a device_v2 automatically

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createLoraDevicesV1Payload** | [**CreateLoraDevicesV1Payload**](CreateLoraDevicesV1Payload.md)|  | 

### Return type

[**ArduinoLoradevicev1**](ArduinoLoradevicev1.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.arduino.loradevicev1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

