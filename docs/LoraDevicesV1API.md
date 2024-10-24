# \LoraDevicesV1API

All URIs are relative to *https://api2.arduino.cc*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LoraDevicesV1Create**](LoraDevicesV1API.md#LoraDevicesV1Create) | **Put** /iot/v1/lora-devices/ | create lora_devices_v1



## LoraDevicesV1Create

> ArduinoLoradevicev1 LoraDevicesV1Create(ctx).CreateLoraDevicesV1Payload(createLoraDevicesV1Payload).XOrganization(xOrganization).Execute()

create lora_devices_v1



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
	createLoraDevicesV1Payload := *openapiclient.NewCreateLoraDevicesV1Payload("App_example", "Eui_example", "FrequencyPlan_example", "Name_example", "Type_example", "UserId_example") // CreateLoraDevicesV1Payload | 
	xOrganization := "xOrganization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoraDevicesV1API.LoraDevicesV1Create(context.Background()).CreateLoraDevicesV1Payload(createLoraDevicesV1Payload).XOrganization(xOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoraDevicesV1API.LoraDevicesV1Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoraDevicesV1Create`: ArduinoLoradevicev1
	fmt.Fprintf(os.Stdout, "Response from `LoraDevicesV1API.LoraDevicesV1Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoraDevicesV1CreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLoraDevicesV1Payload** | [**CreateLoraDevicesV1Payload**](CreateLoraDevicesV1Payload.md) |  | 
 **xOrganization** | **string** |  | 

### Return type

[**ArduinoLoradevicev1**](ArduinoLoradevicev1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.arduino.loradevicev1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

