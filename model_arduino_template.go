/*
Arduino IoT Cloud API

Provides a set of endpoints to manage Arduino IoT Cloud **Devices**, **Things**, **Properties** and **Timeseries**. This API can be called just with any HTTP Client, or using one of these clients:  * [Javascript NPM package](https://www.npmjs.com/package/@arduino/arduino-iot-client)  * [Python PYPI Package](https://pypi.org/project/arduino-iot-client/)  * [Golang Module](https://github.com/arduino/iot-client-go)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ArduinoTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoTemplate{}

// ArduinoTemplate ArduinoTemplate media type (default view)
type ArduinoTemplate struct {
	Dashboards []string `json:"dashboards,omitempty"`
	// ArduinoThingresultCollection is the media type for an array of ArduinoThingresult (default view)
	Things []ArduinoThingresult `json:"things"`
	Triggers []string `json:"triggers,omitempty"`
}

type _ArduinoTemplate ArduinoTemplate

// NewArduinoTemplate instantiates a new ArduinoTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoTemplate(things []ArduinoThingresult) *ArduinoTemplate {
	this := ArduinoTemplate{}
	this.Things = things
	return &this
}

// NewArduinoTemplateWithDefaults instantiates a new ArduinoTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoTemplateWithDefaults() *ArduinoTemplate {
	this := ArduinoTemplate{}
	return &this
}

// GetDashboards returns the Dashboards field value if set, zero value otherwise.
func (o *ArduinoTemplate) GetDashboards() []string {
	if o == nil || IsNil(o.Dashboards) {
		var ret []string
		return ret
	}
	return o.Dashboards
}

// GetDashboardsOk returns a tuple with the Dashboards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoTemplate) GetDashboardsOk() ([]string, bool) {
	if o == nil || IsNil(o.Dashboards) {
		return nil, false
	}
	return o.Dashboards, true
}

// HasDashboards returns a boolean if a field has been set.
func (o *ArduinoTemplate) HasDashboards() bool {
	if o != nil && !IsNil(o.Dashboards) {
		return true
	}

	return false
}

// SetDashboards gets a reference to the given []string and assigns it to the Dashboards field.
func (o *ArduinoTemplate) SetDashboards(v []string) {
	o.Dashboards = v
}

// GetThings returns the Things field value
func (o *ArduinoTemplate) GetThings() []ArduinoThingresult {
	if o == nil {
		var ret []ArduinoThingresult
		return ret
	}

	return o.Things
}

// GetThingsOk returns a tuple with the Things field value
// and a boolean to check if the value has been set.
func (o *ArduinoTemplate) GetThingsOk() ([]ArduinoThingresult, bool) {
	if o == nil {
		return nil, false
	}
	return o.Things, true
}

// SetThings sets field value
func (o *ArduinoTemplate) SetThings(v []ArduinoThingresult) {
	o.Things = v
}

// GetTriggers returns the Triggers field value if set, zero value otherwise.
func (o *ArduinoTemplate) GetTriggers() []string {
	if o == nil || IsNil(o.Triggers) {
		var ret []string
		return ret
	}
	return o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoTemplate) GetTriggersOk() ([]string, bool) {
	if o == nil || IsNil(o.Triggers) {
		return nil, false
	}
	return o.Triggers, true
}

// HasTriggers returns a boolean if a field has been set.
func (o *ArduinoTemplate) HasTriggers() bool {
	if o != nil && !IsNil(o.Triggers) {
		return true
	}

	return false
}

// SetTriggers gets a reference to the given []string and assigns it to the Triggers field.
func (o *ArduinoTemplate) SetTriggers(v []string) {
	o.Triggers = v
}

func (o ArduinoTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dashboards) {
		toSerialize["dashboards"] = o.Dashboards
	}
	toSerialize["things"] = o.Things
	if !IsNil(o.Triggers) {
		toSerialize["triggers"] = o.Triggers
	}
	return toSerialize, nil
}

func (o *ArduinoTemplate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"things",
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

	varArduinoTemplate := _ArduinoTemplate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varArduinoTemplate)

	if err != nil {
		return err
	}

	*o = ArduinoTemplate(varArduinoTemplate)

	return err
}

type NullableArduinoTemplate struct {
	value *ArduinoTemplate
	isSet bool
}

func (v NullableArduinoTemplate) Get() *ArduinoTemplate {
	return v.value
}

func (v *NullableArduinoTemplate) Set(val *ArduinoTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoTemplate(val *ArduinoTemplate) *NullableArduinoTemplate {
	return &NullableArduinoTemplate{value: val, isSet: true}
}

func (v NullableArduinoTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


