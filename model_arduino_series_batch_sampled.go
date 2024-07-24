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

// checks if the ArduinoSeriesBatchSampled type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoSeriesBatchSampled{}

// ArduinoSeriesBatchSampled ArduinoSeriesBatchSampled media type (default view)
type ArduinoSeriesBatchSampled struct {
	// Response version
	RespVersion int64 `json:"resp_version"`
	// Responses of the request
	Responses []ArduinoSeriesSampledResponse `json:"responses"`
}

// NewArduinoSeriesBatchSampled instantiates a new ArduinoSeriesBatchSampled object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoSeriesBatchSampled(respVersion int64, responses []ArduinoSeriesSampledResponse) *ArduinoSeriesBatchSampled {
	this := ArduinoSeriesBatchSampled{}
	this.RespVersion = respVersion
	this.Responses = responses
	return &this
}

// NewArduinoSeriesBatchSampledWithDefaults instantiates a new ArduinoSeriesBatchSampled object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoSeriesBatchSampledWithDefaults() *ArduinoSeriesBatchSampled {
	this := ArduinoSeriesBatchSampled{}
	return &this
}

// GetRespVersion returns the RespVersion field value
func (o *ArduinoSeriesBatchSampled) GetRespVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RespVersion
}

// GetRespVersionOk returns a tuple with the RespVersion field value
// and a boolean to check if the value has been set.
func (o *ArduinoSeriesBatchSampled) GetRespVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RespVersion, true
}

// SetRespVersion sets field value
func (o *ArduinoSeriesBatchSampled) SetRespVersion(v int64) {
	o.RespVersion = v
}

// GetResponses returns the Responses field value
func (o *ArduinoSeriesBatchSampled) GetResponses() []ArduinoSeriesSampledResponse {
	if o == nil {
		var ret []ArduinoSeriesSampledResponse
		return ret
	}

	return o.Responses
}

// GetResponsesOk returns a tuple with the Responses field value
// and a boolean to check if the value has been set.
func (o *ArduinoSeriesBatchSampled) GetResponsesOk() ([]ArduinoSeriesSampledResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Responses, true
}

// SetResponses sets field value
func (o *ArduinoSeriesBatchSampled) SetResponses(v []ArduinoSeriesSampledResponse) {
	o.Responses = v
}

func (o ArduinoSeriesBatchSampled) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoSeriesBatchSampled) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resp_version"] = o.RespVersion
	toSerialize["responses"] = o.Responses
	return toSerialize, nil
}

type NullableArduinoSeriesBatchSampled struct {
	value *ArduinoSeriesBatchSampled
	isSet bool
}

func (v NullableArduinoSeriesBatchSampled) Get() *ArduinoSeriesBatchSampled {
	return v.value
}

func (v *NullableArduinoSeriesBatchSampled) Set(val *ArduinoSeriesBatchSampled) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoSeriesBatchSampled) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoSeriesBatchSampled) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoSeriesBatchSampled(val *ArduinoSeriesBatchSampled) *NullableArduinoSeriesBatchSampled {
	return &NullableArduinoSeriesBatchSampled{value: val, isSet: true}
}

func (v NullableArduinoSeriesBatchSampled) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoSeriesBatchSampled) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


