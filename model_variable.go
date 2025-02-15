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

// checks if the Variable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Variable{}

// Variable struct for Variable
type Variable struct {
	// The template expression that extracts the value from the respective entity
	Attribute string `json:"attribute"`
	// Type of the entity being referenced
	Entity string `json:"entity"`
	// The ID of the referenced entity
	Id *string `json:"id,omitempty"`
	// Name of the variable as referenced by the expression
	Placeholder string `json:"placeholder"`
	// The ID of the property referenced entity
	PropertyId *string `json:"property_id,omitempty"`
	// The ID of the thing referenced entity
	ThingId *string `json:"thing_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Variable Variable

// NewVariable instantiates a new Variable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariable(attribute string, entity string, placeholder string) *Variable {
	this := Variable{}
	this.Attribute = attribute
	this.Entity = entity
	this.Placeholder = placeholder
	return &this
}

// NewVariableWithDefaults instantiates a new Variable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariableWithDefaults() *Variable {
	this := Variable{}
	return &this
}

// GetAttribute returns the Attribute field value
func (o *Variable) GetAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value
// and a boolean to check if the value has been set.
func (o *Variable) GetAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attribute, true
}

// SetAttribute sets field value
func (o *Variable) SetAttribute(v string) {
	o.Attribute = v
}

// GetEntity returns the Entity field value
func (o *Variable) GetEntity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Entity
}

// GetEntityOk returns a tuple with the Entity field value
// and a boolean to check if the value has been set.
func (o *Variable) GetEntityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entity, true
}

// SetEntity sets field value
func (o *Variable) SetEntity(v string) {
	o.Entity = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Variable) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variable) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Variable) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Variable) SetId(v string) {
	o.Id = &v
}

// GetPlaceholder returns the Placeholder field value
func (o *Variable) GetPlaceholder() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Placeholder
}

// GetPlaceholderOk returns a tuple with the Placeholder field value
// and a boolean to check if the value has been set.
func (o *Variable) GetPlaceholderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Placeholder, true
}

// SetPlaceholder sets field value
func (o *Variable) SetPlaceholder(v string) {
	o.Placeholder = v
}

// GetPropertyId returns the PropertyId field value if set, zero value otherwise.
func (o *Variable) GetPropertyId() string {
	if o == nil || IsNil(o.PropertyId) {
		var ret string
		return ret
	}
	return *o.PropertyId
}

// GetPropertyIdOk returns a tuple with the PropertyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variable) GetPropertyIdOk() (*string, bool) {
	if o == nil || IsNil(o.PropertyId) {
		return nil, false
	}
	return o.PropertyId, true
}

// HasPropertyId returns a boolean if a field has been set.
func (o *Variable) HasPropertyId() bool {
	if o != nil && !IsNil(o.PropertyId) {
		return true
	}

	return false
}

// SetPropertyId gets a reference to the given string and assigns it to the PropertyId field.
func (o *Variable) SetPropertyId(v string) {
	o.PropertyId = &v
}

// GetThingId returns the ThingId field value if set, zero value otherwise.
func (o *Variable) GetThingId() string {
	if o == nil || IsNil(o.ThingId) {
		var ret string
		return ret
	}
	return *o.ThingId
}

// GetThingIdOk returns a tuple with the ThingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variable) GetThingIdOk() (*string, bool) {
	if o == nil || IsNil(o.ThingId) {
		return nil, false
	}
	return o.ThingId, true
}

// HasThingId returns a boolean if a field has been set.
func (o *Variable) HasThingId() bool {
	if o != nil && !IsNil(o.ThingId) {
		return true
	}

	return false
}

// SetThingId gets a reference to the given string and assigns it to the ThingId field.
func (o *Variable) SetThingId(v string) {
	o.ThingId = &v
}

func (o Variable) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Variable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attribute"] = o.Attribute
	toSerialize["entity"] = o.Entity
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["placeholder"] = o.Placeholder
	if !IsNil(o.PropertyId) {
		toSerialize["property_id"] = o.PropertyId
	}
	if !IsNil(o.ThingId) {
		toSerialize["thing_id"] = o.ThingId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Variable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"attribute",
		"entity",
		"placeholder",
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

	varVariable := _Variable{}

	err = json.Unmarshal(data, &varVariable)

	if err != nil {
		return err
	}

	*o = Variable(varVariable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "attribute")
		delete(additionalProperties, "entity")
		delete(additionalProperties, "id")
		delete(additionalProperties, "placeholder")
		delete(additionalProperties, "property_id")
		delete(additionalProperties, "thing_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVariable struct {
	value *Variable
	isSet bool
}

func (v NullableVariable) Get() *Variable {
	return v.value
}

func (v *NullableVariable) Set(val *Variable) {
	v.value = val
	v.isSet = true
}

func (v NullableVariable) IsSet() bool {
	return v.isSet
}

func (v *NullableVariable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariable(val *Variable) *NullableVariable {
	return &NullableVariable{value: val, isSet: true}
}

func (v NullableVariable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


