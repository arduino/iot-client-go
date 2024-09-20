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

// checks if the BatchQueryRequestMediaV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchQueryRequestMediaV1{}

// BatchQueryRequestMediaV1 struct for BatchQueryRequestMediaV1
type BatchQueryRequestMediaV1 struct {
	// Aggregation statistic function. For numeric values, AVG statistic is used by default. PCT_X compute the Xth approximate percentile (e.g. PCT_95 is the 95th approximate percentile). For boolean, BOOL_OR statistic is used as default.
	Aggregation *string `json:"aggregation,omitempty"`
	// From timestamp
	From time.Time `json:"from"`
	// Resolution in seconds (max allowed: 86400)
	Interval *int64 `json:"interval,omitempty"`
	// Data selection query (e.g. property.2a99729d-2556-4220-a139-023348a1e6b5 or thing.95717675-4786-4ffc-afcc-799777755391)
	Q string `json:"q"`
	// Maximum number of values returned after data aggregation, if any (default: 300, limit: 1000 - 10000 in case of thing query)
	SeriesLimit *int64 `json:"series_limit,omitempty"`
	// To timestamp
	To time.Time `json:"to"`
}

// NewBatchQueryRequestMediaV1 instantiates a new BatchQueryRequestMediaV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchQueryRequestMediaV1(from time.Time, q string, to time.Time) *BatchQueryRequestMediaV1 {
	this := BatchQueryRequestMediaV1{}
	this.From = from
	this.Q = q
	this.To = to
	return &this
}

// NewBatchQueryRequestMediaV1WithDefaults instantiates a new BatchQueryRequestMediaV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchQueryRequestMediaV1WithDefaults() *BatchQueryRequestMediaV1 {
	this := BatchQueryRequestMediaV1{}
	return &this
}

// GetAggregation returns the Aggregation field value if set, zero value otherwise.
func (o *BatchQueryRequestMediaV1) GetAggregation() string {
	if o == nil || IsNil(o.Aggregation) {
		var ret string
		return ret
	}
	return *o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchQueryRequestMediaV1) GetAggregationOk() (*string, bool) {
	if o == nil || IsNil(o.Aggregation) {
		return nil, false
	}
	return o.Aggregation, true
}

// HasAggregation returns a boolean if a field has been set.
func (o *BatchQueryRequestMediaV1) HasAggregation() bool {
	if o != nil && !IsNil(o.Aggregation) {
		return true
	}

	return false
}

// SetAggregation gets a reference to the given string and assigns it to the Aggregation field.
func (o *BatchQueryRequestMediaV1) SetAggregation(v string) {
	o.Aggregation = &v
}

// GetFrom returns the From field value
func (o *BatchQueryRequestMediaV1) GetFrom() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *BatchQueryRequestMediaV1) GetFromOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *BatchQueryRequestMediaV1) SetFrom(v time.Time) {
	o.From = v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *BatchQueryRequestMediaV1) GetInterval() int64 {
	if o == nil || IsNil(o.Interval) {
		var ret int64
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchQueryRequestMediaV1) GetIntervalOk() (*int64, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *BatchQueryRequestMediaV1) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given int64 and assigns it to the Interval field.
func (o *BatchQueryRequestMediaV1) SetInterval(v int64) {
	o.Interval = &v
}

// GetQ returns the Q field value
func (o *BatchQueryRequestMediaV1) GetQ() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Q
}

// GetQOk returns a tuple with the Q field value
// and a boolean to check if the value has been set.
func (o *BatchQueryRequestMediaV1) GetQOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Q, true
}

// SetQ sets field value
func (o *BatchQueryRequestMediaV1) SetQ(v string) {
	o.Q = v
}

// GetSeriesLimit returns the SeriesLimit field value if set, zero value otherwise.
func (o *BatchQueryRequestMediaV1) GetSeriesLimit() int64 {
	if o == nil || IsNil(o.SeriesLimit) {
		var ret int64
		return ret
	}
	return *o.SeriesLimit
}

// GetSeriesLimitOk returns a tuple with the SeriesLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchQueryRequestMediaV1) GetSeriesLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.SeriesLimit) {
		return nil, false
	}
	return o.SeriesLimit, true
}

// HasSeriesLimit returns a boolean if a field has been set.
func (o *BatchQueryRequestMediaV1) HasSeriesLimit() bool {
	if o != nil && !IsNil(o.SeriesLimit) {
		return true
	}

	return false
}

// SetSeriesLimit gets a reference to the given int64 and assigns it to the SeriesLimit field.
func (o *BatchQueryRequestMediaV1) SetSeriesLimit(v int64) {
	o.SeriesLimit = &v
}

// GetTo returns the To field value
func (o *BatchQueryRequestMediaV1) GetTo() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *BatchQueryRequestMediaV1) GetToOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *BatchQueryRequestMediaV1) SetTo(v time.Time) {
	o.To = v
}

func (o BatchQueryRequestMediaV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchQueryRequestMediaV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Aggregation) {
		toSerialize["aggregation"] = o.Aggregation
	}
	toSerialize["from"] = o.From
	if !IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	toSerialize["q"] = o.Q
	if !IsNil(o.SeriesLimit) {
		toSerialize["series_limit"] = o.SeriesLimit
	}
	toSerialize["to"] = o.To
	return toSerialize, nil
}

type NullableBatchQueryRequestMediaV1 struct {
	value *BatchQueryRequestMediaV1
	isSet bool
}

func (v NullableBatchQueryRequestMediaV1) Get() *BatchQueryRequestMediaV1 {
	return v.value
}

func (v *NullableBatchQueryRequestMediaV1) Set(val *BatchQueryRequestMediaV1) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchQueryRequestMediaV1) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchQueryRequestMediaV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchQueryRequestMediaV1(val *BatchQueryRequestMediaV1) *NullableBatchQueryRequestMediaV1 {
	return &NullableBatchQueryRequestMediaV1{value: val, isSet: true}
}

func (v NullableBatchQueryRequestMediaV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchQueryRequestMediaV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


