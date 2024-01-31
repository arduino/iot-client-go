# ArduinoProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | Creation date of the property | [optional] 
**DeletedAt** | Pointer to **time.Time** | Delete date of the property | [optional] 
**Href** | **string** | The api reference of this property | 
**Id** | **string** | The id of the property | 
**LastValue** | Pointer to **interface{}** | Last value of this property | [optional] 
**LinkedToTrigger** | Pointer to **bool** | Indicates if the property is involved in the activation of at least a trigger | [optional] 
**MaxValue** | Pointer to **float64** | Maximum value of this property | [optional] 
**MinValue** | Pointer to **float64** | Minimum value of this property | [optional] 
**Name** | **string** | The friendly name of the property | 
**Permission** | **string** | The permission of the property | 
**Persist** | Pointer to **bool** | If true, data will persist into a timeseries database | [optional] 
**SyncId** | Pointer to **string** | The id of the sync pool | [optional] 
**Tag** | Pointer to **int64** | The integer id of the property | [optional] 
**ThingId** | **string** | The id of the thing | 
**ThingName** | Pointer to **string** | The name of the associated thing | [optional] 
**Type** | **string** | The type of the property | 
**UpdateParameter** | Pointer to **float64** | The update frequency in seconds, or the amount of the property has to change in order to trigger an update | [optional] 
**UpdateStrategy** | **string** | The update strategy for the property value | 
**UpdatedAt** | Pointer to **time.Time** | Update date of the property | [optional] 
**ValueUpdatedAt** | Pointer to **time.Time** | Last update timestamp of this property | [optional] 
**VariableName** | Pointer to **string** | The sketch variable name of the property | [optional] 

## Methods

### NewArduinoProperty

`func NewArduinoProperty(href string, id string, name string, permission string, thingId string, type_ string, updateStrategy string, ) *ArduinoProperty`

NewArduinoProperty instantiates a new ArduinoProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoPropertyWithDefaults

`func NewArduinoPropertyWithDefaults() *ArduinoProperty`

NewArduinoPropertyWithDefaults instantiates a new ArduinoProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ArduinoProperty) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArduinoProperty) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArduinoProperty) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ArduinoProperty) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ArduinoProperty) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ArduinoProperty) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ArduinoProperty) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ArduinoProperty) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetHref

`func (o *ArduinoProperty) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ArduinoProperty) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ArduinoProperty) SetHref(v string)`

SetHref sets Href field to given value.


### GetId

`func (o *ArduinoProperty) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArduinoProperty) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArduinoProperty) SetId(v string)`

SetId sets Id field to given value.


### GetLastValue

`func (o *ArduinoProperty) GetLastValue() interface{}`

GetLastValue returns the LastValue field if non-nil, zero value otherwise.

### GetLastValueOk

`func (o *ArduinoProperty) GetLastValueOk() (*interface{}, bool)`

GetLastValueOk returns a tuple with the LastValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastValue

`func (o *ArduinoProperty) SetLastValue(v interface{})`

SetLastValue sets LastValue field to given value.

### HasLastValue

`func (o *ArduinoProperty) HasLastValue() bool`

HasLastValue returns a boolean if a field has been set.

### SetLastValueNil

`func (o *ArduinoProperty) SetLastValueNil(b bool)`

 SetLastValueNil sets the value for LastValue to be an explicit nil

### UnsetLastValue
`func (o *ArduinoProperty) UnsetLastValue()`

UnsetLastValue ensures that no value is present for LastValue, not even an explicit nil
### GetLinkedToTrigger

`func (o *ArduinoProperty) GetLinkedToTrigger() bool`

GetLinkedToTrigger returns the LinkedToTrigger field if non-nil, zero value otherwise.

### GetLinkedToTriggerOk

`func (o *ArduinoProperty) GetLinkedToTriggerOk() (*bool, bool)`

GetLinkedToTriggerOk returns a tuple with the LinkedToTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedToTrigger

`func (o *ArduinoProperty) SetLinkedToTrigger(v bool)`

SetLinkedToTrigger sets LinkedToTrigger field to given value.

### HasLinkedToTrigger

`func (o *ArduinoProperty) HasLinkedToTrigger() bool`

HasLinkedToTrigger returns a boolean if a field has been set.

### GetMaxValue

`func (o *ArduinoProperty) GetMaxValue() float64`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *ArduinoProperty) GetMaxValueOk() (*float64, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *ArduinoProperty) SetMaxValue(v float64)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *ArduinoProperty) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### GetMinValue

