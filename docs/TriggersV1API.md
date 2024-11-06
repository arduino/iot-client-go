# \TriggersV1API

All URIs are relative to *https://api2.arduino.cc*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionsV1Create**](TriggersV1API.md#ActionsV1Create) | **Post** /iot/v1/actions | create actions_v1
[**ActionsV1Delete**](TriggersV1API.md#ActionsV1Delete) | **Delete** /iot/v1/actions/{id} | delete actions_v1
[**ActionsV1List**](TriggersV1API.md#ActionsV1List) | **Get** /iot/v1/actions | list actions_v1
[**ActionsV1Show**](TriggersV1API.md#ActionsV1Show) | **Get** /iot/v1/actions/{id} | show actions_v1
[**ActionsV1Update**](TriggersV1API.md#ActionsV1Update) | **Put** /iot/v1/actions/{id} | update actions_v1
[**TriggersV1Create**](TriggersV1API.md#TriggersV1Create) | **Put** /iot/v1/triggers | create triggers_v1
[**TriggersV1Delete**](TriggersV1API.md#TriggersV1Delete) | **Delete** /iot/v1/triggers/{id} | delete triggers_v1
[**TriggersV1List**](TriggersV1API.md#TriggersV1List) | **Get** /iot/v1/triggers | list triggers_v1
[**TriggersV1Patch**](TriggersV1API.md#TriggersV1Patch) | **Patch** /iot/v1/triggers/{id} | patch triggers_v1
[**TriggersV1Show**](TriggersV1API.md#TriggersV1Show) | **Get** /iot/v1/triggers/{id} | show triggers_v1
[**TriggersV1Template**](TriggersV1API.md#TriggersV1Template) | **Get** /iot/v1/triggers/{id}/template | template triggers_v1
[**TriggersV1Update**](TriggersV1API.md#TriggersV1Update) | **Post** /iot/v1/triggers/{id} | update triggers_v1



## ActionsV1Create

> ArduinoAction ActionsV1Create(ctx).CreateAction(createAction).XOrganization(xOrganization).Execute()

create actions_v1



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
	createAction := *openapiclient.NewCreateAction("NOTIFY-EMAIL", "Name_example") // CreateAction | 
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TriggersV1API.ActionsV1Create(context.Background()).CreateAction(createAction).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggersV1API.ActionsV1Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionsV1Create`: ArduinoAction
	fmt.Fprintf(os.Stdout, "Response from `TriggersV1API.ActionsV1Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActionsV1CreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAction** | [**CreateAction**](CreateAction.md) |  | 
 **xOrganization** | **string** |  | 

### Return type

[**ArduinoAction**](ArduinoAction.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.arduino.action+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsV1Delete

> ActionsV1Delete(ctx, id).XOrganization(xOrganization).Execute()

delete actions_v1



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
	id := "id_example" // string | The id of the action
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TriggersV1API.ActionsV1Delete(context.Background(), id).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggersV1API.ActionsV1Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the action | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsV1DeleteRequest struct via the builder pattern


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


## ActionsV1List

> []ArduinoAction ActionsV1List(ctx).XOrganization(xOrganization).Execute()

list actions_v1



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
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TriggersV1API.ActionsV1List(context.Background()).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggersV1API.ActionsV1List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionsV1List`: []ArduinoAction
	fmt.Fprintf(os.Stdout, "Response from `TriggersV1API.ActionsV1List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActionsV1ListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xOrganization** | **string** |  | 

### Return type

[**[]ArduinoAction**](ArduinoAction.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.action+json; type=collection, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsV1Show

> ArduinoAction ActionsV1Show(ctx, id).XOrganization(xOrganization).Execute()

show actions_v1



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
	id := "id_example" // string | The id of the action
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TriggersV1API.ActionsV1Show(context.Background(), id).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggersV1API.ActionsV1Show``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionsV1Show`: ArduinoAction
	fmt.Fprintf(os.Stdout, "Response from `TriggersV1API.ActionsV1Show`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the action | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsV1ShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrganization** | **string** |  | 

### Return type

[**ArduinoAction**](ArduinoAction.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.action+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsV1Update

> ArduinoAction ActionsV1Update(ctx, id).UpdateAction(updateAction).XOrganization(xOrganization).Execute()

update actions_v1



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
	id := "id_example" // string | The id of the action
	updateAction := *openapiclient.NewUpdateAction() // UpdateAction | 
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TriggersV1API.ActionsV1Update(context.Background(), id).UpdateAction(updateAction).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggersV1API.ActionsV1Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionsV1Update`: ArduinoAction
	fmt.Fprintf(os.Stdout, "Response from `TriggersV1API.ActionsV1Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the action | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsV1UpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAction** | [**UpdateAction**](UpdateAction.md) |  | 
 **xOrganization** | **string** |  | 

### Return type

[**ArduinoAction**](ArduinoAction.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.arduino.action+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggersV1Create

> ArduinoTrigger TriggersV1Create(ctx).Trigger(trigger).XOrganization(xOrganization).Execute()

create triggers_v1



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
	trigger := *openapiclient.NewTrigger() // Trigger | 
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TriggersV1API.TriggersV1Create(context.Background()).Trigger(trigger).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggersV1API.TriggersV1Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TriggersV1Create`: ArduinoTrigger
	fmt.Fprintf(os.Stdout, "Response from `TriggersV1API.TriggersV1Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTriggersV1CreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trigger** | [**Trigger**](Trigger.md) |  | 
 **xOrganization** | **string** |  | 

### Return type

[**ArduinoTrigger**](ArduinoTrigger.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.arduino.trigger+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggersV1Delete

> TriggersV1Delete(ctx, id).Force(force).XOrganization(xOrganization).Execute()

delete triggers_v1



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
	id := "id_example" // string | The id of the trigger
	force := true // bool | If true, hard delete the trigger (optional) (default to false)
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TriggersV1API.TriggersV1Delete(context.Background(), id).Force(force).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggersV1API.TriggersV1Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the trigger | 

### Other Parameters

Other parameters are passed through a pointer to a apiTriggersV1DeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | If true, hard delete the trigger | [default to false]
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


## TriggersV1List

> []ArduinoTrigger TriggersV1List(ctx).DeviceId(deviceId).PropertyId(propertyId).ShowDeleted(showDeleted).SourceType(sourceType).XOrganization(xOrganization).Execute()

list triggers_v1



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
	deviceId := "deviceId_example" // string | The id of the device associated with the triggers (mutually exclusive with 'property_id') (optional)
	propertyId := "propertyId_example" // string | The id of the property associated with the triggers (mutually exclusive with 'device_id') (optional)
	showDeleted := true // bool | If true, shows the soft deleted triggers (optional) (default to false)
	sourceType := "sourceType_example" // string | The source type of the trigger, could be PROPERTY, DEVICE_INCLUDE or DEVICE_EXCLUDE (optional)
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TriggersV1API.TriggersV1List(context.Background()).DeviceId(deviceId).PropertyId(propertyId).ShowDeleted(showDeleted).SourceType(sourceType).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggersV1API.TriggersV1List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TriggersV1List`: []ArduinoTrigger
	fmt.Fprintf(os.Stdout, "Response from `TriggersV1API.TriggersV1List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTriggersV1ListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | The id of the device associated with the triggers (mutually exclusive with &#39;property_id&#39;) | 
 **propertyId** | **string** | The id of the property associated with the triggers (mutually exclusive with &#39;device_id&#39;) | 
 **showDeleted** | **bool** | If true, shows the soft deleted triggers | [default to false]
 **sourceType** | **string** | The source type of the trigger, could be PROPERTY, DEVICE_INCLUDE or DEVICE_EXCLUDE | 
 **xOrganization** | **string** |  | 

### Return type

[**[]ArduinoTrigger**](ArduinoTrigger.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.trigger+json; type=collection, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggersV1Patch

> ArduinoTrigger TriggersV1Patch(ctx, id).Trigger(trigger).XOrganization(xOrganization).Execute()

patch triggers_v1



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
	id := "id_example" // string | The id of the trigger
	trigger := *openapiclient.NewTrigger() // Trigger | 
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TriggersV1API.TriggersV1Patch(context.Background(), id).Trigger(trigger).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggersV1API.TriggersV1Patch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TriggersV1Patch`: ArduinoTrigger
	fmt.Fprintf(os.Stdout, "Response from `TriggersV1API.TriggersV1Patch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the trigger | 

### Other Parameters

Other parameters are passed through a pointer to a apiTriggersV1PatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trigger** | [**Trigger**](Trigger.md) |  | 
 **xOrganization** | **string** |  | 

### Return type

[**ArduinoTrigger**](ArduinoTrigger.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.arduino.trigger+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggersV1Show

> ArduinoTriggerWithLinkedEntities TriggersV1Show(ctx, id).XOrganization(xOrganization).Execute()

show triggers_v1



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
	id := "id_example" // string | The id of the trigger
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TriggersV1API.TriggersV1Show(context.Background(), id).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggersV1API.TriggersV1Show``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TriggersV1Show`: ArduinoTriggerWithLinkedEntities
	fmt.Fprintf(os.Stdout, "Response from `TriggersV1API.TriggersV1Show`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the trigger | 

### Other Parameters

Other parameters are passed through a pointer to a apiTriggersV1ShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrganization** | **string** |  | 

### Return type

[**ArduinoTriggerWithLinkedEntities**](ArduinoTriggerWithLinkedEntities.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.trigger_with_linked_entities+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggersV1Template

> ArduinoTriggerTemplate TriggersV1Template(ctx, id).XOrganization(xOrganization).Execute()

template triggers_v1



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
	id := "id_example" // string | The id of the trigger
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TriggersV1API.TriggersV1Template(context.Background(), id).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggersV1API.TriggersV1Template``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TriggersV1Template`: ArduinoTriggerTemplate
	fmt.Fprintf(os.Stdout, "Response from `TriggersV1API.TriggersV1Template`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the trigger | 

### Other Parameters

Other parameters are passed through a pointer to a apiTriggersV1TemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrganization** | **string** |  | 

### Return type

[**ArduinoTriggerTemplate**](ArduinoTriggerTemplate.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.trigger_template+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggersV1Update

> ArduinoTrigger TriggersV1Update(ctx, id).Trigger(trigger).XOrganization(xOrganization).Execute()

update triggers_v1



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
	id := "id_example" // string | The id of the trigger
	trigger := *openapiclient.NewTrigger() // Trigger | 
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TriggersV1API.TriggersV1Update(context.Background(), id).Trigger(trigger).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggersV1API.TriggersV1Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TriggersV1Update`: ArduinoTrigger
	fmt.Fprintf(os.Stdout, "Response from `TriggersV1API.TriggersV1Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the trigger | 

### Other Parameters

Other parameters are passed through a pointer to a apiTriggersV1UpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trigger** | [**Trigger**](Trigger.md) |  | 
 **xOrganization** | **string** |  | 

### Return type

[**ArduinoTrigger**](ArduinoTrigger.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.arduino.trigger+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

