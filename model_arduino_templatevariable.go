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

// checks if the ArduinoTemplatevariable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoTemplatevariable{}

// ArduinoTemplatevariable ArduinoTemplatevariable media type (default view)
type ArduinoTemplatevariable struct {
	// The name of the variable
	Name string `json:"name"`
	// The permission of the linked variable
	Permission string `json:"permission"`
	// The name of the related thing
	ThingId string `json:"thing_id"`
	ThingTimezone *ArduinoTimezone `json:"thing_timezone,omitempty"`
	// The type of the variable
	Type string `json:"type"`
	// The name of the variable in the code
	VariableId string `json:"variable_id"`
}

// NewArduinoTemplatevariable instantiates a new ArduinoTemplatevariable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoTemplatevariable(name string, permission string, thingId string, type_ string, variableId string) *ArduinoTemplatevariable {
	this := ArduinoTemplatevariable{}
	this.Name = name
	this.Permission = permission
	this.ThingId = thingId
	this.Type = type_
	this.VariableId = variableId
	return &this
}

// NewArduinoTemplatevariableWithDefaults instantiates a new ArduinoTemplatevariable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoTemplatevariableWithDefaults() *ArduinoTemplatevariable {
	this := ArduinoTemplatevariable{}
	return &this
}

// GetName returns the Name field value
func (o *ArduinoTemplatevariable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ArduinoTemplatevariable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ArduinoTemplatevariable) SetName(v string) {
	o.Name = v
}

// GetPermission returns the Permission field value
func (o *ArduinoTemplatevariable) GetPermission() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value
// and a boolean to check if the value has been set.
func (o *ArduinoTemplatevariable) GetPermissionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permission, true
}

// SetPermission sets field value
func (o *ArduinoTemplatevariable) SetPermission(v string) {
	o.Permission = v
}

// GetThingId returns the ThingId field value
func (o *ArduinoTemplatevariable) GetThingId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThingId
}

// GetThingIdOk returns a tuple with the ThingId field value
// and a boolean to check if the value has been set.
func (o *ArduinoTemplatevariable) GetThingIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThingId, true
}

// SetThingId sets field value
func (o *ArduinoTemplatevariable) SetThingId(v string) {
	o.ThingId = v
}

// GetThingTimezone returns the ThingTimezone field value if set, zero value otherwise.
func (o *ArduinoTemplatevariable) GetThingTimezone() ArduinoTimezone {
	if o == nil || IsNil(o.ThingTimezone) {
		var ret ArduinoTimezone
		return ret
	}
	return *o.ThingTimezone
}

// GetThingTimezoneOk returns a tuple with the ThingTimezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoTemplatevariable) GetThingTimezoneOk() (*ArduinoTimezone, bool) {
	if o == nil || IsNil(o.ThingTimezone) {
		return nil, false
	}
	return o.ThingTimezone, true
}

// HasThingTimezone returns a boolean if a field has been set.
func (o *ArduinoTemplatevariable) HasThingTimezone() bool {
	if o != nil && !IsNil(o.ThingTimezone) {
		return true
	}

	return false
}

// SetThingTimezone gets a reference to the given ArduinoTimezone and assigns it to the ThingTimezone field.
func (o *ArduinoTemplatevariable) SetThingTimezone(v ArduinoTimezone) {
	o.ThingTimezone = &v
}

// GetType returns the Type field value
func (o *ArduinoTemplatevariable) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ArduinoTemplatevariable) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ArduinoTemplatevariable) SetType(v string) {
	o.Type = v
}

// GetVariableId returns the VariableId field value
func (o *ArduinoTemplatevariable) GetVariableId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VariableId
}

// GetVariableIdOk returns a tuple with the VariableId field value
// and a boolean to check if the value has been set.
func (o *ArduinoTemplatevariable) GetVariableIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VariableId, true
}

// SetVariableId sets field value
func (o *ArduinoTemplatevariable) SetVariableId(v string) {
	o.VariableId = v
}

func (o ArduinoTemplatevariable) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoTemplatevariable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["permission"] = o.Permission
	toSerialize["thing_id"] = o.ThingId
	if !IsNil(o.ThingTimezone) {
		toSerialize["thing_timezone"] = o.ThingTimezone
	}
	toSerialize["type"] = o.Type
	toSerialize["variable_id"] = o.VariableId
	return toSerialize, nil
}

type NullableArduinoTemplatevariable struct {
	value *ArduinoTemplatevariable
	isSet bool
}

func (v NullableArduinoTemplatevariable) Get() *ArduinoTemplatevariable {
	return v.value
}

func (v *NullableArduinoTemplatevariable) Set(val *ArduinoTemplatevariable) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoTemplatevariable) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoTemplatevariable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoTemplatevariable(val *ArduinoTemplatevariable) *NullableArduinoTemplatevariable {
	return &NullableArduinoTemplatevariable{value: val, isSet: true}
}

func (v NullableArduinoTemplatevariable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoTemplatevariable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

