# ArduinoThingtemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceMetadata** | Pointer to [**ArduinoDevicev2templatedevice**](ArduinoDevicev2templatedevice.md) |  | [optional] 
**Id** | Pointer to **string** | The friendly id of the thing | [optional] 
**Name** | **string** | The friendly name of the thing | 
**OrganizationId** | Pointer to **string** | Id of the organization the thing belongs to | [optional] 
**SketchTemplate** | Pointer to **string** | The ID of the template&#39;s sketch | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Tags of the thing | [optional] 
**Timezone** | **string** | Time zone of the thing | 
**Variables** | Pointer to [**[]ArduinoTemplateproperty**](ArduinoTemplateproperty.md) | ArduinoTemplatepropertyCollection is the media type for an array of ArduinoTemplateproperty (default view) | [optional] 
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

### GetDeviceMetadata

`func (o *ArduinoThingtemplate) GetDeviceMetadata() ArduinoDevicev2templatedevice`

GetDeviceMetadata returns the DeviceMetadata field if non-nil, zero value otherwise.

### GetDeviceMetadataOk

`func (o *ArduinoThingtemplate) GetDeviceMetadataOk() (*ArduinoDevicev2templatedevice, bool)`

GetDeviceMetadataOk returns a tuple with the DeviceMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMetadata

`func (o *ArduinoThingtemplate) SetDeviceMetadata(v ArduinoDevicev2templatedevice)`

SetDeviceMetadata sets DeviceMetadata field to given value.

### HasDeviceMetadata

`func (o *ArduinoThingtemplate) HasDeviceMetadata() bool`

HasDeviceMetadata returns a boolean if a field has been set.

### GetId

`func (o *ArduinoThingtemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArduinoThingtemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArduinoThingtemplate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ArduinoThingtemplate) HasId() bool`

HasId returns a boolean if a field has been set.

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

### GetSketchTemplate

`func (o *ArduinoThingtemplate) GetSketchTemplate() string`

GetSketchTemplate returns the SketchTemplate field if non-nil, zero value otherwise.

### GetSketchTemplateOk

`func (o *ArduinoThingtemplate) GetSketchTemplateOk() (*string, bool)`

GetSketchTemplateOk returns a tuple with the SketchTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSketchTemplate

`func (o *ArduinoThingtemplate) SetSketchTemplate(v string)`

SetSketchTemplate sets SketchTemplate field to given value.

### HasSketchTemplate

`func (o *ArduinoThingtemplate) HasSketchTemplate() bool`

HasSketchTemplate returns a boolean if a field has been set.

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


### GetVariables

`func (o *ArduinoThingtemplate) GetVariables() []ArduinoTemplateproperty`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ArduinoThingtemplate) GetVariablesOk() (*[]ArduinoTemplateproperty, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ArduinoThingtemplate) SetVariables(v []ArduinoTemplateproperty)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *ArduinoThingtemplate) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

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


