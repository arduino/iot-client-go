# \DevicesV2API

All URIs are relative to *https://api2.arduino.cc/iot*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevicesV2Create**](DevicesV2API.md#DevicesV2Create) | **Put** /v2/devices | create devices_v2
[**DevicesV2Delete**](DevicesV2API.md#DevicesV2Delete) | **Delete** /v2/devices/{id} | delete devices_v2
[**DevicesV2GetEvents**](DevicesV2API.md#DevicesV2GetEvents) | **Get** /v2/devices/{id}/events | getEvents devices_v2
[**DevicesV2GetProperties**](DevicesV2API.md#DevicesV2GetProperties) | **Get** /v2/devices/{id}/properties | getProperties devices_v2
[**DevicesV2GetStatusEvents**](DevicesV2API.md#DevicesV2GetStatusEvents) | **Get** /v2/devices/{id}/status | GetStatusEvents devices_v2
[**DevicesV2List**](DevicesV2API.md#DevicesV2List) | **Get** /v2/devices | list devices_v2
[**DevicesV2Show**](DevicesV2API.md#DevicesV2Show) | **Get** /v2/devices/{id} | show devices_v2
[**DevicesV2Timeseries**](DevicesV2API.md#DevicesV2Timeseries) | **Get** /v2/devices/{id}/properties/{pid} | timeseries devices_v2
[**DevicesV2Update**](DevicesV2API.md#DevicesV2Update) | **Post** /v2/devices/{id} | update devices_v2
[**DevicesV2UpdateProperties**](DevicesV2API.md#DevicesV2UpdateProperties) | **Put** /v2/devices/{id}/properties | updateProperties devices_v2



## DevicesV2Create

> ArduinoDevicev2 DevicesV2Create(ctx).CreateDevicesV2Payload(createDevicesV2Payload).XOrganization(xOrganization).Execute()

create devices_v2



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
	createDevicesV2Payload := *openapiclient.NewCreateDevicesV2Payload("Type_example") // CreateDevicesV2Payload | DeviceV2 describes a device.
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesV2API.DevicesV2Create(context.Background()).CreateDevicesV2Payload(createDevicesV2Payload).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesV2API.DevicesV2Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DevicesV2Create`: ArduinoDevicev2
	fmt.Fprintf(os.Stdout, "Response from `DevicesV2API.DevicesV2Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDevicesV2CreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDevicesV2Payload** | [**CreateDevicesV2Payload**](CreateDevicesV2Payload.md) | DeviceV2 describes a device. | 
 **xOrganization** | **string** |  | 

### Return type

[**ArduinoDevicev2**](ArduinoDevicev2.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.arduino.devicev2+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2Delete

> DevicesV2Delete(ctx, id).XOrganization(xOrganization).Execute()

delete devices_v2



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
	id := "id_example" // string | The id of the device
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevicesV2API.DevicesV2Delete(context.Background(), id).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesV2API.DevicesV2Delete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDevicesV2DeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## DevicesV2GetEvents

> ArduinoDevicev2EventProperties DevicesV2GetEvents(ctx, id).Limit(limit).Start(start).XOrganization(xOrganization).Execute()

getEvents devices_v2



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
	id := "id_example" // string | The id of the device
	limit := int32(56) // int32 | The number of events to select (optional)
	start := "start_example" // string | The time at which to start selecting events (optional)
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesV2API.DevicesV2GetEvents(context.Background(), id).Limit(limit).Start(start).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesV2API.DevicesV2GetEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DevicesV2GetEvents`: ArduinoDevicev2EventProperties
	fmt.Fprintf(os.Stdout, "Response from `DevicesV2API.DevicesV2GetEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesV2GetEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The number of events to select | 
 **start** | **string** | The time at which to start selecting events | 
 **xOrganization** | **string** |  | 

### Return type

[**ArduinoDevicev2EventProperties**](ArduinoDevicev2EventProperties.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.devicev2.event.properties+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2GetProperties

> ArduinoDevicev2properties DevicesV2GetProperties(ctx, id).ShowDeleted(showDeleted).XOrganization(xOrganization).Execute()

getProperties devices_v2



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
	id := "id_example" // string | The id of the device
	showDeleted := true // bool | If true, shows the soft deleted properties (optional) (default to false)
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesV2API.DevicesV2GetProperties(context.Background(), id).ShowDeleted(showDeleted).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesV2API.DevicesV2GetProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DevicesV2GetProperties`: ArduinoDevicev2properties
	fmt.Fprintf(os.Stdout, "Response from `DevicesV2API.DevicesV2GetProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesV2GetPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **showDeleted** | **bool** | If true, shows the soft deleted properties | [default to false]
 **xOrganization** | **string** |  | 

### Return type

[**ArduinoDevicev2properties**](ArduinoDevicev2properties.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.devicev2properties+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2GetStatusEvents

> ArduinoDevicev2StatusEvents DevicesV2GetStatusEvents(ctx, id).Limit(limit).Start(start).XOrganization(xOrganization).Execute()

GetStatusEvents devices_v2



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
	id := "id_example" // string | The id of the device
	limit := int32(56) // int32 | The number of events to select (optional) (default to 30)
	start := "start_example" // string | The time at which to start selecting events (optional)
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesV2API.DevicesV2GetStatusEvents(context.Background(), id).Limit(limit).Start(start).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesV2API.DevicesV2GetStatusEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DevicesV2GetStatusEvents`: ArduinoDevicev2StatusEvents
	fmt.Fprintf(os.Stdout, "Response from `DevicesV2API.DevicesV2GetStatusEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesV2GetStatusEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The number of events to select | [default to 30]
 **start** | **string** | The time at which to start selecting events | 
 **xOrganization** | **string** |  | 

### Return type

[**ArduinoDevicev2StatusEvents**](ArduinoDevicev2StatusEvents.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.devicev2.status.events+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2List

> []ArduinoDevicev2 DevicesV2List(ctx).AcrossUserIds(acrossUserIds).Serial(serial).Tags(tags).XOrganization(xOrganization).Execute()

list devices_v2



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
	acrossUserIds := true // bool | If true, returns all the devices (optional) (default to false)
	serial := "serial_example" // string | Filter by device serial number (optional)
	tags := []string{"Inner_example"} // []string | Filter by tags (optional)
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesV2API.DevicesV2List(context.Background()).AcrossUserIds(acrossUserIds).Serial(serial).Tags(tags).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesV2API.DevicesV2List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DevicesV2List`: []ArduinoDevicev2
	fmt.Fprintf(os.Stdout, "Response from `DevicesV2API.DevicesV2List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDevicesV2ListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acrossUserIds** | **bool** | If true, returns all the devices | [default to false]
 **serial** | **string** | Filter by device serial number | 
 **tags** | **[]string** | Filter by tags | 
 **xOrganization** | **string** |  | 

### Return type

[**[]ArduinoDevicev2**](ArduinoDevicev2.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.devicev2+json; type=collection, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2Show

> ArduinoDevicev2 DevicesV2Show(ctx, id).XOrganization(xOrganization).Execute()

show devices_v2



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
	id := "id_example" // string | The id of the device
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesV2API.DevicesV2Show(context.Background(), id).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesV2API.DevicesV2Show``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DevicesV2Show`: ArduinoDevicev2
	fmt.Fprintf(os.Stdout, "Response from `DevicesV2API.DevicesV2Show`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesV2ShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrganization** | **string** |  | 

### Return type

[**ArduinoDevicev2**](ArduinoDevicev2.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.devicev2+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2Timeseries

> ArduinoDevicev2propertyvalues DevicesV2Timeseries(ctx, id, pid).Limit(limit).Start(start).XOrganization(xOrganization).Execute()

timeseries devices_v2



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
	id := "id_example" // string | The id of the device
	pid := "pid_example" // string | The id of the property
	limit := int32(56) // int32 | The number of properties to select (optional)
	start := "start_example" // string | The time at which to start selecting properties (optional)
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesV2API.DevicesV2Timeseries(context.Background(), id, pid).Limit(limit).Start(start).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesV2API.DevicesV2Timeseries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DevicesV2Timeseries`: ArduinoDevicev2propertyvalues
	fmt.Fprintf(os.Stdout, "Response from `DevicesV2API.DevicesV2Timeseries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the device | 
**pid** | **string** | The id of the property | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesV2TimeseriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The number of properties to select | 
 **start** | **string** | The time at which to start selecting properties | 
 **xOrganization** | **string** |  | 

### Return type

[**ArduinoDevicev2propertyvalues**](ArduinoDevicev2propertyvalues.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.devicev2propertyvalues+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2Update

> ArduinoDevicev2 DevicesV2Update(ctx, id).Devicev2(devicev2).XOrganization(xOrganization).Execute()

update devices_v2



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
	id := "id_example" // string | The id of the device
	devicev2 := *openapiclient.NewDevicev2() // Devicev2 | DeviceV2 describes a device.
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesV2API.DevicesV2Update(context.Background(), id).Devicev2(devicev2).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesV2API.DevicesV2Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DevicesV2Update`: ArduinoDevicev2
	fmt.Fprintf(os.Stdout, "Response from `DevicesV2API.DevicesV2Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesV2UpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **devicev2** | [**Devicev2**](Devicev2.md) | DeviceV2 describes a device. | 
 **xOrganization** | **string** |  | 

### Return type

[**ArduinoDevicev2**](ArduinoDevicev2.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.arduino.devicev2+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2UpdateProperties

> DevicesV2UpdateProperties(ctx, id).PropertiesValues(propertiesValues).XOrganization(xOrganization).Execute()

updateProperties devices_v2



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
	id := "id_example" // string | The id of the device
	propertiesValues := *openapiclient.NewPropertiesValues([]openapiclient.PropertiesValue{*openapiclient.NewPropertiesValue("Name_example", "Type_example", interface{}(123))}) // PropertiesValues | 
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevicesV2API.DevicesV2UpdateProperties(context.Background(), id).PropertiesValues(propertiesValues).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesV2API.DevicesV2UpdateProperties``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDevicesV2UpdatePropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **propertiesValues** | [**PropertiesValues**](PropertiesValues.md) |  | 
 **xOrganization** | **string** |  | 

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

