/*
Arduino IoT Cloud API

 Provides a set of endpoints to manage Arduino IoT Cloud **Devices**, **Things**, **Properties** and **Timeseries**. This API can be called just with any HTTP Client, or using one of these clients:  * [Javascript NPM package](https://www.npmjs.com/package/@arduino/arduino-iot-client)  * [Python PYPI Package](https://pypi.org/project/arduino-iot-client/)  * [Golang Module](https://github.com/arduino/iot-client-go)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iot

import (
	"encoding/json"
)

// checks if the ThingClone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThingClone{}

// ThingClone Payload to clone a new thing from an existing one
type ThingClone struct {
	// Include tags in clone procedure
	IncludeTags *bool `json:"include_tags,omitempty"`
	// The friendly name of the thing
	Name string `json:"name"`
}

// NewThingClone instantiates a new ThingClone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThingClone(name string) *ThingClone {
	this := ThingClone{}
	this.Name = name
	return &this
}

// NewThingCloneWithDefaults instantiates a new ThingClone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThingCloneWithDefaults() *ThingClone {
	this := ThingClone{}
	return &this
}

// GetIncludeTags returns the IncludeTags field value if set, zero value otherwise.
func (o *ThingClone) GetIncludeTags() bool {
	if o == nil || IsNil(o.IncludeTags) {
		var ret bool
		return ret
	}
	return *o.IncludeTags
}

// GetIncludeTagsOk returns a tuple with the IncludeTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingClone) GetIncludeTagsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeTags) {
		return nil, false
	}
	return o.IncludeTags, true
}

// HasIncludeTags returns a boolean if a field has been set.
func (o *ThingClone) HasIncludeTags() bool {
	if o != nil && !IsNil(o.IncludeTags) {
		return true
	}

	return false
}

// SetIncludeTags gets a reference to the given bool and assigns it to the IncludeTags field.
func (o *ThingClone) SetIncludeTags(v bool) {
	o.IncludeTags = &v
}

// GetName returns the Name field value
func (o *ThingClone) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ThingClone) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ThingClone) SetName(v string) {
	o.Name = v
}

func (o ThingClone) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThingClone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IncludeTags) {
		toSerialize["include_tags"] = o.IncludeTags
	}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableThingClone struct {
	value *ThingClone
	isSet bool
}

func (v NullableThingClone) Get() *ThingClone {
	return v.value
}

func (v *NullableThingClone) Set(val *ThingClone) {
	v.value = val
	v.isSet = true
}

func (v NullableThingClone) IsSet() bool {
	return v.isSet
}

func (v *NullableThingClone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThingClone(val *ThingClone) *NullableThingClone {
	return &NullableThingClone{value: val, isSet: true}
}

func (v NullableThingClone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThingClone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


