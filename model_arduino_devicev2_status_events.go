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

// checks if the ArduinoDevicev2StatusEvents type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoDevicev2StatusEvents{}

// ArduinoDevicev2StatusEvents ArduinoDevicev2StatusEvents media type (default view)
type ArduinoDevicev2StatusEvents struct {
	// ArduinoDevicev2StatusEventCollection is the media type for an array of ArduinoDevicev2StatusEvent (default view)
	Events []ArduinoDevicev2StatusEvent `json:"events"`
	// The id of the device
	Id string `json:"id"`
}

type _ArduinoDevicev2StatusEvents ArduinoDevicev2StatusEvents

// NewArduinoDevicev2StatusEvents instantiates a new ArduinoDevicev2StatusEvents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoDevicev2StatusEvents(events []ArduinoDevicev2StatusEvent, id string) *ArduinoDevicev2StatusEvents {
	this := ArduinoDevicev2StatusEvents{}
	this.Events = events
	this.Id = id
	return &this
}

// NewArduinoDevicev2StatusEventsWithDefaults instantiates a new ArduinoDevicev2StatusEvents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoDevicev2StatusEventsWithDefaults() *ArduinoDevicev2StatusEvents {
	this := ArduinoDevicev2StatusEvents{}
	return &this
}

// GetEvents returns the Events field value
func (o *ArduinoDevicev2StatusEvents) GetEvents() []ArduinoDevicev2StatusEvent {
	if o == nil {
		var ret []ArduinoDevicev2StatusEvent
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *ArduinoDevicev2StatusEvents) GetEventsOk() ([]ArduinoDevicev2StatusEvent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *ArduinoDevicev2StatusEvents) SetEvents(v []ArduinoDevicev2StatusEvent) {
	o.Events = v
}

// GetId returns the Id field value
func (o *ArduinoDevicev2StatusEvents) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ArduinoDevicev2StatusEvents) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ArduinoDevicev2StatusEvents) SetId(v string) {
	o.Id = v
}

func (o ArduinoDevicev2StatusEvents) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoDevicev2StatusEvents) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["events"] = o.Events
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *ArduinoDevicev2StatusEvents) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"events",
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

	varArduinoDevicev2StatusEvents := _ArduinoDevicev2StatusEvents{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varArduinoDevicev2StatusEvents)

	if err != nil {
		return err
	}

	*o = ArduinoDevicev2StatusEvents(varArduinoDevicev2StatusEvents)

	return err
}

type NullableArduinoDevicev2StatusEvents struct {
	value *ArduinoDevicev2StatusEvents
	isSet bool
}

func (v NullableArduinoDevicev2StatusEvents) Get() *ArduinoDevicev2StatusEvents {
	return v.value
}

func (v *NullableArduinoDevicev2StatusEvents) Set(val *ArduinoDevicev2StatusEvents) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoDevicev2StatusEvents) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoDevicev2StatusEvents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoDevicev2StatusEvents(val *ArduinoDevicev2StatusEvents) *NullableArduinoDevicev2StatusEvents {
	return &NullableArduinoDevicev2StatusEvents{value: val, isSet: true}
}

func (v NullableArduinoDevicev2StatusEvents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoDevicev2StatusEvents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


