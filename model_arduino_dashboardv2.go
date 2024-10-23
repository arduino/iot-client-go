/*
Arduino IoT Cloud API

Provides a set of endpoints to manage Arduino IoT Cloud **Devices**, **Things**, **Properties** and **Timeseries**. This API can be called just with any HTTP Client, or using one of these clients:  * [Javascript NPM package](https://www.npmjs.com/package/@arduino/arduino-iot-client)  * [Python PYPI Package](https://pypi.org/project/arduino-iot-client/)  * [Golang Module](https://github.com/arduino/iot-client-go)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the ArduinoDashboardv2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoDashboardv2{}

// ArduinoDashboardv2 Dashboard is a collection of widgets (default view)
type ArduinoDashboardv2 struct {
	// The cover image of the dashboard
	CoverImage *string `json:"cover_image,omitempty"`
	CreatedBy *ArduinoDashboardowner `json:"created_by,omitempty"`
	// The friendly name of the dashboard
	Id string `json:"id"`
	// The friendly name of the dashboard
	Name string `json:"name"`
	// Id of the organization the dashboard belongs to
	OrganizationId *string `json:"organization_id,omitempty"`
	SharedBy *ArduinoDashboardshare `json:"shared_by,omitempty"`
	// ArduinoDashboardshareCollection is the media type for an array of ArduinoDashboardshare (default view)
	SharedWith []ArduinoDashboardshare `json:"shared_with,omitempty"`
	// Last update date
	UpdatedAt time.Time `json:"updated_at"`
	// ArduinoWidgetv2Collection is the media type for an array of ArduinoWidgetv2 (default view)
	Widgets []ArduinoWidgetv2 `json:"widgets,omitempty"`
}

type _ArduinoDashboardv2 ArduinoDashboardv2

// NewArduinoDashboardv2 instantiates a new ArduinoDashboardv2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoDashboardv2(id string, name string, updatedAt time.Time) *ArduinoDashboardv2 {
	this := ArduinoDashboardv2{}
	this.Id = id
	this.Name = name
	this.UpdatedAt = updatedAt
	return &this
}

// NewArduinoDashboardv2WithDefaults instantiates a new ArduinoDashboardv2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoDashboardv2WithDefaults() *ArduinoDashboardv2 {
	this := ArduinoDashboardv2{}
	return &this
}

// GetCoverImage returns the CoverImage field value if set, zero value otherwise.
func (o *ArduinoDashboardv2) GetCoverImage() string {
	if o == nil || IsNil(o.CoverImage) {
		var ret string
		return ret
	}
	return *o.CoverImage
}

// GetCoverImageOk returns a tuple with the CoverImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoDashboardv2) GetCoverImageOk() (*string, bool) {
	if o == nil || IsNil(o.CoverImage) {
		return nil, false
	}
	return o.CoverImage, true
}

// HasCoverImage returns a boolean if a field has been set.
func (o *ArduinoDashboardv2) HasCoverImage() bool {
	if o != nil && !IsNil(o.CoverImage) {
		return true
	}

	return false
}

// SetCoverImage gets a reference to the given string and assigns it to the CoverImage field.
func (o *ArduinoDashboardv2) SetCoverImage(v string) {
	o.CoverImage = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ArduinoDashboardv2) GetCreatedBy() ArduinoDashboardowner {
	if o == nil || IsNil(o.CreatedBy) {
		var ret ArduinoDashboardowner
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoDashboardv2) GetCreatedByOk() (*ArduinoDashboardowner, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ArduinoDashboardv2) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given ArduinoDashboardowner and assigns it to the CreatedBy field.
func (o *ArduinoDashboardv2) SetCreatedBy(v ArduinoDashboardowner) {
	o.CreatedBy = &v
}

// GetId returns the Id field value
func (o *ArduinoDashboardv2) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ArduinoDashboardv2) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ArduinoDashboardv2) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ArduinoDashboardv2) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ArduinoDashboardv2) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ArduinoDashboardv2) SetName(v string) {
	o.Name = v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *ArduinoDashboardv2) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoDashboardv2) GetOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *ArduinoDashboardv2) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *ArduinoDashboardv2) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetSharedBy returns the SharedBy field value if set, zero value otherwise.
func (o *ArduinoDashboardv2) GetSharedBy() ArduinoDashboardshare {
	if o == nil || IsNil(o.SharedBy) {
		var ret ArduinoDashboardshare
		return ret
	}
	return *o.SharedBy
}

// GetSharedByOk returns a tuple with the SharedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoDashboardv2) GetSharedByOk() (*ArduinoDashboardshare, bool) {
	if o == nil || IsNil(o.SharedBy) {
		return nil, false
	}
	return o.SharedBy, true
}

// HasSharedBy returns a boolean if a field has been set.
func (o *ArduinoDashboardv2) HasSharedBy() bool {
	if o != nil && !IsNil(o.SharedBy) {
		return true
	}

	return false
}

// SetSharedBy gets a reference to the given ArduinoDashboardshare and assigns it to the SharedBy field.
func (o *ArduinoDashboardv2) SetSharedBy(v ArduinoDashboardshare) {
	o.SharedBy = &v
}

// GetSharedWith returns the SharedWith field value if set, zero value otherwise.
func (o *ArduinoDashboardv2) GetSharedWith() []ArduinoDashboardshare {
	if o == nil || IsNil(o.SharedWith) {
		var ret []ArduinoDashboardshare
		return ret
	}
	return o.SharedWith
}

// GetSharedWithOk returns a tuple with the SharedWith field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoDashboardv2) GetSharedWithOk() ([]ArduinoDashboardshare, bool) {
	if o == nil || IsNil(o.SharedWith) {
		return nil, false
	}
	return o.SharedWith, true
}

// HasSharedWith returns a boolean if a field has been set.
func (o *ArduinoDashboardv2) HasSharedWith() bool {
	if o != nil && !IsNil(o.SharedWith) {
		return true
	}

	return false
}

// SetSharedWith gets a reference to the given []ArduinoDashboardshare and assigns it to the SharedWith field.
func (o *ArduinoDashboardv2) SetSharedWith(v []ArduinoDashboardshare) {
	o.SharedWith = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ArduinoDashboardv2) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ArduinoDashboardv2) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ArduinoDashboardv2) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetWidgets returns the Widgets field value if set, zero value otherwise.
func (o *ArduinoDashboardv2) GetWidgets() []ArduinoWidgetv2 {
	if o == nil || IsNil(o.Widgets) {
		var ret []ArduinoWidgetv2
		return ret
	}
	return o.Widgets
}

// GetWidgetsOk returns a tuple with the Widgets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoDashboardv2) GetWidgetsOk() ([]ArduinoWidgetv2, bool) {
	if o == nil || IsNil(o.Widgets) {
		return nil, false
	}
	return o.Widgets, true
}

// HasWidgets returns a boolean if a field has been set.
func (o *ArduinoDashboardv2) HasWidgets() bool {
	if o != nil && !IsNil(o.Widgets) {
		return true
	}

	return false
}

// SetWidgets gets a reference to the given []ArduinoWidgetv2 and assigns it to the Widgets field.
func (o *ArduinoDashboardv2) SetWidgets(v []ArduinoWidgetv2) {
	o.Widgets = v
}

func (o ArduinoDashboardv2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoDashboardv2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CoverImage) {
		toSerialize["cover_image"] = o.CoverImage
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["created_by"] = o.CreatedBy
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.OrganizationId) {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if !IsNil(o.SharedBy) {
		toSerialize["shared_by"] = o.SharedBy
	}
	if !IsNil(o.SharedWith) {
		toSerialize["shared_with"] = o.SharedWith
	}
	toSerialize["updated_at"] = o.UpdatedAt
	if !IsNil(o.Widgets) {
		toSerialize["widgets"] = o.Widgets
	}
	return toSerialize, nil
}

func (o *ArduinoDashboardv2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"updated_at",
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

	varArduinoDashboardv2 := _ArduinoDashboardv2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varArduinoDashboardv2)

	if err != nil {
		return err
	}

	*o = ArduinoDashboardv2(varArduinoDashboardv2)

	return err
}

type NullableArduinoDashboardv2 struct {
	value *ArduinoDashboardv2
	isSet bool
}

func (v NullableArduinoDashboardv2) Get() *ArduinoDashboardv2 {
	return v.value
}

func (v *NullableArduinoDashboardv2) Set(val *ArduinoDashboardv2) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoDashboardv2) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoDashboardv2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoDashboardv2(val *ArduinoDashboardv2) *NullableArduinoDashboardv2 {
	return &NullableArduinoDashboardv2{value: val, isSet: true}
}

func (v NullableArduinoDashboardv2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoDashboardv2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


