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

// checks if the ArduinoTemplateproperty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoTemplateproperty{}

// ArduinoTemplateproperty ArduinoTemplateproperty media type (default view)
type ArduinoTemplateproperty struct {
	// The friendly id of the property
	Id *string `json:"id,omitempty"`
	// The friendly name of the property
	Name string `json:"name"`
	// The permission of the property
	Permission string `json:"permission"`
	// The type of the property
	Type string `json:"type"`
	// The update frequency in seconds, or the amount of the property has to change in order to trigger an update
	UpdateParameter *float64 `json:"update_parameter,omitempty"`
	// The update strategy for the property value
	UpdateStrategy string `json:"update_strategy"`
	// The sketch variable name of the property
	VariableName *string `json:"variable_name,omitempty"`
}

// NewArduinoTemplateproperty instantiates a new ArduinoTemplateproperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoTemplateproperty(name string, permission string, type_ string, updateStrategy string) *ArduinoTemplateproperty {
	this := ArduinoTemplateproperty{}
	this.Name = name
	this.Permission = permission
	this.Type = type_
	this.UpdateStrategy = updateStrategy
	return &this
}

// NewArduinoTemplatepropertyWithDefaults instantiates a new ArduinoTemplateproperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoTemplatepropertyWithDefaults() *ArduinoTemplateproperty {
	this := ArduinoTemplateproperty{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ArduinoTemplateproperty) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoTemplateproperty) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ArduinoTemplateproperty) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ArduinoTemplateproperty) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *ArduinoTemplateproperty) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ArduinoTemplateproperty) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ArduinoTemplateproperty) SetName(v string) {
	o.Name = v
}

// GetPermission returns the Permission field value
func (o *ArduinoTemplateproperty) GetPermission() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value
// and a boolean to check if the value has been set.
func (o *ArduinoTemplateproperty) GetPermissionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permission, true
}

// SetPermission sets field value
func (o *ArduinoTemplateproperty) SetPermission(v string) {
	o.Permission = v
}

// GetType returns the Type field value
func (o *ArduinoTemplateproperty) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ArduinoTemplateproperty) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ArduinoTemplateproperty) SetType(v string) {
	o.Type = v
}

// GetUpdateParameter returns the UpdateParameter field value if set, zero value otherwise.
func (o *ArduinoTemplateproperty) GetUpdateParameter() float64 {
	if o == nil || IsNil(o.UpdateParameter) {
		var ret float64
		return ret
	}
	return *o.UpdateParameter
}

// GetUpdateParameterOk returns a tuple with the UpdateParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoTemplateproperty) GetUpdateParameterOk() (*float64, bool) {
	if o == nil || IsNil(o.UpdateParameter) {
		return nil, false
	}
	return o.UpdateParameter, true
}

// HasUpdateParameter returns a boolean if a field has been set.
func (o *ArduinoTemplateproperty) HasUpdateParameter() bool {
	if o != nil && !IsNil(o.UpdateParameter) {
		return true
	}

	return false
}

// SetUpdateParameter gets a reference to the given float64 and assigns it to the UpdateParameter field.
func (o *ArduinoTemplateproperty) SetUpdateParameter(v float64) {
	o.UpdateParameter = &v
}

// GetUpdateStrategy returns the UpdateStrategy field value
func (o *ArduinoTemplateproperty) GetUpdateStrategy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdateStrategy
}

// GetUpdateStrategyOk returns a tuple with the UpdateStrategy field value
// and a boolean to check if the value has been set.
func (o *ArduinoTemplateproperty) GetUpdateStrategyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdateStrategy, true
}

// SetUpdateStrategy sets field value
func (o *ArduinoTemplateproperty) SetUpdateStrategy(v string) {
	o.UpdateStrategy = v
}

// GetVariableName returns the VariableName field value if set, zero value otherwise.
func (o *ArduinoTemplateproperty) GetVariableName() string {
	if o == nil || IsNil(o.VariableName) {
		var ret string
		return ret
	}
	return *o.VariableName
}

// GetVariableNameOk returns a tuple with the VariableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoTemplateproperty) GetVariableNameOk() (*string, bool) {
	if o == nil || IsNil(o.VariableName) {
		return nil, false
	}
	return o.VariableName, true
}

// HasVariableName returns a boolean if a field has been set.
func (o *ArduinoTemplateproperty) HasVariableName() bool {
	if o != nil && !IsNil(o.VariableName) {
		return true
	}

	return false
}

// SetVariableName gets a reference to the given string and assigns it to the VariableName field.
func (o *ArduinoTemplateproperty) SetVariableName(v string) {
	o.VariableName = &v
}

func (o ArduinoTemplateproperty) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoTemplateproperty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	toSerialize["permission"] = o.Permission
	toSerialize["type"] = o.Type
	if !IsNil(o.UpdateParameter) {
		toSerialize["update_parameter"] = o.UpdateParameter
	}
	toSerialize["update_strategy"] = o.UpdateStrategy
	if !IsNil(o.VariableName) {
		toSerialize["variable_name"] = o.VariableName
	}
	return toSerialize, nil
}

type NullableArduinoTemplateproperty struct {
	value *ArduinoTemplateproperty
	isSet bool
}

func (v NullableArduinoTemplateproperty) Get() *ArduinoTemplateproperty {
	return v.value
}

func (v *NullableArduinoTemplateproperty) Set(val *ArduinoTemplateproperty) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoTemplateproperty) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoTemplateproperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoTemplateproperty(val *ArduinoTemplateproperty) *NullableArduinoTemplateproperty {
	return &NullableArduinoTemplateproperty{value: val, isSet: true}
}

func (v NullableArduinoTemplateproperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoTemplateproperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


