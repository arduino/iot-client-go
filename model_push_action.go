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

// checks if the PushAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PushAction{}

// PushAction struct for PushAction
type PushAction struct {
	Body BodyExpression `json:"body"`
	Delivery PushDeliveryOpts `json:"delivery"`
	Title TitleExpression `json:"title"`
	AdditionalProperties map[string]interface{}
}

type _PushAction PushAction

// NewPushAction instantiates a new PushAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPushAction(body BodyExpression, delivery PushDeliveryOpts, title TitleExpression) *PushAction {
	this := PushAction{}
	this.Body = body
	this.Delivery = delivery
	this.Title = title
	return &this
}

// NewPushActionWithDefaults instantiates a new PushAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPushActionWithDefaults() *PushAction {
	this := PushAction{}
	return &this
}

// GetBody returns the Body field value
func (o *PushAction) GetBody() BodyExpression {
	if o == nil {
		var ret BodyExpression
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *PushAction) GetBodyOk() (*BodyExpression, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *PushAction) SetBody(v BodyExpression) {
	o.Body = v
}

// GetDelivery returns the Delivery field value
func (o *PushAction) GetDelivery() PushDeliveryOpts {
	if o == nil {
		var ret PushDeliveryOpts
		return ret
	}

	return o.Delivery
}

// GetDeliveryOk returns a tuple with the Delivery field value
// and a boolean to check if the value has been set.
func (o *PushAction) GetDeliveryOk() (*PushDeliveryOpts, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Delivery, true
}

// SetDelivery sets field value
func (o *PushAction) SetDelivery(v PushDeliveryOpts) {
	o.Delivery = v
}

// GetTitle returns the Title field value
func (o *PushAction) GetTitle() TitleExpression {
	if o == nil {
		var ret TitleExpression
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *PushAction) GetTitleOk() (*TitleExpression, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *PushAction) SetTitle(v TitleExpression) {
	o.Title = v
}

func (o PushAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PushAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["body"] = o.Body
	toSerialize["delivery"] = o.Delivery
	toSerialize["title"] = o.Title

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PushAction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"body",
		"delivery",
		"title",
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

	varPushAction := _PushAction{}

	err = json.Unmarshal(data, &varPushAction)

	if err != nil {
		return err
	}

	*o = PushAction(varPushAction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "body")
		delete(additionalProperties, "delivery")
		delete(additionalProperties, "title")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePushAction struct {
	value *PushAction
	isSet bool
}

func (v NullablePushAction) Get() *PushAction {
	return v.value
}

func (v *NullablePushAction) Set(val *PushAction) {
	v.value = val
	v.isSet = true
}

func (v NullablePushAction) IsSet() bool {
	return v.isSet
}

func (v *NullablePushAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePushAction(val *PushAction) *NullablePushAction {
	return &NullablePushAction{value: val, isSet: true}
}

func (v NullablePushAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePushAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


