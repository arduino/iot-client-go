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

// checks if the Widgetlink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Widgetlink{}

// Widgetlink struct for Widgetlink
type Widgetlink struct {
	Variables []string `json:"variables,omitempty"`
}

// NewWidgetlink instantiates a new Widgetlink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWidgetlink() *Widgetlink {
	this := Widgetlink{}
	return &this
}

// NewWidgetlinkWithDefaults instantiates a new Widgetlink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWidgetlinkWithDefaults() *Widgetlink {
	this := Widgetlink{}
	return &this
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *Widgetlink) GetVariables() []string {
	if o == nil || IsNil(o.Variables) {
		var ret []string
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Widgetlink) GetVariablesOk() ([]string, bool) {
	if o == nil || IsNil(o.Variables) {
		return nil, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *Widgetlink) HasVariables() bool {
	if o != nil && !IsNil(o.Variables) {
		return true
	}

	return false
}

// SetVariables gets a reference to the given []string and assigns it to the Variables field.
func (o *Widgetlink) SetVariables(v []string) {
	o.Variables = v
}

func (o Widgetlink) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Widgetlink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Variables) {
		toSerialize["variables"] = o.Variables
	}
	return toSerialize, nil
}

type NullableWidgetlink struct {
	value *Widgetlink
	isSet bool
}

func (v NullableWidgetlink) Get() *Widgetlink {
	return v.value
}

func (v *NullableWidgetlink) Set(val *Widgetlink) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetlink) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetlink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetlink(val *Widgetlink) *NullableWidgetlink {
	return &NullableWidgetlink{value: val, isSet: true}
}

func (v NullableWidgetlink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetlink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


