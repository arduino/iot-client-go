# ArduinoCompressedv2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorityKeyIdentifier** | Pointer to **string** | The Authority Key Identifier of the certificate | [optional] 
**NotAfter** | **time.Time** | The ending date of the certificate | 
**NotBefore** | **time.Time** | The starting date of the certificate | 
**Serial** | **string** | The serial number of the certificate | 
**Signature** | **string** | The signature of the certificate | 
**SignatureAsn1X** | **string** | The ASN1 X component of certificate signature | 
**SignatureAsn1Y** | **string** | The ASN1 Y component of certificate signature | 

## Methods

### NewArduinoCompressedv2

`func NewArduinoCompressedv2(notAfter time.Time, notBefore time.Time, serial string, signature string, signatureAsn1X string, signatureAsn1Y string, ) *ArduinoCompressedv2`

NewArduinoCompressedv2 instantiates a new ArduinoCompressedv2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoCompressedv2WithDefaults

`func NewArduinoCompressedv2WithDefaults() *ArduinoCompressedv2`

NewArduinoCompressedv2WithDefaults instantiates a new ArduinoCompressedv2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorityKeyIdentifier

`func (o *ArduinoCompressedv2) GetAuthorityKeyIdentifier() string`

GetAuthorityKeyIdentifier returns the AuthorityKeyIdentifier field if non-nil, zero value otherwise.

### GetAuthorityKeyIdentifierOk

`func (o *ArduinoCompressedv2) GetAuthorityKeyIdentifierOk() (*string, bool)`

GetAuthorityKeyIdentifierOk returns a tuple with the AuthorityKeyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityKeyIdentifier

`func (o *ArduinoCompressedv2) SetAuthorityKeyIdentifier(v string)`

SetAuthorityKeyIdentifier sets AuthorityKeyIdentifier field to given value.

### HasAuthorityKeyIdentifier

`func (o *ArduinoCompressedv2) HasAuthorityKeyIdentifier() bool`

HasAuthorityKeyIdentifier returns a boolean if a field has been set.

### GetNotAfter

`func (o *ArduinoCompressedv2) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *ArduinoCompressedv2) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *ArduinoCompressedv2) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.


### GetNotBefore

`func (o *ArduinoCompressedv2) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *ArduinoCompressedv2) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *ArduinoCompressedv2) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.


### GetSerial

`func (o *ArduinoCompressedv2) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *ArduinoCompressedv2) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *ArduinoCompressedv2) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetSignature

`func (o *ArduinoCompressedv2) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *ArduinoCompressedv2) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *ArduinoCompressedv2) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetSignatureAsn1X

`func (o *ArduinoCompressedv2) GetSignatureAsn1X() string`

GetSignatureAsn1X returns the SignatureAsn1X field if non-nil, zero value otherwise.

### GetSignatureAsn1XOk

`func (o *ArduinoCompressedv2) GetSignatureAsn1XOk() (*string, bool)`

GetSignatureAsn1XOk returns a tuple with the SignatureAsn1X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAsn1X

`func (o *ArduinoCompressedv2) SetSignatureAsn1X(v string)`

SetSignatureAsn1X sets SignatureAsn1X field to given value.


### GetSignatureAsn1Y

`func (o *ArduinoCompressedv2) GetSignatureAsn1Y() string`

GetSignatureAsn1Y returns the SignatureAsn1Y field if non-nil, zero value otherwise.

### GetSignatureAsn1YOk

`func (o *ArduinoCompressedv2) GetSignatureAsn1YOk() (*string, bool)`

GetSignatureAsn1YOk returns a tuple with the SignatureAsn1Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAsn1Y

`func (o *ArduinoCompressedv2) SetSignatureAsn1Y(v string)`

SetSignatureAsn1Y sets SignatureAsn1Y field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


