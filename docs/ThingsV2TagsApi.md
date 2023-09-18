# \ThingsV2TagsApi

All URIs are relative to *https://api2.arduino.cc/iot*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ThingsV2TagsDelete**](ThingsV2TagsApi.md#ThingsV2TagsDelete) | **Delete** /v2/things/{id}/tags/{key} | delete things_v2_tags
[**ThingsV2TagsList**](ThingsV2TagsApi.md#ThingsV2TagsList) | **Get** /v2/things/{id}/tags | list things_v2_tags
[**ThingsV2TagsUpsert**](ThingsV2TagsApi.md#ThingsV2TagsUpsert) | **Put** /v2/things/{id}/tags | upsert things_v2_tags



## ThingsV2TagsDelete

> ThingsV2TagsDelete(ctx, id, key).Execute()

delete things_v2_tags



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
    id := "id_example" // string | The id of the thing
    key := "key_example" // string | The key of the tag

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ThingsV2TagsApi.ThingsV2TagsDelete(context.Background(), id, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThingsV2TagsApi.ThingsV2TagsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the thing | 
**key** | **string** | The key of the tag | 

### Other Parameters

Other parameters are passed through a pointer to a apiThingsV2TagsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## ThingsV2TagsList

> ArduinoTags ThingsV2TagsList(ctx, id).Execute()

list things_v2_tags



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
    id := "id_example" // string | The id of the thing

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThingsV2TagsApi.ThingsV2TagsList(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThingsV2TagsApi.ThingsV2TagsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ThingsV2TagsList`: ArduinoTags
    fmt.Fprintf(os.Stdout, "Response from `ThingsV2TagsApi.ThingsV2TagsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the thing | 

### Other Parameters

Other parameters are passed through a pointer to a apiThingsV2TagsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ArduinoTags**](ArduinoTags.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsV2TagsUpsert

> ThingsV2TagsUpsert(ctx, id).Tag(tag).Execute()

upsert things_v2_tags



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
    id := "id_example" // string | The id of the thing
    tag := *openapiclient.NewTag("Key_example", "Value_example") // Tag | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ThingsV2TagsApi.ThingsV2TagsUpsert(context.Background(), id).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThingsV2TagsApi.ThingsV2TagsUpsert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the thing | 

### Other Parameters

Other parameters are passed through a pointer to a apiThingsV2TagsUpsertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tag** | [**Tag**](Tag.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

