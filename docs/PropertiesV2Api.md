# \PropertiesV2Api

All URIs are relative to *https://api2.arduino.cc/iot*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PropertiesV2Create**](PropertiesV2Api.md#PropertiesV2Create) | **Put** /v2/things/{id}/properties | create properties_v2
[**PropertiesV2Delete**](PropertiesV2Api.md#PropertiesV2Delete) | **Delete** /v2/things/{id}/properties/{pid} | delete properties_v2
[**PropertiesV2List**](PropertiesV2Api.md#PropertiesV2List) | **Get** /v2/things/{id}/properties | list properties_v2
[**PropertiesV2Publish**](PropertiesV2Api.md#PropertiesV2Publish) | **Put** /v2/things/{id}/properties/{pid}/publish | publish properties_v2
[**PropertiesV2Show**](PropertiesV2Api.md#PropertiesV2Show) | **Get** /v2/things/{id}/properties/{pid} | show properties_v2
[**PropertiesV2Timeseries**](PropertiesV2Api.md#PropertiesV2Timeseries) | **Get** /v2/things/{id}/properties/{pid}/timeseries | timeseries properties_v2
[**PropertiesV2Update**](PropertiesV2Api.md#PropertiesV2Update) | **Post** /v2/things/{id}/properties/{pid} | update properties_v2



## PropertiesV2Create

> ArduinoProperty PropertiesV2Create(ctx, id).Property(property).Execute()

create properties_v2



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
    property := *openapiclient.NewProperty("Name_example", "Permission_example", "Type_example", "UpdateStrategy_example") // Property | PropertyPayload describes a property of a thing. No field is mandatory

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PropertiesV2Api.PropertiesV2Create(context.Background(), id).Property(property).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropertiesV2Api.PropertiesV2Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PropertiesV2Create`: ArduinoProperty
    fmt.Fprintf(os.Stdout, "Response from `PropertiesV2Api.PropertiesV2Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the thing | 

### Other Parameters

Other parameters are passed through a pointer to a apiPropertiesV2CreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **property** | [**Property**](Property.md) | PropertyPayload describes a property of a thing. No field is mandatory | 

### Return type

[**ArduinoProperty**](ArduinoProperty.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PropertiesV2Delete

> PropertiesV2Delete(ctx, id, pid).Force(force).Execute()

delete properties_v2



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
    pid := "pid_example" // string | The id of the property
    force := true // bool | If true, hard delete the property (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PropertiesV2Api.PropertiesV2Delete(context.Background(), id, pid).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropertiesV2Api.PropertiesV2Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the thing | 
**pid** | **string** | The id of the property | 

### Other Parameters

Other parameters are passed through a pointer to a apiPropertiesV2DeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **bool** | If true, hard delete the property | [default to false]

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


## PropertiesV2List

> []ArduinoProperty PropertiesV2List(ctx, id).ShowDeleted(showDeleted).Execute()

list properties_v2



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
    showDeleted := true // bool | If true, shows the soft deleted properties (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PropertiesV2Api.PropertiesV2List(context.Background(), id).ShowDeleted(showDeleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropertiesV2Api.PropertiesV2List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PropertiesV2List`: []ArduinoProperty
    fmt.Fprintf(os.Stdout, "Response from `PropertiesV2Api.PropertiesV2List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the thing | 

### Other Parameters

Other parameters are passed through a pointer to a apiPropertiesV2ListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **showDeleted** | **bool** | If true, shows the soft deleted properties | [default to false]

### Return type

[**[]ArduinoProperty**](ArduinoProperty.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PropertiesV2Publish

> PropertiesV2Publish(ctx, id, pid).PropertyValue(propertyValue).Execute()

publish properties_v2



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
    pid := "pid_example" // string | The id of the property
    propertyValue := *openapiclient.NewPropertyValue(interface{}(123)) // PropertyValue | PropertyValuePayload describes a property value

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PropertiesV2Api.PropertiesV2Publish(context.Background(), id, pid).PropertyValue(propertyValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropertiesV2Api.PropertiesV2Publish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the thing | 
**pid** | **string** | The id of the property | 

### Other Parameters

Other parameters are passed through a pointer to a apiPropertiesV2PublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **propertyValue** | [**PropertyValue**](PropertyValue.md) | PropertyValuePayload describes a property value | 

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


## PropertiesV2Show

> ArduinoProperty PropertiesV2Show(ctx, id, pid).ShowDeleted(showDeleted).Execute()

show properties_v2



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
    pid := "pid_example" // string | The id of the property
    showDeleted := true // bool | If true, shows the soft deleted properties (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PropertiesV2Api.PropertiesV2Show(context.Background(), id, pid).ShowDeleted(showDeleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropertiesV2Api.PropertiesV2Show``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PropertiesV2Show`: ArduinoProperty
    fmt.Fprintf(os.Stdout, "Response from `PropertiesV2Api.PropertiesV2Show`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the thing | 
**pid** | **string** | The id of the property | 

### Other Parameters

Other parameters are passed through a pointer to a apiPropertiesV2ShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **showDeleted** | **bool** | If true, shows the soft deleted properties | [default to false]

### Return type

[**ArduinoProperty**](ArduinoProperty.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PropertiesV2Timeseries

> ArduinoTimeseriesmedia PropertiesV2Timeseries(ctx, id, pid).Desc(desc).From(from).Interval(interval).To(to).Execute()

timeseries properties_v2



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
    pid := "pid_example" // string | ID of a numerical property
    desc := true // bool | Whether data's ordering (by time) should be descending (optional) (default to false)
    from := "from_example" // string | Get data with a timestamp >= to this date (default: 2 weeks ago, min: 1842-01-01T00:00:00Z, max: 2242-01-01T00:00:00Z) (optional)
    interval := int32(56) // int32 | Binning interval in seconds (defaut: the smallest possible value compatibly with the limit of 1000 data points in the response) (optional)
    to := "to_example" // string | Get data with a timestamp < to this date (default: now, min: 1842-01-01T00:00:00Z, max: 2242-01-01T00:00:00Z) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PropertiesV2Api.PropertiesV2Timeseries(context.Background(), id, pid).Desc(desc).From(from).Interval(interval).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropertiesV2Api.PropertiesV2Timeseries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PropertiesV2Timeseries`: ArduinoTimeseriesmedia
    fmt.Fprintf(os.Stdout, "Response from `PropertiesV2Api.PropertiesV2Timeseries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the thing | 
**pid** | **string** | ID of a numerical property | 

### Other Parameters

Other parameters are passed through a pointer to a apiPropertiesV2TimeseriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **desc** | **bool** | Whether data&#39;s ordering (by time) should be descending | [default to false]
 **from** | **string** | Get data with a timestamp &gt;&#x3D; to this date (default: 2 weeks ago, min: 1842-01-01T00:00:00Z, max: 2242-01-01T00:00:00Z) | 
 **interval** | **int32** | Binning interval in seconds (defaut: the smallest possible value compatibly with the limit of 1000 data points in the response) | 
 **to** | **string** | Get data with a timestamp &lt; to this date (default: now, min: 1842-01-01T00:00:00Z, max: 2242-01-01T00:00:00Z) | 

### Return type

[**ArduinoTimeseriesmedia**](ArduinoTimeseriesmedia.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PropertiesV2Update

> ArduinoProperty PropertiesV2Update(ctx, id, pid).Property(property).Execute()

update properties_v2



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
    pid := "pid_example" // string | The id of the property
    property := *openapiclient.NewProperty("Name_example", "Permission_example", "Type_example", "UpdateStrategy_example") // Property | PropertyPayload describes a property of a thing. No field is mandatory

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PropertiesV2Api.PropertiesV2Update(context.Background(), id, pid).Property(property).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropertiesV2Api.PropertiesV2Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PropertiesV2Update`: ArduinoProperty
    fmt.Fprintf(os.Stdout, "Response from `PropertiesV2Api.PropertiesV2Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the thing | 
**pid** | **string** | The id of the property | 

### Other Parameters

Other parameters are passed through a pointer to a apiPropertiesV2UpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **property** | [**Property**](Property.md) | PropertyPayload describes a property of a thing. No field is mandatory | 

### Return type

[**ArduinoProperty**](ArduinoProperty.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

