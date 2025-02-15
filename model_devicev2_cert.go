/*
Arduino IoT Cloud API

Provides a set of endpoints to manage Arduino IoT Cloud **Devices**, **Things**, **Properties** and **Timeseries**. This API can be called just with any HTTP Client, or using one of these clients:  * [Javascript NPM package](https://www.npmjs.com/package/@arduino/arduino-iot-client)  * [Python PYPI Package](https://pypi.org/project/arduino-iot-client/)  * [Golang Module](https://github.com/arduino/iot-client-go)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
)

// checks if the Devicev2Cert type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Devicev2Cert{}

// Devicev2Cert struct for Devicev2Cert
type Devicev2Cert struct {
	// The Certification Authority you want to use
	Ca *string `json:"ca,omitempty"`
	// The certificate request in pem format
	Csr *string `json:"csr,omitempty"`
	// Whether the certificate is enabled
	Enabled *bool `json:"enabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Devicev2Cert Devicev2Cert

// NewDevicev2Cert instantiates a new Devicev2Cert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicev2Cert() *Devicev2Cert {
	this := Devicev2Cert{}
	return &this
}

// NewDevicev2CertWithDefaults instantiates a new Devicev2Cert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicev2CertWithDefaults() *Devicev2Cert {
	this := Devicev2Cert{}
	return &this
}

// GetCa returns the Ca field value if set, zero value otherwise.
func (o *Devicev2Cert) GetCa() string {
	if o == nil || IsNil(o.Ca) {
		var ret string
		return ret
	}
	return *o.Ca
}

// GetCaOk returns a tuple with the Ca field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Devicev2Cert) GetCaOk() (*string, bool) {
	if o == nil || IsNil(o.Ca) {
		return nil, false
	}
	return o.Ca, true
}

// HasCa returns a boolean if a field has been set.
func (o *Devicev2Cert) HasCa() bool {
	if o != nil && !IsNil(o.Ca) {
		return true
	}

	return false
}

// SetCa gets a reference to the given string and assigns it to the Ca field.
func (o *Devicev2Cert) SetCa(v string) {
	o.Ca = &v
}

// GetCsr returns the Csr field value if set, zero value otherwise.
func (o *Devicev2Cert) GetCsr() string {
	if o == nil || IsNil(o.Csr) {
		var ret string
		return ret
	}
	return *o.Csr
}

// GetCsrOk returns a tuple with the Csr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Devicev2Cert) GetCsrOk() (*string, bool) {
	if o == nil || IsNil(o.Csr) {
		return nil, false
	}
	return o.Csr, true
}

// HasCsr returns a boolean if a field has been set.
func (o *Devicev2Cert) HasCsr() bool {
	if o != nil && !IsNil(o.Csr) {
		return true
	}

	return false
}

// SetCsr gets a reference to the given string and assigns it to the Csr field.
func (o *Devicev2Cert) SetCsr(v string) {
	o.Csr = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Devicev2Cert) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Devicev2Cert) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Devicev2Cert) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Devicev2Cert) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o Devicev2Cert) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Devicev2Cert) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ca) {
		toSerialize["ca"] = o.Ca
	}
	if !IsNil(o.Csr) {
		toSerialize["csr"] = o.Csr
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Devicev2Cert) UnmarshalJSON(data []byte) (err error) {
	varDevicev2Cert := _Devicev2Cert{}

	err = json.Unmarshal(data, &varDevicev2Cert)

	if err != nil {
		return err
	}

	*o = Devicev2Cert(varDevicev2Cert)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ca")
		delete(additionalProperties, "csr")
		delete(additionalProperties, "enabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDevicev2Cert struct {
	value *Devicev2Cert
	isSet bool
}

func (v NullableDevicev2Cert) Get() *Devicev2Cert {
	return v.value
}

func (v *NullableDevicev2Cert) Set(val *Devicev2Cert) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicev2Cert) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicev2Cert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicev2Cert(val *Devicev2Cert) *NullableDevicev2Cert {
	return &NullableDevicev2Cert{value: val, isSet: true}
}

func (v NullableDevicev2Cert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicev2Cert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


