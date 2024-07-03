# ThingCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assistant** | Pointer to **string** | The kind of voice assistant the thing is connected to, it can be ALEXA | GOOGLE | NONE | [optional] 
**DeviceId** | Pointer to **string** | The arn of the associated device | [optional] 
**Id** | Pointer to **string** | The id of the thing | [optional] 
**Name** | Pointer to **string** | The friendly name of the thing | [optional] 
**Properties** | Pointer to [**[]Property**](Property.md) | The properties of the thing | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Optional set of tags | [optional] 
**Timezone** | Pointer to **string** | A time zone name Check /v2/timezones for a list of valid names. | [optional] [default to "America/New_York"]
**WebhookActive** | Pointer to **bool** | Webhook uri | [optional] 
**WebhookUri** | Pointer to **string** | Webhook uri | [optional] 

## Methods

### NewThingCreate

`func NewThingCreate() *ThingCreate`

NewThingCreate instantiates a new ThingCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThingCreateWithDefaults

`func NewThingCreateWithDefaults() *ThingCreate`

NewThingCreateWithDefaults instantiates a new ThingCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssistant

`func (o *ThingCreate) GetAssistant() string`

GetAssistant returns the Assistant field if non-nil, zero value otherwise.

### GetAssistantOk

`func (o *ThingCreate) GetAssistantOk() (*string, bool)`

GetAssistantOk returns a tuple with the Assistant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistant

`func (o *ThingCreate) SetAssistant(v string)`

SetAssistant sets Assistant field to given value.

### HasAssistant

`func (o *ThingCreate) HasAssistant() bool`

HasAssistant returns a boolean if a field has been set.

### GetDeviceId

`func (o *ThingCreate) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *ThingCreate) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *ThingCreate) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *ThingCreate) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetId

`func (o *ThingCreate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThingCreate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThingCreate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ThingCreate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ThingCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ThingCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ThingCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ThingCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProperties

`func (o *ThingCreate) GetProperties() []Property`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ThingCreate) GetPropertiesOk() (*[]Property, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ThingCreate) SetProperties(v []Property)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ThingCreate) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetTags

`func (o *ThingCreate) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ThingCreate) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ThingCreate) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ThingCreate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimezone

`func (o *ThingCreate) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ThingCreate) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ThingCreate) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ThingCreate) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetWebhookActive

`func (o *ThingCreate) GetWebhookActive() bool`

GetWebhookActive returns the WebhookActive field if non-nil, zero value otherwise.

### GetWebhookActiveOk

`func (o *ThingCreate) GetWebhookActiveOk() (*bool, bool)`

GetWebhookActiveOk returns a tuple with the WebhookActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookActive

`func (o *ThingCreate) SetWebhookActive(v bool)`

SetWebhookActive sets WebhookActive field to given value.

### HasWebhookActive

`func (o *ThingCreate) HasWebhookActive() bool`

HasWebhookActive returns a boolean if a field has been set.

### GetWebhookUri

`func (o *ThingCreate) GetWebhookUri() string`

GetWebhookUri returns the WebhookUri field if non-nil, zero value otherwise.

### GetWebhookUriOk

`func (o *ThingCreate) GetWebhookUriOk() (*string, bool)`

GetWebhookUriOk returns a tuple with the WebhookUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUri

`func (o *ThingCreate) SetWebhookUri(v string)`

SetWebhookUri sets WebhookUri field to given value.

### HasWebhookUri

`func (o *ThingCreate) HasWebhookUri() bool`

HasWebhookUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


