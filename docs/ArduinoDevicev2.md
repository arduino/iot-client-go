# ArduinoDevicev2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BleMac** | Pointer to **string** |  | [optional] 
**ConnectionType** | Pointer to **string** | The type of the connections selected by the user when multiple connections are available | [optional] 
**CreatedAt** | Pointer to **time.Time** | Creation date of the device | [optional] 
**DeletedAt** | Pointer to **time.Time** | Deletion date of the trigger | [optional] 
**DeviceStatus** | Pointer to **string** | The connection status of the device | [optional] 
**Events** | Pointer to [**[]ArduinoDevicev2SimpleProperties**](ArduinoDevicev2SimpleProperties.md) | ArduinoDevicev2SimplePropertiesCollection is the media type for an array of ArduinoDevicev2SimpleProperties (default view) | [optional] 
**Fqbn** | Pointer to **string** | The fully qualified board name | [optional] 
**Href** | **string** | The api reference of this device | 
**Id** | **string** | The arn of the device | 
**IssuerCa** | Pointer to **string** |  | [optional] 
**Label** | **string** | The label of the device | 
**LastActivityAt** | Pointer to **time.Time** | Last activity date | [optional] 
**LatestWifiFwVersion** | Pointer to **string** | The latest version of the NINA/WIFI101 firmware available for this device | [optional] 
**LibVersion** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | The metadata of the device | [optional] 
**Name** | **string** | The friendly name of the device | 
**NoSketch** | Pointer to **bool** | True if the device type can not have an associated sketch | [optional] 
**OrganizationId** | Pointer to **string** | Id of the organization the device belongs to | [optional] 
**OtaAvailable** | Pointer to **bool** | True if the device type is ready to receive OTA updated | [optional] 
**OtaCompatible** | Pointer to **bool** | True if the device type is OTA compatible | [optional] 
**RequiredWifiFwVersion** | Pointer to **string** | The required version of the NINA/WIFI101 firmware needed by IoT Cloud | [optional] 
**Serial** | **string** | The serial uuid of the device | 
**Tags** | Pointer to **map[string]interface{}** | Tags belonging to the device | [optional] 
**Thing** | Pointer to [**ArduinoThing**](ArduinoThing.md) |  | [optional] 
**Type** | **string** | The type of the device | 
**UniqueHardwareId** | Pointer to **string** | The unique hardware id of the device | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Update date of the trigger | [optional] 
**UserId** | **string** | The id of the user | 
**Webhooks** | Pointer to [**[]ArduinoDevicev2Webhook**](ArduinoDevicev2Webhook.md) | ArduinoDevicev2WebhookCollection is the media type for an array of ArduinoDevicev2Webhook (default view) | [optional] 
**WifiFwVersion** | Pointer to **string** | The version of the NINA/WIFI101 firmware running on the device | [optional] 

## Methods

### NewArduinoDevicev2

`func NewArduinoDevicev2(href string, id string, label string, name string, serial string, type_ string, userId string, ) *ArduinoDevicev2`

NewArduinoDevicev2 instantiates a new ArduinoDevicev2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoDevicev2WithDefaults

`func NewArduinoDevicev2WithDefaults() *ArduinoDevicev2`

NewArduinoDevicev2WithDefaults instantiates a new ArduinoDevicev2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBleMac

`func (o *ArduinoDevicev2) GetBleMac() string`

GetBleMac returns the BleMac field if non-nil, zero value otherwise.

### GetBleMacOk

`func (o *ArduinoDevicev2) GetBleMacOk() (*string, bool)`

GetBleMacOk returns a tuple with the BleMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBleMac

`func (o *ArduinoDevicev2) SetBleMac(v string)`

SetBleMac sets BleMac field to given value.

### HasBleMac

`func (o *ArduinoDevicev2) HasBleMac() bool`

HasBleMac returns a boolean if a field has been set.

### GetConnectionType

