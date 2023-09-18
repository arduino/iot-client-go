# \LoraFreqPlanV1Api

All URIs are relative to *https://api2.arduino.cc/iot*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LoraFreqPlanV1List**](LoraFreqPlanV1Api.md#LoraFreqPlanV1List) | **Get** /v1/lora-freq-plans/ | list lora_freq_plan_v1



## LoraFreqPlanV1List

> ArduinoLorafreqplansv1 LoraFreqPlanV1List(ctx).Execute()

list lora_freq_plan_v1



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoraFreqPlanV1Api.LoraFreqPlanV1List(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoraFreqPlanV1Api.LoraFreqPlanV1List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoraFreqPlanV1List`: ArduinoLorafreqplansv1
    fmt.Fprintf(os.Stdout, "Response from `LoraFreqPlanV1Api.LoraFreqPlanV1List`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLoraFreqPlanV1ListRequest struct via the builder pattern


### Return type

[**ArduinoLorafreqplansv1**](ArduinoLorafreqplansv1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

