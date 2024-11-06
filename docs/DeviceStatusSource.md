# DeviceStatusSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Criteria** | **string** | The matching criteria of the trigger, this allows to interpret device_ids as an inclusion or exclusion list | 
**DeviceIds** | Pointer to **[]string** | A list of device IDs to be included in or excluded from the trigger (see criteria). Mutually exclusive with property_id. | [optional] 
**GracePeriodOffline** | Pointer to **int32** | Amount of seconds the trigger will wait before the device will be considered disconnected (requires &#39;device_id&#39;) | [optional] 
**GracePeriodOnline** | Pointer to **int32** | Amount of seconds the trigger will wait before the device will be considered connected (requires &#39;device_id&#39;) | [optional] 

## Methods

### NewDeviceStatusSource

`func NewDeviceStatusSource(criteria string, ) *DeviceStatusSource`

NewDeviceStatusSource instantiates a new DeviceStatusSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceStatusSourceWithDefaults

`func NewDeviceStatusSourceWithDefaults() *DeviceStatusSource`

NewDeviceStatusSourceWithDefaults instantiates a new DeviceStatusSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteria

`func (o *DeviceStatusSource) GetCriteria() string`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *DeviceStatusSource) GetCriteriaOk() (*string, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *DeviceStatusSource) SetCriteria(v string)`

SetCriteria sets Criteria field to given value.


### GetDeviceIds

`func (o *DeviceStatusSource) GetDeviceIds() []string`

GetDeviceIds returns the DeviceIds field if non-nil, zero value otherwise.

### GetDeviceIdsOk

`func (o *DeviceStatusSource) GetDeviceIdsOk() (*[]string, bool)`

GetDeviceIdsOk returns a tuple with the DeviceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIds

`func (o *DeviceStatusSource) SetDeviceIds(v []string)`

SetDeviceIds sets DeviceIds field to given value.

### HasDeviceIds

`func (o *DeviceStatusSource) HasDeviceIds() bool`

HasDeviceIds returns a boolean if a field has been set.

### GetGracePeriodOffline

`func (o *DeviceStatusSource) GetGracePeriodOffline() int32`

GetGracePeriodOffline returns the GracePeriodOffline field if non-nil, zero value otherwise.

### GetGracePeriodOfflineOk

`func (o *DeviceStatusSource) GetGracePeriodOfflineOk() (*int32, bool)`

GetGracePeriodOfflineOk returns a tuple with the GracePeriodOffline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodOffline

`func (o *DeviceStatusSource) SetGracePeriodOffline(v int32)`

SetGracePeriodOffline sets GracePeriodOffline field to given value.

### HasGracePeriodOffline

`func (o *DeviceStatusSource) HasGracePeriodOffline() bool`

HasGracePeriodOffline returns a boolean if a field has been set.

### GetGracePeriodOnline

`func (o *DeviceStatusSource) GetGracePeriodOnline() int32`

GetGracePeriodOnline returns the GracePeriodOnline field if non-nil, zero value otherwise.

### GetGracePeriodOnlineOk

`func (o *DeviceStatusSource) GetGracePeriodOnlineOk() (*int32, bool)`

GetGracePeriodOnlineOk returns a tuple with the GracePeriodOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodOnline

`func (o *DeviceStatusSource) SetGracePeriodOnline(v int32)`

SetGracePeriodOnline sets GracePeriodOnline field to given value.

### HasGracePeriodOnline

`func (o *DeviceStatusSource) HasGracePeriodOnline() bool`

HasGracePeriodOnline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


