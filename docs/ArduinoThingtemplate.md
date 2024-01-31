# ArduinoThingtemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The friendly name of the thing | 
**OrganizationId** | Pointer to **string** | Id of the organization the thing belongs to | [optional] 
**Properties** | Pointer to [**[]ArduinoTemplateproperty**](ArduinoTemplateproperty.md) | ArduinoTemplatepropertyCollection is the media type for an array of ArduinoTemplateproperty (default view) | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Tags of the thing | [optional] 
**Timezone** | **string** | Time zone of the thing | 
**WebhookUri** | Pointer to **string** | Webhook uri | [optional] 

## Methods

### NewArduinoThingtemplate

`func NewArduinoThingtemplate(name string, timezone string, ) *ArduinoThingtemplate`

NewArduinoThingtemplate instantiates a new ArduinoThingtemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoThingtemplateWithDefaults

`func NewArduinoThingtemplateWithDefaults() *ArduinoThingtemplate`

NewArduinoThingtemplateWithDefaults instantiates a new ArduinoThingtemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ArduinoThingtemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArduinoThingtemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArduinoThingtemplate) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *ArduinoThingtemplate) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ArduinoThingtemplate) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ArduinoThingtemplate) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ArduinoThingtemplate) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetProperties

`func (o *ArduinoThingtemplate) GetProperties() []ArduinoTemplateproperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ArduinoThingtemplate) GetPropertiesOk() (*[]ArduinoTemplateproperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ArduinoThingtemplate) SetProperties(v []ArduinoTemplateproperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ArduinoThingtemplate) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetTags

`func (o *ArduinoThingtemplate) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ArduinoThingtemplate) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ArduinoThingtemplate) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ArduinoThingtemplate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimezone

`func (o *ArduinoThingtemplate) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ArduinoThingtemplate) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ArduinoThingtemplate) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetWebhookUri

`func (o *ArduinoThingtemplate) GetWebhookUri() string`

GetWebhookUri returns the WebhookUri field if non-nil, zero value otherwise.

### GetWebhookUriOk

`func (o *ArduinoThingtemplate) GetWebhookUriOk() (*string, bool)`

GetWebhookUriOk returns a tuple with the WebhookUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUri

`func (o *ArduinoThingtemplate) SetWebhookUri(v string)`

SetWebhookUri sets WebhookUri field to given value.

### HasWebhookUri

`func (o *ArduinoThingtemplate) HasWebhookUri() bool`

HasWebhookUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


