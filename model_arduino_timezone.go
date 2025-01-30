/*
Arduino IoT Cloud API

Provides a set of endpoints to manage Arduino IoT Cloud **Devices**, **Things**, **Properties** and **Timeseries**. This API can be called just with any HTTP Client, or using one of these clients:  * [Javascript NPM package](https://www.npmjs.com/package/@arduino/arduino-iot-client)  * [Python PYPI Package](https://pypi.org/project/arduino-iot-client/)  * [Golang Module](https://github.com/arduino/iot-client-go)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the ArduinoTimezone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoTimezone{}

// ArduinoTimezone ArduinoTimezone media type (default view)
type ArduinoTimezone struct {
	// Name of the time zone.
	Name string `json:"name"`
	// Current UTC DST offset in seconds.
	Offset int64 `json:"offset"`
	// Date until the offset is valid.
	Until time.Time `json:"until"`
	AdditionalProperties map[string]interface{}
}

type _ArduinoTimezone ArduinoTimezone

// NewArduinoTimezone instantiates a new ArduinoTimezone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoTimezone(name string, offset int64, until time.Time) *ArduinoTimezone {
	this := ArduinoTimezone{}
	this.Name = name
	this.Offset = offset
	this.Until = until
	return &this
}

// NewArduinoTimezoneWithDefaults instantiates a new ArduinoTimezone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoTimezoneWithDefaults() *ArduinoTimezone {
	this := ArduinoTimezone{}
	return &this
}

// GetName returns the Name field value
func (o *ArduinoTimezone) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ArduinoTimezone) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ArduinoTimezone) SetName(v string) {
	o.Name = v
}

// GetOffset returns the Offset field value
func (o *ArduinoTimezone) GetOffset() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *ArduinoTimezone) GetOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *ArduinoTimezone) SetOffset(v int64) {
	o.Offset = v
}

// GetUntil returns the Until field value
func (o *ArduinoTimezone) GetUntil() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Until
}

// GetUntilOk returns a tuple with the Until field value
// and a boolean to check if the value has been set.
func (o *ArduinoTimezone) GetUntilOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Until, true
}

// SetUntil sets field value
func (o *ArduinoTimezone) SetUntil(v time.Time) {
	o.Until = v
}

func (o ArduinoTimezone) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoTimezone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["offset"] = o.Offset
	toSerialize["until"] = o.Until

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ArduinoTimezone) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"offset",
		"until",
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

	varArduinoTimezone := _ArduinoTimezone{}

	err = json.Unmarshal(data, &varArduinoTimezone)

	if err != nil {
		return err
	}

	*o = ArduinoTimezone(varArduinoTimezone)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "offset")
		delete(additionalProperties, "until")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableArduinoTimezone struct {
	value *ArduinoTimezone
	isSet bool
}

func (v NullableArduinoTimezone) Get() *ArduinoTimezone {
	return v.value
}

func (v *NullableArduinoTimezone) Set(val *ArduinoTimezone) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoTimezone) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoTimezone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoTimezone(val *ArduinoTimezone) *NullableArduinoTimezone {
	return &NullableArduinoTimezone{value: val, isSet: true}
}

func (v NullableArduinoTimezone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoTimezone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


