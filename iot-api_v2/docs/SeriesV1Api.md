# \SeriesV1Api

All URIs are relative to *http://api-dev.arduino.cc*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SeriesV1BatchQuery**](SeriesV1Api.md#SeriesV1BatchQuery) | **Post** /iot/v1/series/batch_query | batch_query series_v1
[**SeriesV1BatchQueryRaw**](SeriesV1Api.md#SeriesV1BatchQueryRaw) | **Post** /iot/v1/series/batch_query_raw | batch_query_raw series_v1



## SeriesV1BatchQuery

> ArduinoSeriesBatch SeriesV1BatchQuery(ctx, batchQueryRequestsMediaV1)
batch_query series_v1

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


## SeriesV1BatchQueryRaw

> ArduinoSeriesRawBatch SeriesV1BatchQueryRaw(ctx, batchQueryRawRequestsMediaV1)
batch_query_raw series_v1

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

