# \DevicesV2PassApi

All URIs are relative to *https://api2.arduino.cc/iot*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevicesV2PassCheck**](DevicesV2PassApi.md#DevicesV2PassCheck) | **Post** /v2/devices/{id}/pass | check devices_v2_pass
[**DevicesV2PassDelete**](DevicesV2PassApi.md#DevicesV2PassDelete) | **Delete** /v2/devices/{id}/pass | delete devices_v2_pass
[**DevicesV2PassGet**](DevicesV2PassApi.md#DevicesV2PassGet) | **Get** /v2/devices/{id}/pass | get devices_v2_pass
[**DevicesV2PassSet**](DevicesV2PassApi.md#DevicesV2PassSet) | **Put** /v2/devices/{id}/pass | set devices_v2_pass



## DevicesV2PassCheck

> DevicesV2PassCheck(ctx, id).CheckDevicesV2PassPayload(checkDevicesV2PassPayload).Execute()

check devices_v2_pass



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
    checkDevicesV2PassPayload := *openapiclient.NewCheckDevicesV2PassPayload("Password_example") // CheckDevicesV2PassPayload | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DevicesV2PassApi.DevicesV2PassCheck(context.Background(), id).CheckDevicesV2PassPayload(checkDevicesV2PassPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesV2PassApi.DevicesV2PassCheck``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDevicesV2PassCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **checkDevicesV2PassPayload** | [**CheckDevicesV2PassPayload**](CheckDevicesV2PassPayload.md) |  | 

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


## DevicesV2PassDelete

> DevicesV2PassDelete(ctx, id).Execute()

delete devices_v2_pass



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DevicesV2PassApi.DevicesV2PassDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesV2PassApi.DevicesV2PassDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDevicesV2PassDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DevicesV2PassGet

> ArduinoDevicev2Pass DevicesV2PassGet(ctx, id).SuggestedPassword(suggestedPassword).Execute()

get devices_v2_pass



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
    suggestedPassword := true // bool | If true, return a suggested password (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesV2PassApi.DevicesV2PassGet(context.Background(), id).SuggestedPassword(suggestedPassword).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesV2PassApi.DevicesV2PassGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesV2PassGet`: ArduinoDevicev2Pass
    fmt.Fprintf(os.Stdout, "Response from `DevicesV2PassApi.DevicesV2PassGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesV2PassGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **suggestedPassword** | **bool** | If true, return a suggested password | [default to false]

### Return type

[**ArduinoDevicev2Pass**](ArduinoDevicev2Pass.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.devicev2.pass+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2PassSet

> ArduinoDevicev2Pass DevicesV2PassSet(ctx, id).Devicev2Pass(devicev2Pass).Execute()

set devices_v2_pass



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
    devicev2Pass := *openapiclient.NewDevicev2Pass() // Devicev2Pass | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesV2PassApi.DevicesV2PassSet(context.Background(), id).Devicev2Pass(devicev2Pass).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesV2PassApi.DevicesV2PassSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesV2PassSet`: ArduinoDevicev2Pass
    fmt.Fprintf(os.Stdout, "Response from `DevicesV2PassApi.DevicesV2PassSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesV2PassSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **devicev2Pass** | [**Devicev2Pass**](Devicev2Pass.md) |  | 

### Return type

[**ArduinoDevicev2Pass**](ArduinoDevicev2Pass.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.arduino.devicev2.pass+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

