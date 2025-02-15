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

// checks if the ArduinoDashboardowner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoDashboardowner{}

// ArduinoDashboardowner ArduinoDashboardowner media type (default view)
type ArduinoDashboardowner struct {
	// The userID of the user who created the dashboard
	UserId string `json:"user_id"`
	// The username of the user who created the dashboard
	Username *string `json:"username,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ArduinoDashboardowner ArduinoDashboardowner

// NewArduinoDashboardowner instantiates a new ArduinoDashboardowner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoDashboardowner(userId string) *ArduinoDashboardowner {
	this := ArduinoDashboardowner{}
	this.UserId = userId
	return &this
}

// NewArduinoDashboardownerWithDefaults instantiates a new ArduinoDashboardowner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoDashboardownerWithDefaults() *ArduinoDashboardowner {
	this := ArduinoDashboardowner{}
	return &this
}

// GetUserId returns the UserId field value
func (o *ArduinoDashboardowner) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *ArduinoDashboardowner) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *ArduinoDashboardowner) SetUserId(v string) {
	o.UserId = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ArduinoDashboardowner) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoDashboardowner) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ArduinoDashboardowner) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ArduinoDashboardowner) SetUsername(v string) {
	o.Username = &v
}

func (o ArduinoDashboardowner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoDashboardowner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user_id"] = o.UserId
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ArduinoDashboardowner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user_id",
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

	varArduinoDashboardowner := _ArduinoDashboardowner{}

	err = json.Unmarshal(data, &varArduinoDashboardowner)

	if err != nil {
		return err
	}

	*o = ArduinoDashboardowner(varArduinoDashboardowner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableArduinoDashboardowner struct {
	value *ArduinoDashboardowner
	isSet bool
}

func (v NullableArduinoDashboardowner) Get() *ArduinoDashboardowner {
	return v.value
}

func (v *NullableArduinoDashboardowner) Set(val *ArduinoDashboardowner) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoDashboardowner) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoDashboardowner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoDashboardowner(val *ArduinoDashboardowner) *NullableArduinoDashboardowner {
	return &NullableArduinoDashboardowner{value: val, isSet: true}
}

func (v NullableArduinoDashboardowner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoDashboardowner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


