# \NetworkCredentialsV1Api

All URIs are relative to *https://api2.arduino.cc/iot*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NetworkCredentialsV1Show**](NetworkCredentialsV1Api.md#NetworkCredentialsV1Show) | **Get** /v1/network_credentials/{type} | show network_credentials_v1
[**NetworkCredentialsV1ShowByDevice**](NetworkCredentialsV1Api.md#NetworkCredentialsV1ShowByDevice) | **Get** /v1/network_credentials/{type}/connections | showByDevice network_credentials_v1



## NetworkCredentialsV1Show

> []ArduinoCredentialsv1 NetworkCredentialsV1Show(ctx, type_).Connection(connection).Execute()

show network_credentials_v1



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/arduino/iot-client-go/v2"
)

func main() {
    type_ := "type__example" // string | Device type
    connection := "connection_example" // string | Connection used by the device (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkCredentialsV1Api.NetworkCredentialsV1Show(context.Background(), type_).Connection(connection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkCredentialsV1Api.NetworkCredentialsV1Show``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkCredentialsV1Show`: []ArduinoCredentialsv1
    fmt.Fprintf(os.Stdout, "Response from `NetworkCredentialsV1Api.NetworkCredentialsV1Show`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | Device type | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkCredentialsV1ShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connection** | **string** | Connection used by the device | 

### Return type

[**[]ArduinoCredentialsv1**](ArduinoCredentialsv1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.credentialsv1+json; type=collection

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkCredentialsV1ShowByDevice

> NetworkCredentialsV1ShowByDevice(ctx, type_).Execute()

showByDevice network_credentials_v1



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/arduino/iot-client-go/v2"
)

func main() {
    type_ := "type__example" // string | Device type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NetworkCredentialsV1Api.NetworkCredentialsV1ShowByDevice(context.Background(), type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkCredentialsV1Api.NetworkCredentialsV1ShowByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | Device type | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkCredentialsV1ShowByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

