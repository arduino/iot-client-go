# ArduinoLinkedvariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the linked variable | 
**LastValue** | Pointer to **interface{}** | Last value of the linked property | [optional] 
**LastValueUpdatedAt** | Pointer to **time.Time** | Update date of the last value | [optional] 
**Name** | **string** | The name of the variable | 
**Permission** | **string** | The permission of the linked variable | 
**ThingId** | **string** | The id of the related thing | 
**ThingName** | **string** | The name of the related thing | 
**ThingTimezone** | Pointer to [**ArduinoTimezone**](ArduinoTimezone.md) |  | [optional] 
**Type** | **string** | The type of the variable | 
**VariableName** | **string** | The name of the variable in the code | 

## Methods

### NewArduinoLinkedvariable

`func NewArduinoLinkedvariable(id string, name string, permission string, thingId string, thingName string, type_ string, variableName string, ) *ArduinoLinkedvariable`

NewArduinoLinkedvariable instantiates a new ArduinoLinkedvariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoLinkedvariableWithDefaults

`func NewArduinoLinkedvariableWithDefaults() *ArduinoLinkedvariable`

NewArduinoLinkedvariableWithDefaults instantiates a new ArduinoLinkedvariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArduinoLinkedvariable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArduinoLinkedvariable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArduinoLinkedvariable) SetId(v string)`

SetId sets Id field to given value.


### GetLastValue

`func (o *ArduinoLinkedvariable) GetLastValue() interface{}`

GetLastValue returns the LastValue field if non-nil, zero value otherwise.

### GetLastValueOk

`func (o *ArduinoLinkedvariable) GetLastValueOk() (*interface{}, bool)`

GetLastValueOk returns a tuple with the LastValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastValue

`func (o *ArduinoLinkedvariable) SetLastValue(v interface{})`

SetLastValue sets LastValue field to given value.

### HasLastValue

`func (o *ArduinoLinkedvariable) HasLastValue() bool`

HasLastValue returns a boolean if a field has been set.

### SetLastValueNil

`func (o *ArduinoLinkedvariable) SetLastValueNil(b bool)`

 SetLastValueNil sets the value for LastValue to be an explicit nil

### UnsetLastValue
`func (o *ArduinoLinkedvariable) UnsetLastValue()`

UnsetLastValue ensures that no value is present for LastValue, not even an explicit nil
### GetLastValueUpdatedAt

`func (o *ArduinoLinkedvariable) GetLastValueUpdatedAt() time.Time`

GetLastValueUpdatedAt returns the LastValueUpdatedAt field if non-nil, zero value otherwise.

### GetLastValueUpdatedAtOk

`func (o *ArduinoLinkedvariable) GetLastValueUpdatedAtOk() (*time.Time, bool)`

GetLastValueUpdatedAtOk returns a tuple with the LastValueUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastValueUpdatedAt

`func (o *ArduinoLinkedvariable) SetLastValueUpdatedAt(v time.Time)`

SetLastValueUpdatedAt sets LastValueUpdatedAt field to given value.

### HasLastValueUpdatedAt

`func (o *ArduinoLinkedvariable) HasLastValueUpdatedAt() bool`

HasLastValueUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *ArduinoLinkedvariable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArduinoLinkedvariable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArduinoLinkedvariable) SetName(v string)`

SetName sets Name field to given value.


### GetPermission

`func (o *ArduinoLinkedvariable) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *ArduinoLinkedvariable) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *ArduinoLinkedvariable) SetPermission(v string)`

SetPermission sets Permission field to given value.


### GetThingId

`func (o *ArduinoLinkedvariable) GetThingId() string`

GetThingId returns the ThingId field if non-nil, zero value otherwise.

### GetThingIdOk

`func (o *ArduinoLinkedvariable) GetThingIdOk() (*string, bool)`

GetThingIdOk returns a tuple with the ThingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThingId

`func (o *ArduinoLinkedvariable) SetThingId(v string)`

SetThingId sets ThingId field to given value.


### GetThingName

`func (o *ArduinoLinkedvariable) GetThingName() string`

GetThingName returns the ThingName field if non-nil, zero value otherwise.

### GetThingNameOk

`func (o *ArduinoLinkedvariable) GetThingNameOk() (*string, bool)`

GetThingNameOk returns a tuple with the ThingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThingName

`func (o *ArduinoLinkedvariable) SetThingName(v string)`

SetThingName sets ThingName field to given value.


### GetThingTimezone

`func (o *ArduinoLinkedvariable) GetThingTimezone() ArduinoTimezone`

GetThingTimezone returns the ThingTimezone field if non-nil, zero value otherwise.

### GetThingTimezoneOk

`func (o *ArduinoLinkedvariable) GetThingTimezoneOk() (*ArduinoTimezone, bool)`

GetThingTimezoneOk returns a tuple with the ThingTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThingTimezone

`func (o *ArduinoLinkedvariable) SetThingTimezone(v ArduinoTimezone)`

SetThingTimezone sets ThingTimezone field to given value.

### HasThingTimezone

`func (o *ArduinoLinkedvariable) HasThingTimezone() bool`

HasThingTimezone returns a boolean if a field has been set.

### GetType

`func (o *ArduinoLinkedvariable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArduinoLinkedvariable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArduinoLinkedvariable) SetType(v string)`

SetType sets Type field to given value.


### GetVariableName

`func (o *ArduinoLinkedvariable) GetVariableName() string`

GetVariableName returns the VariableName field if non-nil, zero value otherwise.

### GetVariableNameOk

`func (o *ArduinoLinkedvariable) GetVariableNameOk() (*string, bool)`

GetVariableNameOk returns a tuple with the VariableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableName

`func (o *ArduinoLinkedvariable) SetVariableName(v string)`

SetVariableName sets VariableName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


