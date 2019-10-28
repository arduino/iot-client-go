# \SeriesV2Api

All URIs are relative to *http://api2.arduino.cc/iot*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SeriesV2BatchQuery**](SeriesV2Api.md#SeriesV2BatchQuery) | **Post** /v2/series/batch_query | batch_query series_v2
[**SeriesV2BatchQueryRaw**](SeriesV2Api.md#SeriesV2BatchQueryRaw) | **Post** /v2/series/batch_query_raw | batch_query_raw series_v2
[**SeriesV2BatchQueryRawLastValue**](SeriesV2Api.md#SeriesV2BatchQueryRawLastValue) | **Post** /v2/series/batch_query_raw/lastvalue | batch_query_raw_last_value series_v2



## SeriesV2BatchQuery

> ArduinoSeriesBatch SeriesV2BatchQuery(ctx, batchQueryRequestsMediaV1)

batch_query series_v2

Returns the batch of time-series data

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchQueryRequestsMediaV1** | [**BatchQueryRequestsMediaV1**](BatchQueryRequestsMediaV1.md)|  | 

### Return type

[**ArduinoSeriesBatch**](ArduinoSeriesBatch.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.arduino.series.batch+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SeriesV2BatchQueryRaw

> ArduinoSeriesRawBatch SeriesV2BatchQueryRaw(ctx, batchQueryRawRequestsMediaV1)

batch_query_raw series_v2

Returns the batch of time-series data raw

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchQueryRawRequestsMediaV1** | [**BatchQueryRawRequestsMediaV1**](BatchQueryRawRequestsMediaV1.md)|  | 

### Return type

[**ArduinoSeriesRawBatch**](ArduinoSeriesRawBatch.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.arduino.series.raw.batch+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SeriesV2BatchQueryRawLastValue

> ArduinoSeriesRawBatchLastvalue SeriesV2BatchQueryRawLastValue(ctx, batchLastValueRequestsMediaV1)

batch_query_raw_last_value series_v2

Returns the batch of time-series data raw

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchLastValueRequestsMediaV1** | [**BatchLastValueRequestsMediaV1**](BatchLastValueRequestsMediaV1.md)|  | 

### Return type

[**ArduinoSeriesRawBatchLastvalue**](ArduinoSeriesRawBatchLastvalue.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.arduino.series.raw.batch.lastvalue+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

