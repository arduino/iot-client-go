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

// checks if the ArduinoTriggerTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoTriggerTemplate{}

// ArduinoTriggerTemplate ArduinoTrigger_template media type (default view)
type ArduinoTriggerTemplate struct {
	// A list of actions associated with the trigger
	Actions []ArduinoActionTemplate `json:"actions,omitempty"`
	// Is true if the trigger is enabled
	Active *bool `json:"active,omitempty"`
	// The criteria of the trigger, could be INCLUDE or EXCLUDE
	Criteria *string `json:"criteria,omitempty"`
	// The description of the trigger
	Description *string `json:"description,omitempty"`
	// The amount of seconds the trigger will wait before considering a matching device as offline
	GracePeriodOffline *int64 `json:"grace_period_offline,omitempty"`
	// The amount of seconds the trigger will wait before considering a matching device as online
	GracePeriodOnline *int64 `json:"grace_period_online,omitempty"`
	// The id of the trigger
	Id string `json:"id"`
	// A list of devices the trigger is associated to
	LinkedDevices []ArduinoLinkedDeviceTemplate `json:"linked_devices,omitempty"`
	LinkedProperty *ArduinoLinkedPropertyTemplate `json:"linked_property,omitempty"`
	// The name of the trigger
	Name string `json:"name"`
	// Id of the organization the trigger belongs to
	OrganizationId *string `json:"organization_id,omitempty"`
}

type _ArduinoTriggerTemplate ArduinoTriggerTemplate

// NewArduinoTriggerTemplate instantiates a new ArduinoTriggerTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoTriggerTemplate(id string, name string) *ArduinoTriggerTemplate {
	this := ArduinoTriggerTemplate{}
	this.Id = id
	this.Name = name
	return &this
}

// NewArduinoTriggerTemplateWithDefaults instantiates a new ArduinoTriggerTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoTriggerTemplateWithDefaults() *ArduinoTriggerTemplate {
	this := ArduinoTriggerTemplate{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *ArduinoTriggerTemplate) GetActions() []ArduinoActionTemplate {
	if o == nil || IsNil(o.Actions) {
		var ret []ArduinoActionTemplate
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoTriggerTemplate) GetActionsOk() ([]ArduinoActionTemplate, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *ArduinoTriggerTemplate) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []ArduinoActionTemplate and assigns it to the Actions field.
func (o *ArduinoTriggerTemplate) SetActions(v []ArduinoActionTemplate) {
	o.Actions = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ArduinoTriggerTemplate) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoTriggerTemplate) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ArduinoTriggerTemplate) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ArduinoTriggerTemplate) SetActive(v bool) {
	o.Active = &v
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *ArduinoTriggerTemplate) GetCriteria() string {
	if o == nil || IsNil(o.Criteria) {
		var ret string
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoTriggerTemplate) GetCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *ArduinoTriggerTemplate) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given string and assigns it to the Criteria field.
func (o *ArduinoTriggerTemplate) SetCriteria(v string) {
	o.Criteria = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ArduinoTriggerTemplate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoTriggerTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ArduinoTriggerTemplate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ArduinoTriggerTemplate) SetDescription(v string) {
	o.Description = &v
}

// GetGracePeriodOffline returns the GracePeriodOffline field value if set, zero value otherwise.
func (o *ArduinoTriggerTemplate) GetGracePeriodOffline() int64 {
	if o == nil || IsNil(o.GracePeriodOffline) {
		var ret int64
		return ret
	}
	return *o.GracePeriodOffline
}

// GetGracePeriodOfflineOk returns a tuple with the GracePeriodOffline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoTriggerTemplate) GetGracePeriodOfflineOk() (*int64, bool) {
	if o == nil || IsNil(o.GracePeriodOffline) {
		return nil, false
	}
	return o.GracePeriodOffline, true
}

// HasGracePeriodOffline returns a boolean if a field has been set.
func (o *ArduinoTriggerTemplate) HasGracePeriodOffline() bool {
	if o != nil && !IsNil(o.GracePeriodOffline) {
		return true
	}

	return false
}

// SetGracePeriodOffline gets a reference to the given int64 and assigns it to the GracePeriodOffline field.
func (o *ArduinoTriggerTemplate) SetGracePeriodOffline(v int64) {
	o.GracePeriodOffline = &v
}

// GetGracePeriodOnline returns the GracePeriodOnline field value if set, zero value otherwise.
func (o *ArduinoTriggerTemplate) GetGracePeriodOnline() int64 {
	if o == nil || IsNil(o.GracePeriodOnline) {
		var ret int64
		return ret
	}
	return *o.GracePeriodOnline
}

// GetGracePeriodOnlineOk returns a tuple with the GracePeriodOnline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoTriggerTemplate) GetGracePeriodOnlineOk() (*int64, bool) {
	if o == nil || IsNil(o.GracePeriodOnline) {
		return nil, false
	}
	return o.GracePeriodOnline, true
}

