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

// checks if the ArduinoDevicev2Webhook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoDevicev2Webhook{}

// ArduinoDevicev2Webhook DeviceWebhookV2 describes a webhook associated to the device (default view)
type ArduinoDevicev2Webhook struct {
	// Whether the webhook is active
	Active *bool `json:"active,omitempty"`
	// The uuid of the webhook
	Id string `json:"id"`
	// The uri of the webhook
	Uri string `json:"uri"`
}

// NewArduinoDevicev2Webhook instantiates a new ArduinoDevicev2Webhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoDevicev2Webhook(id string, uri string) *ArduinoDevicev2Webhook {
	this := ArduinoDevicev2Webhook{}
	var active bool = true
	this.Active = &active
	this.Id = id
	this.Uri = uri
	return &this
}

// NewArduinoDevicev2WebhookWithDefaults instantiates a new ArduinoDevicev2Webhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoDevicev2WebhookWithDefaults() *ArduinoDevicev2Webhook {
	this := ArduinoDevicev2Webhook{}
	var active bool = true
	this.Active = &active
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ArduinoDevicev2Webhook) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoDevicev2Webhook) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ArduinoDevicev2Webhook) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ArduinoDevicev2Webhook) SetActive(v bool) {
	o.Active = &v
}

// GetId returns the Id field value
func (o *ArduinoDevicev2Webhook) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ArduinoDevicev2Webhook) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ArduinoDevicev2Webhook) SetId(v string) {
	o.Id = v
}

// GetUri returns the Uri field value
func (o *ArduinoDevicev2Webhook) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *ArduinoDevicev2Webhook) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *ArduinoDevicev2Webhook) SetUri(v string) {
	o.Uri = v
}

func (o ArduinoDevicev2Webhook) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoDevicev2Webhook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	toSerialize["id"] = o.Id
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableArduinoDevicev2Webhook struct {
	value *ArduinoDevicev2Webhook
	isSet bool
}

func (v NullableArduinoDevicev2Webhook) Get() *ArduinoDevicev2Webhook {
	return v.value
}

func (v *NullableArduinoDevicev2Webhook) Set(val *ArduinoDevicev2Webhook) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoDevicev2Webhook) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoDevicev2Webhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoDevicev2Webhook(val *ArduinoDevicev2Webhook) *NullableArduinoDevicev2Webhook {
	return &NullableArduinoDevicev2Webhook{value: val, isSet: true}
}

func (v NullableArduinoDevicev2Webhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoDevicev2Webhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


