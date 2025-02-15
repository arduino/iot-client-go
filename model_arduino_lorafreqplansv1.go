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

// checks if the ArduinoLorafreqplansv1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoLorafreqplansv1{}

// ArduinoLorafreqplansv1 ArduinoLorafreqplansv1 media type (default view)
type ArduinoLorafreqplansv1 struct {
	// The list of frequency plans
	FrequencyPlans []ArduinoLorafreqplanv1 `json:"frequency_plans,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ArduinoLorafreqplansv1 ArduinoLorafreqplansv1

// NewArduinoLorafreqplansv1 instantiates a new ArduinoLorafreqplansv1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoLorafreqplansv1() *ArduinoLorafreqplansv1 {
	this := ArduinoLorafreqplansv1{}
	return &this
}

// NewArduinoLorafreqplansv1WithDefaults instantiates a new ArduinoLorafreqplansv1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoLorafreqplansv1WithDefaults() *ArduinoLorafreqplansv1 {
	this := ArduinoLorafreqplansv1{}
	return &this
}

// GetFrequencyPlans returns the FrequencyPlans field value if set, zero value otherwise.
func (o *ArduinoLorafreqplansv1) GetFrequencyPlans() []ArduinoLorafreqplanv1 {
	if o == nil || IsNil(o.FrequencyPlans) {
		var ret []ArduinoLorafreqplanv1
		return ret
	}
	return o.FrequencyPlans
}

// GetFrequencyPlansOk returns a tuple with the FrequencyPlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoLorafreqplansv1) GetFrequencyPlansOk() ([]ArduinoLorafreqplanv1, bool) {
	if o == nil || IsNil(o.FrequencyPlans) {
		return nil, false
	}
	return o.FrequencyPlans, true
}

// HasFrequencyPlans returns a boolean if a field has been set.
func (o *ArduinoLorafreqplansv1) HasFrequencyPlans() bool {
	if o != nil && !IsNil(o.FrequencyPlans) {
		return true
	}

	return false
}

// SetFrequencyPlans gets a reference to the given []ArduinoLorafreqplanv1 and assigns it to the FrequencyPlans field.
func (o *ArduinoLorafreqplansv1) SetFrequencyPlans(v []ArduinoLorafreqplanv1) {
	o.FrequencyPlans = v
}

func (o ArduinoLorafreqplansv1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoLorafreqplansv1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FrequencyPlans) {
		toSerialize["frequency_plans"] = o.FrequencyPlans
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ArduinoLorafreqplansv1) UnmarshalJSON(data []byte) (err error) {
	varArduinoLorafreqplansv1 := _ArduinoLorafreqplansv1{}

	err = json.Unmarshal(data, &varArduinoLorafreqplansv1)

	if err != nil {
		return err
	}

	*o = ArduinoLorafreqplansv1(varArduinoLorafreqplansv1)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "frequency_plans")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableArduinoLorafreqplansv1 struct {
	value *ArduinoLorafreqplansv1
	isSet bool
}

func (v NullableArduinoLorafreqplansv1) Get() *ArduinoLorafreqplansv1 {
	return v.value
}

func (v *NullableArduinoLorafreqplansv1) Set(val *ArduinoLorafreqplansv1) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoLorafreqplansv1) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoLorafreqplansv1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoLorafreqplansv1(val *ArduinoLorafreqplansv1) *NullableArduinoLorafreqplansv1 {
	return &NullableArduinoLorafreqplansv1{value: val, isSet: true}
}

func (v NullableArduinoLorafreqplansv1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoLorafreqplansv1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


