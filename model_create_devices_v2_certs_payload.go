/*
Arduino IoT Cloud API

Provides a set of endpoints to manage Arduino IoT Cloud **Devices**, **Things**, **Properties** and **Timeseries**. This API can be called just with any HTTP Client, or using one of these clients:  * [Javascript NPM package](https://www.npmjs.com/package/@arduino/arduino-iot-client)  * [Python PYPI Package](https://pypi.org/project/arduino-iot-client/)  * [Golang Module](https://github.com/arduino/iot-client-go)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
	"fmt"
)

// checks if the CreateDevicesV2CertsPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDevicesV2CertsPayload{}

// CreateDevicesV2CertsPayload struct for CreateDevicesV2CertsPayload
type CreateDevicesV2CertsPayload struct {
	// The Certification Authority you want to use
	Ca *string `json:"ca,omitempty"`
	// The certificate request in pem format
	Csr string `json:"csr"`
	// Whether the certificate is enabled
	Enabled bool `json:"enabled"`
	AdditionalProperties map[string]interface{}
}

type _CreateDevicesV2CertsPayload CreateDevicesV2CertsPayload

// NewCreateDevicesV2CertsPayload instantiates a new CreateDevicesV2CertsPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDevicesV2CertsPayload(csr string, enabled bool) *CreateDevicesV2CertsPayload {
	this := CreateDevicesV2CertsPayload{}
	this.Csr = csr
	this.Enabled = enabled
	return &this
}

// NewCreateDevicesV2CertsPayloadWithDefaults instantiates a new CreateDevicesV2CertsPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDevicesV2CertsPayloadWithDefaults() *CreateDevicesV2CertsPayload {
	this := CreateDevicesV2CertsPayload{}
	return &this
}

// GetCa returns the Ca field value if set, zero value otherwise.
func (o *CreateDevicesV2CertsPayload) GetCa() string {
	if o == nil || IsNil(o.Ca) {
		var ret string
		return ret
	}
	return *o.Ca
}

// GetCaOk returns a tuple with the Ca field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDevicesV2CertsPayload) GetCaOk() (*string, bool) {
	if o == nil || IsNil(o.Ca) {
		return nil, false
	}
	return o.Ca, true
}

// HasCa returns a boolean if a field has been set.
func (o *CreateDevicesV2CertsPayload) HasCa() bool {
	if o != nil && !IsNil(o.Ca) {
		return true
	}

	return false
}

// SetCa gets a reference to the given string and assigns it to the Ca field.
func (o *CreateDevicesV2CertsPayload) SetCa(v string) {
	o.Ca = &v
}

// GetCsr returns the Csr field value
func (o *CreateDevicesV2CertsPayload) GetCsr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Csr
}

// GetCsrOk returns a tuple with the Csr field value
// and a boolean to check if the value has been set.
func (o *CreateDevicesV2CertsPayload) GetCsrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Csr, true
}

// SetCsr sets field value
func (o *CreateDevicesV2CertsPayload) SetCsr(v string) {
	o.Csr = v
}

// GetEnabled returns the Enabled field value
func (o *CreateDevicesV2CertsPayload) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CreateDevicesV2CertsPayload) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CreateDevicesV2CertsPayload) SetEnabled(v bool) {
	o.Enabled = v
}

func (o CreateDevicesV2CertsPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDevicesV2CertsPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ca) {
		toSerialize["ca"] = o.Ca
	}
	toSerialize["csr"] = o.Csr
	toSerialize["enabled"] = o.Enabled

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateDevicesV2CertsPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"csr",
		"enabled",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateDevicesV2CertsPayload := _CreateDevicesV2CertsPayload{}

	err = json.Unmarshal(data, &varCreateDevicesV2CertsPayload)

	if err != nil {
		return err
	}

	*o = CreateDevicesV2CertsPayload(varCreateDevicesV2CertsPayload)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ca")
		delete(additionalProperties, "csr")
		delete(additionalProperties, "enabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateDevicesV2CertsPayload struct {
	value *CreateDevicesV2CertsPayload
	isSet bool
}

func (v NullableCreateDevicesV2CertsPayload) Get() *CreateDevicesV2CertsPayload {
	return v.value
}

func (v *NullableCreateDevicesV2CertsPayload) Set(val *CreateDevicesV2CertsPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDevicesV2CertsPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDevicesV2CertsPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDevicesV2CertsPayload(val *CreateDevicesV2CertsPayload) *NullableCreateDevicesV2CertsPayload {
	return &NullableCreateDevicesV2CertsPayload{value: val, isSet: true}
}

func (v NullableCreateDevicesV2CertsPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDevicesV2CertsPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


