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

// checks if the ArduinoDevicev2templatedevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoDevicev2templatedevice{}

// ArduinoDevicev2templatedevice ArduinoDevicev2templatedevice media type (default view)
type ArduinoDevicev2templatedevice struct {
	// The device fqbn
	Fqbn *string `json:"fqbn,omitempty"`
	// The device type name
	Name *string `json:"name,omitempty"`
}

// NewArduinoDevicev2templatedevice instantiates a new ArduinoDevicev2templatedevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoDevicev2templatedevice() *ArduinoDevicev2templatedevice {
	this := ArduinoDevicev2templatedevice{}
	return &this
}

// NewArduinoDevicev2templatedeviceWithDefaults instantiates a new ArduinoDevicev2templatedevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoDevicev2templatedeviceWithDefaults() *ArduinoDevicev2templatedevice {
	this := ArduinoDevicev2templatedevice{}
	return &this
}

// GetFqbn returns the Fqbn field value if set, zero value otherwise.
func (o *ArduinoDevicev2templatedevice) GetFqbn() string {
	if o == nil || IsNil(o.Fqbn) {
		var ret string
		return ret
	}
	return *o.Fqbn
}

// GetFqbnOk returns a tuple with the Fqbn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoDevicev2templatedevice) GetFqbnOk() (*string, bool) {
	if o == nil || IsNil(o.Fqbn) {
		return nil, false
	}
	return o.Fqbn, true
}

// HasFqbn returns a boolean if a field has been set.
func (o *ArduinoDevicev2templatedevice) HasFqbn() bool {
	if o != nil && !IsNil(o.Fqbn) {
		return true
	}

	return false
}

// SetFqbn gets a reference to the given string and assigns it to the Fqbn field.
func (o *ArduinoDevicev2templatedevice) SetFqbn(v string) {
	o.Fqbn = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ArduinoDevicev2templatedevice) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoDevicev2templatedevice) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ArduinoDevicev2templatedevice) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ArduinoDevicev2templatedevice) SetName(v string) {
	o.Name = &v
}

func (o ArduinoDevicev2templatedevice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoDevicev2templatedevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fqbn) {
		toSerialize["fqbn"] = o.Fqbn
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableArduinoDevicev2templatedevice struct {
	value *ArduinoDevicev2templatedevice
	isSet bool
}

func (v NullableArduinoDevicev2templatedevice) Get() *ArduinoDevicev2templatedevice {
	return v.value
}

func (v *NullableArduinoDevicev2templatedevice) Set(val *ArduinoDevicev2templatedevice) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoDevicev2templatedevice) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoDevicev2templatedevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoDevicev2templatedevice(val *ArduinoDevicev2templatedevice) *NullableArduinoDevicev2templatedevice {
	return &NullableArduinoDevicev2templatedevice{value: val, isSet: true}
}

func (v NullableArduinoDevicev2templatedevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoDevicev2templatedevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


