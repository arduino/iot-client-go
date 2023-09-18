# ArduinoDevicev2SimpleProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the property | 
**UpdatedAt** | **time.Time** | Update date of the property | 
**Value** | **interface{}** | Value of the property | 

## Methods

### NewArduinoDevicev2SimpleProperties

`func NewArduinoDevicev2SimpleProperties(name string, updatedAt time.Time, value interface{}, ) *ArduinoDevicev2SimpleProperties`

NewArduinoDevicev2SimpleProperties instantiates a new ArduinoDevicev2SimpleProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoDevicev2SimplePropertiesWithDefaults

`func NewArduinoDevicev2SimplePropertiesWithDefaults() *ArduinoDevicev2SimpleProperties`

NewArduinoDevicev2SimplePropertiesWithDefaults instantiates a new ArduinoDevicev2SimpleProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ArduinoDevicev2SimpleProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArduinoDevicev2SimpleProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArduinoDevicev2SimpleProperties) SetName(v string)`

SetName sets Name field to given value.


### GetUpdatedAt

`func (o *ArduinoDevicev2SimpleProperties) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ArduinoDevicev2SimpleProperties) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ArduinoDevicev2SimpleProperties) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetValue

`func (o *ArduinoDevicev2SimpleProperties) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ArduinoDevicev2SimpleProperties) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ArduinoDevicev2SimpleProperties) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *ArduinoDevicev2SimpleProperties) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ArduinoDevicev2SimpleProperties) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


