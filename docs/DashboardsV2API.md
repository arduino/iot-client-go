# \DashboardsV2API

All URIs are relative to *https://api2.arduino.cc/iot*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DashboardsV2Clone**](DashboardsV2API.md#DashboardsV2Clone) | **Put** /v2/dashboards/{id}/clone | clone dashboards_v2
[**DashboardsV2Create**](DashboardsV2API.md#DashboardsV2Create) | **Post** /v2/dashboards | create dashboards_v2
[**DashboardsV2Delete**](DashboardsV2API.md#DashboardsV2Delete) | **Delete** /v2/dashboards/{id} | delete dashboards_v2
[**DashboardsV2DeleteShare**](DashboardsV2API.md#DashboardsV2DeleteShare) | **Delete** /v2/dashboards/{id}/shares/{user_id} | deleteShare dashboards_v2
[**DashboardsV2Link**](DashboardsV2API.md#DashboardsV2Link) | **Put** /v2/dashboards/{id}/widgets/{widgetId}/variables | link dashboards_v2
[**DashboardsV2List**](DashboardsV2API.md#DashboardsV2List) | **Get** /v2/dashboards | list dashboards_v2
[**DashboardsV2ListShares**](DashboardsV2API.md#DashboardsV2ListShares) | **Get** /v2/dashboards/{id}/shares | listShares dashboards_v2
[**DashboardsV2RequestAccess**](DashboardsV2API.md#DashboardsV2RequestAccess) | **Put** /v2/dashboards/{id}/share_request | requestAccess dashboards_v2
[**DashboardsV2Share**](DashboardsV2API.md#DashboardsV2Share) | **Put** /v2/dashboards/{id}/shares | share dashboards_v2
[**DashboardsV2Show**](DashboardsV2API.md#DashboardsV2Show) | **Get** /v2/dashboards/{id} | show dashboards_v2
[**DashboardsV2Template**](DashboardsV2API.md#DashboardsV2Template) | **Get** /v2/dashboards/{id}/template | template dashboards_v2
[**DashboardsV2Update**](DashboardsV2API.md#DashboardsV2Update) | **Put** /v2/dashboards/{id} | update dashboards_v2



## DashboardsV2Clone

> ArduinoDashboardv2 DashboardsV2Clone(ctx, id).Clone(clone).XOrganization(xOrganization).Execute()

clone dashboards_v2



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
	id := "id_example" // string | The id of the dashboard
	clone := *openapiclient.NewClone() // Clone | 
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsV2API.DashboardsV2Clone(context.Background(), id).Clone(clone).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsV2API.DashboardsV2Clone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DashboardsV2Clone`: ArduinoDashboardv2
	fmt.Fprintf(os.Stdout, "Response from `DashboardsV2API.DashboardsV2Clone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the dashboard | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardsV2CloneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clone** | [**Clone**](Clone.md) |  | 
 **xOrganization** | **string** |  | 

### Return type

[**ArduinoDashboardv2**](ArduinoDashboardv2.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.arduino.dashboardv2+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardsV2Create

> ArduinoDashboardv2 DashboardsV2Create(ctx).Dashboardv2(dashboardv2).XOrganization(xOrganization).Execute()

create dashboards_v2



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
	dashboardv2 := *openapiclient.NewDashboardv2() // Dashboardv2 | DashboardV2Payload describes a dashboard
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsV2API.DashboardsV2Create(context.Background()).Dashboardv2(dashboardv2).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsV2API.DashboardsV2Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DashboardsV2Create`: ArduinoDashboardv2
	fmt.Fprintf(os.Stdout, "Response from `DashboardsV2API.DashboardsV2Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDashboardsV2CreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dashboardv2** | [**Dashboardv2**](Dashboardv2.md) | DashboardV2Payload describes a dashboard | 
 **xOrganization** | **string** |  | 

### Return type

[**ArduinoDashboardv2**](ArduinoDashboardv2.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.arduino.dashboardv2+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardsV2Delete

> DashboardsV2Delete(ctx, id).XOrganization(xOrganization).Execute()

delete dashboards_v2



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
	id := "id_example" // string | The id of the dashboard
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DashboardsV2API.DashboardsV2Delete(context.Background(), id).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsV2API.DashboardsV2Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the dashboard | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardsV2DeleteRequest struct via the builder pattern


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


## DashboardsV2DeleteShare

> DashboardsV2DeleteShare(ctx, id, userId).XOrganization(xOrganization).Execute()

deleteShare dashboards_v2



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
	id := "id_example" // string | The id of the dashboard
	userId := "userId_example" // string | The id of the user
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DashboardsV2API.DashboardsV2DeleteShare(context.Background(), id, userId).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsV2API.DashboardsV2DeleteShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the dashboard | 
**userId** | **string** | The id of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardsV2DeleteShareRequest struct via the builder pattern


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


## DashboardsV2Link

> ArduinoVariableslinks DashboardsV2Link(ctx, id, widgetId).Widgetlink(widgetlink).XOrganization(xOrganization).Execute()

link dashboards_v2



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
	id := "id_example" // string | The id of the dashboard
	widgetId := "widgetId_example" // string | The id of the widget
	widgetlink := *openapiclient.NewWidgetlink() // Widgetlink | 
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsV2API.DashboardsV2Link(context.Background(), id, widgetId).Widgetlink(widgetlink).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsV2API.DashboardsV2Link``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DashboardsV2Link`: ArduinoVariableslinks
	fmt.Fprintf(os.Stdout, "Response from `DashboardsV2API.DashboardsV2Link`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the dashboard | 
**widgetId** | **string** | The id of the widget | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardsV2LinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **widgetlink** | [**Widgetlink**](Widgetlink.md) |  | 
 **xOrganization** | **string** |  | 

### Return type

[**ArduinoVariableslinks**](ArduinoVariableslinks.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.arduino.variableslinks+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardsV2List

> []ArduinoDashboardv2 DashboardsV2List(ctx).Name(name).UserId(userId).XOrganization(xOrganization).Execute()

list dashboards_v2



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
	name := "name_example" // string | The name of the dashboard (optional)
	userId := "userId_example" // string | The user_id of the dashboard's owner (optional)
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsV2API.DashboardsV2List(context.Background()).Name(name).UserId(userId).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsV2API.DashboardsV2List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DashboardsV2List`: []ArduinoDashboardv2
	fmt.Fprintf(os.Stdout, "Response from `DashboardsV2API.DashboardsV2List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDashboardsV2ListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name of the dashboard | 
 **userId** | **string** | The user_id of the dashboard&#39;s owner | 
 **xOrganization** | **string** |  | 

### Return type

[**[]ArduinoDashboardv2**](ArduinoDashboardv2.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.dashboardv2+json; type=collection, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardsV2ListShares

> []ArduinoDashboardshare DashboardsV2ListShares(ctx, id).XOrganization(xOrganization).Execute()

listShares dashboards_v2



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
	id := "id_example" // string | The id of the dashboard
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsV2API.DashboardsV2ListShares(context.Background(), id).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsV2API.DashboardsV2ListShares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DashboardsV2ListShares`: []ArduinoDashboardshare
	fmt.Fprintf(os.Stdout, "Response from `DashboardsV2API.DashboardsV2ListShares`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the dashboard | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardsV2ListSharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrganization** | **string** |  | 

### Return type

[**[]ArduinoDashboardshare**](ArduinoDashboardshare.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.dashboardshare+json; type=collection, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardsV2RequestAccess

> DashboardsV2RequestAccess(ctx, id).Sharerequest(sharerequest).XOrganization(xOrganization).Execute()

requestAccess dashboards_v2



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
	id := "id_example" // string | The id of the dashboard
	sharerequest := *openapiclient.NewSharerequest() // Sharerequest | 
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DashboardsV2API.DashboardsV2RequestAccess(context.Background(), id).Sharerequest(sharerequest).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsV2API.DashboardsV2RequestAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the dashboard | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardsV2RequestAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sharerequest** | [**Sharerequest**](Sharerequest.md) |  | 
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


## DashboardsV2Share

> DashboardsV2Share(ctx, id).Dashboardshare(dashboardshare).XOrganization(xOrganization).Execute()

share dashboards_v2



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
	id := "id_example" // string | The id of the dashboard
	dashboardshare := *openapiclient.NewDashboardshare() // Dashboardshare | 
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DashboardsV2API.DashboardsV2Share(context.Background(), id).Dashboardshare(dashboardshare).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsV2API.DashboardsV2Share``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the dashboard | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardsV2ShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dashboardshare** | [**Dashboardshare**](Dashboardshare.md) |  | 
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


## DashboardsV2Show

> ArduinoDashboardv2 DashboardsV2Show(ctx, id).XOrganization(xOrganization).Execute()

show dashboards_v2



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
	id := "id_example" // string | The id of the dashboard
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsV2API.DashboardsV2Show(context.Background(), id).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsV2API.DashboardsV2Show``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DashboardsV2Show`: ArduinoDashboardv2
	fmt.Fprintf(os.Stdout, "Response from `DashboardsV2API.DashboardsV2Show`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the dashboard | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardsV2ShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrganization** | **string** |  | 

### Return type

[**ArduinoDashboardv2**](ArduinoDashboardv2.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.dashboardv2+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardsV2Template

> ArduinoDashboardv2template DashboardsV2Template(ctx, id).XOrganization(xOrganization).Execute()

template dashboards_v2



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
	id := "id_example" // string | The id of the dashboard
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsV2API.DashboardsV2Template(context.Background(), id).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsV2API.DashboardsV2Template``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DashboardsV2Template`: ArduinoDashboardv2template
	fmt.Fprintf(os.Stdout, "Response from `DashboardsV2API.DashboardsV2Template`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the dashboard | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardsV2TemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrganization** | **string** |  | 

### Return type

[**ArduinoDashboardv2template**](ArduinoDashboardv2template.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.dashboardv2template+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardsV2Update

> ArduinoDashboardv2 DashboardsV2Update(ctx, id).Dashboardv2(dashboardv2).XOrganization(xOrganization).Execute()

update dashboards_v2



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
	id := "id_example" // string | The id of the dashboard
	dashboardv2 := *openapiclient.NewDashboardv2() // Dashboardv2 | DashboardV2Payload describes a dashboard
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsV2API.DashboardsV2Update(context.Background(), id).Dashboardv2(dashboardv2).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsV2API.DashboardsV2Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DashboardsV2Update`: ArduinoDashboardv2
	fmt.Fprintf(os.Stdout, "Response from `DashboardsV2API.DashboardsV2Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the dashboard | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardsV2UpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dashboardv2** | [**Dashboardv2**](Dashboardv2.md) | DashboardV2Payload describes a dashboard | 
 **xOrganization** | **string** |  | 

### Return type

[**ArduinoDashboardv2**](ArduinoDashboardv2.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.arduino.dashboardv2+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

