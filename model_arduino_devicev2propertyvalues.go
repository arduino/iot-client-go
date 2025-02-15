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

// checks if the ArduinoDevicev2propertyvalues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoDevicev2propertyvalues{}

// ArduinoDevicev2propertyvalues ArduinoDevicev2propertyvalues media type (default view)
type ArduinoDevicev2propertyvalues struct {
	Id string `json:"id"`
	LastEvaluatedKey ArduinoDevicev2propertyvaluesLastEvaluatedKey `json:"last_evaluated_key"`
	Name string `json:"name"`
	// ArduinoDevicev2propertyvalueCollection is the media type for an array of ArduinoDevicev2propertyvalue (default view)
	Values []ArduinoDevicev2propertyvalue `json:"values"`
	AdditionalProperties map[string]interface{}
}

type _ArduinoDevicev2propertyvalues ArduinoDevicev2propertyvalues

// NewArduinoDevicev2propertyvalues instantiates a new ArduinoDevicev2propertyvalues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoDevicev2propertyvalues(id string, lastEvaluatedKey ArduinoDevicev2propertyvaluesLastEvaluatedKey, name string, values []ArduinoDevicev2propertyvalue) *ArduinoDevicev2propertyvalues {
	this := ArduinoDevicev2propertyvalues{}
	this.Id = id
	this.LastEvaluatedKey = lastEvaluatedKey
	this.Name = name
	this.Values = values
	return &this
}

// NewArduinoDevicev2propertyvaluesWithDefaults instantiates a new ArduinoDevicev2propertyvalues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoDevicev2propertyvaluesWithDefaults() *ArduinoDevicev2propertyvalues {
	this := ArduinoDevicev2propertyvalues{}
	return &this
}

// GetId returns the Id field value
func (o *ArduinoDevicev2propertyvalues) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ArduinoDevicev2propertyvalues) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ArduinoDevicev2propertyvalues) SetId(v string) {
	o.Id = v
}

// GetLastEvaluatedKey returns the LastEvaluatedKey field value
func (o *ArduinoDevicev2propertyvalues) GetLastEvaluatedKey() ArduinoDevicev2propertyvaluesLastEvaluatedKey {
	if o == nil {
		var ret ArduinoDevicev2propertyvaluesLastEvaluatedKey
		return ret
	}

	return o.LastEvaluatedKey
}

// GetLastEvaluatedKeyOk returns a tuple with the LastEvaluatedKey field value
// and a boolean to check if the value has been set.
func (o *ArduinoDevicev2propertyvalues) GetLastEvaluatedKeyOk() (*ArduinoDevicev2propertyvaluesLastEvaluatedKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastEvaluatedKey, true
}

// SetLastEvaluatedKey sets field value
func (o *ArduinoDevicev2propertyvalues) SetLastEvaluatedKey(v ArduinoDevicev2propertyvaluesLastEvaluatedKey) {
	o.LastEvaluatedKey = v
}

// GetName returns the Name field value
func (o *ArduinoDevicev2propertyvalues) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ArduinoDevicev2propertyvalues) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ArduinoDevicev2propertyvalues) SetName(v string) {
	o.Name = v
}

// GetValues returns the Values field value
func (o *ArduinoDevicev2propertyvalues) GetValues() []ArduinoDevicev2propertyvalue {
	if o == nil {
		var ret []ArduinoDevicev2propertyvalue
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *ArduinoDevicev2propertyvalues) GetValuesOk() ([]ArduinoDevicev2propertyvalue, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *ArduinoDevicev2propertyvalues) SetValues(v []ArduinoDevicev2propertyvalue) {
	o.Values = v
}

func (o ArduinoDevicev2propertyvalues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoDevicev2propertyvalues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["last_evaluated_key"] = o.LastEvaluatedKey
	toSerialize["name"] = o.Name
	toSerialize["values"] = o.Values

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ArduinoDevicev2propertyvalues) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"last_evaluated_key",
		"name",
		"values",
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

	varArduinoDevicev2propertyvalues := _ArduinoDevicev2propertyvalues{}

	err = json.Unmarshal(data, &varArduinoDevicev2propertyvalues)

	if err != nil {
		return err
	}

	*o = ArduinoDevicev2propertyvalues(varArduinoDevicev2propertyvalues)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "last_evaluated_key")
		delete(additionalProperties, "name")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableArduinoDevicev2propertyvalues struct {
	value *ArduinoDevicev2propertyvalues
	isSet bool
}

func (v NullableArduinoDevicev2propertyvalues) Get() *ArduinoDevicev2propertyvalues {
	return v.value
}

func (v *NullableArduinoDevicev2propertyvalues) Set(val *ArduinoDevicev2propertyvalues) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoDevicev2propertyvalues) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoDevicev2propertyvalues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoDevicev2propertyvalues(val *ArduinoDevicev2propertyvalues) *NullableArduinoDevicev2propertyvalues {
	return &NullableArduinoDevicev2propertyvalues{value: val, isSet: true}
}

func (v NullableArduinoDevicev2propertyvalues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoDevicev2propertyvalues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


