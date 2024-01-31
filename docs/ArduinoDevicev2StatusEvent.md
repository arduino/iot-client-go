# ArduinoDevicev2StatusEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdatedAt** | **time.Time** | Update timestamp of the status event | 
**Value** | **string** | The status event of the device | 

## Methods

### NewArduinoDevicev2StatusEvent

`func NewArduinoDevicev2StatusEvent(updatedAt time.Time, value string, ) *ArduinoDevicev2StatusEvent`

NewArduinoDevicev2StatusEvent instantiates a new ArduinoDevicev2StatusEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoDevicev2StatusEventWithDefaults

`func NewArduinoDevicev2StatusEventWithDefaults() *ArduinoDevicev2StatusEvent`

NewArduinoDevicev2StatusEventWithDefaults instantiates a new ArduinoDevicev2StatusEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdatedAt

`func (o *ArduinoDevicev2StatusEvent) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ArduinoDevicev2StatusEvent) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ArduinoDevicev2StatusEvent) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetValue

`func (o *ArduinoDevicev2StatusEvent) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ArduinoDevicev2StatusEvent) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ArduinoDevicev2StatusEvent) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


