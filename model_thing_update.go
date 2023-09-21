/*
Arduino IoT Cloud API

 Provides a set of endpoints to manage Arduino IoT Cloud **Devices**, **Things**, **Properties** and **Timeseries**. This API can be called just with any HTTP Client, or using one of these clients:  * [Javascript NPM package](https://www.npmjs.com/package/@arduino/arduino-iot-client)  * [Python PYPI Package](https://pypi.org/project/arduino-iot-client/)  * [Golang Module](https://github.com/arduino/iot-client-go)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iot

import (
	"encoding/json"
)

// checks if the ThingUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThingUpdate{}

// ThingUpdate Payload to update an existing thing
type ThingUpdate struct {
	// The arn of the associated device
	DeviceId *string `json:"device_id,omitempty"`
	// The id of the thing
	Id *string `json:"id,omitempty"`
	// The friendly name of the thing
	Name *string `json:"name,omitempty"`
	// The properties of the thing
	Properties []Property `json:"properties,omitempty"`
	// A time zone name. Check /v2/timezones for a list of valid names.
	Timezone *string `json:"timezone,omitempty"`
	// Webhook uri
	WebhookActive *bool `json:"webhook_active,omitempty"`
	// Webhook uri
	WebhookUri *string `json:"webhook_uri,omitempty"`
}

// NewThingUpdate instantiates a new ThingUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThingUpdate() *ThingUpdate {
	this := ThingUpdate{}
	return &this
}

// NewThingUpdateWithDefaults instantiates a new ThingUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThingUpdateWithDefaults() *ThingUpdate {
	this := ThingUpdate{}
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *ThingUpdate) GetDeviceId() string {
	if o == nil || IsNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingUpdate) GetDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *ThingUpdate) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *ThingUpdate) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ThingUpdate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingUpdate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ThingUpdate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ThingUpdate) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ThingUpdate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingUpdate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ThingUpdate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ThingUpdate) SetName(v string) {
	o.Name = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *ThingUpdate) GetProperties() []Property {
	if o == nil || IsNil(o.Properties) {
		var ret []Property
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingUpdate) GetPropertiesOk() ([]Property, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *ThingUpdate) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []Property and assigns it to the Properties field.
func (o *ThingUpdate) SetProperties(v []Property) {
	o.Properties = v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *ThingUpdate) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingUpdate) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *ThingUpdate) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *ThingUpdate) SetTimezone(v string) {
	o.Timezone = &v
}

// GetWebhookActive returns the WebhookActive field value if set, zero value otherwise.
func (o *ThingUpdate) GetWebhookActive() bool {
	if o == nil || IsNil(o.WebhookActive) {
		var ret bool
		return ret
	}
	return *o.WebhookActive
}

// GetWebhookActiveOk returns a tuple with the WebhookActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingUpdate) GetWebhookActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.WebhookActive) {
		return nil, false
	}
	return o.WebhookActive, true
}

// HasWebhookActive returns a boolean if a field has been set.
func (o *ThingUpdate) HasWebhookActive() bool {
	if o != nil && !IsNil(o.WebhookActive) {
		return true
	}

	return false
}

// SetWebhookActive gets a reference to the given bool and assigns it to the WebhookActive field.
func (o *ThingUpdate) SetWebhookActive(v bool) {
	o.WebhookActive = &v
}

// GetWebhookUri returns the WebhookUri field value if set, zero value otherwise.
func (o *ThingUpdate) GetWebhookUri() string {
	if o == nil || IsNil(o.WebhookUri) {
		var ret string
		return ret
	}
	return *o.WebhookUri
}

// GetWebhookUriOk returns a tuple with the WebhookUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingUpdate) GetWebhookUriOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookUri) {
		return nil, false
	}
	return o.WebhookUri, true
}

// HasWebhookUri returns a boolean if a field has been set.
func (o *ThingUpdate) HasWebhookUri() bool {
	if o != nil && !IsNil(o.WebhookUri) {
		return true
	}

	return false
}

// SetWebhookUri gets a reference to the given string and assigns it to the WebhookUri field.
func (o *ThingUpdate) SetWebhookUri(v string) {
	o.WebhookUri = &v
}

func (o ThingUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThingUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceId) {
		toSerialize["device_id"] = o.DeviceId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if !IsNil(o.WebhookActive) {
		toSerialize["webhook_active"] = o.WebhookActive
	}
	if !IsNil(o.WebhookUri) {
		toSerialize["webhook_uri"] = o.WebhookUri
	}
	return toSerialize, nil
}

type NullableThingUpdate struct {
	value *ThingUpdate
	isSet bool
}

func (v NullableThingUpdate) Get() *ThingUpdate {
	return v.value
}

func (v *NullableThingUpdate) Set(val *ThingUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableThingUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableThingUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThingUpdate(val *ThingUpdate) *NullableThingUpdate {
	return &NullableThingUpdate{value: val, isSet: true}
}

func (v NullableThingUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThingUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


