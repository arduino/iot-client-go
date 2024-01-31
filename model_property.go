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

// checks if the Property type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Property{}

// Property PropertyPayload describes a property of a thing. No field is mandatory
type Property struct {
	// Maximum value of this property
	MaxValue *float64 `json:"max_value,omitempty"`
	// Minimum value of this property
	MinValue *float64 `json:"min_value,omitempty"`
	// The friendly name of the property
	Name string `json:"name"`
	// The permission of the property
	Permission string `json:"permission"`
	// If true, data will persist into a timeseries database
	Persist *bool `json:"persist,omitempty"`
	// The integer id of the property
	Tag *int64 `json:"tag,omitempty"`
	// The type of the property
	Type string `json:"type"`
	// The update frequency in seconds, or the amount of the property has to change in order to trigger an update
	UpdateParameter *float64 `json:"update_parameter,omitempty"`
	// The update strategy for the property value
	UpdateStrategy string `json:"update_strategy"`
	// The  sketch variable name of the property
	VariableName *string `json:"variable_name,omitempty"`
}

// NewProperty instantiates a new Property object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProperty(name string, permission string, type_ string, updateStrategy string) *Property {
	this := Property{}
	this.Name = name
	this.Permission = permission
	var persist bool = true
	this.Persist = &persist
	this.Type = type_
	this.UpdateStrategy = updateStrategy
	return &this
}

// NewPropertyWithDefaults instantiates a new Property object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyWithDefaults() *Property {
	this := Property{}
	var persist bool = true
	this.Persist = &persist
	return &this
}

// GetMaxValue returns the MaxValue field value if set, zero value otherwise.
func (o *Property) GetMaxValue() float64 {
	if o == nil || IsNil(o.MaxValue) {
		var ret float64
		return ret
	}
	return *o.MaxValue
}

// GetMaxValueOk returns a tuple with the MaxValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetMaxValueOk() (*float64, bool) {
	if o == nil || IsNil(o.MaxValue) {
		return nil, false
	}
	return o.MaxValue, true
}

// HasMaxValue returns a boolean if a field has been set.
func (o *Property) HasMaxValue() bool {
	if o != nil && !IsNil(o.MaxValue) {
		return true
	}

	return false
}

// SetMaxValue gets a reference to the given float64 and assigns it to the MaxValue field.
func (o *Property) SetMaxValue(v float64) {
	o.MaxValue = &v
}

// GetMinValue returns the MinValue field value if set, zero value otherwise.
func (o *Property) GetMinValue() float64 {
	if o == nil || IsNil(o.MinValue) {
		var ret float64
		return ret
	}
	return *o.MinValue
}

// GetMinValueOk returns a tuple with the MinValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetMinValueOk() (*float64, bool) {
	if o == nil || IsNil(o.MinValue) {
		return nil, false
	}
	return o.MinValue, true
}

// HasMinValue returns a boolean if a field has been set.
func (o *Property) HasMinValue() bool {
	if o != nil && !IsNil(o.MinValue) {
		return true
	}

	return false
}

// SetMinValue gets a reference to the given float64 and assigns it to the MinValue field.
func (o *Property) SetMinValue(v float64) {
	o.MinValue = &v
}

// GetName returns the Name field value
func (o *Property) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Property) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Property) SetName(v string) {
	o.Name = v
}

// GetPermission returns the Permission field value
func (o *Property) GetPermission() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value
// and a boolean to check if the value has been set.
func (o *Property) GetPermissionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permission, true
}

// SetPermission sets field value
func (o *Property) SetPermission(v string) {
	o.Permission = v
}

// GetPersist returns the Persist field value if set, zero value otherwise.
func (o *Property) GetPersist() bool {
	if o == nil || IsNil(o.Persist) {
		var ret bool
		return ret
	}
	return *o.Persist
}

// GetPersistOk returns a tuple with the Persist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetPersistOk() (*bool, bool) {
	if o == nil || IsNil(o.Persist) {
		return nil, false
	}
	return o.Persist, true
}

// HasPersist returns a boolean if a field has been set.
func (o *Property) HasPersist() bool {
	if o != nil && !IsNil(o.Persist) {
		return true
	}

	return false
}

// SetPersist gets a reference to the given bool and assigns it to the Persist field.
func (o *Property) SetPersist(v bool) {
	o.Persist = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *Property) GetTag() int64 {
	if o == nil || IsNil(o.Tag) {
		var ret int64
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetTagOk() (*int64, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *Property) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given int64 and assigns it to the Tag field.
func (o *Property) SetTag(v int64) {
	o.Tag = &v
}

// GetType returns the Type field value
func (o *Property) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Property) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Property) SetType(v string) {
	o.Type = v
}

// GetUpdateParameter returns the UpdateParameter field value if set, zero value otherwise.
func (o *Property) GetUpdateParameter() float64 {
	if o == nil || IsNil(o.UpdateParameter) {
		var ret float64
		return ret
	}
	return *o.UpdateParameter
}

// GetUpdateParameterOk returns a tuple with the UpdateParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetUpdateParameterOk() (*float64, bool) {
	if o == nil || IsNil(o.UpdateParameter) {
		return nil, false
	}
	return o.UpdateParameter, true
}

// HasUpdateParameter returns a boolean if a field has been set.
func (o *Property) HasUpdateParameter() bool {
	if o != nil && !IsNil(o.UpdateParameter) {
		return true
	}

	return false
}

// SetUpdateParameter gets a reference to the given float64 and assigns it to the UpdateParameter field.
func (o *Property) SetUpdateParameter(v float64) {
	o.UpdateParameter = &v
}

// GetUpdateStrategy returns the UpdateStrategy field value
func (o *Property) GetUpdateStrategy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdateStrategy
}

// GetUpdateStrategyOk returns a tuple with the UpdateStrategy field value
// and a boolean to check if the value has been set.
func (o *Property) GetUpdateStrategyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdateStrategy, true
}

// SetUpdateStrategy sets field value
func (o *Property) SetUpdateStrategy(v string) {
	o.UpdateStrategy = v
}

// GetVariableName returns the VariableName field value if set, zero value otherwise.
func (o *Property) GetVariableName() string {
	if o == nil || IsNil(o.VariableName) {
		var ret string
		return ret
	}
	return *o.VariableName
}

// GetVariableNameOk returns a tuple with the VariableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetVariableNameOk() (*string, bool) {
	if o == nil || IsNil(o.VariableName) {
		return nil, false
	}
	return o.VariableName, true
}

// HasVariableName returns a boolean if a field has been set.
func (o *Property) HasVariableName() bool {
	if o != nil && !IsNil(o.VariableName) {
		return true
	}

	return false
}

// SetVariableName gets a reference to the given string and assigns it to the VariableName field.
func (o *Property) SetVariableName(v string) {
	o.VariableName = &v
}

func (o Property) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Property) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxValue) {
		toSerialize["max_value"] = o.MaxValue
	}
	if !IsNil(o.MinValue) {
		toSerialize["min_value"] = o.MinValue
	}
	toSerialize["name"] = o.Name
	toSerialize["permission"] = o.Permission
	if !IsNil(o.Persist) {
		toSerialize["persist"] = o.Persist
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
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

type NullableProperty struct {
	value *Property
	isSet bool
}

func (v NullableProperty) Get() *Property {
	return v.value
}

func (v *NullableProperty) Set(val *Property) {
	v.value = val
	v.isSet = true
}

func (v NullableProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProperty(val *Property) *NullableProperty {
	return &NullableProperty{value: val, isSet: true}
}

func (v NullableProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


