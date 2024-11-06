# ThingUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assistant** | Pointer to **string** | The kind of voice assistant the thing is connected to, it can be ALEXA | GOOGLE | NONE | [optional] 
**DeviceId** | Pointer to **string** | The arn of the associated device | [optional] 
**Id** | Pointer to **string** | The id of the thing | [optional] 
**Name** | Pointer to **string** | The friendly name of the thing | [optional] 
**Properties** | Pointer to [**[]Property**](Property.md) | The properties of the thing | [optional] 
**SoftDeleted** | Pointer to **bool** | If false, restore the thing from the soft deletion | [optional] [default to false]
**Timezone** | Pointer to **string** | A time zone name. Check /v2/timezones for a list of valid names. | [optional] 
**WebhookActive** | Pointer to **bool** | Webhook uri | [optional] 
**WebhookUri** | Pointer to **string** | Webhook uri | [optional] 

## Methods

### NewThingUpdate

`func NewThingUpdate() *ThingUpdate`

NewThingUpdate instantiates a new ThingUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThingUpdateWithDefaults

`func NewThingUpdateWithDefaults() *ThingUpdate`

NewThingUpdateWithDefaults instantiates a new ThingUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssistant

`func (o *ThingUpdate) GetAssistant() string`

GetAssistant returns the Assistant field if non-nil, zero value otherwise.

### GetAssistantOk

`func (o *ThingUpdate) GetAssistantOk() (*string, bool)`

GetAssistantOk returns a tuple with the Assistant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistant

`func (o *ThingUpdate) SetAssistant(v string)`

SetAssistant sets Assistant field to given value.

### HasAssistant

`func (o *ThingUpdate) HasAssistant() bool`

HasAssistant returns a boolean if a field has been set.

### GetDeviceId

`func (o *ThingUpdate) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *ThingUpdate) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *ThingUpdate) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *ThingUpdate) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetId

`func (o *ThingUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThingUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThingUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ThingUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ThingUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ThingUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ThingUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ThingUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProperties

`func (o *ThingUpdate) GetProperties() []Property`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ThingUpdate) GetPropertiesOk() (*[]Property, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ThingUpdate) SetProperties(v []Property)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ThingUpdate) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetSoftDeleted

`func (o *ThingUpdate) GetSoftDeleted() bool`

GetSoftDeleted returns the SoftDeleted field if non-nil, zero value otherwise.

### GetSoftDeletedOk

`func (o *ThingUpdate) GetSoftDeletedOk() (*bool, bool)`

GetSoftDeletedOk returns a tuple with the SoftDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftDeleted

`func (o *ThingUpdate) SetSoftDeleted(v bool)`

SetSoftDeleted sets SoftDeleted field to given value.

### HasSoftDeleted

`func (o *ThingUpdate) HasSoftDeleted() bool`

HasSoftDeleted returns a boolean if a field has been set.

### GetTimezone

`func (o *ThingUpdate) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ThingUpdate) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ThingUpdate) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ThingUpdate) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetWebhookActive

`func (o *ThingUpdate) GetWebhookActive() bool`

GetWebhookActive returns the WebhookActive field if non-nil, zero value otherwise.

### GetWebhookActiveOk

`func (o *ThingUpdate) GetWebhookActiveOk() (*bool, bool)`

GetWebhookActiveOk returns a tuple with the WebhookActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookActive

`func (o *ThingUpdate) SetWebhookActive(v bool)`

SetWebhookActive sets WebhookActive field to given value.

### HasWebhookActive

`func (o *ThingUpdate) HasWebhookActive() bool`

HasWebhookActive returns a boolean if a field has been set.

### GetWebhookUri

`func (o *ThingUpdate) GetWebhookUri() string`

GetWebhookUri returns the WebhookUri field if non-nil, zero value otherwise.

### GetWebhookUriOk

`func (o *ThingUpdate) GetWebhookUriOk() (*string, bool)`

GetWebhookUriOk returns a tuple with the WebhookUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUri

`func (o *ThingUpdate) SetWebhookUri(v string)`

SetWebhookUri sets WebhookUri field to given value.

### HasWebhookUri

`func (o *ThingUpdate) HasWebhookUri() bool`

HasWebhookUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


