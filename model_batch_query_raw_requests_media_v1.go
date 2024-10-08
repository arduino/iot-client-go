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

// checks if the BatchQueryRawRequestsMediaV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchQueryRawRequestsMediaV1{}

// BatchQueryRawRequestsMediaV1 struct for BatchQueryRawRequestsMediaV1
type BatchQueryRawRequestsMediaV1 struct {
	// Requests
	Requests []BatchQueryRawRequestMediaV1 `json:"requests"`
	// Response version
	RespVersion int64 `json:"resp_version"`
}

// NewBatchQueryRawRequestsMediaV1 instantiates a new BatchQueryRawRequestsMediaV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchQueryRawRequestsMediaV1(requests []BatchQueryRawRequestMediaV1, respVersion int64) *BatchQueryRawRequestsMediaV1 {
	this := BatchQueryRawRequestsMediaV1{}
	this.Requests = requests
	this.RespVersion = respVersion
	return &this
}

// NewBatchQueryRawRequestsMediaV1WithDefaults instantiates a new BatchQueryRawRequestsMediaV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchQueryRawRequestsMediaV1WithDefaults() *BatchQueryRawRequestsMediaV1 {
	this := BatchQueryRawRequestsMediaV1{}
	return &this
}

// GetRequests returns the Requests field value
func (o *BatchQueryRawRequestsMediaV1) GetRequests() []BatchQueryRawRequestMediaV1 {
	if o == nil {
		var ret []BatchQueryRawRequestMediaV1
		return ret
	}

	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *BatchQueryRawRequestsMediaV1) GetRequestsOk() ([]BatchQueryRawRequestMediaV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Requests, true
}

// SetRequests sets field value
func (o *BatchQueryRawRequestsMediaV1) SetRequests(v []BatchQueryRawRequestMediaV1) {
	o.Requests = v
}

// GetRespVersion returns the RespVersion field value
func (o *BatchQueryRawRequestsMediaV1) GetRespVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RespVersion
}

// GetRespVersionOk returns a tuple with the RespVersion field value
// and a boolean to check if the value has been set.
func (o *BatchQueryRawRequestsMediaV1) GetRespVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RespVersion, true
}

// SetRespVersion sets field value
func (o *BatchQueryRawRequestsMediaV1) SetRespVersion(v int64) {
	o.RespVersion = v
}

func (o BatchQueryRawRequestsMediaV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchQueryRawRequestsMediaV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["requests"] = o.Requests
	toSerialize["resp_version"] = o.RespVersion
	return toSerialize, nil
}

type NullableBatchQueryRawRequestsMediaV1 struct {
	value *BatchQueryRawRequestsMediaV1
	isSet bool
}

func (v NullableBatchQueryRawRequestsMediaV1) Get() *BatchQueryRawRequestsMediaV1 {
	return v.value
}

func (v *NullableBatchQueryRawRequestsMediaV1) Set(val *BatchQueryRawRequestsMediaV1) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchQueryRawRequestsMediaV1) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchQueryRawRequestsMediaV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchQueryRawRequestsMediaV1(val *BatchQueryRawRequestsMediaV1) *NullableBatchQueryRawRequestsMediaV1 {
	return &NullableBatchQueryRawRequestsMediaV1{value: val, isSet: true}
}

func (v NullableBatchQueryRawRequestsMediaV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchQueryRawRequestsMediaV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


