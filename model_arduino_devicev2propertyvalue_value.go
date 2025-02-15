/*
Arduino IoT Cloud API

Provides a set of endpoints to manage Arduino IoT Cloud **Devices**, **Things**, **Properties** and **Timeseries**. This API can be called just with any HTTP Client, or using one of these clients:  * [Javascript NPM package](https://www.npmjs.com/package/@arduino/arduino-iot-client)  * [Python PYPI Package](https://pypi.org/project/arduino-iot-client/)  * [Golang Module](https://github.com/arduino/iot-client-go)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
)

// checks if the ArduinoDevicev2propertyvalueValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoDevicev2propertyvalueValue{}

// ArduinoDevicev2propertyvalueValue struct for ArduinoDevicev2propertyvalueValue
type ArduinoDevicev2propertyvalueValue struct {
	Payload *string `json:"payload,omitempty"`
	Seqno *int64 `json:"seqno,omitempty"`
	Statistics *ArduinoDevicev2propertyvalueValueStatistics `json:"statistics,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ArduinoDevicev2propertyvalueValue ArduinoDevicev2propertyvalueValue

// NewArduinoDevicev2propertyvalueValue instantiates a new ArduinoDevicev2propertyvalueValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoDevicev2propertyvalueValue() *ArduinoDevicev2propertyvalueValue {
	this := ArduinoDevicev2propertyvalueValue{}
	return &this
}

// NewArduinoDevicev2propertyvalueValueWithDefaults instantiates a new ArduinoDevicev2propertyvalueValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoDevicev2propertyvalueValueWithDefaults() *ArduinoDevicev2propertyvalueValue {
	this := ArduinoDevicev2propertyvalueValue{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *ArduinoDevicev2propertyvalueValue) GetPayload() string {
	if o == nil || IsNil(o.Payload) {
		var ret string
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoDevicev2propertyvalueValue) GetPayloadOk() (*string, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *ArduinoDevicev2propertyvalueValue) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given string and assigns it to the Payload field.
func (o *ArduinoDevicev2propertyvalueValue) SetPayload(v string) {
	o.Payload = &v
}

// GetSeqno returns the Seqno field value if set, zero value otherwise.
func (o *ArduinoDevicev2propertyvalueValue) GetSeqno() int64 {
	if o == nil || IsNil(o.Seqno) {
		var ret int64
		return ret
	}
	return *o.Seqno
}

// GetSeqnoOk returns a tuple with the Seqno field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoDevicev2propertyvalueValue) GetSeqnoOk() (*int64, bool) {
	if o == nil || IsNil(o.Seqno) {
		return nil, false
	}
	return o.Seqno, true
}

// HasSeqno returns a boolean if a field has been set.
func (o *ArduinoDevicev2propertyvalueValue) HasSeqno() bool {
	if o != nil && !IsNil(o.Seqno) {
		return true
	}

	return false
}

// SetSeqno gets a reference to the given int64 and assigns it to the Seqno field.
func (o *ArduinoDevicev2propertyvalueValue) SetSeqno(v int64) {
	o.Seqno = &v
}

// GetStatistics returns the Statistics field value if set, zero value otherwise.
func (o *ArduinoDevicev2propertyvalueValue) GetStatistics() ArduinoDevicev2propertyvalueValueStatistics {
	if o == nil || IsNil(o.Statistics) {
		var ret ArduinoDevicev2propertyvalueValueStatistics
		return ret
	}
	return *o.Statistics
}

// GetStatisticsOk returns a tuple with the Statistics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoDevicev2propertyvalueValue) GetStatisticsOk() (*ArduinoDevicev2propertyvalueValueStatistics, bool) {
	if o == nil || IsNil(o.Statistics) {
		return nil, false
	}
	return o.Statistics, true
}

// HasStatistics returns a boolean if a field has been set.
func (o *ArduinoDevicev2propertyvalueValue) HasStatistics() bool {
	if o != nil && !IsNil(o.Statistics) {
		return true
	}

	return false
}

// SetStatistics gets a reference to the given ArduinoDevicev2propertyvalueValueStatistics and assigns it to the Statistics field.
func (o *ArduinoDevicev2propertyvalueValue) SetStatistics(v ArduinoDevicev2propertyvalueValueStatistics) {
	o.Statistics = &v
}

func (o ArduinoDevicev2propertyvalueValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoDevicev2propertyvalueValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Seqno) {
		toSerialize["seqno"] = o.Seqno
	}
	if !IsNil(o.Statistics) {
		toSerialize["statistics"] = o.Statistics
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ArduinoDevicev2propertyvalueValue) UnmarshalJSON(data []byte) (err error) {
	varArduinoDevicev2propertyvalueValue := _ArduinoDevicev2propertyvalueValue{}

	err = json.Unmarshal(data, &varArduinoDevicev2propertyvalueValue)

	if err != nil {
		return err
	}

	*o = ArduinoDevicev2propertyvalueValue(varArduinoDevicev2propertyvalueValue)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "payload")
		delete(additionalProperties, "seqno")
		delete(additionalProperties, "statistics")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableArduinoDevicev2propertyvalueValue struct {
	value *ArduinoDevicev2propertyvalueValue
	isSet bool
}

func (v NullableArduinoDevicev2propertyvalueValue) Get() *ArduinoDevicev2propertyvalueValue {
	return v.value
}

func (v *NullableArduinoDevicev2propertyvalueValue) Set(val *ArduinoDevicev2propertyvalueValue) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoDevicev2propertyvalueValue) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoDevicev2propertyvalueValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoDevicev2propertyvalueValue(val *ArduinoDevicev2propertyvalueValue) *NullableArduinoDevicev2propertyvalueValue {
	return &NullableArduinoDevicev2propertyvalueValue{value: val, isSet: true}
}

func (v NullableArduinoDevicev2propertyvalueValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoDevicev2propertyvalueValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


