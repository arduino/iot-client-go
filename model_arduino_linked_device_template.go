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

// checks if the ArduinoLinkedDeviceTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoLinkedDeviceTemplate{}

// ArduinoLinkedDeviceTemplate ArduinoLinked_device_template media type (default view)
type ArduinoLinkedDeviceTemplate struct {
	// The thing the device is associated to
	ThingId string `json:"thing_id"`
}

type _ArduinoLinkedDeviceTemplate ArduinoLinkedDeviceTemplate

// NewArduinoLinkedDeviceTemplate instantiates a new ArduinoLinkedDeviceTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoLinkedDeviceTemplate(thingId string) *ArduinoLinkedDeviceTemplate {
	this := ArduinoLinkedDeviceTemplate{}
	this.ThingId = thingId
	return &this
}

// NewArduinoLinkedDeviceTemplateWithDefaults instantiates a new ArduinoLinkedDeviceTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoLinkedDeviceTemplateWithDefaults() *ArduinoLinkedDeviceTemplate {
	this := ArduinoLinkedDeviceTemplate{}
	return &this
}

// GetThingId returns the ThingId field value
func (o *ArduinoLinkedDeviceTemplate) GetThingId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThingId
}

// GetThingIdOk returns a tuple with the ThingId field value
// and a boolean to check if the value has been set.
func (o *ArduinoLinkedDeviceTemplate) GetThingIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThingId, true
}

// SetThingId sets field value
func (o *ArduinoLinkedDeviceTemplate) SetThingId(v string) {
	o.ThingId = v
}

func (o ArduinoLinkedDeviceTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoLinkedDeviceTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["thing_id"] = o.ThingId
	return toSerialize, nil
}

func (o *ArduinoLinkedDeviceTemplate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"thing_id",
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

	varArduinoLinkedDeviceTemplate := _ArduinoLinkedDeviceTemplate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varArduinoLinkedDeviceTemplate)

	if err != nil {
		return err
	}

	*o = ArduinoLinkedDeviceTemplate(varArduinoLinkedDeviceTemplate)

	return err
}

type NullableArduinoLinkedDeviceTemplate struct {
	value *ArduinoLinkedDeviceTemplate
	isSet bool
}

func (v NullableArduinoLinkedDeviceTemplate) Get() *ArduinoLinkedDeviceTemplate {
	return v.value
}

func (v *NullableArduinoLinkedDeviceTemplate) Set(val *ArduinoLinkedDeviceTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoLinkedDeviceTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoLinkedDeviceTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoLinkedDeviceTemplate(val *ArduinoLinkedDeviceTemplate) *NullableArduinoLinkedDeviceTemplate {
	return &NullableArduinoLinkedDeviceTemplate{value: val, isSet: true}
}

func (v NullableArduinoLinkedDeviceTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoLinkedDeviceTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


