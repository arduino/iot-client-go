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
	"bytes"
	"fmt"
)

// checks if the ArduinoDevicev2SimpleProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoDevicev2SimpleProperties{}

// ArduinoDevicev2SimpleProperties ArduinoDevicev2SimpleProperties media type (default view)
type ArduinoDevicev2SimpleProperties struct {
	// The name of the property
	Name string `json:"name"`
	// Update date of the property
	UpdatedAt time.Time `json:"updated_at"`
	// Value of the property
	Value interface{} `json:"value"`
}

type _ArduinoDevicev2SimpleProperties ArduinoDevicev2SimpleProperties

// NewArduinoDevicev2SimpleProperties instantiates a new ArduinoDevicev2SimpleProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoDevicev2SimpleProperties(name string, updatedAt time.Time, value interface{}) *ArduinoDevicev2SimpleProperties {
	this := ArduinoDevicev2SimpleProperties{}
	this.Name = name
	this.UpdatedAt = updatedAt
	this.Value = value
	return &this
}

// NewArduinoDevicev2SimplePropertiesWithDefaults instantiates a new ArduinoDevicev2SimpleProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoDevicev2SimplePropertiesWithDefaults() *ArduinoDevicev2SimpleProperties {
	this := ArduinoDevicev2SimpleProperties{}
	return &this
}

// GetName returns the Name field value
func (o *ArduinoDevicev2SimpleProperties) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ArduinoDevicev2SimpleProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ArduinoDevicev2SimpleProperties) SetName(v string) {
	o.Name = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ArduinoDevicev2SimpleProperties) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ArduinoDevicev2SimpleProperties) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ArduinoDevicev2SimpleProperties) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ArduinoDevicev2SimpleProperties) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArduinoDevicev2SimpleProperties) GetValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ArduinoDevicev2SimpleProperties) SetValue(v interface{}) {
	o.Value = v
}

func (o ArduinoDevicev2SimpleProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoDevicev2SimpleProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["updated_at"] = o.UpdatedAt
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

func (o *ArduinoDevicev2SimpleProperties) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"updated_at",
		"value",
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

	varArduinoDevicev2SimpleProperties := _ArduinoDevicev2SimpleProperties{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varArduinoDevicev2SimpleProperties)

	if err != nil {
		return err
	}

	*o = ArduinoDevicev2SimpleProperties(varArduinoDevicev2SimpleProperties)

	return err
}

type NullableArduinoDevicev2SimpleProperties struct {
	value *ArduinoDevicev2SimpleProperties
	isSet bool
}

func (v NullableArduinoDevicev2SimpleProperties) Get() *ArduinoDevicev2SimpleProperties {
	return v.value
}

func (v *NullableArduinoDevicev2SimpleProperties) Set(val *ArduinoDevicev2SimpleProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoDevicev2SimpleProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoDevicev2SimpleProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoDevicev2SimpleProperties(val *ArduinoDevicev2SimpleProperties) *NullableArduinoDevicev2SimpleProperties {
	return &NullableArduinoDevicev2SimpleProperties{value: val, isSet: true}
}

func (v NullableArduinoDevicev2SimpleProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoDevicev2SimpleProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


