/*
Arduino IoT Cloud API

Provides a set of endpoints to manage Arduino IoT Cloud **Devices**, **Things**, **Properties** and **Timeseries**. This API can be called just with any HTTP Client, or using one of these clients:  * [Javascript NPM package](https://www.npmjs.com/package/@arduino/arduino-iot-client)  * [Python PYPI Package](https://pypi.org/project/arduino-iot-client/)  * [Golang Module](https://github.com/arduino/iot-client-go)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
)

// checks if the ThingCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThingCreate{}

// ThingCreate Payload to create a new thing
type ThingCreate struct {
	// The kind of voice assistant the thing is connected to, it can be ALEXA | GOOGLE | NONE
	Assistant *string `json:"assistant,omitempty"`
	// The arn of the associated device
	DeviceId *string `json:"device_id,omitempty"`
	// The id of the thing
	Id *string `json:"id,omitempty"`
	// The friendly name of the thing
	Name *string `json:"name,omitempty" validate:"regexp=^[a-zA-Z0-9_. -]+$"`
	// The properties of the thing
	Properties []Property `json:"properties,omitempty"`
	// Optional set of tags
	Tags []Tag `json:"tags,omitempty"`
	// A time zone name Check /v2/timezones for a list of valid names.
	Timezone *string `json:"timezone,omitempty"`
	// Webhook uri
	WebhookActive *bool `json:"webhook_active,omitempty"`
	// Webhook uri
	WebhookUri *string `json:"webhook_uri,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ThingCreate ThingCreate

// NewThingCreate instantiates a new ThingCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThingCreate() *ThingCreate {
	this := ThingCreate{}
	var timezone string = "America/New_York"
	this.Timezone = &timezone
	return &this
}

// NewThingCreateWithDefaults instantiates a new ThingCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThingCreateWithDefaults() *ThingCreate {
	this := ThingCreate{}
	var timezone string = "America/New_York"
	this.Timezone = &timezone
	return &this
}

// GetAssistant returns the Assistant field value if set, zero value otherwise.
func (o *ThingCreate) GetAssistant() string {
	if o == nil || IsNil(o.Assistant) {
		var ret string
		return ret
	}
	return *o.Assistant
}

// GetAssistantOk returns a tuple with the Assistant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingCreate) GetAssistantOk() (*string, bool) {
	if o == nil || IsNil(o.Assistant) {
		return nil, false
	}
	return o.Assistant, true
}

// HasAssistant returns a boolean if a field has been set.
func (o *ThingCreate) HasAssistant() bool {
	if o != nil && !IsNil(o.Assistant) {
		return true
	}

	return false
}

// SetAssistant gets a reference to the given string and assigns it to the Assistant field.
func (o *ThingCreate) SetAssistant(v string) {
	o.Assistant = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *ThingCreate) GetDeviceId() string {
	if o == nil || IsNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingCreate) GetDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *ThingCreate) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *ThingCreate) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ThingCreate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingCreate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ThingCreate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ThingCreate) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ThingCreate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingCreate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ThingCreate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ThingCreate) SetName(v string) {
	o.Name = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *ThingCreate) GetProperties() []Property {
	if o == nil || IsNil(o.Properties) {
		var ret []Property
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingCreate) GetPropertiesOk() ([]Property, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *ThingCreate) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []Property and assigns it to the Properties field.
func (o *ThingCreate) SetProperties(v []Property) {
	o.Properties = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ThingCreate) GetTags() []Tag {
	if o == nil || IsNil(o.Tags) {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingCreate) GetTagsOk() ([]Tag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ThingCreate) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *ThingCreate) SetTags(v []Tag) {
	o.Tags = v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *ThingCreate) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingCreate) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *ThingCreate) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *ThingCreate) SetTimezone(v string) {
	o.Timezone = &v
}

// GetWebhookActive returns the WebhookActive field value if set, zero value otherwise.
func (o *ThingCreate) GetWebhookActive() bool {
	if o == nil || IsNil(o.WebhookActive) {
		var ret bool
		return ret
	}
	return *o.WebhookActive
}

// GetWebhookActiveOk returns a tuple with the WebhookActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingCreate) GetWebhookActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.WebhookActive) {
		return nil, false
	}
	return o.WebhookActive, true
}

// HasWebhookActive returns a boolean if a field has been set.
func (o *ThingCreate) HasWebhookActive() bool {
	if o != nil && !IsNil(o.WebhookActive) {
		return true
	}

	return false
}

// SetWebhookActive gets a reference to the given bool and assigns it to the WebhookActive field.
func (o *ThingCreate) SetWebhookActive(v bool) {
	o.WebhookActive = &v
}

// GetWebhookUri returns the WebhookUri field value if set, zero value otherwise.
func (o *ThingCreate) GetWebhookUri() string {
	if o == nil || IsNil(o.WebhookUri) {
		var ret string
		return ret
	}
	return *o.WebhookUri
}

// GetWebhookUriOk returns a tuple with the WebhookUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingCreate) GetWebhookUriOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookUri) {
		return nil, false
	}
	return o.WebhookUri, true
}

// HasWebhookUri returns a boolean if a field has been set.
func (o *ThingCreate) HasWebhookUri() bool {
	if o != nil && !IsNil(o.WebhookUri) {
		return true
	}

	return false
}

// SetWebhookUri gets a reference to the given string and assigns it to the WebhookUri field.
func (o *ThingCreate) SetWebhookUri(v string) {
	o.WebhookUri = &v
}

func (o ThingCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThingCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assistant) {
		toSerialize["assistant"] = o.Assistant
	}
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
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ThingCreate) UnmarshalJSON(data []byte) (err error) {
	varThingCreate := _ThingCreate{}

	err = json.Unmarshal(data, &varThingCreate)

	if err != nil {
		return err
	}

	*o = ThingCreate(varThingCreate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "assistant")
		delete(additionalProperties, "device_id")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "properties")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "timezone")
		delete(additionalProperties, "webhook_active")
		delete(additionalProperties, "webhook_uri")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableThingCreate struct {
	value *ThingCreate
	isSet bool
}

func (v NullableThingCreate) Get() *ThingCreate {
	return v.value
}

func (v *NullableThingCreate) Set(val *ThingCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableThingCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableThingCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThingCreate(val *ThingCreate) *NullableThingCreate {
	return &NullableThingCreate{value: val, isSet: true}
}

func (v NullableThingCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThingCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


