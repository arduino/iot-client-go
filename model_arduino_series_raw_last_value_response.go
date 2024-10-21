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
	"bytes"
	"fmt"
)

// checks if the ArduinoSeriesRawLastValueResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoSeriesRawLastValueResponse{}

// ArduinoSeriesRawLastValueResponse ArduinoSeriesRawLastValueResponse media type (default view)
type ArduinoSeriesRawLastValueResponse struct {
	// Total number of values in the array 'values'
	CountValues int64 `json:"count_values"`
	// Property id
	PropertyId string `json:"property_id"`
	// Thing id
	ThingId string `json:"thing_id"`
	// Timestamp in RFC3339
	Times []time.Time `json:"times"`
	// Values can be in Float, String, Bool, Object
	Values []interface{} `json:"values"`
}

type _ArduinoSeriesRawLastValueResponse ArduinoSeriesRawLastValueResponse

// NewArduinoSeriesRawLastValueResponse instantiates a new ArduinoSeriesRawLastValueResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoSeriesRawLastValueResponse(countValues int64, propertyId string, thingId string, times []time.Time, values []interface{}) *ArduinoSeriesRawLastValueResponse {
	this := ArduinoSeriesRawLastValueResponse{}
	this.CountValues = countValues
	this.PropertyId = propertyId
	this.ThingId = thingId
	this.Times = times
	this.Values = values
	return &this
}

// NewArduinoSeriesRawLastValueResponseWithDefaults instantiates a new ArduinoSeriesRawLastValueResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoSeriesRawLastValueResponseWithDefaults() *ArduinoSeriesRawLastValueResponse {
	this := ArduinoSeriesRawLastValueResponse{}
	return &this
}

// GetCountValues returns the CountValues field value
func (o *ArduinoSeriesRawLastValueResponse) GetCountValues() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CountValues
}

// GetCountValuesOk returns a tuple with the CountValues field value
// and a boolean to check if the value has been set.
func (o *ArduinoSeriesRawLastValueResponse) GetCountValuesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountValues, true
}

// SetCountValues sets field value
func (o *ArduinoSeriesRawLastValueResponse) SetCountValues(v int64) {
	o.CountValues = v
}

// GetPropertyId returns the PropertyId field value
func (o *ArduinoSeriesRawLastValueResponse) GetPropertyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PropertyId
}

// GetPropertyIdOk returns a tuple with the PropertyId field value
// and a boolean to check if the value has been set.
func (o *ArduinoSeriesRawLastValueResponse) GetPropertyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PropertyId, true
}

// SetPropertyId sets field value
func (o *ArduinoSeriesRawLastValueResponse) SetPropertyId(v string) {
	o.PropertyId = v
}

// GetThingId returns the ThingId field value
func (o *ArduinoSeriesRawLastValueResponse) GetThingId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThingId
}

// GetThingIdOk returns a tuple with the ThingId field value
// and a boolean to check if the value has been set.
func (o *ArduinoSeriesRawLastValueResponse) GetThingIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThingId, true
}

// SetThingId sets field value
func (o *ArduinoSeriesRawLastValueResponse) SetThingId(v string) {
	o.ThingId = v
}

// GetTimes returns the Times field value
func (o *ArduinoSeriesRawLastValueResponse) GetTimes() []time.Time {
	if o == nil {
		var ret []time.Time
		return ret
	}

	return o.Times
}

// GetTimesOk returns a tuple with the Times field value
// and a boolean to check if the value has been set.
func (o *ArduinoSeriesRawLastValueResponse) GetTimesOk() ([]time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Times, true
}

// SetTimes sets field value
func (o *ArduinoSeriesRawLastValueResponse) SetTimes(v []time.Time) {
	o.Times = v
}

// GetValues returns the Values field value
func (o *ArduinoSeriesRawLastValueResponse) GetValues() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *ArduinoSeriesRawLastValueResponse) GetValuesOk() ([]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *ArduinoSeriesRawLastValueResponse) SetValues(v []interface{}) {
	o.Values = v
}

func (o ArduinoSeriesRawLastValueResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoSeriesRawLastValueResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count_values"] = o.CountValues
	toSerialize["property_id"] = o.PropertyId
	toSerialize["thing_id"] = o.ThingId
	toSerialize["times"] = o.Times
	toSerialize["values"] = o.Values
	return toSerialize, nil
}

func (o *ArduinoSeriesRawLastValueResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count_values",
		"property_id",
		"thing_id",
		"times",
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

	varArduinoSeriesRawLastValueResponse := _ArduinoSeriesRawLastValueResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varArduinoSeriesRawLastValueResponse)

	if err != nil {
		return err
	}

	*o = ArduinoSeriesRawLastValueResponse(varArduinoSeriesRawLastValueResponse)

	return err
}

type NullableArduinoSeriesRawLastValueResponse struct {
	value *ArduinoSeriesRawLastValueResponse
	isSet bool
}

func (v NullableArduinoSeriesRawLastValueResponse) Get() *ArduinoSeriesRawLastValueResponse {
	return v.value
}

func (v *NullableArduinoSeriesRawLastValueResponse) Set(val *ArduinoSeriesRawLastValueResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoSeriesRawLastValueResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoSeriesRawLastValueResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoSeriesRawLastValueResponse(val *ArduinoSeriesRawLastValueResponse) *NullableArduinoSeriesRawLastValueResponse {
	return &NullableArduinoSeriesRawLastValueResponse{value: val, isSet: true}
}

func (v NullableArduinoSeriesRawLastValueResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoSeriesRawLastValueResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


