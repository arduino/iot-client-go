# CreateDevicesV2Payload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BleMac** | Pointer to **string** |  | [optional] 
**ConnectionType** | Pointer to **string** | The type of the connections selected by the user when multiple connections are available | [optional] 
**Fqbn** | Pointer to **string** | The fully qualified board name | [optional] 
**Name** | Pointer to **string** | The friendly name of the device | [optional] 
**Serial** | Pointer to **string** | The serial uuid of the device | [optional] 
**SoftDeleted** | Pointer to **bool** | If false, restore the thing from the soft deletion | [optional] [default to false]
**Type** | **string** | The type of the device | 
**UniqueHardwareId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** | The user_id associated to the device. If absent it will be inferred from the authentication header | [optional] 
**WifiFwVersion** | Pointer to **string** | The version of the NINA/WIFI101 firmware running on the device | [optional] 

## Methods

### NewCreateDevicesV2Payload

`func NewCreateDevicesV2Payload(type_ string, ) *CreateDevicesV2Payload`

NewCreateDevicesV2Payload instantiates a new CreateDevicesV2Payload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDevicesV2PayloadWithDefaults

`func NewCreateDevicesV2PayloadWithDefaults() *CreateDevicesV2Payload`

NewCreateDevicesV2PayloadWithDefaults instantiates a new CreateDevicesV2Payload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBleMac

`func (o *CreateDevicesV2Payload) GetBleMac() string`

GetBleMac returns the BleMac field if non-nil, zero value otherwise.

### GetBleMacOk

`func (o *CreateDevicesV2Payload) GetBleMacOk() (*string, bool)`

GetBleMacOk returns a tuple with the BleMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBleMac

`func (o *CreateDevicesV2Payload) SetBleMac(v string)`

SetBleMac sets BleMac field to given value.

### HasBleMac

`func (o *CreateDevicesV2Payload) HasBleMac() bool`

HasBleMac returns a boolean if a field has been set.

### GetConnectionType

`func (o *CreateDevicesV2Payload) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *CreateDevicesV2Payload) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *CreateDevicesV2Payload) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *CreateDevicesV2Payload) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetFqbn

`func (o *CreateDevicesV2Payload) GetFqbn() string`

GetFqbn returns the Fqbn field if non-nil, zero value otherwise.

### GetFqbnOk

`func (o *CreateDevicesV2Payload) GetFqbnOk() (*string, bool)`

GetFqbnOk returns a tuple with the Fqbn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqbn

`func (o *CreateDevicesV2Payload) SetFqbn(v string)`

SetFqbn sets Fqbn field to given value.

### HasFqbn

`func (o *CreateDevicesV2Payload) HasFqbn() bool`

HasFqbn returns a boolean if a field has been set.

### GetName

`func (o *CreateDevicesV2Payload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDevicesV2Payload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDevicesV2Payload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateDevicesV2Payload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSerial

`func (o *CreateDevicesV2Payload) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *CreateDevicesV2Payload) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *CreateDevicesV2Payload) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *CreateDevicesV2Payload) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSoftDeleted

`func (o *CreateDevicesV2Payload) GetSoftDeleted() bool`

GetSoftDeleted returns the SoftDeleted field if non-nil, zero value otherwise.

### GetSoftDeletedOk

`func (o *CreateDevicesV2Payload) GetSoftDeletedOk() (*bool, bool)`

GetSoftDeletedOk returns a tuple with the SoftDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftDeleted

`func (o *CreateDevicesV2Payload) SetSoftDeleted(v bool)`

SetSoftDeleted sets SoftDeleted field to given value.

### HasSoftDeleted

`func (o *CreateDevicesV2Payload) HasSoftDeleted() bool`

HasSoftDeleted returns a boolean if a field has been set.

### GetType

`func (o *CreateDevicesV2Payload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateDevicesV2Payload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateDevicesV2Payload) SetType(v string)`

SetType sets Type field to given value.


### GetUniqueHardwareId

`func (o *CreateDevicesV2Payload) GetUniqueHardwareId() string`

GetUniqueHardwareId returns the UniqueHardwareId field if non-nil, zero value otherwise.

### GetUniqueHardwareIdOk

`func (o *CreateDevicesV2Payload) GetUniqueHardwareIdOk() (*string, bool)`

GetUniqueHardwareIdOk returns a tuple with the UniqueHardwareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueHardwareId

`func (o *CreateDevicesV2Payload) SetUniqueHardwareId(v string)`

SetUniqueHardwareId sets UniqueHardwareId field to given value.

### HasUniqueHardwareId

`func (o *CreateDevicesV2Payload) HasUniqueHardwareId() bool`

HasUniqueHardwareId returns a boolean if a field has been set.

### GetUserId

`func (o *CreateDevicesV2Payload) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateDevicesV2Payload) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateDevicesV2Payload) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CreateDevicesV2Payload) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetWifiFwVersion

`func (o *CreateDevicesV2Payload) GetWifiFwVersion() string`

GetWifiFwVersion returns the WifiFwVersion field if non-nil, zero value otherwise.

### GetWifiFwVersionOk

`func (o *CreateDevicesV2Payload) GetWifiFwVersionOk() (*string, bool)`

GetWifiFwVersionOk returns a tuple with the WifiFwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiFwVersion

`func (o *CreateDevicesV2Payload) SetWifiFwVersion(v string)`

SetWifiFwVersion sets WifiFwVersion field to given value.

### HasWifiFwVersion

`func (o *CreateDevicesV2Payload) HasWifiFwVersion() bool`

HasWifiFwVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


