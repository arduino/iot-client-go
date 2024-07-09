# ArduinoThingresult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** | UUID of the attached device | [optional] 
**Id** | **string** | UUID of the created Thing | 
**SketchId** | **string** | UUID of the created Sketch | 

## Methods

### NewArduinoThingresult

`func NewArduinoThingresult(id string, sketchId string, ) *ArduinoThingresult`

NewArduinoThingresult instantiates a new ArduinoThingresult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoThingresultWithDefaults

`func NewArduinoThingresultWithDefaults() *ArduinoThingresult`

NewArduinoThingresultWithDefaults instantiates a new ArduinoThingresult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *ArduinoThingresult) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *ArduinoThingresult) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *ArduinoThingresult) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *ArduinoThingresult) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetId

`func (o *ArduinoThingresult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArduinoThingresult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArduinoThingresult) SetId(v string)`

SetId sets Id field to given value.


### GetSketchId

`func (o *ArduinoThingresult) GetSketchId() string`

GetSketchId returns the SketchId field if non-nil, zero value otherwise.

### GetSketchIdOk

`func (o *ArduinoThingresult) GetSketchIdOk() (*string, bool)`

GetSketchIdOk returns a tuple with the SketchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSketchId

`func (o *ArduinoThingresult) SetSketchId(v string)`

SetSketchId sets SketchId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


