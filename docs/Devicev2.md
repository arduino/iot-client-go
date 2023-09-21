# Devicev2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionType** | Pointer to **string** | The type of the connections selected by the user when multiple connections are available | [optional] 
**Fqbn** | Pointer to **string** | The fully qualified board name | [optional] 
**Name** | Pointer to **string** | The friendly name of the device | [optional] 
**Serial** | Pointer to **string** | The serial uuid of the device | [optional] 
**Type** | Pointer to **string** | The type of the device | [optional] 
**UserId** | Pointer to **string** | The user_id associated to the device. If absent it will be inferred from the authentication header | [optional] 
**WifiFwVersion** | Pointer to **string** | The version of the NINA/WIFI101 firmware running on the device | [optional] 

## Methods

### NewDevicev2

`func NewDevicev2() *Devicev2`

NewDevicev2 instantiates a new Devicev2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicev2WithDefaults

`func NewDevicev2WithDefaults() *Devicev2`

NewDevicev2WithDefaults instantiates a new Devicev2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionType

`func (o *Devicev2) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *Devicev2) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *Devicev2) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *Devicev2) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetFqbn

`func (o *Devicev2) GetFqbn() string`

GetFqbn returns the Fqbn field if non-nil, zero value otherwise.

### GetFqbnOk

`func (o *Devicev2) GetFqbnOk() (*string, bool)`

GetFqbnOk returns a tuple with the Fqbn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqbn

`func (o *Devicev2) SetFqbn(v string)`

SetFqbn sets Fqbn field to given value.

### HasFqbn

`func (o *Devicev2) HasFqbn() bool`

HasFqbn returns a boolean if a field has been set.

### GetName

`func (o *Devicev2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Devicev2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Devicev2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Devicev2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSerial

`func (o *Devicev2) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *Devicev2) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *Devicev2) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *Devicev2) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetType

`func (o *Devicev2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Devicev2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Devicev2) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Devicev2) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserId

`func (o *Devicev2) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Devicev2) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Devicev2) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Devicev2) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetWifiFwVersion

`func (o *Devicev2) GetWifiFwVersion() string`

GetWifiFwVersion returns the WifiFwVersion field if non-nil, zero value otherwise.

### GetWifiFwVersionOk

`func (o *Devicev2) GetWifiFwVersionOk() (*string, bool)`

GetWifiFwVersionOk returns a tuple with the WifiFwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiFwVersion

`func (o *Devicev2) SetWifiFwVersion(v string)`

SetWifiFwVersion sets WifiFwVersion field to given value.

### HasWifiFwVersion

`func (o *Devicev2) HasWifiFwVersion() bool`

HasWifiFwVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