`func (o *ArduinoDevicev2) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *ArduinoDevicev2) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *ArduinoDevicev2) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *ArduinoDevicev2) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ArduinoDevicev2) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArduinoDevicev2) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArduinoDevicev2) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ArduinoDevicev2) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ArduinoDevicev2) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ArduinoDevicev2) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ArduinoDevicev2) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ArduinoDevicev2) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDeviceStatus

`func (o *ArduinoDevicev2) GetDeviceStatus() string`

GetDeviceStatus returns the DeviceStatus field if non-nil, zero value otherwise.

### GetDeviceStatusOk

`func (o *ArduinoDevicev2) GetDeviceStatusOk() (*string, bool)`

GetDeviceStatusOk returns a tuple with the DeviceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceStatus

`func (o *ArduinoDevicev2) SetDeviceStatus(v string)`

SetDeviceStatus sets DeviceStatus field to given value.

### HasDeviceStatus

`func (o *ArduinoDevicev2) HasDeviceStatus() bool`

HasDeviceStatus returns a boolean if a field has been set.

### GetEvents

`func (o *ArduinoDevicev2) GetEvents() []ArduinoDevicev2SimpleProperties`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ArduinoDevicev2) GetEventsOk() (*[]ArduinoDevicev2SimpleProperties, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ArduinoDevicev2) SetEvents(v []ArduinoDevicev2SimpleProperties)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ArduinoDevicev2) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetFqbn

`func (o *ArduinoDevicev2) GetFqbn() string`

GetFqbn returns the Fqbn field if non-nil, zero value otherwise.

### GetFqbnOk

`func (o *ArduinoDevicev2) GetFqbnOk() (*string, bool)`

GetFqbnOk returns a tuple with the Fqbn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqbn

`func (o *ArduinoDevicev2) SetFqbn(v string)`

SetFqbn sets Fqbn field to given value.

### HasFqbn

`func (o *ArduinoDevicev2) HasFqbn() bool`

HasFqbn returns a boolean if a field has been set.

### GetHref

`func (o *ArduinoDevicev2) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ArduinoDevicev2) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ArduinoDevicev2) SetHref(v string)`

SetHref sets Href field to given value.


### GetId

`func (o *ArduinoDevicev2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArduinoDevicev2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArduinoDevicev2) SetId(v string)`

SetId sets Id field to given value.


### GetIssuerCa

`func (o *ArduinoDevicev2) GetIssuerCa() string`

GetIssuerCa returns the IssuerCa field if non-nil, zero value otherwise.

### GetIssuerCaOk

`func (o *ArduinoDevicev2) GetIssuerCaOk() (*string, bool)`

GetIssuerCaOk returns a tuple with the IssuerCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCa

`func (o *ArduinoDevicev2) SetIssuerCa(v string)`

SetIssuerCa sets IssuerCa field to given value.

### HasIssuerCa

`func (o *ArduinoDevicev2) HasIssuerCa() bool`

HasIssuerCa returns a boolean if a field has been set.

### GetLabel

`func (o *ArduinoDevicev2) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ArduinoDevicev2) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ArduinoDevicev2) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetLastActivityAt

`func (o *ArduinoDevicev2) GetLastActivityAt() time.Time`

GetLastActivityAt returns the LastActivityAt field if non-nil, zero value otherwise.

### GetLastActivityAtOk

`func (o *ArduinoDevicev2) GetLastActivityAtOk() (*time.Time, bool)`

GetLastActivityAtOk returns a tuple with the LastActivityAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityAt

`func (o *ArduinoDevicev2) SetLastActivityAt(v time.Time)`

SetLastActivityAt sets LastActivityAt field to given value.

### HasLastActivityAt

`func (o *ArduinoDevicev2) HasLastActivityAt() bool`

HasLastActivityAt returns a boolean if a field has been set.

### GetLatestWifiFwVersion

