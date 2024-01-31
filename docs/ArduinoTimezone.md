# ArduinoTimezone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the time zone. | 
**Offset** | **int64** | Current UTC DST offset in seconds. | 
**Until** | **time.Time** | Date until the offset is valid. | 

## Methods

### NewArduinoTimezone

`func NewArduinoTimezone(name string, offset int64, until time.Time, ) *ArduinoTimezone`

NewArduinoTimezone instantiates a new ArduinoTimezone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoTimezoneWithDefaults

`func NewArduinoTimezoneWithDefaults() *ArduinoTimezone`

NewArduinoTimezoneWithDefaults instantiates a new ArduinoTimezone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ArduinoTimezone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArduinoTimezone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArduinoTimezone) SetName(v string)`

SetName sets Name field to given value.


### GetOffset

`func (o *ArduinoTimezone) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ArduinoTimezone) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ArduinoTimezone) SetOffset(v int64)`

SetOffset sets Offset field to given value.


### GetUntil

`func (o *ArduinoTimezone) GetUntil() time.Time`

GetUntil returns the Until field if non-nil, zero value otherwise.

### GetUntilOk

`func (o *ArduinoTimezone) GetUntilOk() (*time.Time, bool)`

GetUntilOk returns a tuple with the Until field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntil

`func (o *ArduinoTimezone) SetUntil(v time.Time)`

SetUntil sets Until field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


