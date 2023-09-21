# ArduinoSeriesRawResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountValues** | **int64** | Total number of values in the array &#39;values&#39; | 
**FromDate** | **time.Time** | From date | 
**Message** | Pointer to **string** | If the response is different than &#39;ok&#39; | [optional] [default to ""]
**Query** | **string** | Query of for the data | 
**RespVersion** | **int64** | Response version | 
**Series** | [**BatchQueryRawResponseSeriesMediaV1**](BatchQueryRawResponseSeriesMediaV1.md) |  | 
**SeriesLimit** | Pointer to **int64** | Max of values | [optional] 
**Sort** | **string** | Sorting | 
**Status** | **string** | Status of the response | 
**Times** | [**[]time.Time**](time.Time.md) | Timestamp in RFC3339 | 
**ToDate** | **time.Time** | To date | 
**Values** | **[]interface{}** | Values can be in Float, String, Bool, Object | 

## Methods

### NewArduinoSeriesRawResponse

`func NewArduinoSeriesRawResponse(countValues int64, fromDate time.Time, query string, respVersion int64, series BatchQueryRawResponseSeriesMediaV1, sort string, status string, times []time.Time, toDate time.Time, values []interface{}, ) *ArduinoSeriesRawResponse`

NewArduinoSeriesRawResponse instantiates a new ArduinoSeriesRawResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoSeriesRawResponseWithDefaults

`func NewArduinoSeriesRawResponseWithDefaults() *ArduinoSeriesRawResponse`

NewArduinoSeriesRawResponseWithDefaults instantiates a new ArduinoSeriesRawResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountValues

`func (o *ArduinoSeriesRawResponse) GetCountValues() int64`

GetCountValues returns the CountValues field if non-nil, zero value otherwise.

### GetCountValuesOk

`func (o *ArduinoSeriesRawResponse) GetCountValuesOk() (*int64, bool)`

GetCountValuesOk returns a tuple with the CountValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountValues

`func (o *ArduinoSeriesRawResponse) SetCountValues(v int64)`

SetCountValues sets CountValues field to given value.


### GetFromDate

`func (o *ArduinoSeriesRawResponse) GetFromDate() time.Time`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *ArduinoSeriesRawResponse) GetFromDateOk() (*time.Time, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *ArduinoSeriesRawResponse) SetFromDate(v time.Time)`

SetFromDate sets FromDate field to given value.


### GetMessage

`func (o *ArduinoSeriesRawResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ArduinoSeriesRawResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ArduinoSeriesRawResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ArduinoSeriesRawResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetQuery

`func (o *ArduinoSeriesRawResponse) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ArduinoSeriesRawResponse) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ArduinoSeriesRawResponse) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetRespVersion

`func (o *ArduinoSeriesRawResponse) GetRespVersion() int64`

GetRespVersion returns the RespVersion field if non-nil, zero value otherwise.

### GetRespVersionOk

`func (o *ArduinoSeriesRawResponse) GetRespVersionOk() (*int64, bool)`

GetRespVersionOk returns a tuple with the RespVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespVersion

`func (o *ArduinoSeriesRawResponse) SetRespVersion(v int64)`

SetRespVersion sets RespVersion field to given value.


### GetSeries

`func (o *ArduinoSeriesRawResponse) GetSeries() BatchQueryRawResponseSeriesMediaV1`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *ArduinoSeriesRawResponse) GetSeriesOk() (*BatchQueryRawResponseSeriesMediaV1, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *ArduinoSeriesRawResponse) SetSeries(v BatchQueryRawResponseSeriesMediaV1)`

SetSeries sets Series field to given value.


### GetSeriesLimit

`func (o *ArduinoSeriesRawResponse) GetSeriesLimit() int64`

GetSeriesLimit returns the SeriesLimit field if non-nil, zero value otherwise.

### GetSeriesLimitOk

`func (o *ArduinoSeriesRawResponse) GetSeriesLimitOk() (*int64, bool)`

GetSeriesLimitOk returns a tuple with the SeriesLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesLimit

`func (o *ArduinoSeriesRawResponse) SetSeriesLimit(v int64)`

SetSeriesLimit sets SeriesLimit field to given value.

### HasSeriesLimit

`func (o *ArduinoSeriesRawResponse) HasSeriesLimit() bool`

HasSeriesLimit returns a boolean if a field has been set.

### GetSort

`func (o *ArduinoSeriesRawResponse) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *ArduinoSeriesRawResponse) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *ArduinoSeriesRawResponse) SetSort(v string)`

SetSort sets Sort field to given value.


### GetStatus

`func (o *ArduinoSeriesRawResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArduinoSeriesRawResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArduinoSeriesRawResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimes

`func (o *ArduinoSeriesRawResponse) GetTimes() []time.Time`

GetTimes returns the Times field if non-nil, zero value otherwise.

### GetTimesOk

`func (o *ArduinoSeriesRawResponse) GetTimesOk() (*[]time.Time, bool)`

GetTimesOk returns a tuple with the Times field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimes

`func (o *ArduinoSeriesRawResponse) SetTimes(v []time.Time)`

SetTimes sets Times field to given value.


### GetToDate

`func (o *ArduinoSeriesRawResponse) GetToDate() time.Time`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *ArduinoSeriesRawResponse) GetToDateOk() (*time.Time, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *ArduinoSeriesRawResponse) SetToDate(v time.Time)`

SetToDate sets ToDate field to given value.


### GetValues

`func (o *ArduinoSeriesRawResponse) GetValues() []interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ArduinoSeriesRawResponse) GetValuesOk() (*[]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ArduinoSeriesRawResponse) SetValues(v []interface{})`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


