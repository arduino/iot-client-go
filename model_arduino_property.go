/*
Arduino IoT Cloud API

 Provides a set of endpoints to manage Arduino IoT Cloud **Devices**, **Things**, **Properties** and **Timeseries**. This API can be called just with any HTTP Client, or using one of these clients:  * [Javascript NPM package](https://www.npmjs.com/package/@arduino/arduino-iot-client)  * [Python PYPI Package](https://pypi.org/project/arduino-iot-client/)  * [Golang Module](https://github.com/arduino/iot-client-go)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iot

import (
	"encoding/json"
	"time"
)

// checks if the ArduinoProperty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoProperty{}

// ArduinoProperty ArduinoProperty media type (default view)
type ArduinoProperty struct {
	// Creation date of the property
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Delete date of the property
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// The api reference of this property
	Href string `json:"href"`
	// The id of the property
	Id string `json:"id"`
	// Last value of this property
	LastValue interface{} `json:"last_value,omitempty"`
	// Maximum value of this property
	MaxValue *float64 `json:"max_value,omitempty"`
	// Minimum value of this property
	MinValue *float64 `json:"min_value,omitempty"`
	// The friendly name of the property
	Name string `json:"name"`
	// The permission of the property
	Permission string `json:"permission"`
	// If true, data will persist into a timeseries database
	Persist *bool `json:"persist,omitempty"`
	// The id of the sync pool
	SyncId *string `json:"sync_id,omitempty"`
	// The integer id of the property
	Tag *float64 `json:"tag,omitempty"`
	// The id of the thing
	ThingId string `json:"thing_id"`
	// The name of the associated thing
	ThingName *string `json:"thing_name,omitempty"`
	// The type of the property
	Type string `json:"type"`
	// The update frequency in seconds, or the amount of the property has to change in order to trigger an update
	UpdateParameter *float64 `json:"update_parameter,omitempty"`
	// The update strategy for the property value
	UpdateStrategy string `json:"update_strategy"`
	// Update date of the property
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Last update timestamp of this property
	ValueUpdatedAt *time.Time `json:"value_updated_at,omitempty"`
	// The sketch variable name of the property
	VariableName *string `json:"variable_name,omitempty"`
}

// NewArduinoProperty instantiates a new ArduinoProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoProperty(href string, id string, name string, permission string, thingId string, type_ string, updateStrategy string) *ArduinoProperty {
	this := ArduinoProperty{}
	this.Href = href
	this.Id = id
	this.Name = name
	this.Permission = permission
	this.ThingId = thingId
	this.Type = type_
	this.UpdateStrategy = updateStrategy
	return &this
}

// NewArduinoPropertyWithDefaults instantiates a new ArduinoProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoPropertyWithDefaults() *ArduinoProperty {
	this := ArduinoProperty{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ArduinoProperty) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoProperty) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ArduinoProperty) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ArduinoProperty) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *ArduinoProperty) GetDeletedAt() time.Time {
	if o == nil || IsNil(o.DeletedAt) {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoProperty) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ArduinoProperty) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *ArduinoProperty) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

// GetHref returns the Href field value
func (o *ArduinoProperty) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *ArduinoProperty) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *ArduinoProperty) SetHref(v string) {
	o.Href = v
}

// GetId returns the Id field value
func (o *ArduinoProperty) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ArduinoProperty) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ArduinoProperty) SetId(v string) {
	o.Id = v
}

// GetLastValue returns the LastValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArduinoProperty) GetLastValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.LastValue
}

// GetLastValueOk returns a tuple with the LastValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArduinoProperty) GetLastValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.LastValue) {
		return nil, false
	}
	return &o.LastValue, true
}

// HasLastValue returns a boolean if a field has been set.
func (o *ArduinoProperty) HasLastValue() bool {
	if o != nil && IsNil(o.LastValue) {
		return true
	}

	return false
}

// SetLastValue gets a reference to the given interface{} and assigns it to the LastValue field.
func (o *ArduinoProperty) SetLastValue(v interface{}) {
	o.LastValue = v
}

// GetMaxValue returns the MaxValue field value if set, zero value otherwise.
func (o *ArduinoProperty) GetMaxValue() float64 {
	if o == nil || IsNil(o.MaxValue) {
		var ret float64
		return ret
	}
	return *o.MaxValue
}

// GetMaxValueOk returns a tuple with the MaxValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoProperty) GetMaxValueOk() (*float64, bool) {
	if o == nil || IsNil(o.MaxValue) {
		return nil, false
	}
	return o.MaxValue, true
}

// HasMaxValue returns a boolean if a field has been set.
func (o *ArduinoProperty) HasMaxValue() bool {
	if o != nil && !IsNil(o.MaxValue) {
		return true
	}

	return false
}

// SetMaxValue gets a reference to the given float64 and assigns it to the MaxValue field.
func (o *ArduinoProperty) SetMaxValue(v float64) {
	o.MaxValue = &v
}

// GetMinValue returns the MinValue field value if set, zero value otherwise.
func (o *ArduinoProperty) GetMinValue() float64 {
	if o == nil || IsNil(o.MinValue) {
		var ret float64
		return ret
	}
	return *o.MinValue
}

// GetMinValueOk returns a tuple with the MinValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoProperty) GetMinValueOk() (*float64, bool) {
	if o == nil || IsNil(o.MinValue) {
		return nil, false
	}
	return o.MinValue, true
}

// HasMinValue returns a boolean if a field has been set.
func (o *ArduinoProperty) HasMinValue() bool {
	if o != nil && !IsNil(o.MinValue) {
		return true
	}

	return false
}

// SetMinValue gets a reference to the given float64 and assigns it to the MinValue field.
func (o *ArduinoProperty) SetMinValue(v float64) {
	o.MinValue = &v
}

// GetName returns the Name field value
func (o *ArduinoProperty) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ArduinoProperty) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ArduinoProperty) SetName(v string) {
	o.Name = v
}

// GetPermission returns the Permission field value
func (o *ArduinoProperty) GetPermission() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value
// and a boolean to check if the value has been set.
func (o *ArduinoProperty) GetPermissionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permission, true
}

// SetPermission sets field value
func (o *ArduinoProperty) SetPermission(v string) {
	o.Permission = v
}

// GetPersist returns the Persist field value if set, zero value otherwise.
func (o *ArduinoProperty) GetPersist() bool {
	if o == nil || IsNil(o.Persist) {
		var ret bool
		return ret
	}
	return *o.Persist
}

// GetPersistOk returns a tuple with the Persist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoProperty) GetPersistOk() (*bool, bool) {
	if o == nil || IsNil(o.Persist) {
		return nil, false
	}
	return o.Persist, true
}

// HasPersist returns a boolean if a field has been set.
func (o *ArduinoProperty) HasPersist() bool {
	if o != nil && !IsNil(o.Persist) {
		return true
	}

	return false
}

// SetPersist gets a reference to the given bool and assigns it to the Persist field.
func (o *ArduinoProperty) SetPersist(v bool) {
	o.Persist = &v
}

// GetSyncId returns the SyncId field value if set, zero value otherwise.
func (o *ArduinoProperty) GetSyncId() string {
	if o == nil || IsNil(o.SyncId) {
		var ret string
		return ret
	}
	return *o.SyncId
}

// GetSyncIdOk returns a tuple with the SyncId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoProperty) GetSyncIdOk() (*string, bool) {
	if o == nil || IsNil(o.SyncId) {
		return nil, false
	}
	return o.SyncId, true
}

// HasSyncId returns a boolean if a field has been set.
func (o *ArduinoProperty) HasSyncId() bool {
	if o != nil && !IsNil(o.SyncId) {
		return true
	}

	return false
}

// SetSyncId gets a reference to the given string and assigns it to the SyncId field.
func (o *ArduinoProperty) SetSyncId(v string) {
	o.SyncId = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *ArduinoProperty) GetTag() float64 {
	if o == nil || IsNil(o.Tag) {
		var ret float64
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoProperty) GetTagOk() (*float64, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *ArduinoProperty) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given float64 and assigns it to the Tag field.
func (o *ArduinoProperty) SetTag(v float64) {
	o.Tag = &v
}

// GetThingId returns the ThingId field value
func (o *ArduinoProperty) GetThingId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThingId
}

// GetThingIdOk returns a tuple with the ThingId field value
// and a boolean to check if the value has been set.
func (o *ArduinoProperty) GetThingIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThingId, true
}

// SetThingId sets field value
func (o *ArduinoProperty) SetThingId(v string) {
	o.ThingId = v
}

// GetThingName returns the ThingName field value if set, zero value otherwise.
func (o *ArduinoProperty) GetThingName() string {
	if o == nil || IsNil(o.ThingName) {
		var ret string
		return ret
	}
	return *o.ThingName
}

// GetThingNameOk returns a tuple with the ThingName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoProperty) GetThingNameOk() (*string, bool) {
	if o == nil || IsNil(o.ThingName) {
		return nil, false
	}
	return o.ThingName, true
}

// HasThingName returns a boolean if a field has been set.
func (o *ArduinoProperty) HasThingName() bool {
	if o != nil && !IsNil(o.ThingName) {
		return true
	}

	return false
}

// SetThingName gets a reference to the given string and assigns it to the ThingName field.
func (o *ArduinoProperty) SetThingName(v string) {
	o.ThingName = &v
}

// GetType returns the Type field value
func (o *ArduinoProperty) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ArduinoProperty) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ArduinoProperty) SetType(v string) {
	o.Type = v
}

// GetUpdateParameter returns the UpdateParameter field value if set, zero value otherwise.
func (o *ArduinoProperty) GetUpdateParameter() float64 {
	if o == nil || IsNil(o.UpdateParameter) {
		var ret float64
		return ret
	}
	return *o.UpdateParameter
}

// GetUpdateParameterOk returns a tuple with the UpdateParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoProperty) GetUpdateParameterOk() (*float64, bool) {
	if o == nil || IsNil(o.UpdateParameter) {
		return nil, false
	}
	return o.UpdateParameter, true
}

// HasUpdateParameter returns a boolean if a field has been set.
func (o *ArduinoProperty) HasUpdateParameter() bool {
	if o != nil && !IsNil(o.UpdateParameter) {
		return true
	}

	return false
}

// SetUpdateParameter gets a reference to the given float64 and assigns it to the UpdateParameter field.
func (o *ArduinoProperty) SetUpdateParameter(v float64) {
	o.UpdateParameter = &v
}

// GetUpdateStrategy returns the UpdateStrategy field value
func (o *ArduinoProperty) GetUpdateStrategy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdateStrategy
}

// GetUpdateStrategyOk returns a tuple with the UpdateStrategy field value
// and a boolean to check if the value has been set.
func (o *ArduinoProperty) GetUpdateStrategyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdateStrategy, true
}

// SetUpdateStrategy sets field value
func (o *ArduinoProperty) SetUpdateStrategy(v string) {
	o.UpdateStrategy = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ArduinoProperty) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoProperty) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ArduinoProperty) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ArduinoProperty) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetValueUpdatedAt returns the ValueUpdatedAt field value if set, zero value otherwise.
func (o *ArduinoProperty) GetValueUpdatedAt() time.Time {
	if o == nil || IsNil(o.ValueUpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.ValueUpdatedAt
}

// GetValueUpdatedAtOk returns a tuple with the ValueUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoProperty) GetValueUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ValueUpdatedAt) {
		return nil, false
	}
	return o.ValueUpdatedAt, true
}

// HasValueUpdatedAt returns a boolean if a field has been set.
func (o *ArduinoProperty) HasValueUpdatedAt() bool {
	if o != nil && !IsNil(o.ValueUpdatedAt) {
		return true
	}

	return false
}

// SetValueUpdatedAt gets a reference to the given time.Time and assigns it to the ValueUpdatedAt field.
func (o *ArduinoProperty) SetValueUpdatedAt(v time.Time) {
	o.ValueUpdatedAt = &v
}

// GetVariableName returns the VariableName field value if set, zero value otherwise.
func (o *ArduinoProperty) GetVariableName() string {
	if o == nil || IsNil(o.VariableName) {
		var ret string
		return ret
	}
	return *o.VariableName
}

// GetVariableNameOk returns a tuple with the VariableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoProperty) GetVariableNameOk() (*string, bool) {
	if o == nil || IsNil(o.VariableName) {
		return nil, false
	}
	return o.VariableName, true
}

// HasVariableName returns a boolean if a field has been set.
func (o *ArduinoProperty) HasVariableName() bool {
	if o != nil && !IsNil(o.VariableName) {
		return true
	}

	return false
}

// SetVariableName gets a reference to the given string and assigns it to the VariableName field.
func (o *ArduinoProperty) SetVariableName(v string) {
	o.VariableName = &v
}

func (o ArduinoProperty) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoProperty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	toSerialize["href"] = o.Href
	toSerialize["id"] = o.Id
	if o.LastValue != nil {
		toSerialize["last_value"] = o.LastValue
	}
	if !IsNil(o.MaxValue) {
		toSerialize["max_value"] = o.MaxValue
	}
	if !IsNil(o.MinValue) {
		toSerialize["min_value"] = o.MinValue
	}
	toSerialize["name"] = o.Name
	toSerialize["permission"] = o.Permission
	if !IsNil(o.Persist) {
		toSerialize["persist"] = o.Persist
	}
	if !IsNil(o.SyncId) {
		toSerialize["sync_id"] = o.SyncId
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	toSerialize["thing_id"] = o.ThingId
	if !IsNil(o.ThingName) {
		toSerialize["thing_name"] = o.ThingName
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.UpdateParameter) {
		toSerialize["update_parameter"] = o.UpdateParameter
	}
	toSerialize["update_strategy"] = o.UpdateStrategy
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.ValueUpdatedAt) {
		toSerialize["value_updated_at"] = o.ValueUpdatedAt
	}
	if !IsNil(o.VariableName) {
		toSerialize["variable_name"] = o.VariableName
	}
	return toSerialize, nil
}

type NullableArduinoProperty struct {
	value *ArduinoProperty
	isSet bool
}

func (v NullableArduinoProperty) Get() *ArduinoProperty {
	return v.value
}

func (v *NullableArduinoProperty) Set(val *ArduinoProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoProperty(val *ArduinoProperty) *NullableArduinoProperty {
	return &NullableArduinoProperty{value: val, isSet: true}
}

func (v NullableArduinoProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


