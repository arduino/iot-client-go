# ArduinoSeriesSampledResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**Values** | **[]interface{}** | Values in Float | 

## Methods

### NewArduinoSeriesSampledResponse

`func NewArduinoSeriesSampledResponse(countValues int64, fromDate time.Time, interval int64, query string, respVersion int64, status string, times []time.Time, toDate time.Time, values []interface{}, ) *ArduinoSeriesSampledResponse`

NewArduinoSeriesSampledResponse instantiates a new ArduinoSeriesSampledResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoSeriesSampledResponseWithDefaults

`func NewArduinoSeriesSampledResponseWithDefaults() *ArduinoSeriesSampledResponse`

NewArduinoSeriesSampledResponseWithDefaults instantiates a new ArduinoSeriesSampledResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountValues

`func (o *ArduinoSeriesSampledResponse) GetCountValues() int64`

GetCountValues returns the CountValues field if non-nil, zero value otherwise.

### GetCountValuesOk

`func (o *ArduinoSeriesSampledResponse) GetCountValuesOk() (*int64, bool)`

GetCountValuesOk returns a tuple with the CountValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountValues

`func (o *ArduinoSeriesSampledResponse) SetCountValues(v int64)`

SetCountValues sets CountValues field to given value.


### GetFromDate

`func (o *ArduinoSeriesSampledResponse) GetFromDate() time.Time`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *ArduinoSeriesSampledResponse) GetFromDateOk() (*time.Time, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *ArduinoSeriesSampledResponse) SetFromDate(v time.Time)`

SetFromDate sets FromDate field to given value.


### GetInterval

`func (o *ArduinoSeriesSampledResponse) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *ArduinoSeriesSampledResponse) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *ArduinoSeriesSampledResponse) SetInterval(v int64)`

SetInterval sets Interval field to given value.


### GetMessage

`func (o *ArduinoSeriesSampledResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ArduinoSeriesSampledResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ArduinoSeriesSampledResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ArduinoSeriesSampledResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetQuery

`func (o *ArduinoSeriesSampledResponse) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ArduinoSeriesSampledResponse) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ArduinoSeriesSampledResponse) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetRespVersion

`func (o *ArduinoSeriesSampledResponse) GetRespVersion() int64`

GetRespVersion returns the RespVersion field if non-nil, zero value otherwise.

### GetRespVersionOk

`func (o *ArduinoSeriesSampledResponse) GetRespVersionOk() (*int64, bool)`

GetRespVersionOk returns a tuple with the RespVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespVersion

`func (o *ArduinoSeriesSampledResponse) SetRespVersion(v int64)`

SetRespVersion sets RespVersion field to given value.


### GetSeriesLimit

`func (o *ArduinoSeriesSampledResponse) GetSeriesLimit() int64`

GetSeriesLimit returns the SeriesLimit field if non-nil, zero value otherwise.

### GetSeriesLimitOk

`func (o *ArduinoSeriesSampledResponse) GetSeriesLimitOk() (*int64, bool)`

GetSeriesLimitOk returns a tuple with the SeriesLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesLimit

`func (o *ArduinoSeriesSampledResponse) SetSeriesLimit(v int64)`

SetSeriesLimit sets SeriesLimit field to given value.

### HasSeriesLimit

`func (o *ArduinoSeriesSampledResponse) HasSeriesLimit() bool`

HasSeriesLimit returns a boolean if a field has been set.

### GetStatus

`func (o *ArduinoSeriesSampledResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArduinoSeriesSampledResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArduinoSeriesSampledResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimes

`func (o *ArduinoSeriesSampledResponse) GetTimes() []time.Time`

GetTimes returns the Times field if non-nil, zero value otherwise.

### GetTimesOk

`func (o *ArduinoSeriesSampledResponse) GetTimesOk() (*[]time.Time, bool)`

GetTimesOk returns a tuple with the Times field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimes

`func (o *ArduinoSeriesSampledResponse) SetTimes(v []time.Time)`

SetTimes sets Times field to given value.


### GetToDate

`func (o *ArduinoSeriesSampledResponse) GetToDate() time.Time`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *ArduinoSeriesSampledResponse) GetToDateOk() (*time.Time, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *ArduinoSeriesSampledResponse) SetToDate(v time.Time)`

SetToDate sets ToDate field to given value.


### GetValues

`func (o *ArduinoSeriesSampledResponse) GetValues() []interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ArduinoSeriesSampledResponse) GetValuesOk() (*[]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ArduinoSeriesSampledResponse) SetValues(v []interface{})`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


