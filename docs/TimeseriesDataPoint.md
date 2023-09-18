# TimeseriesDataPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **time.Time** | Binning timestamp | 
**Value** | **float64** | Avg value on the binning interval | 

## Methods

### NewTimeseriesDataPoint

`func NewTimeseriesDataPoint(time time.Time, value float64, ) *TimeseriesDataPoint`

NewTimeseriesDataPoint instantiates a new TimeseriesDataPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeseriesDataPointWithDefaults

`func NewTimeseriesDataPointWithDefaults() *TimeseriesDataPoint`

NewTimeseriesDataPointWithDefaults instantiates a new TimeseriesDataPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *TimeseriesDataPoint) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *TimeseriesDataPoint) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *TimeseriesDataPoint) SetTime(v time.Time)`

SetTime sets Time field to given value.


### GetValue

`func (o *TimeseriesDataPoint) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TimeseriesDataPoint) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TimeseriesDataPoint) SetValue(v float64)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


