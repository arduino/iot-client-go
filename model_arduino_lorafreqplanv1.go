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

// checks if the ArduinoLorafreqplanv1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoLorafreqplanv1{}

// ArduinoLorafreqplanv1 ArduinoLorafreqplanv1 media type (default view)
type ArduinoLorafreqplanv1 struct {
	// Frequency plan only for advanced users
	Advanced bool `json:"advanced"`
	// The ID of the frequency paln
	Id string `json:"id"`
	// The name of the frequency plan
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _ArduinoLorafreqplanv1 ArduinoLorafreqplanv1

// NewArduinoLorafreqplanv1 instantiates a new ArduinoLorafreqplanv1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoLorafreqplanv1(advanced bool, id string, name string) *ArduinoLorafreqplanv1 {
	this := ArduinoLorafreqplanv1{}
	this.Advanced = advanced
	this.Id = id
	this.Name = name
	return &this
}

// NewArduinoLorafreqplanv1WithDefaults instantiates a new ArduinoLorafreqplanv1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoLorafreqplanv1WithDefaults() *ArduinoLorafreqplanv1 {
	this := ArduinoLorafreqplanv1{}
	return &this
}

// GetAdvanced returns the Advanced field value
func (o *ArduinoLorafreqplanv1) GetAdvanced() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Advanced
}

// GetAdvancedOk returns a tuple with the Advanced field value
// and a boolean to check if the value has been set.
func (o *ArduinoLorafreqplanv1) GetAdvancedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Advanced, true
}

// SetAdvanced sets field value
func (o *ArduinoLorafreqplanv1) SetAdvanced(v bool) {
	o.Advanced = v
}

// GetId returns the Id field value
func (o *ArduinoLorafreqplanv1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ArduinoLorafreqplanv1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ArduinoLorafreqplanv1) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ArduinoLorafreqplanv1) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ArduinoLorafreqplanv1) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ArduinoLorafreqplanv1) SetName(v string) {
	o.Name = v
}

func (o ArduinoLorafreqplanv1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoLorafreqplanv1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["advanced"] = o.Advanced
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ArduinoLorafreqplanv1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"advanced",
		"id",
		"name",
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

	varArduinoLorafreqplanv1 := _ArduinoLorafreqplanv1{}

	err = json.Unmarshal(data, &varArduinoLorafreqplanv1)

	if err != nil {
		return err
	}

	*o = ArduinoLorafreqplanv1(varArduinoLorafreqplanv1)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "advanced")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableArduinoLorafreqplanv1 struct {
	value *ArduinoLorafreqplanv1
	isSet bool
}

func (v NullableArduinoLorafreqplanv1) Get() *ArduinoLorafreqplanv1 {
	return v.value
}

func (v *NullableArduinoLorafreqplanv1) Set(val *ArduinoLorafreqplanv1) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoLorafreqplanv1) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoLorafreqplanv1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoLorafreqplanv1(val *ArduinoLorafreqplanv1) *NullableArduinoLorafreqplanv1 {
	return &NullableArduinoLorafreqplanv1{value: val, isSet: true}
}

func (v NullableArduinoLorafreqplanv1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoLorafreqplanv1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


