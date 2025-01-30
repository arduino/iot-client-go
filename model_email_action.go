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

// checks if the EmailAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailAction{}

// EmailAction struct for EmailAction
type EmailAction struct {
	Body BodyExpression `json:"body"`
	Delivery EmailDeliveryOpts `json:"delivery"`
	Subject TitleExpression `json:"subject"`
	AdditionalProperties map[string]interface{}
}

type _EmailAction EmailAction

// NewEmailAction instantiates a new EmailAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailAction(body BodyExpression, delivery EmailDeliveryOpts, subject TitleExpression) *EmailAction {
	this := EmailAction{}
	this.Body = body
	this.Delivery = delivery
	this.Subject = subject
	return &this
}

// NewEmailActionWithDefaults instantiates a new EmailAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailActionWithDefaults() *EmailAction {
	this := EmailAction{}
	return &this
}

// GetBody returns the Body field value
func (o *EmailAction) GetBody() BodyExpression {
	if o == nil {
		var ret BodyExpression
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *EmailAction) GetBodyOk() (*BodyExpression, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *EmailAction) SetBody(v BodyExpression) {
	o.Body = v
}

// GetDelivery returns the Delivery field value
func (o *EmailAction) GetDelivery() EmailDeliveryOpts {
	if o == nil {
		var ret EmailDeliveryOpts
		return ret
	}

	return o.Delivery
}

// GetDeliveryOk returns a tuple with the Delivery field value
// and a boolean to check if the value has been set.
func (o *EmailAction) GetDeliveryOk() (*EmailDeliveryOpts, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Delivery, true
}

// SetDelivery sets field value
func (o *EmailAction) SetDelivery(v EmailDeliveryOpts) {
	o.Delivery = v
}

// GetSubject returns the Subject field value
func (o *EmailAction) GetSubject() TitleExpression {
	if o == nil {
		var ret TitleExpression
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *EmailAction) GetSubjectOk() (*TitleExpression, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *EmailAction) SetSubject(v TitleExpression) {
	o.Subject = v
}

func (o EmailAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["body"] = o.Body
	toSerialize["delivery"] = o.Delivery
	toSerialize["subject"] = o.Subject

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EmailAction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"body",
		"delivery",
		"subject",
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

	varEmailAction := _EmailAction{}

	err = json.Unmarshal(data, &varEmailAction)

	if err != nil {
		return err
	}

	*o = EmailAction(varEmailAction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "body")
		delete(additionalProperties, "delivery")
		delete(additionalProperties, "subject")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEmailAction struct {
	value *EmailAction
	isSet bool
}

func (v NullableEmailAction) Get() *EmailAction {
	return v.value
}

func (v *NullableEmailAction) Set(val *EmailAction) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailAction) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailAction(val *EmailAction) *NullableEmailAction {
	return &NullableEmailAction{value: val, isSet: true}
}

func (v NullableEmailAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


