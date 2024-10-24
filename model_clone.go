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

// checks if the Clone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Clone{}

// Clone struct for Clone
type Clone struct {
	// The overrides to apply to the cloned dashboard. An override is a tuple of ids: the id of the thing to override and the id of the new thing to link
	Overrides []Override `json:"overrides,omitempty"`
}

// NewClone instantiates a new Clone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClone() *Clone {
	this := Clone{}
	return &this
}

// NewCloneWithDefaults instantiates a new Clone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloneWithDefaults() *Clone {
	this := Clone{}
	return &this
}

// GetOverrides returns the Overrides field value if set, zero value otherwise.
func (o *Clone) GetOverrides() []Override {
	if o == nil || IsNil(o.Overrides) {
		var ret []Override
		return ret
	}
	return o.Overrides
}

// GetOverridesOk returns a tuple with the Overrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Clone) GetOverridesOk() ([]Override, bool) {
	if o == nil || IsNil(o.Overrides) {
		return nil, false
	}
	return o.Overrides, true
}

// HasOverrides returns a boolean if a field has been set.
func (o *Clone) HasOverrides() bool {
	if o != nil && !IsNil(o.Overrides) {
		return true
	}

	return false
}

// SetOverrides gets a reference to the given []Override and assigns it to the Overrides field.
func (o *Clone) SetOverrides(v []Override) {
	o.Overrides = v
}

func (o Clone) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Clone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Overrides) {
		toSerialize["overrides"] = o.Overrides
	}
	return toSerialize, nil
}

type NullableClone struct {
	value *Clone
	isSet bool
}

func (v NullableClone) Get() *Clone {
	return v.value
}

func (v *NullableClone) Set(val *Clone) {
	v.value = val
	v.isSet = true
}

func (v NullableClone) IsSet() bool {
	return v.isSet
}

func (v *NullableClone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClone(val *Clone) *NullableClone {
	return &NullableClone{value: val, isSet: true}
}

func (v NullableClone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


