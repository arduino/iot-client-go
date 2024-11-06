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

// checks if the ArduinoAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoAction{}

// ArduinoAction ArduinoAction media type (default view)
type ArduinoAction struct {
	// Id of the user who created the action
	CreatedBy *string `json:"created_by,omitempty"`
	// The description of the action
	Description *string `json:"description,omitempty"`
	Email *EmailAction `json:"email,omitempty"`
	// The id of the action
	Id *string `json:"id,omitempty"`
	// The kind of the action
	Kind *string `json:"kind,omitempty"`
	// The name of the action
	Name *string `json:"name,omitempty"`
	// Id of the organization the trigger belongs to
	OrganizationId *string `json:"organization_id,omitempty"`
	PushNotification *PushAction `json:"push_notification,omitempty"`
	// Id of the trigger the action is associated to
	TriggerId *string `json:"trigger_id,omitempty"`
}

// NewArduinoAction instantiates a new ArduinoAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoAction() *ArduinoAction {
	this := ArduinoAction{}
	return &this
}

// NewArduinoActionWithDefaults instantiates a new ArduinoAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoActionWithDefaults() *ArduinoAction {
	this := ArduinoAction{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ArduinoAction) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoAction) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ArduinoAction) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *ArduinoAction) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ArduinoAction) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoAction) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ArduinoAction) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ArduinoAction) SetDescription(v string) {
	o.Description = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ArduinoAction) GetEmail() EmailAction {
	if o == nil || IsNil(o.Email) {
		var ret EmailAction
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoAction) GetEmailOk() (*EmailAction, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ArduinoAction) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given EmailAction and assigns it to the Email field.
func (o *ArduinoAction) SetEmail(v EmailAction) {
	o.Email = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ArduinoAction) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoAction) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ArduinoAction) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ArduinoAction) SetId(v string) {
	o.Id = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ArduinoAction) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoAction) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ArduinoAction) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ArduinoAction) SetKind(v string) {
	o.Kind = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ArduinoAction) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoAction) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ArduinoAction) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ArduinoAction) SetName(v string) {
	o.Name = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *ArduinoAction) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoAction) GetOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *ArduinoAction) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *ArduinoAction) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetPushNotification returns the PushNotification field value if set, zero value otherwise.
func (o *ArduinoAction) GetPushNotification() PushAction {
	if o == nil || IsNil(o.PushNotification) {
		var ret PushAction
		return ret
	}
	return *o.PushNotification
}

// GetPushNotificationOk returns a tuple with the PushNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoAction) GetPushNotificationOk() (*PushAction, bool) {
	if o == nil || IsNil(o.PushNotification) {
		return nil, false
	}
	return o.PushNotification, true
}

// HasPushNotification returns a boolean if a field has been set.
func (o *ArduinoAction) HasPushNotification() bool {
	if o != nil && !IsNil(o.PushNotification) {
		return true
	}

	return false
}

// SetPushNotification gets a reference to the given PushAction and assigns it to the PushNotification field.
func (o *ArduinoAction) SetPushNotification(v PushAction) {
	o.PushNotification = &v
}

// GetTriggerId returns the TriggerId field value if set, zero value otherwise.
func (o *ArduinoAction) GetTriggerId() string {
	if o == nil || IsNil(o.TriggerId) {
		var ret string
		return ret
	}
	return *o.TriggerId
}

// GetTriggerIdOk returns a tuple with the TriggerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoAction) GetTriggerIdOk() (*string, bool) {
	if o == nil || IsNil(o.TriggerId) {
		return nil, false
	}
	return o.TriggerId, true
}

// HasTriggerId returns a boolean if a field has been set.
func (o *ArduinoAction) HasTriggerId() bool {
	if o != nil && !IsNil(o.TriggerId) {
		return true
	}

	return false
}

// SetTriggerId gets a reference to the given string and assigns it to the TriggerId field.
func (o *ArduinoAction) SetTriggerId(v string) {
	o.TriggerId = &v
}

func (o ArduinoAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedBy) {
		toSerialize["created_by"] = o.CreatedBy
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
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
	if !IsNil(o.TriggerId) {
		toSerialize["trigger_id"] = o.TriggerId
	}
	return toSerialize, nil
}

type NullableArduinoAction struct {
	value *ArduinoAction
	isSet bool
}

func (v NullableArduinoAction) Get() *ArduinoAction {
	return v.value
}

func (v *NullableArduinoAction) Set(val *ArduinoAction) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoAction) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoAction(val *ArduinoAction) *NullableArduinoAction {
	return &NullableArduinoAction{value: val, isSet: true}
}

func (v NullableArduinoAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


