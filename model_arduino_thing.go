/*
Arduino IoT Cloud API

Provides a set of endpoints to manage Arduino IoT Cloud **Devices**, **Things**, **Properties** and **Timeseries**. This API can be called just with any HTTP Client, or using one of these clients:  * [Javascript NPM package](https://www.npmjs.com/package/@arduino/arduino-iot-client)  * [Python PYPI Package](https://pypi.org/project/arduino-iot-client/)  * [Golang Module](https://github.com/arduino/iot-client-go)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the ArduinoThing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoThing{}

// ArduinoThing ArduinoThing media type (default view)
type ArduinoThing struct {
	// The kind of voice assistant the thing is connected to, it can be ALEXA | GOOGLE | NONE
	Assistant *string `json:"assistant,omitempty"`
	// Creation date of the thing
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Delete date of the thing
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// The fqbn of the attached device, if any
	DeviceFqbn *string `json:"device_fqbn,omitempty"`
	// The id of the device
	DeviceId *string `json:"device_id,omitempty"`
	// The name of the attached device, if any
	DeviceName *string `json:"device_name,omitempty"`
	// The type of the attached device, if any
	DeviceType *string `json:"device_type,omitempty"`
	// The api reference of this thing
	Href string `json:"href"`
	// The id of the thing
	Id string `json:"id"`
	// The friendly name of the thing
	Name string `json:"name"`
	// Id of the organization the thing belongs to
	OrganizationId *string `json:"organization_id,omitempty"`
	// ArduinoPropertyCollection is the media type for an array of ArduinoProperty (default view)
	Properties []ArduinoProperty `json:"properties,omitempty"`
	// The number of properties of the thing
	PropertiesCount *int64 `json:"properties_count,omitempty"`
	// The id of the attached sketch
	SketchId *string `json:"sketch_id,omitempty"`
	// Tags of the thing
	Tags map[string]interface{} `json:"tags,omitempty"`
	// Time zone of the thing
	Timezone string `json:"timezone"`
	// Update date of the thing
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The user id of the owner
	UserId string `json:"user_id"`
	// Webhook uri
	WebhookActive *bool `json:"webhook_active,omitempty"`
	// Webhook uri
	WebhookUri *string `json:"webhook_uri,omitempty"`
}

type _ArduinoThing ArduinoThing

// NewArduinoThing instantiates a new ArduinoThing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoThing(href string, id string, name string, timezone string, userId string) *ArduinoThing {
	this := ArduinoThing{}
	this.Href = href
	this.Id = id
	this.Name = name
	this.Timezone = timezone
	this.UserId = userId
	return &this
}

// NewArduinoThingWithDefaults instantiates a new ArduinoThing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoThingWithDefaults() *ArduinoThing {
	this := ArduinoThing{}
	return &this
}

// GetAssistant returns the Assistant field value if set, zero value otherwise.
func (o *ArduinoThing) GetAssistant() string {
	if o == nil || IsNil(o.Assistant) {
		var ret string
		return ret
	}
	return *o.Assistant
}

// GetAssistantOk returns a tuple with the Assistant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoThing) GetAssistantOk() (*string, bool) {
	if o == nil || IsNil(o.Assistant) {
		return nil, false
	}
	return o.Assistant, true
}

// HasAssistant returns a boolean if a field has been set.
func (o *ArduinoThing) HasAssistant() bool {
	if o != nil && !IsNil(o.Assistant) {
		return true
	}

	return false
}

// SetAssistant gets a reference to the given string and assigns it to the Assistant field.
func (o *ArduinoThing) SetAssistant(v string) {
	o.Assistant = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ArduinoThing) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoThing) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ArduinoThing) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ArduinoThing) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *ArduinoThing) GetDeletedAt() time.Time {
	if o == nil || IsNil(o.DeletedAt) {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoThing) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ArduinoThing) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *ArduinoThing) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

// GetDeviceFqbn returns the DeviceFqbn field value if set, zero value otherwise.
func (o *ArduinoThing) GetDeviceFqbn() string {
	if o == nil || IsNil(o.DeviceFqbn) {
		var ret string
		return ret
	}
	return *o.DeviceFqbn
}

// GetDeviceFqbnOk returns a tuple with the DeviceFqbn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoThing) GetDeviceFqbnOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceFqbn) {
		return nil, false
	}
	return o.DeviceFqbn, true
}

// HasDeviceFqbn returns a boolean if a field has been set.
func (o *ArduinoThing) HasDeviceFqbn() bool {
	if o != nil && !IsNil(o.DeviceFqbn) {
		return true
	}

	return false
}

// SetDeviceFqbn gets a reference to the given string and assigns it to the DeviceFqbn field.
func (o *ArduinoThing) SetDeviceFqbn(v string) {
	o.DeviceFqbn = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *ArduinoThing) GetDeviceId() string {
	if o == nil || IsNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoThing) GetDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *ArduinoThing) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *ArduinoThing) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *ArduinoThing) GetDeviceName() string {
	if o == nil || IsNil(o.DeviceName) {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoThing) GetDeviceNameOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceName) {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *ArduinoThing) HasDeviceName() bool {
	if o != nil && !IsNil(o.DeviceName) {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *ArduinoThing) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *ArduinoThing) GetDeviceType() string {
	if o == nil || IsNil(o.DeviceType) {
		var ret string
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoThing) GetDeviceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceType) {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *ArduinoThing) HasDeviceType() bool {
	if o != nil && !IsNil(o.DeviceType) {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given string and assigns it to the DeviceType field.
func (o *ArduinoThing) SetDeviceType(v string) {
	o.DeviceType = &v
}

// GetHref returns the Href field value
func (o *ArduinoThing) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *ArduinoThing) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *ArduinoThing) SetHref(v string) {
	o.Href = v
}

// GetId returns the Id field value
func (o *ArduinoThing) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ArduinoThing) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ArduinoThing) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ArduinoThing) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ArduinoThing) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ArduinoThing) SetName(v string) {
	o.Name = v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *ArduinoThing) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoThing) GetOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *ArduinoThing) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *ArduinoThing) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *ArduinoThing) GetProperties() []ArduinoProperty {
	if o == nil || IsNil(o.Properties) {
		var ret []ArduinoProperty
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoThing) GetPropertiesOk() ([]ArduinoProperty, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *ArduinoThing) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []ArduinoProperty and assigns it to the Properties field.
func (o *ArduinoThing) SetProperties(v []ArduinoProperty) {
	o.Properties = v
}

// GetPropertiesCount returns the PropertiesCount field value if set, zero value otherwise.
func (o *ArduinoThing) GetPropertiesCount() int64 {
	if o == nil || IsNil(o.PropertiesCount) {
		var ret int64
		return ret
	}
	return *o.PropertiesCount
}

// GetPropertiesCountOk returns a tuple with the PropertiesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoThing) GetPropertiesCountOk() (*int64, bool) {
	if o == nil || IsNil(o.PropertiesCount) {
		return nil, false
	}
	return o.PropertiesCount, true
}

// HasPropertiesCount returns a boolean if a field has been set.
func (o *ArduinoThing) HasPropertiesCount() bool {
	if o != nil && !IsNil(o.PropertiesCount) {
		return true
	}

	return false
}

// SetPropertiesCount gets a reference to the given int64 and assigns it to the PropertiesCount field.
func (o *ArduinoThing) SetPropertiesCount(v int64) {
	o.PropertiesCount = &v
}

// GetSketchId returns the SketchId field value if set, zero value otherwise.
func (o *ArduinoThing) GetSketchId() string {
	if o == nil || IsNil(o.SketchId) {
		var ret string
		return ret
	}
	return *o.SketchId
}

// GetSketchIdOk returns a tuple with the SketchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoThing) GetSketchIdOk() (*string, bool) {
	if o == nil || IsNil(o.SketchId) {
		return nil, false
	}
	return o.SketchId, true
}

// HasSketchId returns a boolean if a field has been set.
func (o *ArduinoThing) HasSketchId() bool {
	if o != nil && !IsNil(o.SketchId) {
		return true
	}

	return false
}

// SetSketchId gets a reference to the given string and assigns it to the SketchId field.
func (o *ArduinoThing) SetSketchId(v string) {
	o.SketchId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ArduinoThing) GetTags() map[string]interface{} {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoThing) GetTagsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Tags) {
		return map[string]interface{}{}, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ArduinoThing) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *ArduinoThing) SetTags(v map[string]interface{}) {
	o.Tags = v
}

// GetTimezone returns the Timezone field value
func (o *ArduinoThing) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *ArduinoThing) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value
func (o *ArduinoThing) SetTimezone(v string) {
	o.Timezone = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ArduinoThing) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoThing) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ArduinoThing) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ArduinoThing) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUserId returns the UserId field value
func (o *ArduinoThing) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *ArduinoThing) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *ArduinoThing) SetUserId(v string) {
	o.UserId = v
}

// GetWebhookActive returns the WebhookActive field value if set, zero value otherwise.
func (o *ArduinoThing) GetWebhookActive() bool {
	if o == nil || IsNil(o.WebhookActive) {
		var ret bool
		return ret
	}
	return *o.WebhookActive
}

// GetWebhookActiveOk returns a tuple with the WebhookActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoThing) GetWebhookActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.WebhookActive) {
		return nil, false
	}
	return o.WebhookActive, true
}

// HasWebhookActive returns a boolean if a field has been set.
func (o *ArduinoThing) HasWebhookActive() bool {
	if o != nil && !IsNil(o.WebhookActive) {
		return true
	}

	return false
}

// SetWebhookActive gets a reference to the given bool and assigns it to the WebhookActive field.
func (o *ArduinoThing) SetWebhookActive(v bool) {
	o.WebhookActive = &v
}

// GetWebhookUri returns the WebhookUri field value if set, zero value otherwise.
func (o *ArduinoThing) GetWebhookUri() string {
	if o == nil || IsNil(o.WebhookUri) {
		var ret string
		return ret
	}
	return *o.WebhookUri
}

// GetWebhookUriOk returns a tuple with the WebhookUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoThing) GetWebhookUriOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookUri) {
		return nil, false
	}
	return o.WebhookUri, true
}

// HasWebhookUri returns a boolean if a field has been set.
func (o *ArduinoThing) HasWebhookUri() bool {
	if o != nil && !IsNil(o.WebhookUri) {
		return true
	}

	return false
}

// SetWebhookUri gets a reference to the given string and assigns it to the WebhookUri field.
func (o *ArduinoThing) SetWebhookUri(v string) {
	o.WebhookUri = &v
}

func (o ArduinoThing) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoThing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assistant) {
		toSerialize["assistant"] = o.Assistant
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if !IsNil(o.DeviceFqbn) {
		toSerialize["device_fqbn"] = o.DeviceFqbn
	}
	if !IsNil(o.DeviceId) {
		toSerialize["device_id"] = o.DeviceId
	}
	if !IsNil(o.DeviceName) {
		toSerialize["device_name"] = o.DeviceName
	}
	if !IsNil(o.DeviceType) {
		toSerialize["device_type"] = o.DeviceType
	}
	toSerialize["href"] = o.Href
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.OrganizationId) {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.PropertiesCount) {
		toSerialize["properties_count"] = o.PropertiesCount
	}
	if !IsNil(o.SketchId) {
		toSerialize["sketch_id"] = o.SketchId
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["timezone"] = o.Timezone
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	toSerialize["user_id"] = o.UserId
	if !IsNil(o.WebhookActive) {
		toSerialize["webhook_active"] = o.WebhookActive
	}
	if !IsNil(o.WebhookUri) {
		toSerialize["webhook_uri"] = o.WebhookUri
	}
	return toSerialize, nil
}

func (o *ArduinoThing) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"href",
		"id",
		"name",
		"timezone",
		"user_id",
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

	varArduinoThing := _ArduinoThing{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varArduinoThing)

	if err != nil {
		return err
	}

	*o = ArduinoThing(varArduinoThing)

	return err
}

type NullableArduinoThing struct {
	value *ArduinoThing
	isSet bool
}

func (v NullableArduinoThing) Get() *ArduinoThing {
	return v.value
}

func (v *NullableArduinoThing) Set(val *ArduinoThing) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoThing) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoThing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoThing(val *ArduinoThing) *NullableArduinoThing {
	return &NullableArduinoThing{value: val, isSet: true}
}

func (v NullableArduinoThing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoThing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


