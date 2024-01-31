# ArduinoDevicev2StatusEvents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]ArduinoDevicev2StatusEvent**](ArduinoDevicev2StatusEvent.md) | ArduinoDevicev2StatusEventCollection is the media type for an array of ArduinoDevicev2StatusEvent (default view) | 
**Id** | **string** | The id of the device | 

## Methods

### NewArduinoDevicev2StatusEvents

`func NewArduinoDevicev2StatusEvents(events []ArduinoDevicev2StatusEvent, id string, ) *ArduinoDevicev2StatusEvents`

NewArduinoDevicev2StatusEvents instantiates a new ArduinoDevicev2StatusEvents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoDevicev2StatusEventsWithDefaults

`func NewArduinoDevicev2StatusEventsWithDefaults() *ArduinoDevicev2StatusEvents`

NewArduinoDevicev2StatusEventsWithDefaults instantiates a new ArduinoDevicev2StatusEvents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *ArduinoDevicev2StatusEvents) GetEvents() []ArduinoDevicev2StatusEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ArduinoDevicev2StatusEvents) GetEventsOk() (*[]ArduinoDevicev2StatusEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ArduinoDevicev2StatusEvents) SetEvents(v []ArduinoDevicev2StatusEvent)`

SetEvents sets Events field to given value.


### GetId

`func (o *ArduinoDevicev2StatusEvents) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArduinoDevicev2StatusEvents) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArduinoDevicev2StatusEvents) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


