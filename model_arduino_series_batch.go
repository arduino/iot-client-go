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

// checks if the ArduinoSeriesBatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoSeriesBatch{}

// ArduinoSeriesBatch ArduinoSeriesBatch media type (default view)
type ArduinoSeriesBatch struct {
	// Response version
	RespVersion int64 `json:"resp_version"`
	// Responses of the request
	Responses []ArduinoSeriesResponse `json:"responses"`
	AdditionalProperties map[string]interface{}
}

type _ArduinoSeriesBatch ArduinoSeriesBatch

// NewArduinoSeriesBatch instantiates a new ArduinoSeriesBatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoSeriesBatch(respVersion int64, responses []ArduinoSeriesResponse) *ArduinoSeriesBatch {
	this := ArduinoSeriesBatch{}
	this.RespVersion = respVersion
	this.Responses = responses
	return &this
}

// NewArduinoSeriesBatchWithDefaults instantiates a new ArduinoSeriesBatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoSeriesBatchWithDefaults() *ArduinoSeriesBatch {
	this := ArduinoSeriesBatch{}
	return &this
}

// GetRespVersion returns the RespVersion field value
func (o *ArduinoSeriesBatch) GetRespVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RespVersion
}

// GetRespVersionOk returns a tuple with the RespVersion field value
// and a boolean to check if the value has been set.
func (o *ArduinoSeriesBatch) GetRespVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RespVersion, true
}

// SetRespVersion sets field value
func (o *ArduinoSeriesBatch) SetRespVersion(v int64) {
	o.RespVersion = v
}

// GetResponses returns the Responses field value
func (o *ArduinoSeriesBatch) GetResponses() []ArduinoSeriesResponse {
	if o == nil {
		var ret []ArduinoSeriesResponse
		return ret
	}

	return o.Responses
}

// GetResponsesOk returns a tuple with the Responses field value
// and a boolean to check if the value has been set.
func (o *ArduinoSeriesBatch) GetResponsesOk() ([]ArduinoSeriesResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Responses, true
}

// SetResponses sets field value
func (o *ArduinoSeriesBatch) SetResponses(v []ArduinoSeriesResponse) {
	o.Responses = v
}

func (o ArduinoSeriesBatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoSeriesBatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resp_version"] = o.RespVersion
	toSerialize["responses"] = o.Responses

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ArduinoSeriesBatch) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resp_version",
		"responses",
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

	varArduinoSeriesBatch := _ArduinoSeriesBatch{}

	err = json.Unmarshal(data, &varArduinoSeriesBatch)

	if err != nil {
		return err
	}

	*o = ArduinoSeriesBatch(varArduinoSeriesBatch)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "resp_version")
		delete(additionalProperties, "responses")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableArduinoSeriesBatch struct {
	value *ArduinoSeriesBatch
	isSet bool
}

func (v NullableArduinoSeriesBatch) Get() *ArduinoSeriesBatch {
	return v.value
}

func (v *NullableArduinoSeriesBatch) Set(val *ArduinoSeriesBatch) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoSeriesBatch) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoSeriesBatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoSeriesBatch(val *ArduinoSeriesBatch) *NullableArduinoSeriesBatch {
	return &NullableArduinoSeriesBatch{value: val, isSet: true}
}

func (v NullableArduinoSeriesBatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoSeriesBatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


