# PropertyTimedValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **time.Time** | The timestamp of the property value | 
**Value** | **interface{}** | The property value | 

## Methods

### NewPropertyTimedValue

`func NewPropertyTimedValue(timestamp time.Time, value interface{}, ) *PropertyTimedValue`

NewPropertyTimedValue instantiates a new PropertyTimedValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyTimedValueWithDefaults

`func NewPropertyTimedValueWithDefaults() *PropertyTimedValue`

NewPropertyTimedValueWithDefaults instantiates a new PropertyTimedValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *PropertyTimedValue) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PropertyTimedValue) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PropertyTimedValue) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetValue

`func (o *PropertyTimedValue) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PropertyTimedValue) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PropertyTimedValue) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *PropertyTimedValue) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *PropertyTimedValue) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


