# ArduinoLinkedDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | [**ArduinoDevicev2**](ArduinoDevicev2.md) |  | 
**Status** | **string** | The status of the linked device | 

## Methods

### NewArduinoLinkedDevice

`func NewArduinoLinkedDevice(device ArduinoDevicev2, status string, ) *ArduinoLinkedDevice`

NewArduinoLinkedDevice instantiates a new ArduinoLinkedDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoLinkedDeviceWithDefaults

`func NewArduinoLinkedDeviceWithDefaults() *ArduinoLinkedDevice`

NewArduinoLinkedDeviceWithDefaults instantiates a new ArduinoLinkedDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *ArduinoLinkedDevice) GetDevice() ArduinoDevicev2`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *ArduinoLinkedDevice) GetDeviceOk() (*ArduinoDevicev2, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *ArduinoLinkedDevice) SetDevice(v ArduinoDevicev2)`

SetDevice sets Device field to given value.


### GetStatus

`func (o *ArduinoLinkedDevice) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArduinoLinkedDevice) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArduinoLinkedDevice) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