// HasGracePeriodOnline returns a boolean if a field has been set.
func (o *ArduinoTriggerTemplate) HasGracePeriodOnline() bool {
	if o != nil && !IsNil(o.GracePeriodOnline) {
		return true
	}

	return false
}

// SetGracePeriodOnline gets a reference to the given int64 and assigns it to the GracePeriodOnline field.
func (o *ArduinoTriggerTemplate) SetGracePeriodOnline(v int64) {
	o.GracePeriodOnline = &v
}

// GetId returns the Id field value
func (o *ArduinoTriggerTemplate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ArduinoTriggerTemplate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ArduinoTriggerTemplate) SetId(v string) {
	o.Id = v
}

// GetLinkedDevices returns the LinkedDevices field value if set, zero value otherwise.
func (o *ArduinoTriggerTemplate) GetLinkedDevices() []ArduinoLinkedDeviceTemplate {
	if o == nil || IsNil(o.LinkedDevices) {
		var ret []ArduinoLinkedDeviceTemplate
		return ret
	}
	return o.LinkedDevices
}

// GetLinkedDevicesOk returns a tuple with the LinkedDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoTriggerTemplate) GetLinkedDevicesOk() ([]ArduinoLinkedDeviceTemplate, bool) {
	if o == nil || IsNil(o.LinkedDevices) {
		return nil, false
	}
	return o.LinkedDevices, true
}

// HasLinkedDevices returns a boolean if a field has been set.
func (o *ArduinoTriggerTemplate) HasLinkedDevices() bool {
	if o != nil && !IsNil(o.LinkedDevices) {
		return true
	}

	return false
}

// SetLinkedDevices gets a reference to the given []ArduinoLinkedDeviceTemplate and assigns it to the LinkedDevices field.
func (o *ArduinoTriggerTemplate) SetLinkedDevices(v []ArduinoLinkedDeviceTemplate) {
	o.LinkedDevices = v
}

// GetLinkedProperty returns the LinkedProperty field value if set, zero value otherwise.
func (o *ArduinoTriggerTemplate) GetLinkedProperty() ArduinoLinkedPropertyTemplate {
	if o == nil || IsNil(o.LinkedProperty) {
		var ret ArduinoLinkedPropertyTemplate
		return ret
	}
	return *o.LinkedProperty
}

// GetLinkedPropertyOk returns a tuple with the LinkedProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoTriggerTemplate) GetLinkedPropertyOk() (*ArduinoLinkedPropertyTemplate, bool) {
	if o == nil || IsNil(o.LinkedProperty) {
		return nil, false
	}
	return o.LinkedProperty, true
}

// HasLinkedProperty returns a boolean if a field has been set.
func (o *ArduinoTriggerTemplate) HasLinkedProperty() bool {
	if o != nil && !IsNil(o.LinkedProperty) {
		return true
	}

	return false
}

// SetLinkedProperty gets a reference to the given ArduinoLinkedPropertyTemplate and assigns it to the LinkedProperty field.
func (o *ArduinoTriggerTemplate) SetLinkedProperty(v ArduinoLinkedPropertyTemplate) {
	o.LinkedProperty = &v
}

// GetName returns the Name field value
func (o *ArduinoTriggerTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ArduinoTriggerTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ArduinoTriggerTemplate) SetName(v string) {
	o.Name = v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *ArduinoTriggerTemplate) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoTriggerTemplate) GetOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *ArduinoTriggerTemplate) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *ArduinoTriggerTemplate) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

func (o ArduinoTriggerTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoTriggerTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Criteria) {
		toSerialize["criteria"] = o.Criteria
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.GracePeriodOffline) {
		toSerialize["grace_period_offline"] = o.GracePeriodOffline
	}
	if !IsNil(o.GracePeriodOnline) {
		toSerialize["grace_period_online"] = o.GracePeriodOnline
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.LinkedDevices) {
		toSerialize["linked_devices"] = o.LinkedDevices
	}
	if !IsNil(o.LinkedProperty) {
		toSerialize["linked_property"] = o.LinkedProperty
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.OrganizationId) {
		toSerialize["organization_id"] = o.OrganizationId
	}
	return toSerialize, nil
}

func (o *ArduinoTriggerTemplate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
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

	varArduinoTriggerTemplate := _ArduinoTriggerTemplate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varArduinoTriggerTemplate)

	if err != nil {
		return err
	}

	*o = ArduinoTriggerTemplate(varArduinoTriggerTemplate)

	return err
}

type NullableArduinoTriggerTemplate struct {
	value *ArduinoTriggerTemplate
	isSet bool
}

func (v NullableArduinoTriggerTemplate) Get() *ArduinoTriggerTemplate {
	return v.value
}

func (v *NullableArduinoTriggerTemplate) Set(val *ArduinoTriggerTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoTriggerTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoTriggerTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoTriggerTemplate(val *ArduinoTriggerTemplate) *NullableArduinoTriggerTemplate {
	return &NullableArduinoTriggerTemplate{value: val, isSet: true}
}

func (v NullableArduinoTriggerTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoTriggerTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

