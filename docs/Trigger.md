# Trigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]CreateAction**](CreateAction.md) | A list of actions to be associated with the trigger | [optional] 
**Active** | Pointer to **bool** | Is true if the trigger is enabled | [optional] 
**Description** | Pointer to **string** | The description of the trigger | [optional] 
**DeviceStatusSource** | Pointer to [**DeviceStatusSource**](DeviceStatusSource.md) |  | [optional] 
**Id** | Pointer to **string** | The id of the trigger | [optional] 
**Name** | Pointer to **string** | The name of the trigger | [optional] 
**PropertyId** | Pointer to **string** | Id of the property the trigger is associated to (mutually exclusive with &#39;device_status_source&#39;) | [optional] 
**SoftDeleted** | Pointer to **bool** | If false, restore the thing from the soft deletion | [optional] [default to false]

## Methods

### NewTrigger

`func NewTrigger() *Trigger`

NewTrigger instantiates a new Trigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerWithDefaults

`func NewTriggerWithDefaults() *Trigger`

NewTriggerWithDefaults instantiates a new Trigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *Trigger) GetActions() []CreateAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *Trigger) GetActionsOk() (*[]CreateAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *Trigger) SetActions(v []CreateAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *Trigger) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetActive

`func (o *Trigger) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Trigger) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Trigger) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Trigger) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDescription

`func (o *Trigger) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Trigger) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Trigger) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Trigger) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeviceStatusSource

`func (o *Trigger) GetDeviceStatusSource() DeviceStatusSource`

GetDeviceStatusSource returns the DeviceStatusSource field if non-nil, zero value otherwise.

### GetDeviceStatusSourceOk

`func (o *Trigger) GetDeviceStatusSourceOk() (*DeviceStatusSource, bool)`

GetDeviceStatusSourceOk returns a tuple with the DeviceStatusSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceStatusSource

`func (o *Trigger) SetDeviceStatusSource(v DeviceStatusSource)`

SetDeviceStatusSource sets DeviceStatusSource field to given value.

### HasDeviceStatusSource

`func (o *Trigger) HasDeviceStatusSource() bool`

HasDeviceStatusSource returns a boolean if a field has been set.

### GetId

`func (o *Trigger) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Trigger) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Trigger) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Trigger) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Trigger) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Trigger) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Trigger) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Trigger) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPropertyId

`func (o *Trigger) GetPropertyId() string`

GetPropertyId returns the PropertyId field if non-nil, zero value otherwise.

### GetPropertyIdOk

`func (o *Trigger) GetPropertyIdOk() (*string, bool)`

GetPropertyIdOk returns a tuple with the PropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyId

`func (o *Trigger) SetPropertyId(v string)`

SetPropertyId sets PropertyId field to given value.

### HasPropertyId

`func (o *Trigger) HasPropertyId() bool`

HasPropertyId returns a boolean if a field has been set.

### GetSoftDeleted

`func (o *Trigger) GetSoftDeleted() bool`

GetSoftDeleted returns the SoftDeleted field if non-nil, zero value otherwise.

### GetSoftDeletedOk

`func (o *Trigger) GetSoftDeletedOk() (*bool, bool)`

GetSoftDeletedOk returns a tuple with the SoftDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftDeleted

`func (o *Trigger) SetSoftDeleted(v bool)`

SetSoftDeleted sets SoftDeleted field to given value.

### HasSoftDeleted

`func (o *Trigger) HasSoftDeleted() bool`

HasSoftDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


