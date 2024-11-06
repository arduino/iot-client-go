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

// checks if the ArduinoActionTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoActionTemplate{}

// ArduinoActionTemplate ArduinoAction_template media type (default view)
type ArduinoActionTemplate struct {
	// The description of the action
	Description *string `json:"description,omitempty"`
	Email *EmailAction `json:"email,omitempty"`
	// The kind of the action
	Kind *string `json:"kind,omitempty"`
	// The name of the action
	Name *string `json:"name,omitempty"`
	// Id of the organization the trigger belongs to
	OrganizationId *string `json:"organization_id,omitempty"`
	PushNotification *PushAction `json:"push_notification,omitempty"`
}

// NewArduinoActionTemplate instantiates a new ArduinoActionTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoActionTemplate() *ArduinoActionTemplate {
	this := ArduinoActionTemplate{}
	return &this
}

// NewArduinoActionTemplateWithDefaults instantiates a new ArduinoActionTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoActionTemplateWithDefaults() *ArduinoActionTemplate {
	this := ArduinoActionTemplate{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ArduinoActionTemplate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoActionTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ArduinoActionTemplate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ArduinoActionTemplate) SetDescription(v string) {
	o.Description = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ArduinoActionTemplate) GetEmail() EmailAction {
	if o == nil || IsNil(o.Email) {
		var ret EmailAction
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoActionTemplate) GetEmailOk() (*EmailAction, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ArduinoActionTemplate) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given EmailAction and assigns it to the Email field.
func (o *ArduinoActionTemplate) SetEmail(v EmailAction) {
	o.Email = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ArduinoActionTemplate) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoActionTemplate) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ArduinoActionTemplate) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ArduinoActionTemplate) SetKind(v string) {
	o.Kind = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ArduinoActionTemplate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoActionTemplate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ArduinoActionTemplate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ArduinoActionTemplate) SetName(v string) {
	o.Name = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *ArduinoActionTemplate) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoActionTemplate) GetOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *ArduinoActionTemplate) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *ArduinoActionTemplate) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetPushNotification returns the PushNotification field value if set, zero value otherwise.
func (o *ArduinoActionTemplate) GetPushNotification() PushAction {
	if o == nil || IsNil(o.PushNotification) {
		var ret PushAction
		return ret
	}
	return *o.PushNotification
}

// GetPushNotificationOk returns a tuple with the PushNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoActionTemplate) GetPushNotificationOk() (*PushAction, bool) {
	if o == nil || IsNil(o.PushNotification) {
		return nil, false
	}
	return o.PushNotification, true
}

// HasPushNotification returns a boolean if a field has been set.
func (o *ArduinoActionTemplate) HasPushNotification() bool {
	if o != nil && !IsNil(o.PushNotification) {
		return true
	}

	return false
}

// SetPushNotification gets a reference to the given PushAction and assigns it to the PushNotification field.
func (o *ArduinoActionTemplate) SetPushNotification(v PushAction) {
	o.PushNotification = &v
}

func (o ArduinoActionTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoActionTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if !IsNil(o.PushNotification) {
		toSerialize["push_notification"] = o.PushNotification
	}
	return toSerialize, nil
}

type NullableArduinoActionTemplate struct {
	value *ArduinoActionTemplate
	isSet bool
}

func (v NullableArduinoActionTemplate) Get() *ArduinoActionTemplate {
	return v.value
}

func (v *NullableArduinoActionTemplate) Set(val *ArduinoActionTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoActionTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoActionTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoActionTemplate(val *ArduinoActionTemplate) *NullableArduinoActionTemplate {
	return &NullableArduinoActionTemplate{value: val, isSet: true}
}

func (v NullableArduinoActionTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoActionTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


