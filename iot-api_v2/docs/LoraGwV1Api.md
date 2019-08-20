# \LoraGwV1Api

All URIs are relative to *http://api-dev.arduino.cc*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LoraGwV1Claim**](LoraGwV1Api.md#LoraGwV1Claim) | **Post** /iot/v1/lora-gw/{id}/claim | claim lora_gw_v1
[**LoraGwV1Create**](LoraGwV1Api.md#LoraGwV1Create) | **Put** /iot/v1/lora-gw/ | create lora_gw_v1
[**LoraGwV1Delete**](LoraGwV1Api.md#LoraGwV1Delete) | **Delete** /iot/v1/lora-gw/{id} | delete lora_gw_v1
[**LoraGwV1Show**](LoraGwV1Api.md#LoraGwV1Show) | **Get** /iot/v1/lora-gw/{id} | show lora_gw_v1
[**LoraGwV1Unclaim**](LoraGwV1Api.md#LoraGwV1Unclaim) | **Post** /iot/v1/lora-gw/{id}/unclaim | unclaim lora_gw_v1



## LoraGwV1Claim

> ArduinoLoragwv1 LoraGwV1Claim(ctx, id, claimLoraGwV1Payload)
claim lora_gw_v1

Assign a lora gateway to a user, create a device_v1 resource, create an a2a user (if not existing), create the device on a2a

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The mac/serial_number/device_code of the lora gateway. It is printed on the box, and used to connect to a2a. | 
**claimLoraGwV1Payload** | [**ClaimLoraGwV1Payload**](ClaimLoraGwV1Payload.md)|  | 

### Return type

[**ArduinoLoragwv1**](ArduinoLoragwv1.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.arduino.loragwv1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoraGwV1Create

> ArduinoLoragwv1 LoraGwV1Create(ctx, createLoraGwV1Payload)
create lora_gw_v1

Create a new pending lora gateway. Its info are only saved on our database.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createLoraGwV1Payload** | [**CreateLoraGwV1Payload**](CreateLoraGwV1Payload.md)|  | 

### Return type

[**ArduinoLoragwv1**](ArduinoLoragwv1.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.arduino.loragwv1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoraGwV1Delete

> LoraGwV1Delete(ctx, id)
delete lora_gw_v1

Removes a lora gateway

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the lora gateway. It&#39;s calculated from the mac using magic | 

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


## LoraGwV1Show

> ArduinoLoragwv1 LoraGwV1Show(ctx, id)
show lora_gw_v1

Retrieve info about a lora gateway.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the lora gateway. It&#39;s calculated from the mac using magic | 

### Return type

[**ArduinoLoragwv1**](ArduinoLoragwv1.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.loragwv1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoraGwV1Unclaim

> ArduinoLoragwv1 LoraGwV1Unclaim(ctx, id)
unclaim lora_gw_v1

Remove the gateway from a2a, delete the device_v1 resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The mac/serial_number/device_code of the lora gateway. It is printed on the box, and used to connect to a2a. | 

### Return type

[**ArduinoLoragwv1**](ArduinoLoragwv1.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.loragwv1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

