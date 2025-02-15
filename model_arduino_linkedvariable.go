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
	"fmt"
)

// checks if the ArduinoLinkedvariable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoLinkedvariable{}

// ArduinoLinkedvariable ArduinoLinkedvariable media type (default view)
type ArduinoLinkedvariable struct {
	// The id of the linked variable
	Id string `json:"id"`
	// Last value of the linked property
	LastValue interface{} `json:"last_value,omitempty"`
	// Update date of the last value
	LastValueUpdatedAt *time.Time `json:"last_value_updated_at,omitempty"`
	// The name of the variable
	Name string `json:"name"`
	// The permission of the linked variable
	Permission string `json:"permission"`
	// The id of the related thing
	ThingId string `json:"thing_id"`
	// The name of the related thing
	ThingName string `json:"thing_name"`
	ThingTimezone *ArduinoTimezone `json:"thing_timezone,omitempty"`
	// The type of the variable
	Type string `json:"type"`
	// The name of the variable in the code
	VariableName string `json:"variable_name"`
	AdditionalProperties map[string]interface{}
}

type _ArduinoLinkedvariable ArduinoLinkedvariable

// NewArduinoLinkedvariable instantiates a new ArduinoLinkedvariable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoLinkedvariable(id string, name string, permission string, thingId string, thingName string, type_ string, variableName string) *ArduinoLinkedvariable {
	this := ArduinoLinkedvariable{}
	this.Id = id
	this.Name = name
	this.Permission = permission
	this.ThingId = thingId
	this.ThingName = thingName
	this.Type = type_
	this.VariableName = variableName
	return &this
}

// NewArduinoLinkedvariableWithDefaults instantiates a new ArduinoLinkedvariable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoLinkedvariableWithDefaults() *ArduinoLinkedvariable {
	this := ArduinoLinkedvariable{}
	return &this
}

// GetId returns the Id field value
func (o *ArduinoLinkedvariable) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ArduinoLinkedvariable) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ArduinoLinkedvariable) SetId(v string) {
	o.Id = v
}

// GetLastValue returns the LastValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArduinoLinkedvariable) GetLastValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.LastValue
}

// GetLastValueOk returns a tuple with the LastValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArduinoLinkedvariable) GetLastValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.LastValue) {
		return nil, false
	}
	return &o.LastValue, true
}

// HasLastValue returns a boolean if a field has been set.
func (o *ArduinoLinkedvariable) HasLastValue() bool {
	if o != nil && !IsNil(o.LastValue) {
		return true
	}

	return false
}

// SetLastValue gets a reference to the given interface{} and assigns it to the LastValue field.
func (o *ArduinoLinkedvariable) SetLastValue(v interface{}) {
	o.LastValue = v
}

// GetLastValueUpdatedAt returns the LastValueUpdatedAt field value if set, zero value otherwise.
func (o *ArduinoLinkedvariable) GetLastValueUpdatedAt() time.Time {
	if o == nil || IsNil(o.LastValueUpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.LastValueUpdatedAt
}

// GetLastValueUpdatedAtOk returns a tuple with the LastValueUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoLinkedvariable) GetLastValueUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastValueUpdatedAt) {
		return nil, false
	}
	return o.LastValueUpdatedAt, true
}

// HasLastValueUpdatedAt returns a boolean if a field has been set.
func (o *ArduinoLinkedvariable) HasLastValueUpdatedAt() bool {
	if o != nil && !IsNil(o.LastValueUpdatedAt) {
		return true
	}

	return false
}

// SetLastValueUpdatedAt gets a reference to the given time.Time and assigns it to the LastValueUpdatedAt field.
func (o *ArduinoLinkedvariable) SetLastValueUpdatedAt(v time.Time) {
	o.LastValueUpdatedAt = &v
}

// GetName returns the Name field value
func (o *ArduinoLinkedvariable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ArduinoLinkedvariable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ArduinoLinkedvariable) SetName(v string) {
	o.Name = v
}

// GetPermission returns the Permission field value
func (o *ArduinoLinkedvariable) GetPermission() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value
// and a boolean to check if the value has been set.
func (o *ArduinoLinkedvariable) GetPermissionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permission, true
}

// SetPermission sets field value
func (o *ArduinoLinkedvariable) SetPermission(v string) {
	o.Permission = v
}

// GetThingId returns the ThingId field value
func (o *ArduinoLinkedvariable) GetThingId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThingId
}

