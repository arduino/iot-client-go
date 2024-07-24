# \TemplatesApi

All URIs are relative to *https://api2.arduino.cc/iot*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TemplatesApply**](TemplatesApi.md#TemplatesApply) | **Put** /v1/templates | apply templates



## TemplatesApply

> ArduinoTemplate TemplatesApply(ctx).Template(template).XOrganization(xOrganization).Execute()

apply templates



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
    template := *openapiclient.NewTemplate("remote-controlled-lights") // Template | TemplatePayload describes the needed attribute to apply a template
    xOrganization := "xOrganization_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplatesApply(context.Background()).Template(template).XOrganization(xOrganization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplatesApply``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplatesApply`: ArduinoTemplate
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplatesApply`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesApplyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **template** | [**Template**](Template.md) | TemplatePayload describes the needed attribute to apply a template | 
 **xOrganization** | **string** |  | 

### Return type

[**ArduinoTemplate**](ArduinoTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.arduino.template+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

