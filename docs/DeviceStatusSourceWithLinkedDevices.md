# DeviceStatusSourceWithLinkedDevices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Criteria** | **string** | The criteria of the trigger, could be INCLUDE or EXCLUDE | 
**GracePeriodOffline** | Pointer to **int32** | The amount of seconds the trigger will wait before considering a matching device as offline | [optional] 
**GracePeriodOnline** | Pointer to **int32** | The amount of seconds the trigger will wait before considering a matching device as online | [optional] 
**LinkedDevices** | Pointer to [**[]ArduinoLinkedDevice**](ArduinoLinkedDevice.md) | A list of devices the trigger is associated to | [optional] 

## Methods

### NewDeviceStatusSourceWithLinkedDevices

`func NewDeviceStatusSourceWithLinkedDevices(criteria string, ) *DeviceStatusSourceWithLinkedDevices`

NewDeviceStatusSourceWithLinkedDevices instantiates a new DeviceStatusSourceWithLinkedDevices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceStatusSourceWithLinkedDevicesWithDefaults

`func NewDeviceStatusSourceWithLinkedDevicesWithDefaults() *DeviceStatusSourceWithLinkedDevices`

NewDeviceStatusSourceWithLinkedDevicesWithDefaults instantiates a new DeviceStatusSourceWithLinkedDevices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteria

`func (o *DeviceStatusSourceWithLinkedDevices) GetCriteria() string`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *DeviceStatusSourceWithLinkedDevices) GetCriteriaOk() (*string, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *DeviceStatusSourceWithLinkedDevices) SetCriteria(v string)`

SetCriteria sets Criteria field to given value.


### GetGracePeriodOffline

`func (o *DeviceStatusSourceWithLinkedDevices) GetGracePeriodOffline() int32`

GetGracePeriodOffline returns the GracePeriodOffline field if non-nil, zero value otherwise.

### GetGracePeriodOfflineOk

`func (o *DeviceStatusSourceWithLinkedDevices) GetGracePeriodOfflineOk() (*int32, bool)`

GetGracePeriodOfflineOk returns a tuple with the GracePeriodOffline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodOffline

`func (o *DeviceStatusSourceWithLinkedDevices) SetGracePeriodOffline(v int32)`

SetGracePeriodOffline sets GracePeriodOffline field to given value.

### HasGracePeriodOffline

`func (o *DeviceStatusSourceWithLinkedDevices) HasGracePeriodOffline() bool`

HasGracePeriodOffline returns a boolean if a field has been set.

### GetGracePeriodOnline

`func (o *DeviceStatusSourceWithLinkedDevices) GetGracePeriodOnline() int32`

GetGracePeriodOnline returns the GracePeriodOnline field if non-nil, zero value otherwise.

### GetGracePeriodOnlineOk

`func (o *DeviceStatusSourceWithLinkedDevices) GetGracePeriodOnlineOk() (*int32, bool)`

GetGracePeriodOnlineOk returns a tuple with the GracePeriodOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodOnline

`func (o *DeviceStatusSourceWithLinkedDevices) SetGracePeriodOnline(v int32)`

SetGracePeriodOnline sets GracePeriodOnline field to given value.

### HasGracePeriodOnline

`func (o *DeviceStatusSourceWithLinkedDevices) HasGracePeriodOnline() bool`

HasGracePeriodOnline returns a boolean if a field has been set.

### GetLinkedDevices

`func (o *DeviceStatusSourceWithLinkedDevices) GetLinkedDevices() []ArduinoLinkedDevice`

GetLinkedDevices returns the LinkedDevices field if non-nil, zero value otherwise.

### GetLinkedDevicesOk

`func (o *DeviceStatusSourceWithLinkedDevices) GetLinkedDevicesOk() (*[]ArduinoLinkedDevice, bool)`

GetLinkedDevicesOk returns a tuple with the LinkedDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedDevices

`func (o *DeviceStatusSourceWithLinkedDevices) SetLinkedDevices(v []ArduinoLinkedDevice)`

SetLinkedDevices sets LinkedDevices field to given value.

### HasLinkedDevices

`func (o *DeviceStatusSourceWithLinkedDevices) HasLinkedDevices() bool`

HasLinkedDevices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


