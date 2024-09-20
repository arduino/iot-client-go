# \PropertyTypesV1Api

All URIs are relative to *https://api2.arduino.cc/iot*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PropertyTypesV1ListTypes**](PropertyTypesV1Api.md#PropertyTypesV1ListTypes) | **Get** /v1/property_types | listTypes property_types_v1



## PropertyTypesV1ListTypes

> []ArduinoPropertytype PropertyTypesV1ListTypes(ctx).Execute()

listTypes property_types_v1



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PropertyTypesV1Api.PropertyTypesV1ListTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropertyTypesV1Api.PropertyTypesV1ListTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PropertyTypesV1ListTypes`: []ArduinoPropertytype
    fmt.Fprintf(os.Stdout, "Response from `PropertyTypesV1Api.PropertyTypesV1ListTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPropertyTypesV1ListTypesRequest struct via the builder pattern


### Return type

[**[]ArduinoPropertytype**](ArduinoPropertytype.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.propertytype+json; type=collection, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

