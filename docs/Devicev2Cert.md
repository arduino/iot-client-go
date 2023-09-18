# Devicev2Cert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ca** | Pointer to **string** | The Certification Authority you want to use | [optional] 
**Csr** | Pointer to **string** | The certificate request in pem format | [optional] 
**Enabled** | Pointer to **bool** | Whether the certificate is enabled | [optional] 

## Methods

### NewDevicev2Cert

`func NewDevicev2Cert() *Devicev2Cert`

NewDevicev2Cert instantiates a new Devicev2Cert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicev2CertWithDefaults

`func NewDevicev2CertWithDefaults() *Devicev2Cert`

NewDevicev2CertWithDefaults instantiates a new Devicev2Cert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCa

`func (o *Devicev2Cert) GetCa() string`

GetCa returns the Ca field if non-nil, zero value otherwise.

### GetCaOk

`func (o *Devicev2Cert) GetCaOk() (*string, bool)`

GetCaOk returns a tuple with the Ca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCa

`func (o *Devicev2Cert) SetCa(v string)`

SetCa sets Ca field to given value.

### HasCa

`func (o *Devicev2Cert) HasCa() bool`

HasCa returns a boolean if a field has been set.

### GetCsr

`func (o *Devicev2Cert) GetCsr() string`

GetCsr returns the Csr field if non-nil, zero value otherwise.

### GetCsrOk

`func (o *Devicev2Cert) GetCsrOk() (*string, bool)`

GetCsrOk returns a tuple with the Csr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsr

`func (o *Devicev2Cert) SetCsr(v string)`

SetCsr sets Csr field to given value.

### HasCsr

`func (o *Devicev2Cert) HasCsr() bool`

HasCsr returns a boolean if a field has been set.

### GetEnabled

`func (o *Devicev2Cert) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Devicev2Cert) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Devicev2Cert) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Devicev2Cert) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


