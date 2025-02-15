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

// checks if the Sharerequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Sharerequest{}

// Sharerequest struct for Sharerequest
type Sharerequest struct {
	// The message the user want to send to the dashboard owner
	Message *string `json:"message,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Sharerequest Sharerequest

// NewSharerequest instantiates a new Sharerequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharerequest() *Sharerequest {
	this := Sharerequest{}
	return &this
}

// NewSharerequestWithDefaults instantiates a new Sharerequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharerequestWithDefaults() *Sharerequest {
	this := Sharerequest{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Sharerequest) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sharerequest) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Sharerequest) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Sharerequest) SetMessage(v string) {
	o.Message = &v
}

func (o Sharerequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Sharerequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Sharerequest) UnmarshalJSON(data []byte) (err error) {
	varSharerequest := _Sharerequest{}

	err = json.Unmarshal(data, &varSharerequest)

	if err != nil {
		return err
	}

	*o = Sharerequest(varSharerequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "message")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSharerequest struct {
	value *Sharerequest
	isSet bool
}

func (v NullableSharerequest) Get() *Sharerequest {
	return v.value
}

func (v *NullableSharerequest) Set(val *Sharerequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSharerequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSharerequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharerequest(val *Sharerequest) *NullableSharerequest {
	return &NullableSharerequest{value: val, isSet: true}
}

func (v NullableSharerequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharerequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


