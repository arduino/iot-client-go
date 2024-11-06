# ArduinoAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to **string** | Id of the user who created the action | [optional] 
**Description** | Pointer to **string** | The description of the action | [optional] 
**Email** | Pointer to [**EmailAction**](EmailAction.md) |  | [optional] 
**Id** | Pointer to **string** | The id of the action | [optional] 
**Kind** | Pointer to **string** | The kind of the action | [optional] 
**Name** | Pointer to **string** | The name of the action | [optional] 
**OrganizationId** | Pointer to **string** | Id of the organization the trigger belongs to | [optional] 
**PushNotification** | Pointer to [**PushAction**](PushAction.md) |  | [optional] 
**TriggerId** | Pointer to **string** | Id of the trigger the action is associated to | [optional] 

## Methods

### NewArduinoAction

`func NewArduinoAction() *ArduinoAction`

NewArduinoAction instantiates a new ArduinoAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoActionWithDefaults

`func NewArduinoActionWithDefaults() *ArduinoAction`

NewArduinoActionWithDefaults instantiates a new ArduinoAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *ArduinoAction) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ArduinoAction) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ArduinoAction) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ArduinoAction) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *ArduinoAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ArduinoAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ArduinoAction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ArduinoAction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmail

`func (o *ArduinoAction) GetEmail() EmailAction`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ArduinoAction) GetEmailOk() (*EmailAction, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ArduinoAction) SetEmail(v EmailAction)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ArduinoAction) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *ArduinoAction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArduinoAction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArduinoAction) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ArduinoAction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *ArduinoAction) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ArduinoAction) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ArduinoAction) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ArduinoAction) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *ArduinoAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArduinoAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArduinoAction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ArduinoAction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizationId

`func (o *ArduinoAction) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ArduinoAction) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ArduinoAction) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ArduinoAction) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetPushNotification

`func (o *ArduinoAction) GetPushNotification() PushAction`

GetPushNotification returns the PushNotification field if non-nil, zero value otherwise.

### GetPushNotificationOk

`func (o *ArduinoAction) GetPushNotificationOk() (*PushAction, bool)`

GetPushNotificationOk returns a tuple with the PushNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushNotification

`func (o *ArduinoAction) SetPushNotification(v PushAction)`

SetPushNotification sets PushNotification field to given value.

### HasPushNotification

`func (o *ArduinoAction) HasPushNotification() bool`

HasPushNotification returns a boolean if a field has been set.

### GetTriggerId

`func (o *ArduinoAction) GetTriggerId() string`

GetTriggerId returns the TriggerId field if non-nil, zero value otherwise.

### GetTriggerIdOk

`func (o *ArduinoAction) GetTriggerIdOk() (*string, bool)`

GetTriggerIdOk returns a tuple with the TriggerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerId

`func (o *ArduinoAction) SetTriggerId(v string)`

SetTriggerId sets TriggerId field to given value.

### HasTriggerId

`func (o *ArduinoAction) HasTriggerId() bool`

HasTriggerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


