# UpdateAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the trigger | [optional] 
**Email** | Pointer to [**EmailAction**](EmailAction.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the action | [optional] 
**PushNotification** | Pointer to [**PushAction**](PushAction.md) |  | [optional] 
**TriggerId** | Pointer to **string** | Id of the trigger the action is associated to | [optional] 

## Methods

### NewUpdateAction

`func NewUpdateAction() *UpdateAction`

NewUpdateAction instantiates a new UpdateAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateActionWithDefaults

`func NewUpdateActionWithDefaults() *UpdateAction`

NewUpdateActionWithDefaults instantiates a new UpdateAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateAction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateAction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateAction) GetEmail() EmailAction`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateAction) GetEmailOk() (*EmailAction, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateAction) SetEmail(v EmailAction)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateAction) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *UpdateAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPushNotification

`func (o *UpdateAction) GetPushNotification() PushAction`

GetPushNotification returns the PushNotification field if non-nil, zero value otherwise.

### GetPushNotificationOk

`func (o *UpdateAction) GetPushNotificationOk() (*PushAction, bool)`

GetPushNotificationOk returns a tuple with the PushNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushNotification

`func (o *UpdateAction) SetPushNotification(v PushAction)`

SetPushNotification sets PushNotification field to given value.

### HasPushNotification

`func (o *UpdateAction) HasPushNotification() bool`

HasPushNotification returns a boolean if a field has been set.

### GetTriggerId

`func (o *UpdateAction) GetTriggerId() string`

GetTriggerId returns the TriggerId field if non-nil, zero value otherwise.

### GetTriggerIdOk

`func (o *UpdateAction) GetTriggerIdOk() (*string, bool)`

GetTriggerIdOk returns a tuple with the TriggerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerId

`func (o *UpdateAction) SetTriggerId(v string)`

SetTriggerId sets TriggerId field to given value.

### HasTriggerId

`func (o *UpdateAction) HasTriggerId() bool`

HasTriggerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


