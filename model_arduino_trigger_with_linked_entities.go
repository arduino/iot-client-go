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

// checks if the ArduinoTriggerWithLinkedEntities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoTriggerWithLinkedEntities{}

// ArduinoTriggerWithLinkedEntities ArduinoTrigger_with_linked_entities media type (default view)
type ArduinoTriggerWithLinkedEntities struct {
	// A list of actions associated with the trigger
	Actions []ArduinoAction `json:"actions,omitempty"`
	// Is true if the trigger is enabled
	Active *bool `json:"active,omitempty"`
	// Id of the user who last updated the trigger
	CreatedBy *string `json:"created_by,omitempty"`
	// The description of the trigger
	Description *string `json:"description,omitempty"`
	DeviceStatusSource *DeviceStatusSourceWithLinkedDevices `json:"device_status_source,omitempty"`
	// The id of the trigger
	Id string `json:"id"`
	LinkedProperty *ArduinoLinkedProperty `json:"linked_property,omitempty"`
	// The name of the trigger
	Name string `json:"name"`
	// Id of the organization the trigger belongs to
	OrganizationId *string `json:"organization_id,omitempty"`
}

type _ArduinoTriggerWithLinkedEntities ArduinoTriggerWithLinkedEntities

// NewArduinoTriggerWithLinkedEntities instantiates a new ArduinoTriggerWithLinkedEntities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoTriggerWithLinkedEntities(id string, name string) *ArduinoTriggerWithLinkedEntities {
	this := ArduinoTriggerWithLinkedEntities{}
	this.Id = id
	this.Name = name
	return &this
}

// NewArduinoTriggerWithLinkedEntitiesWithDefaults instantiates a new ArduinoTriggerWithLinkedEntities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoTriggerWithLinkedEntitiesWithDefaults() *ArduinoTriggerWithLinkedEntities {
	this := ArduinoTriggerWithLinkedEntities{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *ArduinoTriggerWithLinkedEntities) GetActions() []ArduinoAction {
	if o == nil || IsNil(o.Actions) {
		var ret []ArduinoAction
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoTriggerWithLinkedEntities) GetActionsOk() ([]ArduinoAction, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *ArduinoTriggerWithLinkedEntities) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []ArduinoAction and assigns it to the Actions field.
func (o *ArduinoTriggerWithLinkedEntities) SetActions(v []ArduinoAction) {
	o.Actions = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ArduinoTriggerWithLinkedEntities) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoTriggerWithLinkedEntities) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ArduinoTriggerWithLinkedEntities) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ArduinoTriggerWithLinkedEntities) SetActive(v bool) {
	o.Active = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ArduinoTriggerWithLinkedEntities) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoTriggerWithLinkedEntities) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ArduinoTriggerWithLinkedEntities) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *ArduinoTriggerWithLinkedEntities) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ArduinoTriggerWithLinkedEntities) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoTriggerWithLinkedEntities) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ArduinoTriggerWithLinkedEntities) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ArduinoTriggerWithLinkedEntities) SetDescription(v string) {
	o.Description = &v
}

// GetDeviceStatusSource returns the DeviceStatusSource field value if set, zero value otherwise.
func (o *ArduinoTriggerWithLinkedEntities) GetDeviceStatusSource() DeviceStatusSourceWithLinkedDevices {
	if o == nil || IsNil(o.DeviceStatusSource) {
		var ret DeviceStatusSourceWithLinkedDevices
		return ret
	}
	return *o.DeviceStatusSource
}

// GetDeviceStatusSourceOk returns a tuple with the DeviceStatusSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoTriggerWithLinkedEntities) GetDeviceStatusSourceOk() (*DeviceStatusSourceWithLinkedDevices, bool) {
	if o == nil || IsNil(o.DeviceStatusSource) {
		return nil, false
	}
	return o.DeviceStatusSource, true
}

// HasDeviceStatusSource returns a boolean if a field has been set.
func (o *ArduinoTriggerWithLinkedEntities) HasDeviceStatusSource() bool {
	if o != nil && !IsNil(o.DeviceStatusSource) {
		return true
	}

	return false
}

