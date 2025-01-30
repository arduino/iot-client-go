# CreateClaimedDevicesV2Payload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** | The id of the device | [optional] 
**BleMac** | Pointer to **string** |  | [optional] 
**ConnectionType** | Pointer to **string** | The type of the connections selected by the user when multiple connections are available | [optional] 
**Fqbn** | Pointer to **string** | The fully qualified board name | [optional] 
**Name** | Pointer to **string** | The friendly name of the device | [optional] 
**Serial** | Pointer to **string** | The serial uuid of the device | [optional] 
**Type** | **string** | The type of the device | 
**UniqueHardwareId** | **string** | The unique hardware id of the device | 
**UserId** | **string** | The user_id associated to the device. If absent it will be inferred from the authentication header | 
**WifiFwVersion** | Pointer to **string** | The version of the NINA/WIFI101 firmware running on the device | [optional] 

## Methods

### NewCreateClaimedDevicesV2Payload

`func NewCreateClaimedDevicesV2Payload(type_ string, uniqueHardwareId string, userId string, ) *CreateClaimedDevicesV2Payload`

NewCreateClaimedDevicesV2Payload instantiates a new CreateClaimedDevicesV2Payload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClaimedDevicesV2PayloadWithDefaults

`func NewCreateClaimedDevicesV2PayloadWithDefaults() *CreateClaimedDevicesV2Payload`

NewCreateClaimedDevicesV2PayloadWithDefaults instantiates a new CreateClaimedDevicesV2Payload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *CreateClaimedDevicesV2Payload) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *CreateClaimedDevicesV2Payload) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *CreateClaimedDevicesV2Payload) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *CreateClaimedDevicesV2Payload) HasID() bool`

HasID returns a boolean if a field has been set.

### GetBleMac

`func (o *CreateClaimedDevicesV2Payload) GetBleMac() string`

GetBleMac returns the BleMac field if non-nil, zero value otherwise.

### GetBleMacOk

`func (o *CreateClaimedDevicesV2Payload) GetBleMacOk() (*string, bool)`

GetBleMacOk returns a tuple with the BleMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBleMac

`func (o *CreateClaimedDevicesV2Payload) SetBleMac(v string)`

SetBleMac sets BleMac field to given value.

### HasBleMac

`func (o *CreateClaimedDevicesV2Payload) HasBleMac() bool`

HasBleMac returns a boolean if a field has been set.

### GetConnectionType

`func (o *CreateClaimedDevicesV2Payload) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *CreateClaimedDevicesV2Payload) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *CreateClaimedDevicesV2Payload) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *CreateClaimedDevicesV2Payload) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetFqbn

`func (o *CreateClaimedDevicesV2Payload) GetFqbn() string`

GetFqbn returns the Fqbn field if non-nil, zero value otherwise.

### GetFqbnOk

`func (o *CreateClaimedDevicesV2Payload) GetFqbnOk() (*string, bool)`

GetFqbnOk returns a tuple with the Fqbn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqbn

`func (o *CreateClaimedDevicesV2Payload) SetFqbn(v string)`

SetFqbn sets Fqbn field to given value.

### HasFqbn

`func (o *CreateClaimedDevicesV2Payload) HasFqbn() bool`

HasFqbn returns a boolean if a field has been set.

### GetName

`func (o *CreateClaimedDevicesV2Payload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateClaimedDevicesV2Payload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateClaimedDevicesV2Payload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateClaimedDevicesV2Payload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSerial

`func (o *CreateClaimedDevicesV2Payload) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *CreateClaimedDevicesV2Payload) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *CreateClaimedDevicesV2Payload) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *CreateClaimedDevicesV2Payload) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetType

`func (o *CreateClaimedDevicesV2Payload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateClaimedDevicesV2Payload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateClaimedDevicesV2Payload) SetType(v string)`

SetType sets Type field to given value.


### GetUniqueHardwareId

`func (o *CreateClaimedDevicesV2Payload) GetUniqueHardwareId() string`

GetUniqueHardwareId returns the UniqueHardwareId field if non-nil, zero value otherwise.

### GetUniqueHardwareIdOk

`func (o *CreateClaimedDevicesV2Payload) GetUniqueHardwareIdOk() (*string, bool)`

GetUniqueHardwareIdOk returns a tuple with the UniqueHardwareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueHardwareId

`func (o *CreateClaimedDevicesV2Payload) SetUniqueHardwareId(v string)`

SetUniqueHardwareId sets UniqueHardwareId field to given value.


### GetUserId

`func (o *CreateClaimedDevicesV2Payload) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateClaimedDevicesV2Payload) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateClaimedDevicesV2Payload) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetWifiFwVersion

`func (o *CreateClaimedDevicesV2Payload) GetWifiFwVersion() string`

GetWifiFwVersion returns the WifiFwVersion field if non-nil, zero value otherwise.

### GetWifiFwVersionOk

`func (o *CreateClaimedDevicesV2Payload) GetWifiFwVersionOk() (*string, bool)`

GetWifiFwVersionOk returns a tuple with the WifiFwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiFwVersion

`func (o *CreateClaimedDevicesV2Payload) SetWifiFwVersion(v string)`

SetWifiFwVersion sets WifiFwVersion field to given value.

### HasWifiFwVersion

`func (o *CreateClaimedDevicesV2Payload) HasWifiFwVersion() bool`

HasWifiFwVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


