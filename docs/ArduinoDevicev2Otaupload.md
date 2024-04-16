# ArduinoDevicev2Otaupload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileSha** | Pointer to **string** | SHA256 of the uploaded file | [optional] 
**OtaId** | Pointer to **string** | OTA request id (only available from OTA version 2 and above) | [optional] 
**OtaVersion** | **int64** | OTA version | 
**Status** | Pointer to **string** | OTA request status (only available from OTA version 2 and above) | [optional] 

## Methods

### NewArduinoDevicev2Otaupload

`func NewArduinoDevicev2Otaupload(otaVersion int64, ) *ArduinoDevicev2Otaupload`

NewArduinoDevicev2Otaupload instantiates a new ArduinoDevicev2Otaupload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoDevicev2OtauploadWithDefaults

`func NewArduinoDevicev2OtauploadWithDefaults() *ArduinoDevicev2Otaupload`

NewArduinoDevicev2OtauploadWithDefaults instantiates a new ArduinoDevicev2Otaupload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileSha

`func (o *ArduinoDevicev2Otaupload) GetFileSha() string`

GetFileSha returns the FileSha field if non-nil, zero value otherwise.

### GetFileShaOk

`func (o *ArduinoDevicev2Otaupload) GetFileShaOk() (*string, bool)`

GetFileShaOk returns a tuple with the FileSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSha

`func (o *ArduinoDevicev2Otaupload) SetFileSha(v string)`

SetFileSha sets FileSha field to given value.

### HasFileSha

`func (o *ArduinoDevicev2Otaupload) HasFileSha() bool`

HasFileSha returns a boolean if a field has been set.

### GetOtaId

`func (o *ArduinoDevicev2Otaupload) GetOtaId() string`

GetOtaId returns the OtaId field if non-nil, zero value otherwise.

### GetOtaIdOk

`func (o *ArduinoDevicev2Otaupload) GetOtaIdOk() (*string, bool)`

GetOtaIdOk returns a tuple with the OtaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtaId

`func (o *ArduinoDevicev2Otaupload) SetOtaId(v string)`

SetOtaId sets OtaId field to given value.

### HasOtaId

`func (o *ArduinoDevicev2Otaupload) HasOtaId() bool`

HasOtaId returns a boolean if a field has been set.

### GetOtaVersion

`func (o *ArduinoDevicev2Otaupload) GetOtaVersion() int64`

GetOtaVersion returns the OtaVersion field if non-nil, zero value otherwise.

### GetOtaVersionOk

`func (o *ArduinoDevicev2Otaupload) GetOtaVersionOk() (*int64, bool)`

GetOtaVersionOk returns a tuple with the OtaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtaVersion

`func (o *ArduinoDevicev2Otaupload) SetOtaVersion(v int64)`

SetOtaVersion sets OtaVersion field to given value.


### GetStatus

`func (o *ArduinoDevicev2Otaupload) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArduinoDevicev2Otaupload) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArduinoDevicev2Otaupload) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ArduinoDevicev2Otaupload) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


