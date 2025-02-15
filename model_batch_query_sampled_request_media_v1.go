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

// checks if the BatchQuerySampledRequestMediaV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchQuerySampledRequestMediaV1{}

// BatchQuerySampledRequestMediaV1 struct for BatchQuerySampledRequestMediaV1
type BatchQuerySampledRequestMediaV1 struct {
	// From timestamp (default: now UTC - 24h)
	From *time.Time `json:"from,omitempty"`
	// Resolution in seconds (allowed min:60, max:86400)
	Interval *int32 `json:"interval,omitempty"`
	// Data selection query (e.g. property.2a99729d-2556-4220-a139-023348a1e6b5)
	Q string `json:"q"`
	// Maximum number of values returned after data aggregation, if any (default: 300, limit: 1000)
	SeriesLimit *int64 `json:"series_limit,omitempty"`
	// To timestamp (default: now UTC)
	To *time.Time `json:"to,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BatchQuerySampledRequestMediaV1 BatchQuerySampledRequestMediaV1

// NewBatchQuerySampledRequestMediaV1 instantiates a new BatchQuerySampledRequestMediaV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchQuerySampledRequestMediaV1(q string) *BatchQuerySampledRequestMediaV1 {
	this := BatchQuerySampledRequestMediaV1{}
	var interval int32 = 300
	this.Interval = &interval
	this.Q = q
	return &this
}

// NewBatchQuerySampledRequestMediaV1WithDefaults instantiates a new BatchQuerySampledRequestMediaV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchQuerySampledRequestMediaV1WithDefaults() *BatchQuerySampledRequestMediaV1 {
	this := BatchQuerySampledRequestMediaV1{}
	var interval int32 = 300
	this.Interval = &interval
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *BatchQuerySampledRequestMediaV1) GetFrom() time.Time {
	if o == nil || IsNil(o.From) {
		var ret time.Time
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchQuerySampledRequestMediaV1) GetFromOk() (*time.Time, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *BatchQuerySampledRequestMediaV1) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given time.Time and assigns it to the From field.
func (o *BatchQuerySampledRequestMediaV1) SetFrom(v time.Time) {
	o.From = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *BatchQuerySampledRequestMediaV1) GetInterval() int32 {
	if o == nil || IsNil(o.Interval) {
		var ret int32
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchQuerySampledRequestMediaV1) GetIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *BatchQuerySampledRequestMediaV1) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given int32 and assigns it to the Interval field.
func (o *BatchQuerySampledRequestMediaV1) SetInterval(v int32) {
	o.Interval = &v
}

// GetQ returns the Q field value
func (o *BatchQuerySampledRequestMediaV1) GetQ() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Q
}

// GetQOk returns a tuple with the Q field value
// and a boolean to check if the value has been set.
func (o *BatchQuerySampledRequestMediaV1) GetQOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Q, true
}

// SetQ sets field value
func (o *BatchQuerySampledRequestMediaV1) SetQ(v string) {
	o.Q = v
}

// GetSeriesLimit returns the SeriesLimit field value if set, zero value otherwise.
func (o *BatchQuerySampledRequestMediaV1) GetSeriesLimit() int64 {
	if o == nil || IsNil(o.SeriesLimit) {
		var ret int64
		return ret
	}
	return *o.SeriesLimit
}

// GetSeriesLimitOk returns a tuple with the SeriesLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchQuerySampledRequestMediaV1) GetSeriesLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.SeriesLimit) {
		return nil, false
	}
	return o.SeriesLimit, true
}

// HasSeriesLimit returns a boolean if a field has been set.
func (o *BatchQuerySampledRequestMediaV1) HasSeriesLimit() bool {
	if o != nil && !IsNil(o.SeriesLimit) {
		return true
	}

	return false
}

// SetSeriesLimit gets a reference to the given int64 and assigns it to the SeriesLimit field.
func (o *BatchQuerySampledRequestMediaV1) SetSeriesLimit(v int64) {
	o.SeriesLimit = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *BatchQuerySampledRequestMediaV1) GetTo() time.Time {
	if o == nil || IsNil(o.To) {
		var ret time.Time
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchQuerySampledRequestMediaV1) GetToOk() (*time.Time, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *BatchQuerySampledRequestMediaV1) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given time.Time and assigns it to the To field.
func (o *BatchQuerySampledRequestMediaV1) SetTo(v time.Time) {
	o.To = &v
}

func (o BatchQuerySampledRequestMediaV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchQuerySampledRequestMediaV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	toSerialize["q"] = o.Q
	if !IsNil(o.SeriesLimit) {
		toSerialize["series_limit"] = o.SeriesLimit
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BatchQuerySampledRequestMediaV1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"q",
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

	varBatchQuerySampledRequestMediaV1 := _BatchQuerySampledRequestMediaV1{}

	err = json.Unmarshal(data, &varBatchQuerySampledRequestMediaV1)

	if err != nil {
		return err
	}

	*o = BatchQuerySampledRequestMediaV1(varBatchQuerySampledRequestMediaV1)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "from")
		delete(additionalProperties, "interval")
		delete(additionalProperties, "q")
		delete(additionalProperties, "series_limit")
		delete(additionalProperties, "to")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBatchQuerySampledRequestMediaV1 struct {
	value *BatchQuerySampledRequestMediaV1
	isSet bool
}

func (v NullableBatchQuerySampledRequestMediaV1) Get() *BatchQuerySampledRequestMediaV1 {
	return v.value
}

func (v *NullableBatchQuerySampledRequestMediaV1) Set(val *BatchQuerySampledRequestMediaV1) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchQuerySampledRequestMediaV1) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchQuerySampledRequestMediaV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchQuerySampledRequestMediaV1(val *BatchQuerySampledRequestMediaV1) *NullableBatchQuerySampledRequestMediaV1 {
	return &NullableBatchQuerySampledRequestMediaV1{value: val, isSet: true}
}

func (v NullableBatchQuerySampledRequestMediaV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchQuerySampledRequestMediaV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


