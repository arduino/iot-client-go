# Property

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxValue** | Pointer to **float64** | Maximum value of this property | [optional] 
**MinValue** | Pointer to **float64** | Minimum value of this property | [optional] 
**Name** | **string** | The friendly name of the property | 
**Permission** | **string** | The permission of the property | 
**Persist** | Pointer to **bool** | If true, data will persist into a timeseries database | [optional] [default to true]
**Tag** | Pointer to **float64** | The integer id of the property | [optional] 
**Type** | **string** | The type of the property | 
**UpdateParameter** | Pointer to **float64** | The update frequency in seconds, or the amount of the property has to change in order to trigger an update | [optional] 
**UpdateStrategy** | **string** | The update strategy for the property value | 
**VariableName** | Pointer to **string** | The  sketch variable name of the property | [optional] 

## Methods

### NewProperty

`func NewProperty(name string, permission string, type_ string, updateStrategy string, ) *Property`

NewProperty instantiates a new Property object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyWithDefaults

`func NewPropertyWithDefaults() *Property`

NewPropertyWithDefaults instantiates a new Property object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxValue

`func (o *Property) GetMaxValue() float64`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *Property) GetMaxValueOk() (*float64, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *Property) SetMaxValue(v float64)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *Property) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### GetMinValue

`func (o *Property) GetMinValue() float64`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *Property) GetMinValueOk() (*float64, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *Property) SetMinValue(v float64)`

SetMinValue sets MinValue field to given value.

### HasMinValue

`func (o *Property) HasMinValue() bool`

HasMinValue returns a boolean if a field has been set.

### GetName

`func (o *Property) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Property) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Property) SetName(v string)`

SetName sets Name field to given value.


### GetPermission

`func (o *Property) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *Property) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *Property) SetPermission(v string)`

SetPermission sets Permission field to given value.


### GetPersist

`func (o *Property) GetPersist() bool`

GetPersist returns the Persist field if non-nil, zero value otherwise.

### GetPersistOk

`func (o *Property) GetPersistOk() (*bool, bool)`

GetPersistOk returns a tuple with the Persist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersist

`func (o *Property) SetPersist(v bool)`

SetPersist sets Persist field to given value.

### HasPersist

`func (o *Property) HasPersist() bool`

HasPersist returns a boolean if a field has been set.

### GetTag

`func (o *Property) GetTag() float64`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *Property) GetTagOk() (*float64, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *Property) SetTag(v float64)`

SetTag sets Tag field to given value.

### HasTag

`func (o *Property) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetType

`func (o *Property) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Property) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Property) SetType(v string)`

SetType sets Type field to given value.


### GetUpdateParameter

`func (o *Property) GetUpdateParameter() float64`

GetUpdateParameter returns the UpdateParameter field if non-nil, zero value otherwise.

### GetUpdateParameterOk

`func (o *Property) GetUpdateParameterOk() (*float64, bool)`

GetUpdateParameterOk returns a tuple with the UpdateParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateParameter

`func (o *Property) SetUpdateParameter(v float64)`

SetUpdateParameter sets UpdateParameter field to given value.

### HasUpdateParameter

`func (o *Property) HasUpdateParameter() bool`

HasUpdateParameter returns a boolean if a field has been set.

### GetUpdateStrategy

`func (o *Property) GetUpdateStrategy() string`

GetUpdateStrategy returns the UpdateStrategy field if non-nil, zero value otherwise.

### GetUpdateStrategyOk

`func (o *Property) GetUpdateStrategyOk() (*string, bool)`

GetUpdateStrategyOk returns a tuple with the UpdateStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateStrategy

`func (o *Property) SetUpdateStrategy(v string)`

SetUpdateStrategy sets UpdateStrategy field to given value.


### GetVariableName

`func (o *Property) GetVariableName() string`

GetVariableName returns the VariableName field if non-nil, zero value otherwise.

### GetVariableNameOk

`func (o *Property) GetVariableNameOk() (*string, bool)`

GetVariableNameOk returns a tuple with the VariableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableName

`func (o *Property) SetVariableName(v string)`

SetVariableName sets VariableName field to given value.

### HasVariableName

`func (o *Property) HasVariableName() bool`

HasVariableName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


