# ArduinoDevicev2Cert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ca** | Pointer to **string** | The Certification Authority used to sign the certificate | [optional] 
**Compressed** | [**ArduinoCompressedv2**](ArduinoCompressedv2.md) |  | 
**Der** | **string** | The certificate in DER format | 
**DeviceId** | **string** | The unique identifier of the device | 
**Enabled** | **bool** | Whether the certificate is enabled | [default to true]
**Href** | **string** | The api reference of this cert | 
**Id** | **string** | The unique identifier of the key | 
**Pem** | **string** | The certificate in pem format | 

## Methods

### NewArduinoDevicev2Cert

`func NewArduinoDevicev2Cert(compressed ArduinoCompressedv2, der string, deviceId string, enabled bool, href string, id string, pem string, ) *ArduinoDevicev2Cert`

NewArduinoDevicev2Cert instantiates a new ArduinoDevicev2Cert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoDevicev2CertWithDefaults

`func NewArduinoDevicev2CertWithDefaults() *ArduinoDevicev2Cert`

NewArduinoDevicev2CertWithDefaults instantiates a new ArduinoDevicev2Cert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCa

`func (o *ArduinoDevicev2Cert) GetCa() string`

GetCa returns the Ca field if non-nil, zero value otherwise.

### GetCaOk

`func (o *ArduinoDevicev2Cert) GetCaOk() (*string, bool)`

GetCaOk returns a tuple with the Ca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCa

`func (o *ArduinoDevicev2Cert) SetCa(v string)`

SetCa sets Ca field to given value.

### HasCa

`func (o *ArduinoDevicev2Cert) HasCa() bool`

HasCa returns a boolean if a field has been set.

### GetCompressed

`func (o *ArduinoDevicev2Cert) GetCompressed() ArduinoCompressedv2`

GetCompressed returns the Compressed field if non-nil, zero value otherwise.

### GetCompressedOk

`func (o *ArduinoDevicev2Cert) GetCompressedOk() (*ArduinoCompressedv2, bool)`

GetCompressedOk returns a tuple with the Compressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressed

`func (o *ArduinoDevicev2Cert) SetCompressed(v ArduinoCompressedv2)`

SetCompressed sets Compressed field to given value.


### GetDer

`func (o *ArduinoDevicev2Cert) GetDer() string`

GetDer returns the Der field if non-nil, zero value otherwise.

### GetDerOk

`func (o *ArduinoDevicev2Cert) GetDerOk() (*string, bool)`

GetDerOk returns a tuple with the Der field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDer

`func (o *ArduinoDevicev2Cert) SetDer(v string)`

SetDer sets Der field to given value.


### GetDeviceId

`func (o *ArduinoDevicev2Cert) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *ArduinoDevicev2Cert) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *ArduinoDevicev2Cert) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetEnabled

`func (o *ArduinoDevicev2Cert) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ArduinoDevicev2Cert) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ArduinoDevicev2Cert) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetHref

`func (o *ArduinoDevicev2Cert) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ArduinoDevicev2Cert) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ArduinoDevicev2Cert) SetHref(v string)`

SetHref sets Href field to given value.


### GetId

`func (o *ArduinoDevicev2Cert) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArduinoDevicev2Cert) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArduinoDevicev2Cert) SetId(v string)`

SetId sets Id field to given value.


### GetPem

`func (o *ArduinoDevicev2Cert) GetPem() string`

GetPem returns the Pem field if non-nil, zero value otherwise.

### GetPemOk

`func (o *ArduinoDevicev2Cert) GetPemOk() (*string, bool)`

GetPemOk returns a tuple with the Pem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPem

`func (o *ArduinoDevicev2Cert) SetPem(v string)`

SetPem sets Pem field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


