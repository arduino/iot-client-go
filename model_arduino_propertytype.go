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

// checks if the ArduinoPropertytype type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoPropertytype{}

// ArduinoPropertytype ArduinoPropertytype media type (default view)
type ArduinoPropertytype struct {
	// The voice assistants available for this type
	Assistants []string `json:"assistants,omitempty"`
	// The c++ type we are using for this variable type
	Declaration string `json:"declaration"`
	// Tell if this type is deprecated
	Deprecated bool `json:"deprecated"`
	// Example of use
	Example *string `json:"example,omitempty"`
	// The friendly name of the property type
	Name string `json:"name"`
	// Tell if the type allow a R/W permission
	Rw bool `json:"rw"`
	// The type of property to use if it's deprecated
	SupersededBy *string `json:"supersededBy,omitempty"`
	// The tags related to the type
	Tags []string `json:"tags,omitempty"`
	// The api reference of this type
	Type string `json:"type"`
	// The measure units available for this type
	Units []string `json:"units,omitempty"`
}

// NewArduinoPropertytype instantiates a new ArduinoPropertytype object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoPropertytype(declaration string, deprecated bool, name string, rw bool, type_ string) *ArduinoPropertytype {
	this := ArduinoPropertytype{}
	this.Declaration = declaration
	this.Deprecated = deprecated
	this.Name = name
	this.Rw = rw
	this.Type = type_
	return &this
}

// NewArduinoPropertytypeWithDefaults instantiates a new ArduinoPropertytype object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoPropertytypeWithDefaults() *ArduinoPropertytype {
	this := ArduinoPropertytype{}
	return &this
}

// GetAssistants returns the Assistants field value if set, zero value otherwise.
func (o *ArduinoPropertytype) GetAssistants() []string {
	if o == nil || IsNil(o.Assistants) {
		var ret []string
		return ret
	}
	return o.Assistants
}

// GetAssistantsOk returns a tuple with the Assistants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoPropertytype) GetAssistantsOk() ([]string, bool) {
	if o == nil || IsNil(o.Assistants) {
		return nil, false
	}
	return o.Assistants, true
}

// HasAssistants returns a boolean if a field has been set.
func (o *ArduinoPropertytype) HasAssistants() bool {
	if o != nil && !IsNil(o.Assistants) {
		return true
	}

	return false
}

// SetAssistants gets a reference to the given []string and assigns it to the Assistants field.
func (o *ArduinoPropertytype) SetAssistants(v []string) {
	o.Assistants = v
}

// GetDeclaration returns the Declaration field value
func (o *ArduinoPropertytype) GetDeclaration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Declaration
}

// GetDeclarationOk returns a tuple with the Declaration field value
// and a boolean to check if the value has been set.
func (o *ArduinoPropertytype) GetDeclarationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Declaration, true
}

// SetDeclaration sets field value
func (o *ArduinoPropertytype) SetDeclaration(v string) {
	o.Declaration = v
}

// GetDeprecated returns the Deprecated field value
func (o *ArduinoPropertytype) GetDeprecated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value
// and a boolean to check if the value has been set.
func (o *ArduinoPropertytype) GetDeprecatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deprecated, true
}

// SetDeprecated sets field value
func (o *ArduinoPropertytype) SetDeprecated(v bool) {
	o.Deprecated = v
}

// GetExample returns the Example field value if set, zero value otherwise.
func (o *ArduinoPropertytype) GetExample() string {
	if o == nil || IsNil(o.Example) {
		var ret string
		return ret
	}
	return *o.Example
}

// GetExampleOk returns a tuple with the Example field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoPropertytype) GetExampleOk() (*string, bool) {
	if o == nil || IsNil(o.Example) {
		return nil, false
	}
	return o.Example, true
}

// HasExample returns a boolean if a field has been set.
func (o *ArduinoPropertytype) HasExample() bool {
	if o != nil && !IsNil(o.Example) {
		return true
	}

	return false
}

// SetExample gets a reference to the given string and assigns it to the Example field.
func (o *ArduinoPropertytype) SetExample(v string) {
	o.Example = &v
}

// GetName returns the Name field value
func (o *ArduinoPropertytype) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ArduinoPropertytype) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ArduinoPropertytype) SetName(v string) {
	o.Name = v
}

// GetRw returns the Rw field value
func (o *ArduinoPropertytype) GetRw() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Rw
}

// GetRwOk returns a tuple with the Rw field value
// and a boolean to check if the value has been set.
func (o *ArduinoPropertytype) GetRwOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rw, true
}

// SetRw sets field value
func (o *ArduinoPropertytype) SetRw(v bool) {
	o.Rw = v
}

// GetSupersededBy returns the SupersededBy field value if set, zero value otherwise.
func (o *ArduinoPropertytype) GetSupersededBy() string {
	if o == nil || IsNil(o.SupersededBy) {
		var ret string
		return ret
	}
	return *o.SupersededBy
}

// GetSupersededByOk returns a tuple with the SupersededBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoPropertytype) GetSupersededByOk() (*string, bool) {
	if o == nil || IsNil(o.SupersededBy) {
		return nil, false
	}
	return o.SupersededBy, true
}

// HasSupersededBy returns a boolean if a field has been set.
func (o *ArduinoPropertytype) HasSupersededBy() bool {
	if o != nil && !IsNil(o.SupersededBy) {
		return true
	}

	return false
}

// SetSupersededBy gets a reference to the given string and assigns it to the SupersededBy field.
func (o *ArduinoPropertytype) SetSupersededBy(v string) {
	o.SupersededBy = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ArduinoPropertytype) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoPropertytype) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ArduinoPropertytype) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ArduinoPropertytype) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value
func (o *ArduinoPropertytype) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ArduinoPropertytype) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ArduinoPropertytype) SetType(v string) {
	o.Type = v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *ArduinoPropertytype) GetUnits() []string {
	if o == nil || IsNil(o.Units) {
		var ret []string
		return ret
	}
	return o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoPropertytype) GetUnitsOk() ([]string, bool) {
	if o == nil || IsNil(o.Units) {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *ArduinoPropertytype) HasUnits() bool {
	if o != nil && !IsNil(o.Units) {
		return true
	}

	return false
}

// SetUnits gets a reference to the given []string and assigns it to the Units field.
func (o *ArduinoPropertytype) SetUnits(v []string) {
	o.Units = v
}

func (o ArduinoPropertytype) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoPropertytype) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assistants) {
		toSerialize["assistants"] = o.Assistants
	}
	toSerialize["declaration"] = o.Declaration
	toSerialize["deprecated"] = o.Deprecated
	if !IsNil(o.Example) {
		toSerialize["example"] = o.Example
	}
	toSerialize["name"] = o.Name
	toSerialize["rw"] = o.Rw
	if !IsNil(o.SupersededBy) {
		toSerialize["supersededBy"] = o.SupersededBy
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Units) {
		toSerialize["units"] = o.Units
	}
	return toSerialize, nil
}

type NullableArduinoPropertytype struct {
	value *ArduinoPropertytype
	isSet bool
}

func (v NullableArduinoPropertytype) Get() *ArduinoPropertytype {
	return v.value
}

func (v *NullableArduinoPropertytype) Set(val *ArduinoPropertytype) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoPropertytype) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoPropertytype) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoPropertytype(val *ArduinoPropertytype) *NullableArduinoPropertytype {
	return &NullableArduinoPropertytype{value: val, isSet: true}
}

func (v NullableArduinoPropertytype) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoPropertytype) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


