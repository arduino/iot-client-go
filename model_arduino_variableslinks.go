/*
Arduino IoT Cloud API

 Provides a set of endpoints to manage Arduino IoT Cloud **Devices**, **Things**, **Properties** and **Timeseries**. This API can be called just with any HTTP Client, or using one of these clients:  * [Javascript NPM package](https://www.npmjs.com/package/@arduino/arduino-iot-client)  * [Python PYPI Package](https://pypi.org/project/arduino-iot-client/)  * [Golang Module](https://github.com/arduino/iot-client-go)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

// checks if the ArduinoVariableslinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoVariableslinks{}

// ArduinoVariableslinks ArduinoVariableslinks media type (default view)
type ArduinoVariableslinks struct {
	// The ids of the linked variables
	Variables []string `json:"variables"`
}

// NewArduinoVariableslinks instantiates a new ArduinoVariableslinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoVariableslinks(variables []string) *ArduinoVariableslinks {
	this := ArduinoVariableslinks{}
	this.Variables = variables
	return &this
}

// NewArduinoVariableslinksWithDefaults instantiates a new ArduinoVariableslinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoVariableslinksWithDefaults() *ArduinoVariableslinks {
	this := ArduinoVariableslinks{}
	return &this
}

// GetVariables returns the Variables field value
func (o *ArduinoVariableslinks) GetVariables() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value
// and a boolean to check if the value has been set.
func (o *ArduinoVariableslinks) GetVariablesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Variables, true
}

// SetVariables sets field value
func (o *ArduinoVariableslinks) SetVariables(v []string) {
	o.Variables = v
}

func (o ArduinoVariableslinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoVariableslinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["variables"] = o.Variables
	return toSerialize, nil
}

type NullableArduinoVariableslinks struct {
	value *ArduinoVariableslinks
	isSet bool
}

func (v NullableArduinoVariableslinks) Get() *ArduinoVariableslinks {
	return v.value
}

func (v *NullableArduinoVariableslinks) Set(val *ArduinoVariableslinks) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoVariableslinks) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoVariableslinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoVariableslinks(val *ArduinoVariableslinks) *NullableArduinoVariableslinks {
	return &NullableArduinoVariableslinks{value: val, isSet: true}
}

func (v NullableArduinoVariableslinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoVariableslinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


