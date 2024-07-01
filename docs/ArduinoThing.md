# ArduinoThing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assistant** | Pointer to **string** | The kind of voice assistant the thing is connected to, it can be ALEXA | GOOGLE | NONE | [optional] 
**CreatedAt** | Pointer to **time.Time** | Creation date of the thing | [optional] 
**DeletedAt** | Pointer to **time.Time** | Delete date of the thing | [optional] 
**DeviceFqbn** | Pointer to **string** | The fqbn of the attached device, if any | [optional] 
**DeviceId** | Pointer to **string** | The id of the device | [optional] 
**DeviceName** | Pointer to **string** | The name of the attached device, if any | [optional] 
**DeviceType** | Pointer to **string** | The type of the attached device, if any | [optional] 
**Href** | **string** | The api reference of this thing | 
**Id** | **string** | The id of the thing | 
**Name** | **string** | The friendly name of the thing | 
**OrganizationId** | Pointer to **string** | Id of the organization the thing belongs to | [optional] 
**Properties** | Pointer to [**[]ArduinoProperty**](ArduinoProperty.md) | ArduinoPropertyCollection is the media type for an array of ArduinoProperty (default view) | [optional] 
**PropertiesCount** | Pointer to **int64** | The number of properties of the thing | [optional] 
**SketchId** | Pointer to **string** | The id of the attached sketch | [optional] 
**Tags** | Pointer to **map[string]interface{}** | Tags of the thing | [optional] 
**Timezone** | **string** | Time zone of the thing | 
**UpdatedAt** | Pointer to **time.Time** | Update date of the thing | [optional] 
**UserId** | **string** | The user id of the owner | 
**WebhookActive** | Pointer to **bool** | Webhook uri | [optional] 
**WebhookUri** | Pointer to **string** | Webhook uri | [optional] 

## Methods

### NewArduinoThing

`func NewArduinoThing(href string, id string, name string, timezone string, userId string, ) *ArduinoThing`

NewArduinoThing instantiates a new ArduinoThing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoThingWithDefaults

`func NewArduinoThingWithDefaults() *ArduinoThing`

NewArduinoThingWithDefaults instantiates a new ArduinoThing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssistant

`func (o *ArduinoThing) GetAssistant() string`

GetAssistant returns the Assistant field if non-nil, zero value otherwise.

### GetAssistantOk

`func (o *ArduinoThing) GetAssistantOk() (*string, bool)`

GetAssistantOk returns a tuple with the Assistant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistant

`func (o *ArduinoThing) SetAssistant(v string)`

SetAssistant sets Assistant field to given value.

### HasAssistant

`func (o *ArduinoThing) HasAssistant() bool`

HasAssistant returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ArduinoThing) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArduinoThing) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArduinoThing) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ArduinoThing) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ArduinoThing) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ArduinoThing) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ArduinoThing) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ArduinoThing) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDeviceFqbn

`func (o *ArduinoThing) GetDeviceFqbn() string`

GetDeviceFqbn returns the DeviceFqbn field if non-nil, zero value otherwise.

### GetDeviceFqbnOk

`func (o *ArduinoThing) GetDeviceFqbnOk() (*string, bool)`

GetDeviceFqbnOk returns a tuple with the DeviceFqbn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceFqbn

`func (o *ArduinoThing) SetDeviceFqbn(v string)`

SetDeviceFqbn sets DeviceFqbn field to given value.

### HasDeviceFqbn

`func (o *ArduinoThing) HasDeviceFqbn() bool`

HasDeviceFqbn returns a boolean if a field has been set.

### GetDeviceId

`func (o *ArduinoThing) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *ArduinoThing) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *ArduinoThing) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *ArduinoThing) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceName

`func (o *ArduinoThing) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *ArduinoThing) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *ArduinoThing) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *ArduinoThing) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceType

`func (o *ArduinoThing) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *ArduinoThing) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *ArduinoThing) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *ArduinoThing) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetHref

`func (o *ArduinoThing) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ArduinoThing) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ArduinoThing) SetHref(v string)`

SetHref sets Href field to given value.


### GetId

`func (o *ArduinoThing) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArduinoThing) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArduinoThing) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ArduinoThing) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArduinoThing) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArduinoThing) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *ArduinoThing) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ArduinoThing) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ArduinoThing) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ArduinoThing) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetProperties

`func (o *ArduinoThing) GetProperties() []ArduinoProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ArduinoThing) GetPropertiesOk() (*[]ArduinoProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ArduinoThing) SetProperties(v []ArduinoProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ArduinoThing) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetPropertiesCount

`func (o *ArduinoThing) GetPropertiesCount() int64`

GetPropertiesCount returns the PropertiesCount field if non-nil, zero value otherwise.

### GetPropertiesCountOk

`func (o *ArduinoThing) GetPropertiesCountOk() (*int64, bool)`

GetPropertiesCountOk returns a tuple with the PropertiesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertiesCount

`func (o *ArduinoThing) SetPropertiesCount(v int64)`

SetPropertiesCount sets PropertiesCount field to given value.

### HasPropertiesCount

`func (o *ArduinoThing) HasPropertiesCount() bool`

HasPropertiesCount returns a boolean if a field has been set.

### GetSketchId

`func (o *ArduinoThing) GetSketchId() string`

GetSketchId returns the SketchId field if non-nil, zero value otherwise.

### GetSketchIdOk

`func (o *ArduinoThing) GetSketchIdOk() (*string, bool)`

GetSketchIdOk returns a tuple with the SketchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSketchId

`func (o *ArduinoThing) SetSketchId(v string)`

SetSketchId sets SketchId field to given value.

### HasSketchId

`func (o *ArduinoThing) HasSketchId() bool`

HasSketchId returns a boolean if a field has been set.

### GetTags

`func (o *ArduinoThing) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ArduinoThing) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ArduinoThing) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *ArduinoThing) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimezone

`func (o *ArduinoThing) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ArduinoThing) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ArduinoThing) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetUpdatedAt

`func (o *ArduinoThing) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ArduinoThing) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ArduinoThing) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ArduinoThing) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserId

`func (o *ArduinoThing) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ArduinoThing) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ArduinoThing) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetWebhookActive

`func (o *ArduinoThing) GetWebhookActive() bool`

GetWebhookActive returns the WebhookActive field if non-nil, zero value otherwise.

### GetWebhookActiveOk

`func (o *ArduinoThing) GetWebhookActiveOk() (*bool, bool)`

GetWebhookActiveOk returns a tuple with the WebhookActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookActive

`func (o *ArduinoThing) SetWebhookActive(v bool)`

SetWebhookActive sets WebhookActive field to given value.

### HasWebhookActive

`func (o *ArduinoThing) HasWebhookActive() bool`

HasWebhookActive returns a boolean if a field has been set.

### GetWebhookUri

`func (o *ArduinoThing) GetWebhookUri() string`

GetWebhookUri returns the WebhookUri field if non-nil, zero value otherwise.

### GetWebhookUriOk

`func (o *ArduinoThing) GetWebhookUriOk() (*string, bool)`

GetWebhookUriOk returns a tuple with the WebhookUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUri

`func (o *ArduinoThing) SetWebhookUri(v string)`

SetWebhookUri sets WebhookUri field to given value.

### HasWebhookUri

`func (o *ArduinoThing) HasWebhookUri() bool`

HasWebhookUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


