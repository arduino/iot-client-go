# CreateDevicesV2CertsPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ca** | Pointer to **string** | The Certification Authority you want to use | [optional] 
**Csr** | **string** | The certificate request in pem format | 
**Enabled** | **bool** | Whether the certificate is enabled | 

## Methods

### NewCreateDevicesV2CertsPayload

`func NewCreateDevicesV2CertsPayload(csr string, enabled bool, ) *CreateDevicesV2CertsPayload`

NewCreateDevicesV2CertsPayload instantiates a new CreateDevicesV2CertsPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDevicesV2CertsPayloadWithDefaults

`func NewCreateDevicesV2CertsPayloadWithDefaults() *CreateDevicesV2CertsPayload`

NewCreateDevicesV2CertsPayloadWithDefaults instantiates a new CreateDevicesV2CertsPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCa

`func (o *CreateDevicesV2CertsPayload) GetCa() string`

GetCa returns the Ca field if non-nil, zero value otherwise.

### GetCaOk

`func (o *CreateDevicesV2CertsPayload) GetCaOk() (*string, bool)`

GetCaOk returns a tuple with the Ca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCa

`func (o *CreateDevicesV2CertsPayload) SetCa(v string)`

SetCa sets Ca field to given value.

### HasCa

`func (o *CreateDevicesV2CertsPayload) HasCa() bool`

HasCa returns a boolean if a field has been set.

### GetCsr

`func (o *CreateDevicesV2CertsPayload) GetCsr() string`

GetCsr returns the Csr field if non-nil, zero value otherwise.

### GetCsrOk

`func (o *CreateDevicesV2CertsPayload) GetCsrOk() (*string, bool)`

GetCsrOk returns a tuple with the Csr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsr

`func (o *CreateDevicesV2CertsPayload) SetCsr(v string)`

SetCsr sets Csr field to given value.


### GetEnabled

`func (o *CreateDevicesV2CertsPayload) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateDevicesV2CertsPayload) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateDevicesV2CertsPayload) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


