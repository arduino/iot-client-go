# \DevicesV2OtaApi

All URIs are relative to *https://api2.arduino.cc/iot*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevicesV2OtaSend**](DevicesV2OtaApi.md#DevicesV2OtaSend) | **Put** /v2/devices/{id}/ota | send devices_v2_ota
[**DevicesV2OtaUpload**](DevicesV2OtaApi.md#DevicesV2OtaUpload) | **Post** /v2/devices/{id}/ota | upload devices_v2_ota
[**DevicesV2OtaUrl**](DevicesV2OtaApi.md#DevicesV2OtaUrl) | **Post** /v2/devices/{id}/ota/url | url devices_v2_ota



## DevicesV2OtaSend

> DevicesV2OtaSend(ctx, id).Devicev2Otabinaryurl(devicev2Otabinaryurl).Execute()

send devices_v2_ota



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/arduino/iot-client-go"
)

func main() {
    id := "id_example" // string | The id of the device
    devicev2Otabinaryurl := *openapiclient.NewDevicev2Otabinaryurl("BinaryKey_example") // Devicev2Otabinaryurl | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DevicesV2OtaApi.DevicesV2OtaSend(context.Background(), id).Devicev2Otabinaryurl(devicev2Otabinaryurl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesV2OtaApi.DevicesV2OtaSend``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesV2OtaSendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **devicev2Otabinaryurl** | [**Devicev2Otabinaryurl**](Devicev2Otabinaryurl.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.goa.error+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2OtaUpload

> ArduinoDevicev2Otaupload DevicesV2OtaUpload(ctx, id).OtaFile(otaFile).Async(async).ExpireInMins(expireInMins).Execute()

upload devices_v2_ota



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/arduino/iot-client-go"
)

func main() {
    id := "id_example" // string | The id of the device
    otaFile := os.NewFile(1234, "some_file") // *os.File | OTA file
    async := true // bool | If false, wait for the full OTA process, until it gets a result from the device (optional) (default to true)
    expireInMins := int32(56) // int32 | Binary expire time in minutes, default 10 mins (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesV2OtaApi.DevicesV2OtaUpload(context.Background(), id).OtaFile(otaFile).Async(async).ExpireInMins(expireInMins).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesV2OtaApi.DevicesV2OtaUpload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesV2OtaUpload`: ArduinoDevicev2Otaupload
    fmt.Fprintf(os.Stdout, "Response from `DevicesV2OtaApi.DevicesV2OtaUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesV2OtaUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **otaFile** | ***os.File** | OTA file | 
 **async** | **bool** | If false, wait for the full OTA process, until it gets a result from the device | [default to true]
 **expireInMins** | **int32** | Binary expire time in minutes, default 10 mins | [default to 10]

### Return type

[**ArduinoDevicev2Otaupload**](ArduinoDevicev2Otaupload.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/vnd.arduino.devicev2.otaupload+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2OtaUrl

> DevicesV2OtaUrl(ctx, id).Devicev2Otaurlpyalod(devicev2Otaurlpyalod).Execute()

url devices_v2_ota



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/arduino/iot-client-go"
)

func main() {
    id := "id_example" // string | The id of the device
    devicev2Otaurlpyalod := *openapiclient.NewDevicev2Otaurlpyalod() // Devicev2Otaurlpyalod | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DevicesV2OtaApi.DevicesV2OtaUrl(context.Background(), id).Devicev2Otaurlpyalod(devicev2Otaurlpyalod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesV2OtaApi.DevicesV2OtaUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesV2OtaUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **devicev2Otaurlpyalod** | [**Devicev2Otaurlpyalod**](Devicev2Otaurlpyalod.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.goa.error+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

