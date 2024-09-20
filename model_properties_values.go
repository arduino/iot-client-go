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

// checks if the PropertiesValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PropertiesValues{}

// PropertiesValues struct for PropertiesValues
type PropertiesValues struct {
	// If true, send property values to device's input topic.
	Input *bool `json:"input,omitempty"`
	Properties []PropertiesValue `json:"properties"`
}

// NewPropertiesValues instantiates a new PropertiesValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertiesValues(properties []PropertiesValue) *PropertiesValues {
	this := PropertiesValues{}
	var input bool = false
	this.Input = &input
	this.Properties = properties
	return &this
}

// NewPropertiesValuesWithDefaults instantiates a new PropertiesValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertiesValuesWithDefaults() *PropertiesValues {
	this := PropertiesValues{}
	var input bool = false
	this.Input = &input
	return &this
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *PropertiesValues) GetInput() bool {
	if o == nil || IsNil(o.Input) {
		var ret bool
		return ret
	}
	return *o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertiesValues) GetInputOk() (*bool, bool) {
	if o == nil || IsNil(o.Input) {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *PropertiesValues) HasInput() bool {
	if o != nil && !IsNil(o.Input) {
		return true
	}

	return false
}

// SetInput gets a reference to the given bool and assigns it to the Input field.
func (o *PropertiesValues) SetInput(v bool) {
	o.Input = &v
}

// GetProperties returns the Properties field value
func (o *PropertiesValues) GetProperties() []PropertiesValue {
	if o == nil {
		var ret []PropertiesValue
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *PropertiesValues) GetPropertiesOk() ([]PropertiesValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties, true
}

// SetProperties sets field value
func (o *PropertiesValues) SetProperties(v []PropertiesValue) {
	o.Properties = v
}

func (o PropertiesValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PropertiesValues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Input) {
		toSerialize["input"] = o.Input
	}
	toSerialize["properties"] = o.Properties
	return toSerialize, nil
}

type NullablePropertiesValues struct {
	value *PropertiesValues
	isSet bool
}

func (v NullablePropertiesValues) Get() *PropertiesValues {
	return v.value
}

func (v *NullablePropertiesValues) Set(val *PropertiesValues) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertiesValues) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertiesValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertiesValues(val *PropertiesValues) *NullablePropertiesValues {
	return &NullablePropertiesValues{value: val, isSet: true}
}

func (v NullablePropertiesValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertiesValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


