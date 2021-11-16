# \LoraDevicesV1Api

All URIs are relative to *https://api2.arduino.cc/iot*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LoraDevicesV1Create**](LoraDevicesV1Api.md#LoraDevicesV1Create) | **Put** /v1/lora-devices/ | create lora_devices_v1



## LoraDevicesV1Create

> ArduinoLoradevicev1 LoraDevicesV1Create(ctx, createLoraDevicesV1Payload)

create lora_devices_v1

Create a new lora device. Its info are saved on our database, and on the lora provider network. Creates a device_v2 automatically

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createLoraDevicesV1Payload** | [**CreateLoraDevicesV1Payload**](CreateLoraDevicesV1Payload.md)|  | 

### Return type

[**ArduinoLoradevicev1**](ArduinoLoradevicev1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