`func (o *ArduinoDevicev2) GetLatestWifiFwVersion() string`

GetLatestWifiFwVersion returns the LatestWifiFwVersion field if non-nil, zero value otherwise.

### GetLatestWifiFwVersionOk

`func (o *ArduinoDevicev2) GetLatestWifiFwVersionOk() (*string, bool)`

GetLatestWifiFwVersionOk returns a tuple with the LatestWifiFwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestWifiFwVersion

`func (o *ArduinoDevicev2) SetLatestWifiFwVersion(v string)`

SetLatestWifiFwVersion sets LatestWifiFwVersion field to given value.

### HasLatestWifiFwVersion

`func (o *ArduinoDevicev2) HasLatestWifiFwVersion() bool`

HasLatestWifiFwVersion returns a boolean if a field has been set.

### GetLibVersion

`func (o *ArduinoDevicev2) GetLibVersion() string`

GetLibVersion returns the LibVersion field if non-nil, zero value otherwise.

### GetLibVersionOk

`func (o *ArduinoDevicev2) GetLibVersionOk() (*string, bool)`

GetLibVersionOk returns a tuple with the LibVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibVersion

`func (o *ArduinoDevicev2) SetLibVersion(v string)`

SetLibVersion sets LibVersion field to given value.

### HasLibVersion

`func (o *ArduinoDevicev2) HasLibVersion() bool`

HasLibVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *ArduinoDevicev2) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ArduinoDevicev2) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ArduinoDevicev2) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ArduinoDevicev2) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *ArduinoDevicev2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArduinoDevicev2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArduinoDevicev2) SetName(v string)`

SetName sets Name field to given value.


### GetNoSketch

`func (o *ArduinoDevicev2) GetNoSketch() bool`

GetNoSketch returns the NoSketch field if non-nil, zero value otherwise.

### GetNoSketchOk

`func (o *ArduinoDevicev2) GetNoSketchOk() (*bool, bool)`

GetNoSketchOk returns a tuple with the NoSketch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoSketch

`func (o *ArduinoDevicev2) SetNoSketch(v bool)`

SetNoSketch sets NoSketch field to given value.

### HasNoSketch

`func (o *ArduinoDevicev2) HasNoSketch() bool`

HasNoSketch returns a boolean if a field has been set.

### GetOrganizationId

`func (o *ArduinoDevicev2) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ArduinoDevicev2) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ArduinoDevicev2) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ArduinoDevicev2) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetOtaAvailable

`func (o *ArduinoDevicev2) GetOtaAvailable() bool`

GetOtaAvailable returns the OtaAvailable field if non-nil, zero value otherwise.

### GetOtaAvailableOk

`func (o *ArduinoDevicev2) GetOtaAvailableOk() (*bool, bool)`

GetOtaAvailableOk returns a tuple with the OtaAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtaAvailable

`func (o *ArduinoDevicev2) SetOtaAvailable(v bool)`

SetOtaAvailable sets OtaAvailable field to given value.

### HasOtaAvailable

`func (o *ArduinoDevicev2) HasOtaAvailable() bool`

HasOtaAvailable returns a boolean if a field has been set.

### GetOtaCompatible

`func (o *ArduinoDevicev2) GetOtaCompatible() bool`

GetOtaCompatible returns the OtaCompatible field if non-nil, zero value otherwise.

### GetOtaCompatibleOk

`func (o *ArduinoDevicev2) GetOtaCompatibleOk() (*bool, bool)`

GetOtaCompatibleOk returns a tuple with the OtaCompatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtaCompatible

`func (o *ArduinoDevicev2) SetOtaCompatible(v bool)`

SetOtaCompatible sets OtaCompatible field to given value.

### HasOtaCompatible

`func (o *ArduinoDevicev2) HasOtaCompatible() bool`

HasOtaCompatible returns a boolean if a field has been set.

### GetRequiredWifiFwVersion

