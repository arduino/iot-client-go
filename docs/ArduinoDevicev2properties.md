# ArduinoDevicev2properties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataRetentionDays** | **float64** | How many days the data will be kept | 
**DeviceId** | **string** | The device of the property | 
**Properties** | [**[]ArduinoProperty**](ArduinoProperty.md) | ArduinoPropertyCollection is the media type for an array of ArduinoProperty (default view) | 
**UserId** | **string** | The user id of the owner | 

## Methods

### NewArduinoDevicev2properties

`func NewArduinoDevicev2properties(dataRetentionDays float64, deviceId string, properties []ArduinoProperty, userId string, ) *ArduinoDevicev2properties`

NewArduinoDevicev2properties instantiates a new ArduinoDevicev2properties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoDevicev2propertiesWithDefaults

`func NewArduinoDevicev2propertiesWithDefaults() *ArduinoDevicev2properties`

NewArduinoDevicev2propertiesWithDefaults instantiates a new ArduinoDevicev2properties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataRetentionDays

`func (o *ArduinoDevicev2properties) GetDataRetentionDays() float64`

GetDataRetentionDays returns the DataRetentionDays field if non-nil, zero value otherwise.

### GetDataRetentionDaysOk

`func (o *ArduinoDevicev2properties) GetDataRetentionDaysOk() (*float64, bool)`

GetDataRetentionDaysOk returns a tuple with the DataRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRetentionDays

`func (o *ArduinoDevicev2properties) SetDataRetentionDays(v float64)`

SetDataRetentionDays sets DataRetentionDays field to given value.


### GetDeviceId

`func (o *ArduinoDevicev2properties) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *ArduinoDevicev2properties) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *ArduinoDevicev2properties) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetProperties

`func (o *ArduinoDevicev2properties) GetProperties() []ArduinoProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ArduinoDevicev2properties) GetPropertiesOk() (*[]ArduinoProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ArduinoDevicev2properties) SetProperties(v []ArduinoProperty)`

SetProperties sets Properties field to given value.


### GetUserId

`func (o *ArduinoDevicev2properties) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ArduinoDevicev2properties) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ArduinoDevicev2properties) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


