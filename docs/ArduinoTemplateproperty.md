# ArduinoTemplateproperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The friendly name of the property | 
**Permission** | **string** | The permission of the property | 
**Type** | **string** | The type of the property | 
**UpdateParameter** | Pointer to **float64** | The update frequency in seconds, or the amount of the property has to change in order to trigger an update | [optional] 
**UpdateStrategy** | **string** | The update strategy for the property value | 
**VariableName** | Pointer to **string** | The sketch variable name of the property | [optional] 

## Methods

### NewArduinoTemplateproperty

`func NewArduinoTemplateproperty(name string, permission string, type_ string, updateStrategy string, ) *ArduinoTemplateproperty`

NewArduinoTemplateproperty instantiates a new ArduinoTemplateproperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoTemplatepropertyWithDefaults

`func NewArduinoTemplatepropertyWithDefaults() *ArduinoTemplateproperty`

NewArduinoTemplatepropertyWithDefaults instantiates a new ArduinoTemplateproperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ArduinoTemplateproperty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArduinoTemplateproperty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArduinoTemplateproperty) SetName(v string)`

SetName sets Name field to given value.


### GetPermission

`func (o *ArduinoTemplateproperty) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *ArduinoTemplateproperty) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *ArduinoTemplateproperty) SetPermission(v string)`

SetPermission sets Permission field to given value.


### GetType

`func (o *ArduinoTemplateproperty) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArduinoTemplateproperty) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArduinoTemplateproperty) SetType(v string)`

SetType sets Type field to given value.


### GetUpdateParameter

`func (o *ArduinoTemplateproperty) GetUpdateParameter() float64`

GetUpdateParameter returns the UpdateParameter field if non-nil, zero value otherwise.

### GetUpdateParameterOk

`func (o *ArduinoTemplateproperty) GetUpdateParameterOk() (*float64, bool)`

GetUpdateParameterOk returns a tuple with the UpdateParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateParameter

`func (o *ArduinoTemplateproperty) SetUpdateParameter(v float64)`

SetUpdateParameter sets UpdateParameter field to given value.

### HasUpdateParameter

`func (o *ArduinoTemplateproperty) HasUpdateParameter() bool`

HasUpdateParameter returns a boolean if a field has been set.

### GetUpdateStrategy

`func (o *ArduinoTemplateproperty) GetUpdateStrategy() string`

GetUpdateStrategy returns the UpdateStrategy field if non-nil, zero value otherwise.

### GetUpdateStrategyOk

`func (o *ArduinoTemplateproperty) GetUpdateStrategyOk() (*string, bool)`

GetUpdateStrategyOk returns a tuple with the UpdateStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateStrategy

`func (o *ArduinoTemplateproperty) SetUpdateStrategy(v string)`

SetUpdateStrategy sets UpdateStrategy field to given value.


### GetVariableName

`func (o *ArduinoTemplateproperty) GetVariableName() string`

GetVariableName returns the VariableName field if non-nil, zero value otherwise.

### GetVariableNameOk

`func (o *ArduinoTemplateproperty) GetVariableNameOk() (*string, bool)`

GetVariableNameOk returns a tuple with the VariableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableName

`func (o *ArduinoTemplateproperty) SetVariableName(v string)`

SetVariableName sets VariableName field to given value.

### HasVariableName

`func (o *ArduinoTemplateproperty) HasVariableName() bool`

HasVariableName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


