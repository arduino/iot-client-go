/*
Arduino IoT Cloud API

Provides a set of endpoints to manage Arduino IoT Cloud **Devices**, **Things**, **Properties** and **Timeseries**. This API can be called just with any HTTP Client, or using one of these clients:  * [Javascript NPM package](https://www.npmjs.com/package/@arduino/arduino-iot-client)  * [Python PYPI Package](https://pypi.org/project/arduino-iot-client/)  * [Golang Module](https://github.com/arduino/iot-client-go)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ArduinoTags type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoTags{}

// ArduinoTags ArduinoTags media type (default view)
type ArduinoTags struct {
	Tags []Tag `json:"tags"`
}

type _ArduinoTags ArduinoTags

// NewArduinoTags instantiates a new ArduinoTags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoTags(tags []Tag) *ArduinoTags {
	this := ArduinoTags{}
	this.Tags = tags
	return &this
}

// NewArduinoTagsWithDefaults instantiates a new ArduinoTags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoTagsWithDefaults() *ArduinoTags {
	this := ArduinoTags{}
	return &this
}

// GetTags returns the Tags field value
func (o *ArduinoTags) GetTags() []Tag {
	if o == nil {
		var ret []Tag
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *ArduinoTags) GetTagsOk() ([]Tag, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *ArduinoTags) SetTags(v []Tag) {
	o.Tags = v
}

func (o ArduinoTags) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoTags) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tags"] = o.Tags
	return toSerialize, nil
}

func (o *ArduinoTags) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tags",
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

	varArduinoTags := _ArduinoTags{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varArduinoTags)

	if err != nil {
		return err
	}

	*o = ArduinoTags(varArduinoTags)

	return err
}

type NullableArduinoTags struct {
	value *ArduinoTags
	isSet bool
}

func (v NullableArduinoTags) Get() *ArduinoTags {
	return v.value
}

func (v *NullableArduinoTags) Set(val *ArduinoTags) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoTags) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoTags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoTags(val *ArduinoTags) *NullableArduinoTags {
	return &NullableArduinoTags{value: val, isSet: true}
}

func (v NullableArduinoTags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoTags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


