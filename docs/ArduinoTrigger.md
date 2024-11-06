# ArduinoTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]ArduinoAction**](ArduinoAction.md) | A list of actions associated with the trigger | [optional] 
**Active** | Pointer to **bool** | Is true if the trigger is enabled | [optional] 
**CreatedAt** | Pointer to **time.Time** | Creation date of the trigger | [optional] 
**CreatedBy** | Pointer to **string** | Id of the user who last updated the trigger | [optional] 
**DeletedAt** | Pointer to **time.Time** | Deletion date of the trigger | [optional] 
**Description** | Pointer to **string** | The description of the trigger | [optional] 
**DeviceStatusSource** | Pointer to [**DeviceStatusSource**](DeviceStatusSource.md) |  | [optional] 
**Id** | Pointer to **string** | The id of the trigger | [optional] 
**Name** | **string** | The name of the trigger | 
**OrganizationId** | Pointer to **string** | Id of the organization the trigger belongs to | [optional] 
**PropertyId** | Pointer to **string** | Id of the property the trigger is associated to (mutually exclusive with &#39;device_status_source&#39;) | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Update date of the trigger | [optional] 

## Methods

### NewArduinoTrigger

`func NewArduinoTrigger(name string, ) *ArduinoTrigger`

NewArduinoTrigger instantiates a new ArduinoTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoTriggerWithDefaults

`func NewArduinoTriggerWithDefaults() *ArduinoTrigger`

NewArduinoTriggerWithDefaults instantiates a new ArduinoTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *ArduinoTrigger) GetActions() []ArduinoAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ArduinoTrigger) GetActionsOk() (*[]ArduinoAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ArduinoTrigger) SetActions(v []ArduinoAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ArduinoTrigger) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetActive

`func (o *ArduinoTrigger) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ArduinoTrigger) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ArduinoTrigger) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ArduinoTrigger) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ArduinoTrigger) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArduinoTrigger) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArduinoTrigger) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ArduinoTrigger) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ArduinoTrigger) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ArduinoTrigger) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ArduinoTrigger) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ArduinoTrigger) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ArduinoTrigger) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ArduinoTrigger) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ArduinoTrigger) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ArduinoTrigger) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDescription

`func (o *ArduinoTrigger) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ArduinoTrigger) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ArduinoTrigger) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ArduinoTrigger) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeviceStatusSource

`func (o *ArduinoTrigger) GetDeviceStatusSource() DeviceStatusSource`

GetDeviceStatusSource returns the DeviceStatusSource field if non-nil, zero value otherwise.

### GetDeviceStatusSourceOk

`func (o *ArduinoTrigger) GetDeviceStatusSourceOk() (*DeviceStatusSource, bool)`

GetDeviceStatusSourceOk returns a tuple with the DeviceStatusSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceStatusSource

`func (o *ArduinoTrigger) SetDeviceStatusSource(v DeviceStatusSource)`

SetDeviceStatusSource sets DeviceStatusSource field to given value.

### HasDeviceStatusSource

`func (o *ArduinoTrigger) HasDeviceStatusSource() bool`

HasDeviceStatusSource returns a boolean if a field has been set.

### GetId

`func (o *ArduinoTrigger) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArduinoTrigger) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArduinoTrigger) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ArduinoTrigger) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ArduinoTrigger) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArduinoTrigger) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArduinoTrigger) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *ArduinoTrigger) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ArduinoTrigger) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ArduinoTrigger) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ArduinoTrigger) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetPropertyId

`func (o *ArduinoTrigger) GetPropertyId() string`

GetPropertyId returns the PropertyId field if non-nil, zero value otherwise.

### GetPropertyIdOk

`func (o *ArduinoTrigger) GetPropertyIdOk() (*string, bool)`

GetPropertyIdOk returns a tuple with the PropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyId

`func (o *ArduinoTrigger) SetPropertyId(v string)`

SetPropertyId sets PropertyId field to given value.

### HasPropertyId

`func (o *ArduinoTrigger) HasPropertyId() bool`

HasPropertyId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ArduinoTrigger) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ArduinoTrigger) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ArduinoTrigger) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ArduinoTrigger) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


