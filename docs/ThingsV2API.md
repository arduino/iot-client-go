# \ThingsV2API

All URIs are relative to *https://api2.arduino.cc/iot*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ThingsV2Clone**](ThingsV2API.md#ThingsV2Clone) | **Put** /v2/things/{id}/clone | clone things_v2
[**ThingsV2Create**](ThingsV2API.md#ThingsV2Create) | **Put** /v2/things | create things_v2
[**ThingsV2CreateSketch**](ThingsV2API.md#ThingsV2CreateSketch) | **Put** /v2/things/{id}/sketch | createSketch things_v2
[**ThingsV2Delete**](ThingsV2API.md#ThingsV2Delete) | **Delete** /v2/things/{id} | delete things_v2
[**ThingsV2DeleteSketch**](ThingsV2API.md#ThingsV2DeleteSketch) | **Delete** /v2/things/{id}/sketch | deleteSketch things_v2
[**ThingsV2List**](ThingsV2API.md#ThingsV2List) | **Get** /v2/things | list things_v2
[**ThingsV2Show**](ThingsV2API.md#ThingsV2Show) | **Get** /v2/things/{id} | show things_v2
[**ThingsV2Template**](ThingsV2API.md#ThingsV2Template) | **Get** /v2/things/{id}/template | template things_v2
[**ThingsV2Update**](ThingsV2API.md#ThingsV2Update) | **Post** /v2/things/{id} | update things_v2
[**ThingsV2UpdateSketch**](ThingsV2API.md#ThingsV2UpdateSketch) | **Put** /v2/things/{id}/sketch/{sketchId} | updateSketch things_v2



## ThingsV2Clone

> ArduinoThing ThingsV2Clone(ctx, id).ThingClone(thingClone).XOrganization(xOrganization).Execute()

clone things_v2



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
    id := "id_example" // string | The id of the thing
    thingClone := *openapiclient.NewThingClone("Name_example") // ThingClone | Payload to clone a new thing from an existing one
    xOrganization := "xOrganization_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThingsV2API.ThingsV2Clone(context.Background(), id).ThingClone(thingClone).XOrganization(xOrganization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThingsV2API.ThingsV2Clone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ThingsV2Clone`: ArduinoThing
    fmt.Fprintf(os.Stdout, "Response from `ThingsV2API.ThingsV2Clone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the thing | 

### Other Parameters

Other parameters are passed through a pointer to a apiThingsV2CloneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **thingClone** | [**ThingClone**](ThingClone.md) | Payload to clone a new thing from an existing one | 
 **xOrganization** | **string** |  | 

### Return type

[**ArduinoThing**](ArduinoThing.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.arduino.thing+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsV2Create

> ArduinoThing ThingsV2Create(ctx).ThingCreate(thingCreate).Force(force).XOrganization(xOrganization).Execute()

create things_v2



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
    thingCreate := *openapiclient.NewThingCreate() // ThingCreate | Payload to create a new thing
    force := true // bool | If true, detach device from the other thing, and attach to this thing (optional) (default to false)
    xOrganization := "xOrganization_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThingsV2API.ThingsV2Create(context.Background()).ThingCreate(thingCreate).Force(force).XOrganization(xOrganization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThingsV2API.ThingsV2Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ThingsV2Create`: ArduinoThing
    fmt.Fprintf(os.Stdout, "Response from `ThingsV2API.ThingsV2Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiThingsV2CreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **thingCreate** | [**ThingCreate**](ThingCreate.md) | Payload to create a new thing | 
 **force** | **bool** | If true, detach device from the other thing, and attach to this thing | [default to false]
 **xOrganization** | **string** |  | 

### Return type

[**ArduinoThing**](ArduinoThing.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.arduino.thing+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsV2CreateSketch

> ArduinoThing ThingsV2CreateSketch(ctx, id).ThingSketch(thingSketch).XOrganization(xOrganization).Execute()

createSketch things_v2



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
    id := "id_example" // string | The id of the thing
    thingSketch := *openapiclient.NewThingSketch() // ThingSketch | ThingSketchPayload describes a sketch of a thing
    xOrganization := "xOrganization_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThingsV2API.ThingsV2CreateSketch(context.Background(), id).ThingSketch(thingSketch).XOrganization(xOrganization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThingsV2API.ThingsV2CreateSketch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ThingsV2CreateSketch`: ArduinoThing
    fmt.Fprintf(os.Stdout, "Response from `ThingsV2API.ThingsV2CreateSketch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the thing | 

### Other Parameters

Other parameters are passed through a pointer to a apiThingsV2CreateSketchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **thingSketch** | [**ThingSketch**](ThingSketch.md) | ThingSketchPayload describes a sketch of a thing | 
 **xOrganization** | **string** |  | 

### Return type

[**ArduinoThing**](ArduinoThing.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.arduino.thing+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsV2Delete

> ThingsV2Delete(ctx, id).Force(force).XOrganization(xOrganization).Execute()

delete things_v2



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
    id := "id_example" // string | The id of the thing
    force := true // bool | If true, hard delete the thing (optional) (default to false)
    xOrganization := "xOrganization_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ThingsV2API.ThingsV2Delete(context.Background(), id).Force(force).XOrganization(xOrganization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThingsV2API.ThingsV2Delete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiThingsV2DeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | If true, hard delete the thing | [default to false]
 **xOrganization** | **string** |  | 

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


## ThingsV2DeleteSketch

> ArduinoThing ThingsV2DeleteSketch(ctx, id).XOrganization(xOrganization).Execute()

deleteSketch things_v2

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
    id := "id_example" // string | The id of the thing
    xOrganization := "xOrganization_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThingsV2API.ThingsV2DeleteSketch(context.Background(), id).XOrganization(xOrganization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThingsV2API.ThingsV2DeleteSketch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ThingsV2DeleteSketch`: ArduinoThing
    fmt.Fprintf(os.Stdout, "Response from `ThingsV2API.ThingsV2DeleteSketch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the thing | 

### Other Parameters

Other parameters are passed through a pointer to a apiThingsV2DeleteSketchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrganization** | **string** |  | 

### Return type

[**ArduinoThing**](ArduinoThing.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.thing+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsV2List

> []ArduinoThing ThingsV2List(ctx).AcrossUserIds(acrossUserIds).DeviceId(deviceId).Ids(ids).ShowDeleted(showDeleted).ShowProperties(showProperties).Tags(tags).XOrganization(xOrganization).Execute()

list things_v2



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
    acrossUserIds := true // bool | If true, returns all the things (optional) (default to false)
    deviceId := "deviceId_example" // string | The id of the device you want to filter (optional)
    ids := []string{"Inner_example"} // []string | Filter only the desired things (optional)
    showDeleted := true // bool | If true, shows the soft deleted things (optional) (default to false)
    showProperties := true // bool | If true, returns things with their properties, and last values (optional) (default to false)
    tags := []string{"Inner_example"} // []string | Filter by tags (optional)
    xOrganization := "xOrganization_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThingsV2API.ThingsV2List(context.Background()).AcrossUserIds(acrossUserIds).DeviceId(deviceId).Ids(ids).ShowDeleted(showDeleted).ShowProperties(showProperties).Tags(tags).XOrganization(xOrganization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThingsV2API.ThingsV2List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ThingsV2List`: []ArduinoThing
    fmt.Fprintf(os.Stdout, "Response from `ThingsV2API.ThingsV2List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiThingsV2ListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acrossUserIds** | **bool** | If true, returns all the things | [default to false]
 **deviceId** | **string** | The id of the device you want to filter | 
 **ids** | **[]string** | Filter only the desired things | 
 **showDeleted** | **bool** | If true, shows the soft deleted things | [default to false]
 **showProperties** | **bool** | If true, returns things with their properties, and last values | [default to false]
 **tags** | **[]string** | Filter by tags | 
 **xOrganization** | **string** |  | 

### Return type

[**[]ArduinoThing**](ArduinoThing.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.thing+json; type=collection, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsV2Show

> ArduinoThing ThingsV2Show(ctx, id).ShowDeleted(showDeleted).XOrganization(xOrganization).Execute()

show things_v2



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
    id := "id_example" // string | The id of the thing
    showDeleted := true // bool | If true, shows the soft deleted thing (optional) (default to false)
    xOrganization := "xOrganization_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThingsV2API.ThingsV2Show(context.Background(), id).ShowDeleted(showDeleted).XOrganization(xOrganization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThingsV2API.ThingsV2Show``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ThingsV2Show`: ArduinoThing
    fmt.Fprintf(os.Stdout, "Response from `ThingsV2API.ThingsV2Show`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the thing | 

### Other Parameters

Other parameters are passed through a pointer to a apiThingsV2ShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **showDeleted** | **bool** | If true, shows the soft deleted thing | [default to false]
 **xOrganization** | **string** |  | 

### Return type

[**ArduinoThing**](ArduinoThing.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.thing+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsV2Template

> ArduinoThingtemplate ThingsV2Template(ctx, id).XOrganization(xOrganization).Execute()

template things_v2



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
    id := "id_example" // string | The id of the thing
    xOrganization := "xOrganization_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThingsV2API.ThingsV2Template(context.Background(), id).XOrganization(xOrganization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThingsV2API.ThingsV2Template``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ThingsV2Template`: ArduinoThingtemplate
    fmt.Fprintf(os.Stdout, "Response from `ThingsV2API.ThingsV2Template`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the thing | 

### Other Parameters

Other parameters are passed through a pointer to a apiThingsV2TemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrganization** | **string** |  | 

### Return type

[**ArduinoThingtemplate**](ArduinoThingtemplate.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.thingtemplate+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsV2Update

> ArduinoThing ThingsV2Update(ctx, id).ThingUpdate(thingUpdate).Force(force).XOrganization(xOrganization).Execute()

update things_v2



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
    id := "id_example" // string | The id of the thing
    thingUpdate := *openapiclient.NewThingUpdate() // ThingUpdate | Payload to update an existing thing
    force := true // bool | If true, detach device from the other thing, and attach to this thing (optional) (default to false)
    xOrganization := "xOrganization_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThingsV2API.ThingsV2Update(context.Background(), id).ThingUpdate(thingUpdate).Force(force).XOrganization(xOrganization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThingsV2API.ThingsV2Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ThingsV2Update`: ArduinoThing
    fmt.Fprintf(os.Stdout, "Response from `ThingsV2API.ThingsV2Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the thing | 

### Other Parameters

Other parameters are passed through a pointer to a apiThingsV2UpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **thingUpdate** | [**ThingUpdate**](ThingUpdate.md) | Payload to update an existing thing | 
 **force** | **bool** | If true, detach device from the other thing, and attach to this thing | [default to false]
 **xOrganization** | **string** |  | 

### Return type

[**ArduinoThing**](ArduinoThing.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.arduino.thing+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsV2UpdateSketch

> ArduinoThing ThingsV2UpdateSketch(ctx, id, sketchId).XOrganization(xOrganization).UpdateSketch(updateSketch).Execute()

updateSketch things_v2



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
    id := "id_example" // string | The id of the thing
    sketchId := "sketchId_example" // string | The id of the sketch
    xOrganization := "xOrganization_example" // string |  (optional)
    updateSketch := *openapiclient.NewUpdateSketch() // UpdateSketch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThingsV2API.ThingsV2UpdateSketch(context.Background(), id, sketchId).XOrganization(xOrganization).UpdateSketch(updateSketch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThingsV2API.ThingsV2UpdateSketch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ThingsV2UpdateSketch`: ArduinoThing
    fmt.Fprintf(os.Stdout, "Response from `ThingsV2API.ThingsV2UpdateSketch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the thing | 
**sketchId** | **string** | The id of the sketch | 

### Other Parameters

Other parameters are passed through a pointer to a apiThingsV2UpdateSketchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xOrganization** | **string** |  | 
 **updateSketch** | [**UpdateSketch**](UpdateSketch.md) |  | 

### Return type

[**ArduinoThing**](ArduinoThing.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.arduino.thing+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

