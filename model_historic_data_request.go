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
	"bytes"
	"fmt"
)

// checks if the HistoricDataRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoricDataRequest{}

// HistoricDataRequest struct for HistoricDataRequest
type HistoricDataRequest struct {
	// Get data starting from this date
	From time.Time `json:"from"`
	// IDs of properties
	Properties []string `json:"properties"`
	// Get data up to this date
	To time.Time `json:"to"`
}

type _HistoricDataRequest HistoricDataRequest

// NewHistoricDataRequest instantiates a new HistoricDataRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoricDataRequest(from time.Time, properties []string, to time.Time) *HistoricDataRequest {
	this := HistoricDataRequest{}
	this.From = from
	this.Properties = properties
	this.To = to
	return &this
}

// NewHistoricDataRequestWithDefaults instantiates a new HistoricDataRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoricDataRequestWithDefaults() *HistoricDataRequest {
	this := HistoricDataRequest{}
	return &this
}

// GetFrom returns the From field value
func (o *HistoricDataRequest) GetFrom() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *HistoricDataRequest) GetFromOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *HistoricDataRequest) SetFrom(v time.Time) {
	o.From = v
}

// GetProperties returns the Properties field value
func (o *HistoricDataRequest) GetProperties() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *HistoricDataRequest) GetPropertiesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties, true
}

// SetProperties sets field value
func (o *HistoricDataRequest) SetProperties(v []string) {
	o.Properties = v
}

// GetTo returns the To field value
func (o *HistoricDataRequest) GetTo() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *HistoricDataRequest) GetToOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *HistoricDataRequest) SetTo(v time.Time) {
	o.To = v
}

func (o HistoricDataRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoricDataRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["from"] = o.From
	toSerialize["properties"] = o.Properties
	toSerialize["to"] = o.To
	return toSerialize, nil
}

func (o *HistoricDataRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"from",
		"properties",
		"to",
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

	varHistoricDataRequest := _HistoricDataRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHistoricDataRequest)

	if err != nil {
		return err
	}

	*o = HistoricDataRequest(varHistoricDataRequest)

	return err
}

type NullableHistoricDataRequest struct {
	value *HistoricDataRequest
	isSet bool
}

func (v NullableHistoricDataRequest) Get() *HistoricDataRequest {
	return v.value
}

func (v *NullableHistoricDataRequest) Set(val *HistoricDataRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoricDataRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoricDataRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoricDataRequest(val *HistoricDataRequest) *NullableHistoricDataRequest {
	return &NullableHistoricDataRequest{value: val, isSet: true}
}

func (v NullableHistoricDataRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoricDataRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


