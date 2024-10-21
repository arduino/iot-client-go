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

// checks if the ArduinoLoradevicev1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoLoradevicev1{}

// ArduinoLoradevicev1 ArduinoLoradevicev1 media type (default view)
type ArduinoLoradevicev1 struct {
	// The eui of the app
	AppEui string `json:"app_eui"`
	// The key of the device
	AppKey string `json:"app_key"`
	// The id of the device
	DeviceId string `json:"device_id"`
	// The eui of the lora device
	Eui string `json:"eui"`
}

type _ArduinoLoradevicev1 ArduinoLoradevicev1

// NewArduinoLoradevicev1 instantiates a new ArduinoLoradevicev1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoLoradevicev1(appEui string, appKey string, deviceId string, eui string) *ArduinoLoradevicev1 {
	this := ArduinoLoradevicev1{}
	this.AppEui = appEui
	this.AppKey = appKey
	this.DeviceId = deviceId
	this.Eui = eui
	return &this
}

// NewArduinoLoradevicev1WithDefaults instantiates a new ArduinoLoradevicev1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoLoradevicev1WithDefaults() *ArduinoLoradevicev1 {
	this := ArduinoLoradevicev1{}
	return &this
}

// GetAppEui returns the AppEui field value
func (o *ArduinoLoradevicev1) GetAppEui() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppEui
}

// GetAppEuiOk returns a tuple with the AppEui field value
// and a boolean to check if the value has been set.
func (o *ArduinoLoradevicev1) GetAppEuiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppEui, true
}

// SetAppEui sets field value
func (o *ArduinoLoradevicev1) SetAppEui(v string) {
	o.AppEui = v
}

// GetAppKey returns the AppKey field value
func (o *ArduinoLoradevicev1) GetAppKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppKey
}

// GetAppKeyOk returns a tuple with the AppKey field value
// and a boolean to check if the value has been set.
func (o *ArduinoLoradevicev1) GetAppKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppKey, true
}

// SetAppKey sets field value
func (o *ArduinoLoradevicev1) SetAppKey(v string) {
	o.AppKey = v
}

// GetDeviceId returns the DeviceId field value
func (o *ArduinoLoradevicev1) GetDeviceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value
// and a boolean to check if the value has been set.
func (o *ArduinoLoradevicev1) GetDeviceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceId, true
}

// SetDeviceId sets field value
func (o *ArduinoLoradevicev1) SetDeviceId(v string) {
	o.DeviceId = v
}

// GetEui returns the Eui field value
func (o *ArduinoLoradevicev1) GetEui() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Eui
}

// GetEuiOk returns a tuple with the Eui field value
// and a boolean to check if the value has been set.
func (o *ArduinoLoradevicev1) GetEuiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Eui, true
}

// SetEui sets field value
func (o *ArduinoLoradevicev1) SetEui(v string) {
	o.Eui = v
}

func (o ArduinoLoradevicev1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoLoradevicev1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["app_eui"] = o.AppEui
	toSerialize["app_key"] = o.AppKey
	toSerialize["device_id"] = o.DeviceId
	toSerialize["eui"] = o.Eui
	return toSerialize, nil
}

func (o *ArduinoLoradevicev1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"app_eui",
		"app_key",
		"device_id",
		"eui",
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

	varArduinoLoradevicev1 := _ArduinoLoradevicev1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varArduinoLoradevicev1)

	if err != nil {
		return err
	}

	*o = ArduinoLoradevicev1(varArduinoLoradevicev1)

	return err
}

type NullableArduinoLoradevicev1 struct {
	value *ArduinoLoradevicev1
	isSet bool
}

func (v NullableArduinoLoradevicev1) Get() *ArduinoLoradevicev1 {
	return v.value
}

func (v *NullableArduinoLoradevicev1) Set(val *ArduinoLoradevicev1) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoLoradevicev1) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoLoradevicev1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoLoradevicev1(val *ArduinoLoradevicev1) *NullableArduinoLoradevicev1 {
	return &NullableArduinoLoradevicev1{value: val, isSet: true}
}

func (v NullableArduinoLoradevicev1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoLoradevicev1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


