# \SeriesV2Api

All URIs are relative to *https://api2.arduino.cc/iot*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SeriesV2BatchQuery**](SeriesV2Api.md#SeriesV2BatchQuery) | **Post** /v2/series/batch_query | batch_query series_v2
[**SeriesV2BatchQueryRaw**](SeriesV2Api.md#SeriesV2BatchQueryRaw) | **Post** /v2/series/batch_query_raw | batch_query_raw series_v2
[**SeriesV2BatchQueryRawLastValue**](SeriesV2Api.md#SeriesV2BatchQueryRawLastValue) | **Post** /v2/series/batch_query_raw/lastvalue | batch_query_raw_last_value series_v2
[**SeriesV2BatchQuerySampling**](SeriesV2Api.md#SeriesV2BatchQuerySampling) | **Post** /v2/series/batch_query_sampling | batch_query_sampling series_v2
[**SeriesV2HistoricData**](SeriesV2Api.md#SeriesV2HistoricData) | **Post** /v2/series/historic_data | historic_data series_v2



## SeriesV2BatchQuery

> ArduinoSeriesBatch SeriesV2BatchQuery(ctx).BatchQueryRequestsMediaV1(batchQueryRequestsMediaV1).XOrganization(xOrganization).Execute()

batch_query series_v2



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/arduino/iot-client-go/v2"
)

func main() {
    batchQueryRequestsMediaV1 := *openapiclient.NewBatchQueryRequestsMediaV1([]openapiclient.BatchQueryRequestMediaV1{*openapiclient.NewBatchQueryRequestMediaV1(time.Now(), "Q_example", time.Now())}, int64(123)) // BatchQueryRequestsMediaV1 | 
    xOrganization := "xOrganization_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesV2Api.SeriesV2BatchQuery(context.Background()).BatchQueryRequestsMediaV1(batchQueryRequestsMediaV1).XOrganization(xOrganization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesV2Api.SeriesV2BatchQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SeriesV2BatchQuery`: ArduinoSeriesBatch
    fmt.Fprintf(os.Stdout, "Response from `SeriesV2Api.SeriesV2BatchQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSeriesV2BatchQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchQueryRequestsMediaV1** | [**BatchQueryRequestsMediaV1**](BatchQueryRequestsMediaV1.md) |  | 
 **xOrganization** | **string** |  | 

### Return type

[**ArduinoSeriesBatch**](ArduinoSeriesBatch.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.arduino.series.batch+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SeriesV2BatchQueryRaw

> ArduinoSeriesRawBatch SeriesV2BatchQueryRaw(ctx).BatchQueryRawRequestsMediaV1(batchQueryRawRequestsMediaV1).XOrganization(xOrganization).Execute()

batch_query_raw series_v2



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
    batchQueryRawRequestsMediaV1 := *openapiclient.NewBatchQueryRawRequestsMediaV1([]openapiclient.BatchQueryRawRequestMediaV1{*openapiclient.NewBatchQueryRawRequestMediaV1("Q_example")}, int64(123)) // BatchQueryRawRequestsMediaV1 | 
    xOrganization := "xOrganization_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesV2Api.SeriesV2BatchQueryRaw(context.Background()).BatchQueryRawRequestsMediaV1(batchQueryRawRequestsMediaV1).XOrganization(xOrganization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesV2Api.SeriesV2BatchQueryRaw``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SeriesV2BatchQueryRaw`: ArduinoSeriesRawBatch
    fmt.Fprintf(os.Stdout, "Response from `SeriesV2Api.SeriesV2BatchQueryRaw`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSeriesV2BatchQueryRawRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchQueryRawRequestsMediaV1** | [**BatchQueryRawRequestsMediaV1**](BatchQueryRawRequestsMediaV1.md) |  | 
 **xOrganization** | **string** |  | 

### Return type

[**ArduinoSeriesRawBatch**](ArduinoSeriesRawBatch.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.arduino.series.raw.batch+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SeriesV2BatchQueryRawLastValue

> ArduinoSeriesRawBatchLastvalue SeriesV2BatchQueryRawLastValue(ctx).BatchLastValueRequestsMediaV1(batchLastValueRequestsMediaV1).XOrganization(xOrganization).Execute()

batch_query_raw_last_value series_v2



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
    batchLastValueRequestsMediaV1 := *openapiclient.NewBatchLastValueRequestsMediaV1([]openapiclient.BatchQueryRawLastValueRequestMediaV1{*openapiclient.NewBatchQueryRawLastValueRequestMediaV1("PropertyId_example", "ThingId_example")}) // BatchLastValueRequestsMediaV1 | 
    xOrganization := "xOrganization_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesV2Api.SeriesV2BatchQueryRawLastValue(context.Background()).BatchLastValueRequestsMediaV1(batchLastValueRequestsMediaV1).XOrganization(xOrganization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesV2Api.SeriesV2BatchQueryRawLastValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SeriesV2BatchQueryRawLastValue`: ArduinoSeriesRawBatchLastvalue
    fmt.Fprintf(os.Stdout, "Response from `SeriesV2Api.SeriesV2BatchQueryRawLastValue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSeriesV2BatchQueryRawLastValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchLastValueRequestsMediaV1** | [**BatchLastValueRequestsMediaV1**](BatchLastValueRequestsMediaV1.md) |  | 
 **xOrganization** | **string** |  | 

### Return type

[**ArduinoSeriesRawBatchLastvalue**](ArduinoSeriesRawBatchLastvalue.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.arduino.series.raw.batch.lastvalue+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SeriesV2BatchQuerySampling

> ArduinoSeriesBatchSampled SeriesV2BatchQuerySampling(ctx).BatchQuerySampledRequestsMediaV1(batchQuerySampledRequestsMediaV1).XOrganization(xOrganization).Execute()

batch_query_sampling series_v2



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
    batchQuerySampledRequestsMediaV1 := *openapiclient.NewBatchQuerySampledRequestsMediaV1([]openapiclient.BatchQuerySampledRequestMediaV1{*openapiclient.NewBatchQuerySampledRequestMediaV1("Q_example")}, int64(123)) // BatchQuerySampledRequestsMediaV1 | 
    xOrganization := "xOrganization_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesV2Api.SeriesV2BatchQuerySampling(context.Background()).BatchQuerySampledRequestsMediaV1(batchQuerySampledRequestsMediaV1).XOrganization(xOrganization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesV2Api.SeriesV2BatchQuerySampling``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SeriesV2BatchQuerySampling`: ArduinoSeriesBatchSampled
    fmt.Fprintf(os.Stdout, "Response from `SeriesV2Api.SeriesV2BatchQuerySampling`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSeriesV2BatchQuerySamplingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchQuerySampledRequestsMediaV1** | [**BatchQuerySampledRequestsMediaV1**](BatchQuerySampledRequestsMediaV1.md) |  | 
 **xOrganization** | **string** |  | 

### Return type

[**ArduinoSeriesBatchSampled**](ArduinoSeriesBatchSampled.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.arduino.series.batch.sampled+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SeriesV2HistoricData

> SeriesV2HistoricData(ctx).HistoricDataRequest(historicDataRequest).XOrganization(xOrganization).Execute()

historic_data series_v2



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/arduino/iot-client-go/v2"
)

func main() {
    historicDataRequest := *openapiclient.NewHistoricDataRequest(time.Now(), []string{"Properties_example"}, time.Now()) // HistoricDataRequest | 
    xOrganization := "xOrganization_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SeriesV2Api.SeriesV2HistoricData(context.Background()).HistoricDataRequest(historicDataRequest).XOrganization(xOrganization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesV2Api.SeriesV2HistoricData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSeriesV2HistoricDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **historicDataRequest** | [**HistoricDataRequest**](HistoricDataRequest.md) |  | 
 **xOrganization** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