// SetDeviceStatusSource gets a reference to the given DeviceStatusSourceWithLinkedDevices and assigns it to the DeviceStatusSource field.
func (o *ArduinoTriggerWithLinkedEntities) SetDeviceStatusSource(v DeviceStatusSourceWithLinkedDevices) {
	o.DeviceStatusSource = &v
}

// GetId returns the Id field value
func (o *ArduinoTriggerWithLinkedEntities) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ArduinoTriggerWithLinkedEntities) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ArduinoTriggerWithLinkedEntities) SetId(v string) {
	o.Id = v
}

// GetLinkedProperty returns the LinkedProperty field value if set, zero value otherwise.
func (o *ArduinoTriggerWithLinkedEntities) GetLinkedProperty() ArduinoLinkedProperty {
	if o == nil || IsNil(o.LinkedProperty) {
		var ret ArduinoLinkedProperty
		return ret
	}
	return *o.LinkedProperty
}

// GetLinkedPropertyOk returns a tuple with the LinkedProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoTriggerWithLinkedEntities) GetLinkedPropertyOk() (*ArduinoLinkedProperty, bool) {
	if o == nil || IsNil(o.LinkedProperty) {
		return nil, false
	}
	return o.LinkedProperty, true
}

// HasLinkedProperty returns a boolean if a field has been set.
func (o *ArduinoTriggerWithLinkedEntities) HasLinkedProperty() bool {
	if o != nil && !IsNil(o.LinkedProperty) {
		return true
	}

	return false
}

// SetLinkedProperty gets a reference to the given ArduinoLinkedProperty and assigns it to the LinkedProperty field.
func (o *ArduinoTriggerWithLinkedEntities) SetLinkedProperty(v ArduinoLinkedProperty) {
	o.LinkedProperty = &v
}

// GetName returns the Name field value
func (o *ArduinoTriggerWithLinkedEntities) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ArduinoTriggerWithLinkedEntities) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ArduinoTriggerWithLinkedEntities) SetName(v string) {
	o.Name = v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *ArduinoTriggerWithLinkedEntities) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoTriggerWithLinkedEntities) GetOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *ArduinoTriggerWithLinkedEntities) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *ArduinoTriggerWithLinkedEntities) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

func (o ArduinoTriggerWithLinkedEntities) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoTriggerWithLinkedEntities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["created_by"] = o.CreatedBy
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DeviceStatusSource) {
		toSerialize["device_status_source"] = o.DeviceStatusSource
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.LinkedProperty) {
		toSerialize["linked_property"] = o.LinkedProperty
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.OrganizationId) {
		toSerialize["organization_id"] = o.OrganizationId
	}
	return toSerialize, nil
}

func (o *ArduinoTriggerWithLinkedEntities) UnmarshalJSON(data []byte) (err error) {
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

	varArduinoTriggerWithLinkedEntities := _ArduinoTriggerWithLinkedEntities{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varArduinoTriggerWithLinkedEntities)

	if err != nil {
		return err
	}

	*o = ArduinoTriggerWithLinkedEntities(varArduinoTriggerWithLinkedEntities)

	return err
}

type NullableArduinoTriggerWithLinkedEntities struct {
	value *ArduinoTriggerWithLinkedEntities
	isSet bool
}

func (v NullableArduinoTriggerWithLinkedEntities) Get() *ArduinoTriggerWithLinkedEntities {
	return v.value
}

func (v *NullableArduinoTriggerWithLinkedEntities) Set(val *ArduinoTriggerWithLinkedEntities) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoTriggerWithLinkedEntities) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoTriggerWithLinkedEntities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoTriggerWithLinkedEntities(val *ArduinoTriggerWithLinkedEntities) *NullableArduinoTriggerWithLinkedEntities {
	return &NullableArduinoTriggerWithLinkedEntities{value: val, isSet: true}
}

func (v NullableArduinoTriggerWithLinkedEntities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoTriggerWithLinkedEntities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