// GetThingIdOk returns a tuple with the ThingId field value
// and a boolean to check if the value has been set.
func (o *ArduinoLinkedvariable) GetThingIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThingId, true
}

// SetThingId sets field value
func (o *ArduinoLinkedvariable) SetThingId(v string) {
	o.ThingId = v
}

// GetThingName returns the ThingName field value
func (o *ArduinoLinkedvariable) GetThingName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThingName
}

// GetThingNameOk returns a tuple with the ThingName field value
// and a boolean to check if the value has been set.
func (o *ArduinoLinkedvariable) GetThingNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThingName, true
}

// SetThingName sets field value
func (o *ArduinoLinkedvariable) SetThingName(v string) {
	o.ThingName = v
}

// GetThingTimezone returns the ThingTimezone field value if set, zero value otherwise.
func (o *ArduinoLinkedvariable) GetThingTimezone() ArduinoTimezone {
	if o == nil || IsNil(o.ThingTimezone) {
		var ret ArduinoTimezone
		return ret
	}
	return *o.ThingTimezone
}

// GetThingTimezoneOk returns a tuple with the ThingTimezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoLinkedvariable) GetThingTimezoneOk() (*ArduinoTimezone, bool) {
	if o == nil || IsNil(o.ThingTimezone) {
		return nil, false
	}
	return o.ThingTimezone, true
}

// HasThingTimezone returns a boolean if a field has been set.
func (o *ArduinoLinkedvariable) HasThingTimezone() bool {
	if o != nil && !IsNil(o.ThingTimezone) {
		return true
	}

	return false
}

// SetThingTimezone gets a reference to the given ArduinoTimezone and assigns it to the ThingTimezone field.
func (o *ArduinoLinkedvariable) SetThingTimezone(v ArduinoTimezone) {
	o.ThingTimezone = &v
}

// GetType returns the Type field value
func (o *ArduinoLinkedvariable) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ArduinoLinkedvariable) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ArduinoLinkedvariable) SetType(v string) {
	o.Type = v
}

// GetVariableName returns the VariableName field value
func (o *ArduinoLinkedvariable) GetVariableName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VariableName
}

// GetVariableNameOk returns a tuple with the VariableName field value
// and a boolean to check if the value has been set.
func (o *ArduinoLinkedvariable) GetVariableNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VariableName, true
}

// SetVariableName sets field value
func (o *ArduinoLinkedvariable) SetVariableName(v string) {
	o.VariableName = v
}

func (o ArduinoLinkedvariable) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoLinkedvariable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.LastValue != nil {
		toSerialize["last_value"] = o.LastValue
	}
	if !IsNil(o.LastValueUpdatedAt) {
		toSerialize["last_value_updated_at"] = o.LastValueUpdatedAt
	}
	toSerialize["name"] = o.Name
	toSerialize["permission"] = o.Permission
	toSerialize["thing_id"] = o.ThingId
	toSerialize["thing_name"] = o.ThingName
	if !IsNil(o.ThingTimezone) {
		toSerialize["thing_timezone"] = o.ThingTimezone
	}
	toSerialize["type"] = o.Type
	toSerialize["variable_name"] = o.VariableName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ArduinoLinkedvariable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"permission",
		"thing_id",
		"thing_name",
		"type",
		"variable_name",
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

	varArduinoLinkedvariable := _ArduinoLinkedvariable{}

	err = json.Unmarshal(data, &varArduinoLinkedvariable)

	if err != nil {
		return err
	}

	*o = ArduinoLinkedvariable(varArduinoLinkedvariable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "last_value")
		delete(additionalProperties, "last_value_updated_at")
		delete(additionalProperties, "name")
		delete(additionalProperties, "permission")
		delete(additionalProperties, "thing_id")
		delete(additionalProperties, "thing_name")
		delete(additionalProperties, "thing_timezone")
		delete(additionalProperties, "type")
		delete(additionalProperties, "variable_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableArduinoLinkedvariable struct {
	value *ArduinoLinkedvariable
	isSet bool
}

func (v NullableArduinoLinkedvariable) Get() *ArduinoLinkedvariable {
	return v.value
}

func (v *NullableArduinoLinkedvariable) Set(val *ArduinoLinkedvariable) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoLinkedvariable) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoLinkedvariable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoLinkedvariable(val *ArduinoLinkedvariable) *NullableArduinoLinkedvariable {
	return &NullableArduinoLinkedvariable{value: val, isSet: true}
}

func (v NullableArduinoLinkedvariable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoLinkedvariable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


