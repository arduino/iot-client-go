# ArduinoLinkedProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | [**ArduinoProperty**](ArduinoProperty.md) |  | 
**Status** | **string** | The status of the linked property | 

## Methods

### NewArduinoLinkedProperty

`func NewArduinoLinkedProperty(property ArduinoProperty, status string, ) *ArduinoLinkedProperty`

NewArduinoLinkedProperty instantiates a new ArduinoLinkedProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoLinkedPropertyWithDefaults

`func NewArduinoLinkedPropertyWithDefaults() *ArduinoLinkedProperty`

NewArduinoLinkedPropertyWithDefaults instantiates a new ArduinoLinkedProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *ArduinoLinkedProperty) GetProperty() ArduinoProperty`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *ArduinoLinkedProperty) GetPropertyOk() (*ArduinoProperty, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *ArduinoLinkedProperty) SetProperty(v ArduinoProperty)`

SetProperty sets Property field to given value.


### GetStatus

`func (o *ArduinoLinkedProperty) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArduinoLinkedProperty) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArduinoLinkedProperty) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


