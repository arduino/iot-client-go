# ArduinoSeriesBatchSampled

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RespVersion** | **int64** | Response version | 
**Responses** | [**[]ArduinoSeriesSampledResponse**](ArduinoSeriesSampledResponse.md) | Responses of the request | 

## Methods

### NewArduinoSeriesBatchSampled

`func NewArduinoSeriesBatchSampled(respVersion int64, responses []ArduinoSeriesSampledResponse, ) *ArduinoSeriesBatchSampled`

NewArduinoSeriesBatchSampled instantiates a new ArduinoSeriesBatchSampled object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoSeriesBatchSampledWithDefaults

`func NewArduinoSeriesBatchSampledWithDefaults() *ArduinoSeriesBatchSampled`

NewArduinoSeriesBatchSampledWithDefaults instantiates a new ArduinoSeriesBatchSampled object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRespVersion

`func (o *ArduinoSeriesBatchSampled) GetRespVersion() int64`

GetRespVersion returns the RespVersion field if non-nil, zero value otherwise.

### GetRespVersionOk

`func (o *ArduinoSeriesBatchSampled) GetRespVersionOk() (*int64, bool)`

GetRespVersionOk returns a tuple with the RespVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespVersion

`func (o *ArduinoSeriesBatchSampled) SetRespVersion(v int64)`

SetRespVersion sets RespVersion field to given value.


### GetResponses

`func (o *ArduinoSeriesBatchSampled) GetResponses() []ArduinoSeriesSampledResponse`

GetResponses returns the Responses field if non-nil, zero value otherwise.

### GetResponsesOk

`func (o *ArduinoSeriesBatchSampled) GetResponsesOk() (*[]ArduinoSeriesSampledResponse, bool)`

GetResponsesOk returns a tuple with the Responses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponses

`func (o *ArduinoSeriesBatchSampled) SetResponses(v []ArduinoSeriesSampledResponse)`

SetResponses sets Responses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


