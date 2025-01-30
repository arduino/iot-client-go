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

// checks if the ArduinoDevicev2Otaupload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoDevicev2Otaupload{}

// ArduinoDevicev2Otaupload ArduinoDevicev2Otaupload media type (default view)
type ArduinoDevicev2Otaupload struct {
	// SHA256 of the uploaded file
	FileSha *string `json:"file_sha,omitempty"`
	// OTA request id (only available from OTA version 2 and above)
	OtaId *string `json:"ota_id,omitempty"`
	// OTA version
	OtaVersion int64 `json:"ota_version"`
	// OTA request status (only available from OTA version 2 and above)
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ArduinoDevicev2Otaupload ArduinoDevicev2Otaupload

// NewArduinoDevicev2Otaupload instantiates a new ArduinoDevicev2Otaupload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoDevicev2Otaupload(otaVersion int64) *ArduinoDevicev2Otaupload {
	this := ArduinoDevicev2Otaupload{}
	this.OtaVersion = otaVersion
	return &this
}

// NewArduinoDevicev2OtauploadWithDefaults instantiates a new ArduinoDevicev2Otaupload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoDevicev2OtauploadWithDefaults() *ArduinoDevicev2Otaupload {
	this := ArduinoDevicev2Otaupload{}
	return &this
}

// GetFileSha returns the FileSha field value if set, zero value otherwise.
func (o *ArduinoDevicev2Otaupload) GetFileSha() string {
	if o == nil || IsNil(o.FileSha) {
		var ret string
		return ret
	}
	return *o.FileSha
}

// GetFileShaOk returns a tuple with the FileSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoDevicev2Otaupload) GetFileShaOk() (*string, bool) {
	if o == nil || IsNil(o.FileSha) {
		return nil, false
	}
	return o.FileSha, true
}

// HasFileSha returns a boolean if a field has been set.
func (o *ArduinoDevicev2Otaupload) HasFileSha() bool {
	if o != nil && !IsNil(o.FileSha) {
		return true
	}

	return false
}

// SetFileSha gets a reference to the given string and assigns it to the FileSha field.
func (o *ArduinoDevicev2Otaupload) SetFileSha(v string) {
	o.FileSha = &v
}

// GetOtaId returns the OtaId field value if set, zero value otherwise.
func (o *ArduinoDevicev2Otaupload) GetOtaId() string {
	if o == nil || IsNil(o.OtaId) {
		var ret string
		return ret
	}
	return *o.OtaId
}

// GetOtaIdOk returns a tuple with the OtaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoDevicev2Otaupload) GetOtaIdOk() (*string, bool) {
	if o == nil || IsNil(o.OtaId) {
		return nil, false
	}
	return o.OtaId, true
}

// HasOtaId returns a boolean if a field has been set.
func (o *ArduinoDevicev2Otaupload) HasOtaId() bool {
	if o != nil && !IsNil(o.OtaId) {
		return true
	}

	return false
}

// SetOtaId gets a reference to the given string and assigns it to the OtaId field.
func (o *ArduinoDevicev2Otaupload) SetOtaId(v string) {
	o.OtaId = &v
}

// GetOtaVersion returns the OtaVersion field value
func (o *ArduinoDevicev2Otaupload) GetOtaVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OtaVersion
}

// GetOtaVersionOk returns a tuple with the OtaVersion field value
// and a boolean to check if the value has been set.
func (o *ArduinoDevicev2Otaupload) GetOtaVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OtaVersion, true
}

// SetOtaVersion sets field value
func (o *ArduinoDevicev2Otaupload) SetOtaVersion(v int64) {
	o.OtaVersion = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ArduinoDevicev2Otaupload) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoDevicev2Otaupload) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ArduinoDevicev2Otaupload) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ArduinoDevicev2Otaupload) SetStatus(v string) {
	o.Status = &v
}

func (o ArduinoDevicev2Otaupload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoDevicev2Otaupload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileSha) {
		toSerialize["file_sha"] = o.FileSha
	}
	if !IsNil(o.OtaId) {
		toSerialize["ota_id"] = o.OtaId
	}
	toSerialize["ota_version"] = o.OtaVersion
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ArduinoDevicev2Otaupload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ota_version",
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

	varArduinoDevicev2Otaupload := _ArduinoDevicev2Otaupload{}

	err = json.Unmarshal(data, &varArduinoDevicev2Otaupload)

	if err != nil {
		return err
	}

	*o = ArduinoDevicev2Otaupload(varArduinoDevicev2Otaupload)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "file_sha")
		delete(additionalProperties, "ota_id")
		delete(additionalProperties, "ota_version")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableArduinoDevicev2Otaupload struct {
	value *ArduinoDevicev2Otaupload
	isSet bool
}

func (v NullableArduinoDevicev2Otaupload) Get() *ArduinoDevicev2Otaupload {
	return v.value
}

func (v *NullableArduinoDevicev2Otaupload) Set(val *ArduinoDevicev2Otaupload) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoDevicev2Otaupload) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoDevicev2Otaupload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoDevicev2Otaupload(val *ArduinoDevicev2Otaupload) *NullableArduinoDevicev2Otaupload {
	return &NullableArduinoDevicev2Otaupload{value: val, isSet: true}
}

func (v NullableArduinoDevicev2Otaupload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoDevicev2Otaupload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


