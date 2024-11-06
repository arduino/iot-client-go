# Variable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attribute** | **string** | The template expression that extracts the value from the respective entity | 
**Entity** | **string** | Type of the entity being referenced | 
**Id** | Pointer to **string** | The ID of the referenced entity | [optional] 
**Placeholder** | **string** | Name of the variable as referenced by the expression | 
**PropertyId** | Pointer to **string** | The ID of the property referenced entity | [optional] 
**ThingId** | Pointer to **string** | The ID of the thing referenced entity | [optional] 

## Methods

### NewVariable

`func NewVariable(attribute string, entity string, placeholder string, ) *Variable`

NewVariable instantiates a new Variable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableWithDefaults

`func NewVariableWithDefaults() *Variable`

NewVariableWithDefaults instantiates a new Variable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttribute

`func (o *Variable) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *Variable) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *Variable) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.


### GetEntity

`func (o *Variable) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *Variable) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *Variable) SetEntity(v string)`

SetEntity sets Entity field to given value.


### GetId

`func (o *Variable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Variable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Variable) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Variable) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlaceholder

`func (o *Variable) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *Variable) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *Variable) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.


### GetPropertyId

`func (o *Variable) GetPropertyId() string`

GetPropertyId returns the PropertyId field if non-nil, zero value otherwise.

### GetPropertyIdOk

`func (o *Variable) GetPropertyIdOk() (*string, bool)`

GetPropertyIdOk returns a tuple with the PropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyId

`func (o *Variable) SetPropertyId(v string)`

SetPropertyId sets PropertyId field to given value.

### HasPropertyId

`func (o *Variable) HasPropertyId() bool`

HasPropertyId returns a boolean if a field has been set.

### GetThingId

`func (o *Variable) GetThingId() string`

GetThingId returns the ThingId field if non-nil, zero value otherwise.

### GetThingIdOk

`func (o *Variable) GetThingIdOk() (*string, bool)`

GetThingIdOk returns a tuple with the ThingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThingId

`func (o *Variable) SetThingId(v string)`

SetThingId sets ThingId field to given value.

### HasThingId

`func (o *Variable) HasThingId() bool`

HasThingId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


