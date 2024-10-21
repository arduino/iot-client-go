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

// checks if the BatchQuerySampledRequestsMediaV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchQuerySampledRequestsMediaV1{}

// BatchQuerySampledRequestsMediaV1 struct for BatchQuerySampledRequestsMediaV1
type BatchQuerySampledRequestsMediaV1 struct {
	// Requests
	Requests []BatchQuerySampledRequestMediaV1 `json:"requests"`
	// Response version
	RespVersion int64 `json:"resp_version"`
}

// NewBatchQuerySampledRequestsMediaV1 instantiates a new BatchQuerySampledRequestsMediaV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchQuerySampledRequestsMediaV1(requests []BatchQuerySampledRequestMediaV1, respVersion int64) *BatchQuerySampledRequestsMediaV1 {
	this := BatchQuerySampledRequestsMediaV1{}
	this.Requests = requests
	this.RespVersion = respVersion
	return &this
}

// NewBatchQuerySampledRequestsMediaV1WithDefaults instantiates a new BatchQuerySampledRequestsMediaV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchQuerySampledRequestsMediaV1WithDefaults() *BatchQuerySampledRequestsMediaV1 {
	this := BatchQuerySampledRequestsMediaV1{}
	return &this
}

// GetRequests returns the Requests field value
func (o *BatchQuerySampledRequestsMediaV1) GetRequests() []BatchQuerySampledRequestMediaV1 {
	if o == nil {
		var ret []BatchQuerySampledRequestMediaV1
		return ret
	}

	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *BatchQuerySampledRequestsMediaV1) GetRequestsOk() ([]BatchQuerySampledRequestMediaV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Requests, true
}

// SetRequests sets field value
func (o *BatchQuerySampledRequestsMediaV1) SetRequests(v []BatchQuerySampledRequestMediaV1) {
	o.Requests = v
}

// GetRespVersion returns the RespVersion field value
func (o *BatchQuerySampledRequestsMediaV1) GetRespVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RespVersion
}

// GetRespVersionOk returns a tuple with the RespVersion field value
// and a boolean to check if the value has been set.
func (o *BatchQuerySampledRequestsMediaV1) GetRespVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RespVersion, true
}

// SetRespVersion sets field value
func (o *BatchQuerySampledRequestsMediaV1) SetRespVersion(v int64) {
	o.RespVersion = v
}

func (o BatchQuerySampledRequestsMediaV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchQuerySampledRequestsMediaV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["requests"] = o.Requests
	toSerialize["resp_version"] = o.RespVersion
	return toSerialize, nil
}

type NullableBatchQuerySampledRequestsMediaV1 struct {
	value *BatchQuerySampledRequestsMediaV1
	isSet bool
}

func (v NullableBatchQuerySampledRequestsMediaV1) Get() *BatchQuerySampledRequestsMediaV1 {
	return v.value
}

func (v *NullableBatchQuerySampledRequestsMediaV1) Set(val *BatchQuerySampledRequestsMediaV1) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchQuerySampledRequestsMediaV1) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchQuerySampledRequestsMediaV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchQuerySampledRequestsMediaV1(val *BatchQuerySampledRequestsMediaV1) *NullableBatchQuerySampledRequestsMediaV1 {
	return &NullableBatchQuerySampledRequestsMediaV1{value: val, isSet: true}
}

func (v NullableBatchQuerySampledRequestsMediaV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchQuerySampledRequestsMediaV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


