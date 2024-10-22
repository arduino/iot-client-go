/*
Arduino IoT Cloud API

Provides a set of endpoints to manage Arduino IoT Cloud **Devices**, **Things**, **Properties** and **Timeseries**. This API can be called just with any HTTP Client, or using one of these clients:  * [Javascript NPM package](https://www.npmjs.com/package/@arduino/arduino-iot-client)  * [Python PYPI Package](https://pypi.org/project/arduino-iot-client/)  * [Golang Module](https://github.com/arduino/iot-client-go)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Devicev2Otabinaryurl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Devicev2Otabinaryurl{}

// Devicev2Otabinaryurl struct for Devicev2Otabinaryurl
type Devicev2Otabinaryurl struct {
	// If false, wait for the full OTA process, until it gets a result from the device
	Async *bool `json:"async,omitempty"`
	// The object key of the binary
	BinaryKey string `json:"binary_key" validate:"regexp=^ota\\/[a-zA-Z0-9_-]+\\/[a-zA-Z0-9_-]+.ota$"`
	// Binary expire time in minutes, default 10 mins
	ExpireInMins *int64 `json:"expire_in_mins,omitempty"`
}

type _Devicev2Otabinaryurl Devicev2Otabinaryurl

// NewDevicev2Otabinaryurl instantiates a new Devicev2Otabinaryurl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicev2Otabinaryurl(binaryKey string) *Devicev2Otabinaryurl {
	this := Devicev2Otabinaryurl{}
	var async bool = true
	this.Async = &async
	this.BinaryKey = binaryKey
	var expireInMins int64 = 10
	this.ExpireInMins = &expireInMins
	return &this
}

// NewDevicev2OtabinaryurlWithDefaults instantiates a new Devicev2Otabinaryurl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicev2OtabinaryurlWithDefaults() *Devicev2Otabinaryurl {
	this := Devicev2Otabinaryurl{}
	var async bool = true
	this.Async = &async
	var expireInMins int64 = 10
	this.ExpireInMins = &expireInMins
	return &this
}

// GetAsync returns the Async field value if set, zero value otherwise.
func (o *Devicev2Otabinaryurl) GetAsync() bool {
	if o == nil || IsNil(o.Async) {
		var ret bool
		return ret
	}
	return *o.Async
}

// GetAsyncOk returns a tuple with the Async field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Devicev2Otabinaryurl) GetAsyncOk() (*bool, bool) {
	if o == nil || IsNil(o.Async) {
		return nil, false
	}
	return o.Async, true
}

// HasAsync returns a boolean if a field has been set.
func (o *Devicev2Otabinaryurl) HasAsync() bool {
	if o != nil && !IsNil(o.Async) {
		return true
	}

	return false
}

// SetAsync gets a reference to the given bool and assigns it to the Async field.
func (o *Devicev2Otabinaryurl) SetAsync(v bool) {
	o.Async = &v
}

// GetBinaryKey returns the BinaryKey field value
func (o *Devicev2Otabinaryurl) GetBinaryKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BinaryKey
}

// GetBinaryKeyOk returns a tuple with the BinaryKey field value
// and a boolean to check if the value has been set.
func (o *Devicev2Otabinaryurl) GetBinaryKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BinaryKey, true
}

// SetBinaryKey sets field value
func (o *Devicev2Otabinaryurl) SetBinaryKey(v string) {
	o.BinaryKey = v
}

// GetExpireInMins returns the ExpireInMins field value if set, zero value otherwise.
func (o *Devicev2Otabinaryurl) GetExpireInMins() int64 {
	if o == nil || IsNil(o.ExpireInMins) {
		var ret int64
		return ret
	}
	return *o.ExpireInMins
}

// GetExpireInMinsOk returns a tuple with the ExpireInMins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Devicev2Otabinaryurl) GetExpireInMinsOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpireInMins) {
		return nil, false
	}
	return o.ExpireInMins, true
}

// HasExpireInMins returns a boolean if a field has been set.
func (o *Devicev2Otabinaryurl) HasExpireInMins() bool {
	if o != nil && !IsNil(o.ExpireInMins) {
		return true
	}

	return false
}

// SetExpireInMins gets a reference to the given int64 and assigns it to the ExpireInMins field.
func (o *Devicev2Otabinaryurl) SetExpireInMins(v int64) {
	o.ExpireInMins = &v
}

func (o Devicev2Otabinaryurl) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Devicev2Otabinaryurl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Async) {
		toSerialize["async"] = o.Async
	}
	toSerialize["binary_key"] = o.BinaryKey
	if !IsNil(o.ExpireInMins) {
		toSerialize["expire_in_mins"] = o.ExpireInMins
	}
	return toSerialize, nil
}

func (o *Devicev2Otabinaryurl) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"binary_key",
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

	varDevicev2Otabinaryurl := _Devicev2Otabinaryurl{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDevicev2Otabinaryurl)

	if err != nil {
		return err
	}

	*o = Devicev2Otabinaryurl(varDevicev2Otabinaryurl)

	return err
}

type NullableDevicev2Otabinaryurl struct {
	value *Devicev2Otabinaryurl
	isSet bool
}

func (v NullableDevicev2Otabinaryurl) Get() *Devicev2Otabinaryurl {
	return v.value
}

func (v *NullableDevicev2Otabinaryurl) Set(val *Devicev2Otabinaryurl) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicev2Otabinaryurl) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicev2Otabinaryurl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicev2Otabinaryurl(val *Devicev2Otabinaryurl) *NullableDevicev2Otabinaryurl {
	return &NullableDevicev2Otabinaryurl{value: val, isSet: true}
}

func (v NullableDevicev2Otabinaryurl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicev2Otabinaryurl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


