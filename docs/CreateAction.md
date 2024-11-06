# CreateAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the trigger | [optional] 
**Email** | Pointer to [**EmailAction**](EmailAction.md) |  | [optional] 
**Kind** | **string** | The kind of the action | 
**Name** | **string** | The name of the action | 
**OrganizationId** | Pointer to **string** | Id of the organization the trigger belongs to | [optional] 
**PushNotification** | Pointer to [**PushAction**](PushAction.md) |  | [optional] 
**TriggerId** | Pointer to **string** | Id of the trigger the action is associated to | [optional] 

## Methods

### NewCreateAction

`func NewCreateAction(kind string, name string, ) *CreateAction`

NewCreateAction instantiates a new CreateAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateActionWithDefaults

`func NewCreateActionWithDefaults() *CreateAction`

NewCreateActionWithDefaults instantiates a new CreateAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateAction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmail

`func (o *CreateAction) GetEmail() EmailAction`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateAction) GetEmailOk() (*EmailAction, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateAction) SetEmail(v EmailAction)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateAction) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetKind

`func (o *CreateAction) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateAction) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateAction) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *CreateAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAction) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *CreateAction) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateAction) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateAction) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CreateAction) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetPushNotification

`func (o *CreateAction) GetPushNotification() PushAction`

GetPushNotification returns the PushNotification field if non-nil, zero value otherwise.

### GetPushNotificationOk

`func (o *CreateAction) GetPushNotificationOk() (*PushAction, bool)`

GetPushNotificationOk returns a tuple with the PushNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushNotification

`func (o *CreateAction) SetPushNotification(v PushAction)`

SetPushNotification sets PushNotification field to given value.

### HasPushNotification

`func (o *CreateAction) HasPushNotification() bool`

HasPushNotification returns a boolean if a field has been set.

### GetTriggerId

`func (o *CreateAction) GetTriggerId() string`

GetTriggerId returns the TriggerId field if non-nil, zero value otherwise.

### GetTriggerIdOk

`func (o *CreateAction) GetTriggerIdOk() (*string, bool)`

GetTriggerIdOk returns a tuple with the TriggerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerId

`func (o *CreateAction) SetTriggerId(v string)`

SetTriggerId sets TriggerId field to given value.

### HasTriggerId

`func (o *CreateAction) HasTriggerId() bool`

HasTriggerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


