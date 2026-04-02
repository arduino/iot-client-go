# \DashboardsV3API

All URIs are relative to *https://api2.arduino.cc*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DashboardsV3Clone**](DashboardsV3API.md#DashboardsV3Clone) | **Put** /iot/v3/dashboards/{id}/clone | clone dashboards_v3
[**DashboardsV3Create**](DashboardsV3API.md#DashboardsV3Create) | **Post** /iot/v3/dashboards | create dashboards_v3
[**DashboardsV3List**](DashboardsV3API.md#DashboardsV3List) | **Get** /iot/v3/dashboards | list dashboards_v3
[**DashboardsV3Patch**](DashboardsV3API.md#DashboardsV3Patch) | **Patch** /iot/v3/dashboards/{id} | patch dashboards_v3
[**DashboardsV3Show**](DashboardsV3API.md#DashboardsV3Show) | **Get** /iot/v3/dashboards/{id} | show dashboards_v3
[**DashboardsV3Template**](DashboardsV3API.md#DashboardsV3Template) | **Get** /iot/v3/dashboards/{id}/template | template dashboards_v3
[**DashboardsV3Update**](DashboardsV3API.md#DashboardsV3Update) | **Put** /iot/v3/dashboards/{id} | update dashboards_v3



## DashboardsV3Clone

> ArduinoDashboardv3 DashboardsV3Clone(ctx, id).Clone(clone).XOrganization(xOrganization).Execute()

clone dashboards_v3



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/arduino/iot-client-go/v3"
)

func main() {
	id := "id_example" // string | The id of the dashboard
	clone := *openapiclient.NewClone() // Clone | 
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsV3API.DashboardsV3Clone(context.Background(), id).Clone(clone).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsV3API.DashboardsV3Clone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DashboardsV3Clone`: ArduinoDashboardv3
	fmt.Fprintf(os.Stdout, "Response from `DashboardsV3API.DashboardsV3Clone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the dashboard | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardsV3CloneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clone** | [**Clone**](Clone.md) |  | 
 **xOrganization** | **string** |  | 

### Return type

[**ArduinoDashboardv3**](ArduinoDashboardv3.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.arduino.dashboardv3+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardsV3Create

> ArduinoDashboardv3 DashboardsV3Create(ctx).Dashboardv3(dashboardv3).XOrganization(xOrganization).Execute()

create dashboards_v3



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/arduino/iot-client-go/v3"
)

func main() {
	dashboardv3 := *openapiclient.NewDashboardv3() // Dashboardv3 | DashboardV3Payload describes a dashboard
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsV3API.DashboardsV3Create(context.Background()).Dashboardv3(dashboardv3).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsV3API.DashboardsV3Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DashboardsV3Create`: ArduinoDashboardv3
	fmt.Fprintf(os.Stdout, "Response from `DashboardsV3API.DashboardsV3Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDashboardsV3CreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dashboardv3** | [**Dashboardv3**](Dashboardv3.md) | DashboardV3Payload describes a dashboard | 
 **xOrganization** | **string** |  | 

### Return type

[**ArduinoDashboardv3**](ArduinoDashboardv3.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.arduino.dashboardv3+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardsV3List

> []ArduinoDashboardv3 DashboardsV3List(ctx).Name(name).ThingId(thingId).UserId(userId).XOrganization(xOrganization).Execute()

list dashboards_v3



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/arduino/iot-client-go/v3"
)

func main() {
	name := "name_example" // string | The name of the dashboard (optional)
	thingId := "thingId_example" // string | The thing_id of the dashboard's properties (optional)
	userId := "userId_example" // string | The user_id of the dashboard's owner (optional)
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsV3API.DashboardsV3List(context.Background()).Name(name).ThingId(thingId).UserId(userId).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsV3API.DashboardsV3List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DashboardsV3List`: []ArduinoDashboardv3
	fmt.Fprintf(os.Stdout, "Response from `DashboardsV3API.DashboardsV3List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDashboardsV3ListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name of the dashboard | 
 **thingId** | **string** | The thing_id of the dashboard&#39;s properties | 
 **userId** | **string** | The user_id of the dashboard&#39;s owner | 
 **xOrganization** | **string** |  | 

### Return type

[**[]ArduinoDashboardv3**](ArduinoDashboardv3.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.dashboardv3+json; type=collection, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardsV3Patch

> ArduinoDashboardv3 DashboardsV3Patch(ctx, id).Dashboardv3(dashboardv3).XOrganization(xOrganization).Execute()

patch dashboards_v3



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/arduino/iot-client-go/v3"
)

func main() {
	id := "id_example" // string | The id of the dashboard
	dashboardv3 := *openapiclient.NewDashboardv3() // Dashboardv3 | DashboardV3Payload describes a dashboard
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsV3API.DashboardsV3Patch(context.Background(), id).Dashboardv3(dashboardv3).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsV3API.DashboardsV3Patch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DashboardsV3Patch`: ArduinoDashboardv3
	fmt.Fprintf(os.Stdout, "Response from `DashboardsV3API.DashboardsV3Patch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the dashboard | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardsV3PatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dashboardv3** | [**Dashboardv3**](Dashboardv3.md) | DashboardV3Payload describes a dashboard | 
 **xOrganization** | **string** |  | 

### Return type

[**ArduinoDashboardv3**](ArduinoDashboardv3.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.arduino.dashboardv3+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardsV3Show

> ArduinoDashboardv3 DashboardsV3Show(ctx, id).XOrganization(xOrganization).Execute()

show dashboards_v3



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/arduino/iot-client-go/v3"
)

func main() {
	id := "id_example" // string | The id of the dashboard
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsV3API.DashboardsV3Show(context.Background(), id).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsV3API.DashboardsV3Show``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DashboardsV3Show`: ArduinoDashboardv3
	fmt.Fprintf(os.Stdout, "Response from `DashboardsV3API.DashboardsV3Show`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the dashboard | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardsV3ShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrganization** | **string** |  | 

### Return type

[**ArduinoDashboardv3**](ArduinoDashboardv3.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.dashboardv3+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardsV3Template

> ArduinoDashboardv3template DashboardsV3Template(ctx, id).XOrganization(xOrganization).Execute()

template dashboards_v3



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/arduino/iot-client-go/v3"
)

func main() {
	id := "id_example" // string | The id of the dashboard
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsV3API.DashboardsV3Template(context.Background(), id).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsV3API.DashboardsV3Template``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DashboardsV3Template`: ArduinoDashboardv3template
	fmt.Fprintf(os.Stdout, "Response from `DashboardsV3API.DashboardsV3Template`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the dashboard | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardsV3TemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrganization** | **string** |  | 

### Return type

[**ArduinoDashboardv3template**](ArduinoDashboardv3template.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.dashboardv3template+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardsV3Update

> ArduinoDashboardv3 DashboardsV3Update(ctx, id).Dashboardv3(dashboardv3).XOrganization(xOrganization).Execute()

update dashboards_v3



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/arduino/iot-client-go/v3"
)

func main() {
	id := "id_example" // string | The id of the dashboard
	dashboardv3 := *openapiclient.NewDashboardv3() // Dashboardv3 | DashboardV3Payload describes a dashboard
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsV3API.DashboardsV3Update(context.Background(), id).Dashboardv3(dashboardv3).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsV3API.DashboardsV3Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DashboardsV3Update`: ArduinoDashboardv3
	fmt.Fprintf(os.Stdout, "Response from `DashboardsV3API.DashboardsV3Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the dashboard | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardsV3UpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dashboardv3** | [**Dashboardv3**](Dashboardv3.md) | DashboardV3Payload describes a dashboard | 
 **xOrganization** | **string** |  | 

### Return type

[**ArduinoDashboardv3**](ArduinoDashboardv3.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.arduino.dashboardv3+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

