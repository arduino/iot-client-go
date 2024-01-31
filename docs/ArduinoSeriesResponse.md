# ArduinoSeriesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregation** | Pointer to **string** | Aggregation statistic function. For numeric values, AVG statistic is used by default. PCT_X compute the Xth approximate percentile (e.g. PCT_95 is the 95th approximate percentile). For boolean, BOOL_OR statistic is used as default. | [optional] 
**CountValues** | **int64** | Total number of values in the array &#39;values&#39; | 
**FromDate** | **time.Time** | From date | 
**Interval** | **int64** | Resolution in seconds | 
**Message** | Pointer to **string** | If the response is different than &#39;ok&#39; | [optional] [default to ""]
**Query** | **string** | Query of for the data | 
**RespVersion** | **int64** | Response version | 
**SeriesLimit** | Pointer to **int64** | Maximum number of values returned after data aggregation, if any | [optional] 
**Status** | **string** | Status of the response | 
**Times** | [**[]time.Time**](time.Time.md) | Timestamp in RFC3339 | 
**ToDate** | **time.Time** | To date | 
**Values** | **[]float64** | Values in Float | 

## Methods

### NewArduinoSeriesResponse

`func NewArduinoSeriesResponse(countValues int64, fromDate time.Time, interval int64, query string, respVersion int64, status string, times []time.Time, toDate time.Time, values []float64, ) *ArduinoSeriesResponse`

NewArduinoSeriesResponse instantiates a new ArduinoSeriesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoSeriesResponseWithDefaults

`func NewArduinoSeriesResponseWithDefaults() *ArduinoSeriesResponse`

NewArduinoSeriesResponseWithDefaults instantiates a new ArduinoSeriesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *ArduinoSeriesResponse) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *ArduinoSeriesResponse) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *ArduinoSeriesResponse) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *ArduinoSeriesResponse) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetCountValues

`func (o *ArduinoSeriesResponse) GetCountValues() int64`

GetCountValues returns the CountValues field if non-nil, zero value otherwise.

### GetCountValuesOk

`func (o *ArduinoSeriesResponse) GetCountValuesOk() (*int64, bool)`

GetCountValuesOk returns a tuple with the CountValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountValues

`func (o *ArduinoSeriesResponse) SetCountValues(v int64)`

SetCountValues sets CountValues field to given value.


### GetFromDate

`func (o *ArduinoSeriesResponse) GetFromDate() time.Time`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *ArduinoSeriesResponse) GetFromDateOk() (*time.Time, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *ArduinoSeriesResponse) SetFromDate(v time.Time)`

SetFromDate sets FromDate field to given value.


### GetInterval

`func (o *ArduinoSeriesResponse) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *ArduinoSeriesResponse) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *ArduinoSeriesResponse) SetInterval(v int64)`

SetInterval sets Interval field to given value.


### GetMessage

`func (o *ArduinoSeriesResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ArduinoSeriesResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ArduinoSeriesResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ArduinoSeriesResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetQuery

`func (o *ArduinoSeriesResponse) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ArduinoSeriesResponse) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ArduinoSeriesResponse) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetRespVersion

`func (o *ArduinoSeriesResponse) GetRespVersion() int64`

GetRespVersion returns the RespVersion field if non-nil, zero value otherwise.

### GetRespVersionOk

`func (o *ArduinoSeriesResponse) GetRespVersionOk() (*int64, bool)`

GetRespVersionOk returns a tuple with the RespVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespVersion

`func (o *ArduinoSeriesResponse) SetRespVersion(v int64)`

SetRespVersion sets RespVersion field to given value.


### GetSeriesLimit

`func (o *ArduinoSeriesResponse) GetSeriesLimit() int64`

GetSeriesLimit returns the SeriesLimit field if non-nil, zero value otherwise.

### GetSeriesLimitOk

`func (o *ArduinoSeriesResponse) GetSeriesLimitOk() (*int64, bool)`

GetSeriesLimitOk returns a tuple with the SeriesLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesLimit

`func (o *ArduinoSeriesResponse) SetSeriesLimit(v int64)`

SetSeriesLimit sets SeriesLimit field to given value.

### HasSeriesLimit

`func (o *ArduinoSeriesResponse) HasSeriesLimit() bool`

HasSeriesLimit returns a boolean if a field has been set.

### GetStatus

`func (o *ArduinoSeriesResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArduinoSeriesResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArduinoSeriesResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimes

`func (o *ArduinoSeriesResponse) GetTimes() []time.Time`

GetTimes returns the Times field if non-nil, zero value otherwise.

### GetTimesOk

`func (o *ArduinoSeriesResponse) GetTimesOk() (*[]time.Time, bool)`

GetTimesOk returns a tuple with the Times field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimes

`func (o *ArduinoSeriesResponse) SetTimes(v []time.Time)`

SetTimes sets Times field to given value.


### GetToDate

`func (o *ArduinoSeriesResponse) GetToDate() time.Time`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *ArduinoSeriesResponse) GetToDateOk() (*time.Time, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *ArduinoSeriesResponse) SetToDate(v time.Time)`

SetToDate sets ToDate field to given value.


### GetValues

`func (o *ArduinoSeriesResponse) GetValues() []float64`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ArduinoSeriesResponse) GetValuesOk() (*[]float64, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ArduinoSeriesResponse) SetValues(v []float64)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


