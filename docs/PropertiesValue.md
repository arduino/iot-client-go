# PropertiesValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the property | 
**Type** | **string** | The type of the property | [default to "infer"]
**Value** | **interface{}** | The last value of the property | 

## Methods

### NewPropertiesValue

`func NewPropertiesValue(name string, type_ string, value interface{}, ) *PropertiesValue`

NewPropertiesValue instantiates a new PropertiesValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertiesValueWithDefaults

`func NewPropertiesValueWithDefaults() *PropertiesValue`

NewPropertiesValueWithDefaults instantiates a new PropertiesValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PropertiesValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PropertiesValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PropertiesValue) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *PropertiesValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PropertiesValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PropertiesValue) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *PropertiesValue) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PropertiesValue) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PropertiesValue) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *PropertiesValue) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *PropertiesValue) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


