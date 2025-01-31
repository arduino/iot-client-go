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

// checks if the EmailDeliveryOpts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailDeliveryOpts{}

// EmailDeliveryOpts struct for EmailDeliveryOpts
type EmailDeliveryOpts struct {
	// The \"bcc:\" field of an e-mail
	Bcc []UserRecipient `json:"bcc,omitempty"`
	// The \"cc:\" field of an e-mail
	Cc []UserRecipient `json:"cc,omitempty"`
	// The \"to:\" field of an e-mail
	To []UserRecipient `json:"to"`
	AdditionalProperties map[string]interface{}
}

type _EmailDeliveryOpts EmailDeliveryOpts

// NewEmailDeliveryOpts instantiates a new EmailDeliveryOpts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailDeliveryOpts(to []UserRecipient) *EmailDeliveryOpts {
	this := EmailDeliveryOpts{}
	this.To = to
	return &this
}

// NewEmailDeliveryOptsWithDefaults instantiates a new EmailDeliveryOpts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailDeliveryOptsWithDefaults() *EmailDeliveryOpts {
	this := EmailDeliveryOpts{}
	return &this
}

// GetBcc returns the Bcc field value if set, zero value otherwise.
func (o *EmailDeliveryOpts) GetBcc() []UserRecipient {
	if o == nil || IsNil(o.Bcc) {
		var ret []UserRecipient
		return ret
	}
	return o.Bcc
}

// GetBccOk returns a tuple with the Bcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDeliveryOpts) GetBccOk() ([]UserRecipient, bool) {
	if o == nil || IsNil(o.Bcc) {
		return nil, false
	}
	return o.Bcc, true
}

// HasBcc returns a boolean if a field has been set.
func (o *EmailDeliveryOpts) HasBcc() bool {
	if o != nil && !IsNil(o.Bcc) {
		return true
	}

	return false
}

// SetBcc gets a reference to the given []UserRecipient and assigns it to the Bcc field.
func (o *EmailDeliveryOpts) SetBcc(v []UserRecipient) {
	o.Bcc = v
}

// GetCc returns the Cc field value if set, zero value otherwise.
func (o *EmailDeliveryOpts) GetCc() []UserRecipient {
	if o == nil || IsNil(o.Cc) {
		var ret []UserRecipient
		return ret
	}
	return o.Cc
}

// GetCcOk returns a tuple with the Cc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDeliveryOpts) GetCcOk() ([]UserRecipient, bool) {
	if o == nil || IsNil(o.Cc) {
		return nil, false
	}
	return o.Cc, true
}

// HasCc returns a boolean if a field has been set.
func (o *EmailDeliveryOpts) HasCc() bool {
	if o != nil && !IsNil(o.Cc) {
		return true
	}

	return false
}

// SetCc gets a reference to the given []UserRecipient and assigns it to the Cc field.
func (o *EmailDeliveryOpts) SetCc(v []UserRecipient) {
	o.Cc = v
}

// GetTo returns the To field value
func (o *EmailDeliveryOpts) GetTo() []UserRecipient {
	if o == nil {
		var ret []UserRecipient
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *EmailDeliveryOpts) GetToOk() ([]UserRecipient, bool) {
	if o == nil {
		return nil, false
	}
	return o.To, true
}

// SetTo sets field value
func (o *EmailDeliveryOpts) SetTo(v []UserRecipient) {
	o.To = v
}

func (o EmailDeliveryOpts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailDeliveryOpts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bcc) {
		toSerialize["bcc"] = o.Bcc
	}
	if !IsNil(o.Cc) {
		toSerialize["cc"] = o.Cc
	}
	toSerialize["to"] = o.To

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EmailDeliveryOpts) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varEmailDeliveryOpts := _EmailDeliveryOpts{}

	err = json.Unmarshal(data, &varEmailDeliveryOpts)

	if err != nil {
		return err
	}

	*o = EmailDeliveryOpts(varEmailDeliveryOpts)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bcc")
		delete(additionalProperties, "cc")
		delete(additionalProperties, "to")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEmailDeliveryOpts struct {
	value *EmailDeliveryOpts
	isSet bool
}

func (v NullableEmailDeliveryOpts) Get() *EmailDeliveryOpts {
	return v.value
}

func (v *NullableEmailDeliveryOpts) Set(val *EmailDeliveryOpts) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailDeliveryOpts) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailDeliveryOpts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailDeliveryOpts(val *EmailDeliveryOpts) *NullableEmailDeliveryOpts {
	return &NullableEmailDeliveryOpts{value: val, isSet: true}
}

func (v NullableEmailDeliveryOpts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailDeliveryOpts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


