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

// checks if the UserRecipient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserRecipient{}

// UserRecipient struct for UserRecipient
type UserRecipient struct {
	// The email address of the user
	Email *string `json:"email,omitempty"`
	// The id of the user
	Id string `json:"id"`
	// The username of the user
	Username *string `json:"username,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserRecipient UserRecipient

// NewUserRecipient instantiates a new UserRecipient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserRecipient(id string) *UserRecipient {
	this := UserRecipient{}
	this.Id = id
	return &this
}

// NewUserRecipientWithDefaults instantiates a new UserRecipient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserRecipientWithDefaults() *UserRecipient {
	this := UserRecipient{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserRecipient) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRecipient) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserRecipient) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserRecipient) SetEmail(v string) {
	o.Email = &v
}

// GetId returns the Id field value
func (o *UserRecipient) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserRecipient) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserRecipient) SetId(v string) {
	o.Id = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UserRecipient) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRecipient) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UserRecipient) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UserRecipient) SetUsername(v string) {
	o.Username = &v
}

func (o UserRecipient) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserRecipient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserRecipient) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varUserRecipient := _UserRecipient{}

	err = json.Unmarshal(data, &varUserRecipient)

	if err != nil {
		return err
	}

	*o = UserRecipient(varUserRecipient)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "id")
		delete(additionalProperties, "username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserRecipient struct {
	value *UserRecipient
	isSet bool
}

func (v NullableUserRecipient) Get() *UserRecipient {
	return v.value
}

func (v *NullableUserRecipient) Set(val *UserRecipient) {
	v.value = val
	v.isSet = true
}

func (v NullableUserRecipient) IsSet() bool {
	return v.isSet
}

func (v *NullableUserRecipient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserRecipient(val *UserRecipient) *NullableUserRecipient {
	return &NullableUserRecipient{value: val, isSet: true}
}

func (v NullableUserRecipient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserRecipient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


