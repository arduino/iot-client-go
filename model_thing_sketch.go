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

// checks if the ThingSketch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThingSketch{}

// ThingSketch ThingSketchPayload describes a sketch of a thing
type ThingSketch struct {
	// The autogenerated sketch version
	SketchVersion *string `json:"sketch_version,omitempty"`
}

// NewThingSketch instantiates a new ThingSketch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThingSketch() *ThingSketch {
	this := ThingSketch{}
	return &this
}

// NewThingSketchWithDefaults instantiates a new ThingSketch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThingSketchWithDefaults() *ThingSketch {
	this := ThingSketch{}
	return &this
}

// GetSketchVersion returns the SketchVersion field value if set, zero value otherwise.
func (o *ThingSketch) GetSketchVersion() string {
	if o == nil || IsNil(o.SketchVersion) {
		var ret string
		return ret
	}
	return *o.SketchVersion
}

// GetSketchVersionOk returns a tuple with the SketchVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingSketch) GetSketchVersionOk() (*string, bool) {
	if o == nil || IsNil(o.SketchVersion) {
		return nil, false
	}
	return o.SketchVersion, true
}

// HasSketchVersion returns a boolean if a field has been set.
func (o *ThingSketch) HasSketchVersion() bool {
	if o != nil && !IsNil(o.SketchVersion) {
		return true
	}

	return false
}

// SetSketchVersion gets a reference to the given string and assigns it to the SketchVersion field.
func (o *ThingSketch) SetSketchVersion(v string) {
	o.SketchVersion = &v
}

func (o ThingSketch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThingSketch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SketchVersion) {
		toSerialize["sketch_version"] = o.SketchVersion
	}
	return toSerialize, nil
}

type NullableThingSketch struct {
	value *ThingSketch
	isSet bool
}

func (v NullableThingSketch) Get() *ThingSketch {
	return v.value
}

func (v *NullableThingSketch) Set(val *ThingSketch) {
	v.value = val
	v.isSet = true
}

func (v NullableThingSketch) IsSet() bool {
	return v.isSet
}

func (v *NullableThingSketch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThingSketch(val *ThingSketch) *NullableThingSketch {
	return &NullableThingSketch{value: val, isSet: true}
}

func (v NullableThingSketch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThingSketch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


