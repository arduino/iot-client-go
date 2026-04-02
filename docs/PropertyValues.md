# PropertyValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Properties** | [**[]PropertyDefinition**](PropertyDefinition.md) | The set of properties to publish | 
**ToDevice** | Pointer to **bool** | Handle data direction, simulating data from or to device (default false - data sent by device) | [optional] [default to false]

## Methods

### NewPropertyValues

`func NewPropertyValues(properties []PropertyDefinition, ) *PropertyValues`

NewPropertyValues instantiates a new PropertyValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyValuesWithDefaults

`func NewPropertyValuesWithDefaults() *PropertyValues`

NewPropertyValuesWithDefaults instantiates a new PropertyValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *PropertyValues) GetProperties() []PropertyDefinition`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PropertyValues) GetPropertiesOk() (*[]PropertyDefinition, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PropertyValues) SetProperties(v []PropertyDefinition)`

SetProperties sets Properties field to given value.


### GetToDevice

`func (o *PropertyValues) GetToDevice() bool`

GetToDevice returns the ToDevice field if non-nil, zero value otherwise.

### GetToDeviceOk

`func (o *PropertyValues) GetToDeviceOk() (*bool, bool)`

GetToDeviceOk returns a tuple with the ToDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDevice

`func (o *PropertyValues) SetToDevice(v bool)`

SetToDevice sets ToDevice field to given value.

### HasToDevice

`func (o *PropertyValues) HasToDevice() bool`

HasToDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