`func (o *ArduinoDevicev2) GetRequiredWifiFwVersion() string`

GetRequiredWifiFwVersion returns the RequiredWifiFwVersion field if non-nil, zero value otherwise.

### GetRequiredWifiFwVersionOk

`func (o *ArduinoDevicev2) GetRequiredWifiFwVersionOk() (*string, bool)`

GetRequiredWifiFwVersionOk returns a tuple with the RequiredWifiFwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredWifiFwVersion

`func (o *ArduinoDevicev2) SetRequiredWifiFwVersion(v string)`

SetRequiredWifiFwVersion sets RequiredWifiFwVersion field to given value.

### HasRequiredWifiFwVersion

`func (o *ArduinoDevicev2) HasRequiredWifiFwVersion() bool`

HasRequiredWifiFwVersion returns a boolean if a field has been set.

### GetSerial

`func (o *ArduinoDevicev2) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *ArduinoDevicev2) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *ArduinoDevicev2) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetTags

`func (o *ArduinoDevicev2) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ArduinoDevicev2) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ArduinoDevicev2) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *ArduinoDevicev2) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThing

`func (o *ArduinoDevicev2) GetThing() ArduinoThing`

GetThing returns the Thing field if non-nil, zero value otherwise.

### GetThingOk

`func (o *ArduinoDevicev2) GetThingOk() (*ArduinoThing, bool)`

GetThingOk returns a tuple with the Thing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThing

`func (o *ArduinoDevicev2) SetThing(v ArduinoThing)`

SetThing sets Thing field to given value.

### HasThing

`func (o *ArduinoDevicev2) HasThing() bool`

HasThing returns a boolean if a field has been set.

### GetType

`func (o *ArduinoDevicev2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArduinoDevicev2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArduinoDevicev2) SetType(v string)`

SetType sets Type field to given value.


### GetUniqueHardwareId

`func (o *ArduinoDevicev2) GetUniqueHardwareId() string`

GetUniqueHardwareId returns the UniqueHardwareId field if non-nil, zero value otherwise.

### GetUniqueHardwareIdOk

`func (o *ArduinoDevicev2) GetUniqueHardwareIdOk() (*string, bool)`

GetUniqueHardwareIdOk returns a tuple with the UniqueHardwareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueHardwareId

`func (o *ArduinoDevicev2) SetUniqueHardwareId(v string)`

SetUniqueHardwareId sets UniqueHardwareId field to given value.

### HasUniqueHardwareId

`func (o *ArduinoDevicev2) HasUniqueHardwareId() bool`

HasUniqueHardwareId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ArduinoDevicev2) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ArduinoDevicev2) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ArduinoDevicev2) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ArduinoDevicev2) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserId

`func (o *ArduinoDevicev2) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ArduinoDevicev2) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ArduinoDevicev2) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetWebhooks

`func (o *ArduinoDevicev2) GetWebhooks() []ArduinoDevicev2Webhook`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *ArduinoDevicev2) GetWebhooksOk() (*[]ArduinoDevicev2Webhook, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *ArduinoDevicev2) SetWebhooks(v []ArduinoDevicev2Webhook)`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *ArduinoDevicev2) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.

### GetWifiFwVersion

`func (o *ArduinoDevicev2) GetWifiFwVersion() string`

GetWifiFwVersion returns the WifiFwVersion field if non-nil, zero value otherwise.

### GetWifiFwVersionOk

`func (o *ArduinoDevicev2) GetWifiFwVersionOk() (*string, bool)`

GetWifiFwVersionOk returns a tuple with the WifiFwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiFwVersion

`func (o *ArduinoDevicev2) SetWifiFwVersion(v string)`

SetWifiFwVersion sets WifiFwVersion field to given value.

### HasWifiFwVersion

`func (o *ArduinoDevicev2) HasWifiFwVersion() bool`

HasWifiFwVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


