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

// checks if the ArduinoThingtemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoThingtemplate{}

// ArduinoThingtemplate ArduinoThingtemplate media type (default view)
type ArduinoThingtemplate struct {
	DeviceMetadata *ArduinoDevicev2templatedevice `json:"device_metadata,omitempty"`
	// The friendly id of the thing
	Id *string `json:"id,omitempty"`
	// The friendly name of the thing
	Name string `json:"name"`
	// Id of the organization the thing belongs to
	OrganizationId *string `json:"organization_id,omitempty"`
	// The ID of the template's sketch
	SketchTemplate *string `json:"sketch_template,omitempty"`
	// Tags of the thing
	Tags []Tag `json:"tags,omitempty"`
	// Time zone of the thing
	Timezone string `json:"timezone"`
	// ArduinoTemplatepropertyCollection is the media type for an array of ArduinoTemplateproperty (default view)
	Variables []ArduinoTemplateproperty `json:"variables,omitempty"`
	// Webhook uri
	WebhookUri *string `json:"webhook_uri,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ArduinoThingtemplate ArduinoThingtemplate

// NewArduinoThingtemplate instantiates a new ArduinoThingtemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoThingtemplate(name string, timezone string) *ArduinoThingtemplate {
	this := ArduinoThingtemplate{}
	this.Name = name
	this.Timezone = timezone
	return &this
}

// NewArduinoThingtemplateWithDefaults instantiates a new ArduinoThingtemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoThingtemplateWithDefaults() *ArduinoThingtemplate {
	this := ArduinoThingtemplate{}
	return &this
}

// GetDeviceMetadata returns the DeviceMetadata field value if set, zero value otherwise.
func (o *ArduinoThingtemplate) GetDeviceMetadata() ArduinoDevicev2templatedevice {
	if o == nil || IsNil(o.DeviceMetadata) {
		var ret ArduinoDevicev2templatedevice
		return ret
	}
	return *o.DeviceMetadata
}

// GetDeviceMetadataOk returns a tuple with the DeviceMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoThingtemplate) GetDeviceMetadataOk() (*ArduinoDevicev2templatedevice, bool) {
	if o == nil || IsNil(o.DeviceMetadata) {
		return nil, false
	}
	return o.DeviceMetadata, true
}

// HasDeviceMetadata returns a boolean if a field has been set.
func (o *ArduinoThingtemplate) HasDeviceMetadata() bool {
	if o != nil && !IsNil(o.DeviceMetadata) {
		return true
	}

	return false
}

// SetDeviceMetadata gets a reference to the given ArduinoDevicev2templatedevice and assigns it to the DeviceMetadata field.
func (o *ArduinoThingtemplate) SetDeviceMetadata(v ArduinoDevicev2templatedevice) {
	o.DeviceMetadata = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ArduinoThingtemplate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoThingtemplate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ArduinoThingtemplate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ArduinoThingtemplate) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *ArduinoThingtemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ArduinoThingtemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ArduinoThingtemplate) SetName(v string) {
	o.Name = v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *ArduinoThingtemplate) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoThingtemplate) GetOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *ArduinoThingtemplate) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *ArduinoThingtemplate) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetSketchTemplate returns the SketchTemplate field value if set, zero value otherwise.
func (o *ArduinoThingtemplate) GetSketchTemplate() string {
	if o == nil || IsNil(o.SketchTemplate) {
		var ret string
		return ret
	}
	return *o.SketchTemplate
}

// GetSketchTemplateOk returns a tuple with the SketchTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoThingtemplate) GetSketchTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.SketchTemplate) {
		return nil, false
	}
	return o.SketchTemplate, true
}

// HasSketchTemplate returns a boolean if a field has been set.
func (o *ArduinoThingtemplate) HasSketchTemplate() bool {
	if o != nil && !IsNil(o.SketchTemplate) {
		return true
	}

	return false
}

// SetSketchTemplate gets a reference to the given string and assigns it to the SketchTemplate field.
func (o *ArduinoThingtemplate) SetSketchTemplate(v string) {
	o.SketchTemplate = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ArduinoThingtemplate) GetTags() []Tag {
	if o == nil || IsNil(o.Tags) {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoThingtemplate) GetTagsOk() ([]Tag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ArduinoThingtemplate) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *ArduinoThingtemplate) SetTags(v []Tag) {
	o.Tags = v
}

// GetTimezone returns the Timezone field value
func (o *ArduinoThingtemplate) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *ArduinoThingtemplate) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value
func (o *ArduinoThingtemplate) SetTimezone(v string) {
	o.Timezone = v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *ArduinoThingtemplate) GetVariables() []ArduinoTemplateproperty {
	if o == nil || IsNil(o.Variables) {
		var ret []ArduinoTemplateproperty
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoThingtemplate) GetVariablesOk() ([]ArduinoTemplateproperty, bool) {
	if o == nil || IsNil(o.Variables) {
		return nil, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *ArduinoThingtemplate) HasVariables() bool {
	if o != nil && !IsNil(o.Variables) {
		return true
	}

	return false
}

// SetVariables gets a reference to the given []ArduinoTemplateproperty and assigns it to the Variables field.
func (o *ArduinoThingtemplate) SetVariables(v []ArduinoTemplateproperty) {
	o.Variables = v
}

// GetWebhookUri returns the WebhookUri field value if set, zero value otherwise.
func (o *ArduinoThingtemplate) GetWebhookUri() string {
	if o == nil || IsNil(o.WebhookUri) {
		var ret string
		return ret
	}
	return *o.WebhookUri
}

// GetWebhookUriOk returns a tuple with the WebhookUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoThingtemplate) GetWebhookUriOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookUri) {
		return nil, false
	}
	return o.WebhookUri, true
}

// HasWebhookUri returns a boolean if a field has been set.
func (o *ArduinoThingtemplate) HasWebhookUri() bool {
	if o != nil && !IsNil(o.WebhookUri) {
		return true
	}

	return false
}

// SetWebhookUri gets a reference to the given string and assigns it to the WebhookUri field.
func (o *ArduinoThingtemplate) SetWebhookUri(v string) {
	o.WebhookUri = &v
}

func (o ArduinoThingtemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoThingtemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceMetadata) {
		toSerialize["device_metadata"] = o.DeviceMetadata
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.OrganizationId) {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if !IsNil(o.SketchTemplate) {
		toSerialize["sketch_template"] = o.SketchTemplate
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["timezone"] = o.Timezone
	if !IsNil(o.Variables) {
		toSerialize["variables"] = o.Variables
	}
	if !IsNil(o.WebhookUri) {
		toSerialize["webhook_uri"] = o.WebhookUri
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ArduinoThingtemplate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"timezone",
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

	varArduinoThingtemplate := _ArduinoThingtemplate{}

	err = json.Unmarshal(data, &varArduinoThingtemplate)

	if err != nil {
		return err
	}

	*o = ArduinoThingtemplate(varArduinoThingtemplate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device_metadata")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "organization_id")
		delete(additionalProperties, "sketch_template")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "timezone")
		delete(additionalProperties, "variables")
		delete(additionalProperties, "webhook_uri")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableArduinoThingtemplate struct {
	value *ArduinoThingtemplate
	isSet bool
}

func (v NullableArduinoThingtemplate) Get() *ArduinoThingtemplate {
	return v.value
}

func (v *NullableArduinoThingtemplate) Set(val *ArduinoThingtemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoThingtemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoThingtemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoThingtemplate(val *ArduinoThingtemplate) *NullableArduinoThingtemplate {
	return &NullableArduinoThingtemplate{value: val, isSet: true}
}

func (v NullableArduinoThingtemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoThingtemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


