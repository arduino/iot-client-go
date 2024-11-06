# ArduinoTriggerWithLinkedEntities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]ArduinoAction**](ArduinoAction.md) | A list of actions associated with the trigger | [optional] 
**Active** | Pointer to **bool** | Is true if the trigger is enabled | [optional] 
**CreatedBy** | Pointer to **string** | Id of the user who last updated the trigger | [optional] 
**Description** | Pointer to **string** | The description of the trigger | [optional] 
**DeviceStatusSource** | Pointer to [**DeviceStatusSourceWithLinkedDevices**](DeviceStatusSourceWithLinkedDevices.md) |  | [optional] 
**Id** | **string** | The id of the trigger | 
**LinkedProperty** | Pointer to [**ArduinoLinkedProperty**](ArduinoLinkedProperty.md) |  | [optional] 
**Name** | **string** | The name of the trigger | 
**OrganizationId** | Pointer to **string** | Id of the organization the trigger belongs to | [optional] 

## Methods

### NewArduinoTriggerWithLinkedEntities

`func NewArduinoTriggerWithLinkedEntities(id string, name string, ) *ArduinoTriggerWithLinkedEntities`

NewArduinoTriggerWithLinkedEntities instantiates a new ArduinoTriggerWithLinkedEntities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoTriggerWithLinkedEntitiesWithDefaults

`func NewArduinoTriggerWithLinkedEntitiesWithDefaults() *ArduinoTriggerWithLinkedEntities`

NewArduinoTriggerWithLinkedEntitiesWithDefaults instantiates a new ArduinoTriggerWithLinkedEntities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *ArduinoTriggerWithLinkedEntities) GetActions() []ArduinoAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ArduinoTriggerWithLinkedEntities) GetActionsOk() (*[]ArduinoAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ArduinoTriggerWithLinkedEntities) SetActions(v []ArduinoAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ArduinoTriggerWithLinkedEntities) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetActive

`func (o *ArduinoTriggerWithLinkedEntities) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ArduinoTriggerWithLinkedEntities) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ArduinoTriggerWithLinkedEntities) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ArduinoTriggerWithLinkedEntities) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ArduinoTriggerWithLinkedEntities) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ArduinoTriggerWithLinkedEntities) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ArduinoTriggerWithLinkedEntities) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ArduinoTriggerWithLinkedEntities) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *ArduinoTriggerWithLinkedEntities) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ArduinoTriggerWithLinkedEntities) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ArduinoTriggerWithLinkedEntities) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ArduinoTriggerWithLinkedEntities) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeviceStatusSource

`func (o *ArduinoTriggerWithLinkedEntities) GetDeviceStatusSource() DeviceStatusSourceWithLinkedDevices`

GetDeviceStatusSource returns the DeviceStatusSource field if non-nil, zero value otherwise.

### GetDeviceStatusSourceOk

`func (o *ArduinoTriggerWithLinkedEntities) GetDeviceStatusSourceOk() (*DeviceStatusSourceWithLinkedDevices, bool)`

GetDeviceStatusSourceOk returns a tuple with the DeviceStatusSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceStatusSource

`func (o *ArduinoTriggerWithLinkedEntities) SetDeviceStatusSource(v DeviceStatusSourceWithLinkedDevices)`

SetDeviceStatusSource sets DeviceStatusSource field to given value.

### HasDeviceStatusSource

`func (o *ArduinoTriggerWithLinkedEntities) HasDeviceStatusSource() bool`

HasDeviceStatusSource returns a boolean if a field has been set.

### GetId

`func (o *ArduinoTriggerWithLinkedEntities) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArduinoTriggerWithLinkedEntities) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArduinoTriggerWithLinkedEntities) SetId(v string)`

SetId sets Id field to given value.


### GetLinkedProperty

`func (o *ArduinoTriggerWithLinkedEntities) GetLinkedProperty() ArduinoLinkedProperty`

GetLinkedProperty returns the LinkedProperty field if non-nil, zero value otherwise.

### GetLinkedPropertyOk

`func (o *ArduinoTriggerWithLinkedEntities) GetLinkedPropertyOk() (*ArduinoLinkedProperty, bool)`

GetLinkedPropertyOk returns a tuple with the LinkedProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedProperty

`func (o *ArduinoTriggerWithLinkedEntities) SetLinkedProperty(v ArduinoLinkedProperty)`

SetLinkedProperty sets LinkedProperty field to given value.

### HasLinkedProperty

`func (o *ArduinoTriggerWithLinkedEntities) HasLinkedProperty() bool`

HasLinkedProperty returns a boolean if a field has been set.

### GetName

`func (o *ArduinoTriggerWithLinkedEntities) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArduinoTriggerWithLinkedEntities) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArduinoTriggerWithLinkedEntities) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *ArduinoTriggerWithLinkedEntities) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ArduinoTriggerWithLinkedEntities) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ArduinoTriggerWithLinkedEntities) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ArduinoTriggerWithLinkedEntities) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


