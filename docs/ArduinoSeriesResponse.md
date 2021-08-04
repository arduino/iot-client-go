# ArduinoSeriesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountValues** | **int64** | Total number of values in the array &#39;values&#39; | 
**FromDate** | [**time.Time**](time.Time.md) | From date | 
**Interval** | **int64** | Resolution in seconds | 
**Message** | **string** | If the response is different than &#39;ok&#39; | [optional] [default to ""]
**Query** | **string** | Query of for the data | 
**RespVersion** | **int64** | Response version | 
**SeriesLimit** | **int64** | Max of values | [optional] 
**Status** | **string** | Status of the response | 
**Times** | [**[]time.Time**](time.Time.md) | Timestamp in RFC3339 | 
**ToDate** | [**time.Time**](time.Time.md) | To date | 
**Values** | **[]float64** | Values in Float | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


