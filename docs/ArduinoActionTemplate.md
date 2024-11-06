# ArduinoActionTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the action | [optional] 
**Email** | Pointer to [**EmailAction**](EmailAction.md) |  | [optional] 
**Kind** | Pointer to **string** | The kind of the action | [optional] 
**Name** | Pointer to **string** | The name of the action | [optional] 
**OrganizationId** | Pointer to **string** | Id of the organization the trigger belongs to | [optional] 
**PushNotification** | Pointer to [**PushAction**](PushAction.md) |  | [optional] 

## Methods

### NewArduinoActionTemplate

`func NewArduinoActionTemplate() *ArduinoActionTemplate`

NewArduinoActionTemplate instantiates a new ArduinoActionTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoActionTemplateWithDefaults

`func NewArduinoActionTemplateWithDefaults() *ArduinoActionTemplate`

NewArduinoActionTemplateWithDefaults instantiates a new ArduinoActionTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ArduinoActionTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ArduinoActionTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ArduinoActionTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ArduinoActionTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmail

`func (o *ArduinoActionTemplate) GetEmail() EmailAction`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ArduinoActionTemplate) GetEmailOk() (*EmailAction, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ArduinoActionTemplate) SetEmail(v EmailAction)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ArduinoActionTemplate) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetKind

`func (o *ArduinoActionTemplate) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ArduinoActionTemplate) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ArduinoActionTemplate) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ArduinoActionTemplate) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *ArduinoActionTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArduinoActionTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArduinoActionTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ArduinoActionTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizationId

`func (o *ArduinoActionTemplate) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ArduinoActionTemplate) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ArduinoActionTemplate) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ArduinoActionTemplate) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetPushNotification

`func (o *ArduinoActionTemplate) GetPushNotification() PushAction`

GetPushNotification returns the PushNotification field if non-nil, zero value otherwise.

### GetPushNotificationOk

`func (o *ArduinoActionTemplate) GetPushNotificationOk() (*PushAction, bool)`

GetPushNotificationOk returns a tuple with the PushNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushNotification

`func (o *ArduinoActionTemplate) SetPushNotification(v PushAction)`

SetPushNotification sets PushNotification field to given value.

### HasPushNotification

`func (o *ArduinoActionTemplate) HasPushNotification() bool`

HasPushNotification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


