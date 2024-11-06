# ArduinoTriggerTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]ArduinoActionTemplate**](ArduinoActionTemplate.md) | A list of actions associated with the trigger | [optional] 
**Active** | Pointer to **bool** | Is true if the trigger is enabled | [optional] 
**Criteria** | Pointer to **string** | The criteria of the trigger, could be INCLUDE or EXCLUDE | [optional] 
**Description** | Pointer to **string** | The description of the trigger | [optional] 
**GracePeriodOffline** | Pointer to **int64** | The amount of seconds the trigger will wait before considering a matching device as offline | [optional] 
**GracePeriodOnline** | Pointer to **int64** | The amount of seconds the trigger will wait before considering a matching device as online | [optional] 
**Id** | **string** | The id of the trigger | 
**LinkedDevices** | Pointer to [**[]ArduinoLinkedDeviceTemplate**](ArduinoLinkedDeviceTemplate.md) | A list of devices the trigger is associated to | [optional] 
**LinkedProperty** | Pointer to [**ArduinoLinkedPropertyTemplate**](ArduinoLinkedPropertyTemplate.md) |  | [optional] 
**Name** | **string** | The name of the trigger | 
**OrganizationId** | Pointer to **string** | Id of the organization the trigger belongs to | [optional] 

## Methods

### NewArduinoTriggerTemplate

`func NewArduinoTriggerTemplate(id string, name string, ) *ArduinoTriggerTemplate`

NewArduinoTriggerTemplate instantiates a new ArduinoTriggerTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoTriggerTemplateWithDefaults

`func NewArduinoTriggerTemplateWithDefaults() *ArduinoTriggerTemplate`

NewArduinoTriggerTemplateWithDefaults instantiates a new ArduinoTriggerTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *ArduinoTriggerTemplate) GetActions() []ArduinoActionTemplate`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ArduinoTriggerTemplate) GetActionsOk() (*[]ArduinoActionTemplate, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ArduinoTriggerTemplate) SetActions(v []ArduinoActionTemplate)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ArduinoTriggerTemplate) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetActive

`func (o *ArduinoTriggerTemplate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ArduinoTriggerTemplate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ArduinoTriggerTemplate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ArduinoTriggerTemplate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCriteria

`func (o *ArduinoTriggerTemplate) GetCriteria() string`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *ArduinoTriggerTemplate) GetCriteriaOk() (*string, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *ArduinoTriggerTemplate) SetCriteria(v string)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *ArduinoTriggerTemplate) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetDescription

`func (o *ArduinoTriggerTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ArduinoTriggerTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ArduinoTriggerTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ArduinoTriggerTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGracePeriodOffline

`func (o *ArduinoTriggerTemplate) GetGracePeriodOffline() int64`

GetGracePeriodOffline returns the GracePeriodOffline field if non-nil, zero value otherwise.

### GetGracePeriodOfflineOk

`func (o *ArduinoTriggerTemplate) GetGracePeriodOfflineOk() (*int64, bool)`

GetGracePeriodOfflineOk returns a tuple with the GracePeriodOffline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodOffline

`func (o *ArduinoTriggerTemplate) SetGracePeriodOffline(v int64)`

SetGracePeriodOffline sets GracePeriodOffline field to given value.

### HasGracePeriodOffline

`func (o *ArduinoTriggerTemplate) HasGracePeriodOffline() bool`

HasGracePeriodOffline returns a boolean if a field has been set.

### GetGracePeriodOnline

`func (o *ArduinoTriggerTemplate) GetGracePeriodOnline() int64`

GetGracePeriodOnline returns the GracePeriodOnline field if non-nil, zero value otherwise.

### GetGracePeriodOnlineOk

`func (o *ArduinoTriggerTemplate) GetGracePeriodOnlineOk() (*int64, bool)`

GetGracePeriodOnlineOk returns a tuple with the GracePeriodOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodOnline

`func (o *ArduinoTriggerTemplate) SetGracePeriodOnline(v int64)`

SetGracePeriodOnline sets GracePeriodOnline field to given value.

### HasGracePeriodOnline

`func (o *ArduinoTriggerTemplate) HasGracePeriodOnline() bool`

HasGracePeriodOnline returns a boolean if a field has been set.

### GetId

`func (o *ArduinoTriggerTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArduinoTriggerTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArduinoTriggerTemplate) SetId(v string)`

SetId sets Id field to given value.


### GetLinkedDevices

`func (o *ArduinoTriggerTemplate) GetLinkedDevices() []ArduinoLinkedDeviceTemplate`

GetLinkedDevices returns the LinkedDevices field if non-nil, zero value otherwise.

### GetLinkedDevicesOk

`func (o *ArduinoTriggerTemplate) GetLinkedDevicesOk() (*[]ArduinoLinkedDeviceTemplate, bool)`

GetLinkedDevicesOk returns a tuple with the LinkedDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedDevices

`func (o *ArduinoTriggerTemplate) SetLinkedDevices(v []ArduinoLinkedDeviceTemplate)`

SetLinkedDevices sets LinkedDevices field to given value.

### HasLinkedDevices

`func (o *ArduinoTriggerTemplate) HasLinkedDevices() bool`

HasLinkedDevices returns a boolean if a field has been set.

### GetLinkedProperty

`func (o *ArduinoTriggerTemplate) GetLinkedProperty() ArduinoLinkedPropertyTemplate`

GetLinkedProperty returns the LinkedProperty field if non-nil, zero value otherwise.

### GetLinkedPropertyOk

`func (o *ArduinoTriggerTemplate) GetLinkedPropertyOk() (*ArduinoLinkedPropertyTemplate, bool)`

GetLinkedPropertyOk returns a tuple with the LinkedProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedProperty

`func (o *ArduinoTriggerTemplate) SetLinkedProperty(v ArduinoLinkedPropertyTemplate)`

SetLinkedProperty sets LinkedProperty field to given value.

### HasLinkedProperty

`func (o *ArduinoTriggerTemplate) HasLinkedProperty() bool`

HasLinkedProperty returns a boolean if a field has been set.

### GetName

`func (o *ArduinoTriggerTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArduinoTriggerTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArduinoTriggerTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *ArduinoTriggerTemplate) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ArduinoTriggerTemplate) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ArduinoTriggerTemplate) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ArduinoTriggerTemplate) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


