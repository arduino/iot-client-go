# \DevicesV2TagsApi

All URIs are relative to *https://api2.arduino.cc/iot*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevicesV2TagsDelete**](DevicesV2TagsApi.md#DevicesV2TagsDelete) | **Delete** /v2/devices/{id}/tags/{key} | delete devices_v2_tags
[**DevicesV2TagsList**](DevicesV2TagsApi.md#DevicesV2TagsList) | **Get** /v2/devices/{id}/tags | list devices_v2_tags
[**DevicesV2TagsUpsert**](DevicesV2TagsApi.md#DevicesV2TagsUpsert) | **Put** /v2/devices/{id}/tags | upsert devices_v2_tags



## DevicesV2TagsDelete

> DevicesV2TagsDelete(ctx, id, key).Execute()

delete devices_v2_tags



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
    key := "key_example" // string | The key of the tag

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DevicesV2TagsApi.DevicesV2TagsDelete(context.Background(), id, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesV2TagsApi.DevicesV2TagsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the device | 
**key** | **string** | The key of the tag | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesV2TagsDeleteRequest struct via the builder pattern


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


## DevicesV2TagsList

> ArduinoTags DevicesV2TagsList(ctx, id).Execute()

list devices_v2_tags



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
    resp, r, err := apiClient.DevicesV2TagsApi.DevicesV2TagsList(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesV2TagsApi.DevicesV2TagsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesV2TagsList`: ArduinoTags
    fmt.Fprintf(os.Stdout, "Response from `DevicesV2TagsApi.DevicesV2TagsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesV2TagsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ArduinoTags**](ArduinoTags.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.tags+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2TagsUpsert

> DevicesV2TagsUpsert(ctx, id).Tag(tag).Execute()

upsert devices_v2_tags



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
    tag := *openapiclient.NewTag("Key_example", "Value_example") // Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DevicesV2TagsApi.DevicesV2TagsUpsert(context.Background(), id).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesV2TagsApi.DevicesV2TagsUpsert``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDevicesV2TagsUpsertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tag** | [**Tag**](Tag.md) |  | 

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

