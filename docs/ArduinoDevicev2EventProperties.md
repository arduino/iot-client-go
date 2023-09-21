# ArduinoDevicev2EventProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]ArduinoDevicev2SimpleProperties**](ArduinoDevicev2SimpleProperties.md) | ArduinoDevicev2SimplePropertiesCollection is the media type for an array of ArduinoDevicev2SimpleProperties (default view) | 
**Id** | **string** | The device of the property | 

## Methods

### NewArduinoDevicev2EventProperties

`func NewArduinoDevicev2EventProperties(events []ArduinoDevicev2SimpleProperties, id string, ) *ArduinoDevicev2EventProperties`

NewArduinoDevicev2EventProperties instantiates a new ArduinoDevicev2EventProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoDevicev2EventPropertiesWithDefaults

`func NewArduinoDevicev2EventPropertiesWithDefaults() *ArduinoDevicev2EventProperties`

NewArduinoDevicev2EventPropertiesWithDefaults instantiates a new ArduinoDevicev2EventProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *ArduinoDevicev2EventProperties) GetEvents() []ArduinoDevicev2SimpleProperties`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ArduinoDevicev2EventProperties) GetEventsOk() (*[]ArduinoDevicev2SimpleProperties, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ArduinoDevicev2EventProperties) SetEvents(v []ArduinoDevicev2SimpleProperties)`

SetEvents sets Events field to given value.


### GetId

`func (o *ArduinoDevicev2EventProperties) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArduinoDevicev2EventProperties) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArduinoDevicev2EventProperties) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


