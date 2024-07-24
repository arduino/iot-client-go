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

// checks if the ArduinoDashboardshare type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoDashboardshare{}

// ArduinoDashboardshare ArduinoDashboardshare media type (default view)
type ArduinoDashboardshare struct {
	// The userID of the user you want to share the dashboard with
	UserId string `json:"user_id"`
	// The username of the user you want to share the dashboard with
	Username *string `json:"username,omitempty"`
}

// NewArduinoDashboardshare instantiates a new ArduinoDashboardshare object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoDashboardshare(userId string) *ArduinoDashboardshare {
	this := ArduinoDashboardshare{}
	this.UserId = userId
	return &this
}

// NewArduinoDashboardshareWithDefaults instantiates a new ArduinoDashboardshare object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoDashboardshareWithDefaults() *ArduinoDashboardshare {
	this := ArduinoDashboardshare{}
	return &this
}

// GetUserId returns the UserId field value
func (o *ArduinoDashboardshare) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *ArduinoDashboardshare) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *ArduinoDashboardshare) SetUserId(v string) {
	o.UserId = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ArduinoDashboardshare) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoDashboardshare) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ArduinoDashboardshare) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ArduinoDashboardshare) SetUsername(v string) {
	o.Username = &v
}

func (o ArduinoDashboardshare) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoDashboardshare) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user_id"] = o.UserId
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableArduinoDashboardshare struct {
	value *ArduinoDashboardshare
	isSet bool
}

func (v NullableArduinoDashboardshare) Get() *ArduinoDashboardshare {
	return v.value
}

func (v *NullableArduinoDashboardshare) Set(val *ArduinoDashboardshare) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoDashboardshare) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoDashboardshare) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoDashboardshare(val *ArduinoDashboardshare) *NullableArduinoDashboardshare {
	return &NullableArduinoDashboardshare{value: val, isSet: true}
}

func (v NullableArduinoDashboardshare) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoDashboardshare) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


