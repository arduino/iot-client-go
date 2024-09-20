/*
Arduino IoT Cloud API

 Provides a set of endpoints to manage Arduino IoT Cloud **Devices**, **Things**, **Properties** and **Timeseries**. This API can be called just with any HTTP Client, or using one of these clients:  * [Javascript NPM package](https://www.npmjs.com/package/@arduino/arduino-iot-client)  * [Python PYPI Package](https://pypi.org/project/arduino-iot-client/)  * [Golang Module](https://github.com/arduino/iot-client-go)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"time"
)

// checks if the ArduinoDevicev2propertyvaluesLastEvaluatedKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoDevicev2propertyvaluesLastEvaluatedKey{}

// ArduinoDevicev2propertyvaluesLastEvaluatedKey struct for ArduinoDevicev2propertyvaluesLastEvaluatedKey
type ArduinoDevicev2propertyvaluesLastEvaluatedKey struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewArduinoDevicev2propertyvaluesLastEvaluatedKey instantiates a new ArduinoDevicev2propertyvaluesLastEvaluatedKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoDevicev2propertyvaluesLastEvaluatedKey() *ArduinoDevicev2propertyvaluesLastEvaluatedKey {
	this := ArduinoDevicev2propertyvaluesLastEvaluatedKey{}
	return &this
}

// NewArduinoDevicev2propertyvaluesLastEvaluatedKeyWithDefaults instantiates a new ArduinoDevicev2propertyvaluesLastEvaluatedKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoDevicev2propertyvaluesLastEvaluatedKeyWithDefaults() *ArduinoDevicev2propertyvaluesLastEvaluatedKey {
	this := ArduinoDevicev2propertyvaluesLastEvaluatedKey{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ArduinoDevicev2propertyvaluesLastEvaluatedKey) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoDevicev2propertyvaluesLastEvaluatedKey) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ArduinoDevicev2propertyvaluesLastEvaluatedKey) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ArduinoDevicev2propertyvaluesLastEvaluatedKey) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ArduinoDevicev2propertyvaluesLastEvaluatedKey) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoDevicev2propertyvaluesLastEvaluatedKey) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ArduinoDevicev2propertyvaluesLastEvaluatedKey) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ArduinoDevicev2propertyvaluesLastEvaluatedKey) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ArduinoDevicev2propertyvaluesLastEvaluatedKey) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoDevicev2propertyvaluesLastEvaluatedKey) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ArduinoDevicev2propertyvaluesLastEvaluatedKey) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ArduinoDevicev2propertyvaluesLastEvaluatedKey) SetName(v string) {
	o.Name = &v
}

func (o ArduinoDevicev2propertyvaluesLastEvaluatedKey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoDevicev2propertyvaluesLastEvaluatedKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableArduinoDevicev2propertyvaluesLastEvaluatedKey struct {
	value *ArduinoDevicev2propertyvaluesLastEvaluatedKey
	isSet bool
}

func (v NullableArduinoDevicev2propertyvaluesLastEvaluatedKey) Get() *ArduinoDevicev2propertyvaluesLastEvaluatedKey {
	return v.value
}

func (v *NullableArduinoDevicev2propertyvaluesLastEvaluatedKey) Set(val *ArduinoDevicev2propertyvaluesLastEvaluatedKey) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoDevicev2propertyvaluesLastEvaluatedKey) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoDevicev2propertyvaluesLastEvaluatedKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoDevicev2propertyvaluesLastEvaluatedKey(val *ArduinoDevicev2propertyvaluesLastEvaluatedKey) *NullableArduinoDevicev2propertyvaluesLastEvaluatedKey {
	return &NullableArduinoDevicev2propertyvaluesLastEvaluatedKey{value: val, isSet: true}
}

func (v NullableArduinoDevicev2propertyvaluesLastEvaluatedKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoDevicev2propertyvaluesLastEvaluatedKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


