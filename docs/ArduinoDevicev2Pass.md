# ArduinoDevicev2Pass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Set** | **bool** | Whether the password is set or not | 
**SuggestedPassword** | Pointer to **string** | A random suggested password | [optional] 

## Methods

### NewArduinoDevicev2Pass

`func NewArduinoDevicev2Pass(set bool, ) *ArduinoDevicev2Pass`

NewArduinoDevicev2Pass instantiates a new ArduinoDevicev2Pass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoDevicev2PassWithDefaults

`func NewArduinoDevicev2PassWithDefaults() *ArduinoDevicev2Pass`

NewArduinoDevicev2PassWithDefaults instantiates a new ArduinoDevicev2Pass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSet

`func (o *ArduinoDevicev2Pass) GetSet() bool`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *ArduinoDevicev2Pass) GetSetOk() (*bool, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSet

`func (o *ArduinoDevicev2Pass) SetSet(v bool)`

SetSet sets Set field to given value.


### GetSuggestedPassword

`func (o *ArduinoDevicev2Pass) GetSuggestedPassword() string`

GetSuggestedPassword returns the SuggestedPassword field if non-nil, zero value otherwise.

### GetSuggestedPasswordOk

`func (o *ArduinoDevicev2Pass) GetSuggestedPasswordOk() (*string, bool)`

GetSuggestedPasswordOk returns a tuple with the SuggestedPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedPassword

`func (o *ArduinoDevicev2Pass) SetSuggestedPassword(v string)`

SetSuggestedPassword sets SuggestedPassword field to given value.

### HasSuggestedPassword

`func (o *ArduinoDevicev2Pass) HasSuggestedPassword() bool`

HasSuggestedPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


