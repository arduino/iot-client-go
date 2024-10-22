/*
Arduino IoT Cloud API

Provides a set of endpoints to manage Arduino IoT Cloud **Devices**, **Things**, **Properties** and **Timeseries**. This API can be called just with any HTTP Client, or using one of these clients:  * [Javascript NPM package](https://www.npmjs.com/package/@arduino/arduino-iot-client)  * [Python PYPI Package](https://pypi.org/project/arduino-iot-client/)  * [Golang Module](https://github.com/arduino/iot-client-go)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the BatchQueryRawLastValueRequestMediaV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchQueryRawLastValueRequestMediaV1{}

// BatchQueryRawLastValueRequestMediaV1 struct for BatchQueryRawLastValueRequestMediaV1
type BatchQueryRawLastValueRequestMediaV1 struct {
	// Property id
	PropertyId string `json:"property_id"`
	// Thing id
	ThingId string `json:"thing_id"`
}

type _BatchQueryRawLastValueRequestMediaV1 BatchQueryRawLastValueRequestMediaV1

// NewBatchQueryRawLastValueRequestMediaV1 instantiates a new BatchQueryRawLastValueRequestMediaV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchQueryRawLastValueRequestMediaV1(propertyId string, thingId string) *BatchQueryRawLastValueRequestMediaV1 {
	this := BatchQueryRawLastValueRequestMediaV1{}
	this.PropertyId = propertyId
	this.ThingId = thingId
	return &this
}

// NewBatchQueryRawLastValueRequestMediaV1WithDefaults instantiates a new BatchQueryRawLastValueRequestMediaV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchQueryRawLastValueRequestMediaV1WithDefaults() *BatchQueryRawLastValueRequestMediaV1 {
	this := BatchQueryRawLastValueRequestMediaV1{}
	return &this
}

// GetPropertyId returns the PropertyId field value
func (o *BatchQueryRawLastValueRequestMediaV1) GetPropertyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PropertyId
}

// GetPropertyIdOk returns a tuple with the PropertyId field value
// and a boolean to check if the value has been set.
func (o *BatchQueryRawLastValueRequestMediaV1) GetPropertyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PropertyId, true
}

// SetPropertyId sets field value
func (o *BatchQueryRawLastValueRequestMediaV1) SetPropertyId(v string) {
	o.PropertyId = v
}

// GetThingId returns the ThingId field value
func (o *BatchQueryRawLastValueRequestMediaV1) GetThingId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThingId
}

// GetThingIdOk returns a tuple with the ThingId field value
// and a boolean to check if the value has been set.
func (o *BatchQueryRawLastValueRequestMediaV1) GetThingIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThingId, true
}

// SetThingId sets field value
func (o *BatchQueryRawLastValueRequestMediaV1) SetThingId(v string) {
	o.ThingId = v
}

func (o BatchQueryRawLastValueRequestMediaV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchQueryRawLastValueRequestMediaV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["property_id"] = o.PropertyId
	toSerialize["thing_id"] = o.ThingId
	return toSerialize, nil
}

func (o *BatchQueryRawLastValueRequestMediaV1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"property_id",
		"thing_id",
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

	varBatchQueryRawLastValueRequestMediaV1 := _BatchQueryRawLastValueRequestMediaV1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBatchQueryRawLastValueRequestMediaV1)

	if err != nil {
		return err
	}

	*o = BatchQueryRawLastValueRequestMediaV1(varBatchQueryRawLastValueRequestMediaV1)

	return err
}

type NullableBatchQueryRawLastValueRequestMediaV1 struct {
	value *BatchQueryRawLastValueRequestMediaV1
	isSet bool
}

func (v NullableBatchQueryRawLastValueRequestMediaV1) Get() *BatchQueryRawLastValueRequestMediaV1 {
	return v.value
}

func (v *NullableBatchQueryRawLastValueRequestMediaV1) Set(val *BatchQueryRawLastValueRequestMediaV1) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchQueryRawLastValueRequestMediaV1) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchQueryRawLastValueRequestMediaV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchQueryRawLastValueRequestMediaV1(val *BatchQueryRawLastValueRequestMediaV1) *NullableBatchQueryRawLastValueRequestMediaV1 {
	return &NullableBatchQueryRawLastValueRequestMediaV1{value: val, isSet: true}
}

func (v NullableBatchQueryRawLastValueRequestMediaV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchQueryRawLastValueRequestMediaV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