`func (o *ArduinoProperty) GetMinValue() float64`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *ArduinoProperty) GetMinValueOk() (*float64, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *ArduinoProperty) SetMinValue(v float64)`

SetMinValue sets MinValue field to given value.

### HasMinValue

`func (o *ArduinoProperty) HasMinValue() bool`

HasMinValue returns a boolean if a field has been set.

### GetName

`func (o *ArduinoProperty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArduinoProperty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArduinoProperty) SetName(v string)`

SetName sets Name field to given value.


### GetPermission

`func (o *ArduinoProperty) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *ArduinoProperty) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *ArduinoProperty) SetPermission(v string)`

SetPermission sets Permission field to given value.


### GetPersist

`func (o *ArduinoProperty) GetPersist() bool`

GetPersist returns the Persist field if non-nil, zero value otherwise.

### GetPersistOk

`func (o *ArduinoProperty) GetPersistOk() (*bool, bool)`

GetPersistOk returns a tuple with the Persist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersist

`func (o *ArduinoProperty) SetPersist(v bool)`

SetPersist sets Persist field to given value.

### HasPersist

`func (o *ArduinoProperty) HasPersist() bool`

HasPersist returns a boolean if a field has been set.

### GetSyncId

`func (o *ArduinoProperty) GetSyncId() string`

GetSyncId returns the SyncId field if non-nil, zero value otherwise.

### GetSyncIdOk

`func (o *ArduinoProperty) GetSyncIdOk() (*string, bool)`

GetSyncIdOk returns a tuple with the SyncId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncId

`func (o *ArduinoProperty) SetSyncId(v string)`

SetSyncId sets SyncId field to given value.

### HasSyncId

`func (o *ArduinoProperty) HasSyncId() bool`

HasSyncId returns a boolean if a field has been set.

### GetTag

`func (o *ArduinoProperty) GetTag() int64`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ArduinoProperty) GetTagOk() (*int64, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ArduinoProperty) SetTag(v int64)`

SetTag sets Tag field to given value.

### HasTag

`func (o *ArduinoProperty) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetThingId

`func (o *ArduinoProperty) GetThingId() string`

GetThingId returns the ThingId field if non-nil, zero value otherwise.

### GetThingIdOk

`func (o *ArduinoProperty) GetThingIdOk() (*string, bool)`

GetThingIdOk returns a tuple with the ThingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThingId

`func (o *ArduinoProperty) SetThingId(v string)`

SetThingId sets ThingId field to given value.


### GetThingName

`func (o *ArduinoProperty) GetThingName() string`

GetThingName returns the ThingName field if non-nil, zero value otherwise.

### GetThingNameOk

`func (o *ArduinoProperty) GetThingNameOk() (*string, bool)`

GetThingNameOk returns a tuple with the ThingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThingName

`func (o *ArduinoProperty) SetThingName(v string)`

SetThingName sets ThingName field to given value.

### HasThingName

`func (o *ArduinoProperty) HasThingName() bool`

HasThingName returns a boolean if a field has been set.

### GetType

`func (o *ArduinoProperty) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArduinoProperty) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArduinoProperty) SetType(v string)`

SetType sets Type field to given value.


### GetUpdateParameter

`func (o *ArduinoProperty) GetUpdateParameter() float64`

GetUpdateParameter returns the UpdateParameter field if non-nil, zero value otherwise.

### GetUpdateParameterOk

`func (o *ArduinoProperty) GetUpdateParameterOk() (*float64, bool)`

GetUpdateParameterOk returns a tuple with the UpdateParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateParameter

`func (o *ArduinoProperty) SetUpdateParameter(v float64)`

SetUpdateParameter sets UpdateParameter field to given value.

### HasUpdateParameter

`func (o *ArduinoProperty) HasUpdateParameter() bool`

HasUpdateParameter returns a boolean if a field has been set.

### GetUpdateStrategy

`func (o *ArduinoProperty) GetUpdateStrategy() string`

GetUpdateStrategy returns the UpdateStrategy field if non-nil, zero value otherwise.

### GetUpdateStrategyOk

`func (o *ArduinoProperty) GetUpdateStrategyOk() (*string, bool)`

GetUpdateStrategyOk returns a tuple with the UpdateStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateStrategy

`func (o *ArduinoProperty) SetUpdateStrategy(v string)`

SetUpdateStrategy sets UpdateStrategy field to given value.


### GetUpdatedAt

`func (o *ArduinoProperty) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ArduinoProperty) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ArduinoProperty) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ArduinoProperty) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetValueUpdatedAt

`func (o *ArduinoProperty) GetValueUpdatedAt() time.Time`

GetValueUpdatedAt returns the ValueUpdatedAt field if non-nil, zero value otherwise.

### GetValueUpdatedAtOk

`func (o *ArduinoProperty) GetValueUpdatedAtOk() (*time.Time, bool)`

GetValueUpdatedAtOk returns a tuple with the ValueUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueUpdatedAt

`func (o *ArduinoProperty) SetValueUpdatedAt(v time.Time)`

SetValueUpdatedAt sets ValueUpdatedAt field to given value.

### HasValueUpdatedAt

`func (o *ArduinoProperty) HasValueUpdatedAt() bool`

HasValueUpdatedAt returns a boolean if a field has been set.

### GetVariableName

`func (o *ArduinoProperty) GetVariableName() string`

GetVariableName returns the VariableName field if non-nil, zero value otherwise.

### GetVariableNameOk

`func (o *ArduinoProperty) GetVariableNameOk() (*string, bool)`

GetVariableNameOk returns a tuple with the VariableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableName

`func (o *ArduinoProperty) SetVariableName(v string)`

SetVariableName sets VariableName field to given value.

### HasVariableName

`func (o *ArduinoProperty) HasVariableName() bool`

HasVariableName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


