# ArduinoTemplatevariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the variable | 
**Permission** | **string** | The permission of the linked variable | 
**ThingId** | **string** | The name of the related thing | 
**ThingTimezone** | Pointer to [**ArduinoTimezone**](ArduinoTimezone.md) |  | [optional] 
**Type** | **string** | The type of the variable | 
**VariableId** | **string** | The name of the variable in the code | 

## Methods

### NewArduinoTemplatevariable

`func NewArduinoTemplatevariable(name string, permission string, thingId string, type_ string, variableId string, ) *ArduinoTemplatevariable`

NewArduinoTemplatevariable instantiates a new ArduinoTemplatevariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoTemplatevariableWithDefaults

`func NewArduinoTemplatevariableWithDefaults() *ArduinoTemplatevariable`

NewArduinoTemplatevariableWithDefaults instantiates a new ArduinoTemplatevariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ArduinoTemplatevariable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArduinoTemplatevariable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArduinoTemplatevariable) SetName(v string)`

SetName sets Name field to given value.


### GetPermission

`func (o *ArduinoTemplatevariable) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *ArduinoTemplatevariable) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *ArduinoTemplatevariable) SetPermission(v string)`

SetPermission sets Permission field to given value.


### GetThingId

`func (o *ArduinoTemplatevariable) GetThingId() string`

GetThingId returns the ThingId field if non-nil, zero value otherwise.

### GetThingIdOk

`func (o *ArduinoTemplatevariable) GetThingIdOk() (*string, bool)`

GetThingIdOk returns a tuple with the ThingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThingId

`func (o *ArduinoTemplatevariable) SetThingId(v string)`

SetThingId sets ThingId field to given value.


### GetThingTimezone

`func (o *ArduinoTemplatevariable) GetThingTimezone() ArduinoTimezone`

GetThingTimezone returns the ThingTimezone field if non-nil, zero value otherwise.

### GetThingTimezoneOk

`func (o *ArduinoTemplatevariable) GetThingTimezoneOk() (*ArduinoTimezone, bool)`

GetThingTimezoneOk returns a tuple with the ThingTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThingTimezone

`func (o *ArduinoTemplatevariable) SetThingTimezone(v ArduinoTimezone)`

SetThingTimezone sets ThingTimezone field to given value.

### HasThingTimezone

`func (o *ArduinoTemplatevariable) HasThingTimezone() bool`

HasThingTimezone returns a boolean if a field has been set.

### GetType

`func (o *ArduinoTemplatevariable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArduinoTemplatevariable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArduinoTemplatevariable) SetType(v string)`

SetType sets Type field to given value.


### GetVariableId

`func (o *ArduinoTemplatevariable) GetVariableId() string`

GetVariableId returns the VariableId field if non-nil, zero value otherwise.

### GetVariableIdOk

`func (o *ArduinoTemplatevariable) GetVariableIdOk() (*string, bool)`

GetVariableIdOk returns a tuple with the VariableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableId

`func (o *ArduinoTemplatevariable) SetVariableId(v string)`

SetVariableId sets VariableId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


