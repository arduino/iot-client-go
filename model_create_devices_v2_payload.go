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

// checks if the CreateDevicesV2Payload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDevicesV2Payload{}

// CreateDevicesV2Payload DeviceV2 describes a device.
type CreateDevicesV2Payload struct {
	// The type of the connections selected by the user when multiple connections are available
	ConnectionType *string `json:"connection_type,omitempty"`
	// The fully qualified board name
	Fqbn *string `json:"fqbn,omitempty"`
	// The friendly name of the device
	Name *string `json:"name,omitempty" validate:"regexp=[a-zA-Z0-9_.@-]+"`
	// The serial uuid of the device
	Serial *string `json:"serial,omitempty" validate:"regexp=[a-zA-Z0-9_.@-]+"`
	// The type of the device
	Type string `json:"type"`
	// The user_id associated to the device. If absent it will be inferred from the authentication header
	UserId *string `json:"user_id,omitempty"`
	// The version of the NINA/WIFI101 firmware running on the device
	WifiFwVersion *string `json:"wifi_fw_version,omitempty" validate:"regexp=^(0|[1-9]\\\\d*)\\\\.(0|[1-9]\\\\d*)\\\\.(0|[1-9]\\\\d*)(?:-((?:0|[1-9]\\\\d*|\\\\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\\\\.(?:0|[1-9]\\\\d*|\\\\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\\\\+([0-9a-zA-Z-]+(?:\\\\.[0-9a-zA-Z-]+)*))?$"`
}

type _CreateDevicesV2Payload CreateDevicesV2Payload

// NewCreateDevicesV2Payload instantiates a new CreateDevicesV2Payload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDevicesV2Payload(type_ string) *CreateDevicesV2Payload {
	this := CreateDevicesV2Payload{}
	this.Type = type_
	return &this
}

// NewCreateDevicesV2PayloadWithDefaults instantiates a new CreateDevicesV2Payload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDevicesV2PayloadWithDefaults() *CreateDevicesV2Payload {
	this := CreateDevicesV2Payload{}
	return &this
}

// GetConnectionType returns the ConnectionType field value if set, zero value otherwise.
func (o *CreateDevicesV2Payload) GetConnectionType() string {
	if o == nil || IsNil(o.ConnectionType) {
		var ret string
		return ret
	}
	return *o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDevicesV2Payload) GetConnectionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionType) {
		return nil, false
	}
	return o.ConnectionType, true
}

// HasConnectionType returns a boolean if a field has been set.
func (o *CreateDevicesV2Payload) HasConnectionType() bool {
	if o != nil && !IsNil(o.ConnectionType) {
		return true
	}

	return false
}

// SetConnectionType gets a reference to the given string and assigns it to the ConnectionType field.
func (o *CreateDevicesV2Payload) SetConnectionType(v string) {
	o.ConnectionType = &v
}

// GetFqbn returns the Fqbn field value if set, zero value otherwise.
func (o *CreateDevicesV2Payload) GetFqbn() string {
	if o == nil || IsNil(o.Fqbn) {
		var ret string
		return ret
	}
	return *o.Fqbn
}

// GetFqbnOk returns a tuple with the Fqbn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDevicesV2Payload) GetFqbnOk() (*string, bool) {
	if o == nil || IsNil(o.Fqbn) {
		return nil, false
	}
	return o.Fqbn, true
}

// HasFqbn returns a boolean if a field has been set.
func (o *CreateDevicesV2Payload) HasFqbn() bool {
	if o != nil && !IsNil(o.Fqbn) {
		return true
	}

	return false
}

// SetFqbn gets a reference to the given string and assigns it to the Fqbn field.
func (o *CreateDevicesV2Payload) SetFqbn(v string) {
	o.Fqbn = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateDevicesV2Payload) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDevicesV2Payload) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateDevicesV2Payload) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateDevicesV2Payload) SetName(v string) {
	o.Name = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *CreateDevicesV2Payload) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDevicesV2Payload) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *CreateDevicesV2Payload) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *CreateDevicesV2Payload) SetSerial(v string) {
	o.Serial = &v
}

// GetType returns the Type field value
func (o *CreateDevicesV2Payload) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateDevicesV2Payload) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateDevicesV2Payload) SetType(v string) {
	o.Type = v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *CreateDevicesV2Payload) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDevicesV2Payload) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *CreateDevicesV2Payload) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *CreateDevicesV2Payload) SetUserId(v string) {
	o.UserId = &v
}

// GetWifiFwVersion returns the WifiFwVersion field value if set, zero value otherwise.
func (o *CreateDevicesV2Payload) GetWifiFwVersion() string {
	if o == nil || IsNil(o.WifiFwVersion) {
		var ret string
		return ret
	}
	return *o.WifiFwVersion
}

// GetWifiFwVersionOk returns a tuple with the WifiFwVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDevicesV2Payload) GetWifiFwVersionOk() (*string, bool) {
	if o == nil || IsNil(o.WifiFwVersion) {
		return nil, false
	}
	return o.WifiFwVersion, true
}

// HasWifiFwVersion returns a boolean if a field has been set.
func (o *CreateDevicesV2Payload) HasWifiFwVersion() bool {
	if o != nil && !IsNil(o.WifiFwVersion) {
		return true
	}

	return false
}

// SetWifiFwVersion gets a reference to the given string and assigns it to the WifiFwVersion field.
func (o *CreateDevicesV2Payload) SetWifiFwVersion(v string) {
	o.WifiFwVersion = &v
}

func (o CreateDevicesV2Payload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDevicesV2Payload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConnectionType) {
		toSerialize["connection_type"] = o.ConnectionType
	}
	if !IsNil(o.Fqbn) {
		toSerialize["fqbn"] = o.Fqbn
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.WifiFwVersion) {
		toSerialize["wifi_fw_version"] = o.WifiFwVersion
	}
	return toSerialize, nil
}

func (o *CreateDevicesV2Payload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varCreateDevicesV2Payload := _CreateDevicesV2Payload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateDevicesV2Payload)

	if err != nil {
		return err
	}

	*o = CreateDevicesV2Payload(varCreateDevicesV2Payload)

	return err
}

type NullableCreateDevicesV2Payload struct {
	value *CreateDevicesV2Payload
	isSet bool
}

func (v NullableCreateDevicesV2Payload) Get() *CreateDevicesV2Payload {
	return v.value
}

func (v *NullableCreateDevicesV2Payload) Set(val *CreateDevicesV2Payload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDevicesV2Payload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDevicesV2Payload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDevicesV2Payload(val *CreateDevicesV2Payload) *NullableCreateDevicesV2Payload {
	return &NullableCreateDevicesV2Payload{value: val, isSet: true}
}

func (v NullableCreateDevicesV2Payload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDevicesV2Payload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


