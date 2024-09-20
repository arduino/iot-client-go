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

// checks if the ArduinoDevicev2EventProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoDevicev2EventProperties{}

// ArduinoDevicev2EventProperties ArduinoDevicev2EventProperties media type (default view)
type ArduinoDevicev2EventProperties struct {
	// ArduinoDevicev2SimplePropertiesCollection is the media type for an array of ArduinoDevicev2SimpleProperties (default view)
	Events []ArduinoDevicev2SimpleProperties `json:"events"`
	// The device of the property
	Id string `json:"id"`
}

// NewArduinoDevicev2EventProperties instantiates a new ArduinoDevicev2EventProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoDevicev2EventProperties(events []ArduinoDevicev2SimpleProperties, id string) *ArduinoDevicev2EventProperties {
	this := ArduinoDevicev2EventProperties{}
	this.Events = events
	this.Id = id
	return &this
}

// NewArduinoDevicev2EventPropertiesWithDefaults instantiates a new ArduinoDevicev2EventProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoDevicev2EventPropertiesWithDefaults() *ArduinoDevicev2EventProperties {
	this := ArduinoDevicev2EventProperties{}
	return &this
}

// GetEvents returns the Events field value
func (o *ArduinoDevicev2EventProperties) GetEvents() []ArduinoDevicev2SimpleProperties {
	if o == nil {
		var ret []ArduinoDevicev2SimpleProperties
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *ArduinoDevicev2EventProperties) GetEventsOk() ([]ArduinoDevicev2SimpleProperties, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *ArduinoDevicev2EventProperties) SetEvents(v []ArduinoDevicev2SimpleProperties) {
	o.Events = v
}

// GetId returns the Id field value
func (o *ArduinoDevicev2EventProperties) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ArduinoDevicev2EventProperties) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ArduinoDevicev2EventProperties) SetId(v string) {
	o.Id = v
}

func (o ArduinoDevicev2EventProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoDevicev2EventProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["events"] = o.Events
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableArduinoDevicev2EventProperties struct {
	value *ArduinoDevicev2EventProperties
	isSet bool
}

func (v NullableArduinoDevicev2EventProperties) Get() *ArduinoDevicev2EventProperties {
	return v.value
}

func (v *NullableArduinoDevicev2EventProperties) Set(val *ArduinoDevicev2EventProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoDevicev2EventProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoDevicev2EventProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoDevicev2EventProperties(val *ArduinoDevicev2EventProperties) *NullableArduinoDevicev2EventProperties {
	return &NullableArduinoDevicev2EventProperties{value: val, isSet: true}
}

func (v NullableArduinoDevicev2EventProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoDevicev2EventProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


