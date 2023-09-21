# ArduinoSeriesRawLastValueResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountValues** | **int64** | Total number of values in the array &#39;values&#39; | 
**PropertyId** | **string** | Property id | 
**ThingId** | **string** | Thing id | 
**Times** | [**[]time.Time**](time.Time.md) | Timestamp in RFC3339 | 
**Values** | **[]interface{}** | Values can be in Float, String, Bool, Object | 

## Methods

### NewArduinoSeriesRawLastValueResponse

`func NewArduinoSeriesRawLastValueResponse(countValues int64, propertyId string, thingId string, times []time.Time, values []interface{}, ) *ArduinoSeriesRawLastValueResponse`

NewArduinoSeriesRawLastValueResponse instantiates a new ArduinoSeriesRawLastValueResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoSeriesRawLastValueResponseWithDefaults

`func NewArduinoSeriesRawLastValueResponseWithDefaults() *ArduinoSeriesRawLastValueResponse`

NewArduinoSeriesRawLastValueResponseWithDefaults instantiates a new ArduinoSeriesRawLastValueResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountValues

`func (o *ArduinoSeriesRawLastValueResponse) GetCountValues() int64`

GetCountValues returns the CountValues field if non-nil, zero value otherwise.

### GetCountValuesOk

`func (o *ArduinoSeriesRawLastValueResponse) GetCountValuesOk() (*int64, bool)`

GetCountValuesOk returns a tuple with the CountValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountValues

`func (o *ArduinoSeriesRawLastValueResponse) SetCountValues(v int64)`

SetCountValues sets CountValues field to given value.


### GetPropertyId

`func (o *ArduinoSeriesRawLastValueResponse) GetPropertyId() string`

GetPropertyId returns the PropertyId field if non-nil, zero value otherwise.

### GetPropertyIdOk

`func (o *ArduinoSeriesRawLastValueResponse) GetPropertyIdOk() (*string, bool)`

GetPropertyIdOk returns a tuple with the PropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyId

`func (o *ArduinoSeriesRawLastValueResponse) SetPropertyId(v string)`

SetPropertyId sets PropertyId field to given value.


### GetThingId

`func (o *ArduinoSeriesRawLastValueResponse) GetThingId() string`

GetThingId returns the ThingId field if non-nil, zero value otherwise.

### GetThingIdOk

`func (o *ArduinoSeriesRawLastValueResponse) GetThingIdOk() (*string, bool)`

GetThingIdOk returns a tuple with the ThingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThingId

`func (o *ArduinoSeriesRawLastValueResponse) SetThingId(v string)`

SetThingId sets ThingId field to given value.


### GetTimes

`func (o *ArduinoSeriesRawLastValueResponse) GetTimes() []time.Time`

GetTimes returns the Times field if non-nil, zero value otherwise.

### GetTimesOk

`func (o *ArduinoSeriesRawLastValueResponse) GetTimesOk() (*[]time.Time, bool)`

GetTimesOk returns a tuple with the Times field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimes

`func (o *ArduinoSeriesRawLastValueResponse) SetTimes(v []time.Time)`

SetTimes sets Times field to given value.


### GetValues

`func (o *ArduinoSeriesRawLastValueResponse) GetValues() []interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ArduinoSeriesRawLastValueResponse) GetValuesOk() (*[]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ArduinoSeriesRawLastValueResponse) SetValues(v []interface{})`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


