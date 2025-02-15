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

// checks if the Dashboardv2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dashboardv2{}

// Dashboardv2 DashboardV2Payload describes a dashboard
type Dashboardv2 struct {
	// The cover image of the dashboard
	CoverImage *string `json:"cover_image,omitempty"`
	// The friendly name of the dashboard
	Name *string `json:"name,omitempty" validate:"regexp=[a-zA-Z0-9_.@-]+"`
	// If false, restore the thing from the soft deletion
	SoftDeleted *bool `json:"soft_deleted,omitempty"`
	// Widgets attached to this dashboard
	Widgets []Widget `json:"widgets,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Dashboardv2 Dashboardv2

// NewDashboardv2 instantiates a new Dashboardv2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardv2() *Dashboardv2 {
	this := Dashboardv2{}
	var softDeleted bool = false
	this.SoftDeleted = &softDeleted
	return &this
}

// NewDashboardv2WithDefaults instantiates a new Dashboardv2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardv2WithDefaults() *Dashboardv2 {
	this := Dashboardv2{}
	var softDeleted bool = false
	this.SoftDeleted = &softDeleted
	return &this
}

// GetCoverImage returns the CoverImage field value if set, zero value otherwise.
func (o *Dashboardv2) GetCoverImage() string {
	if o == nil || IsNil(o.CoverImage) {
		var ret string
		return ret
	}
	return *o.CoverImage
}

// GetCoverImageOk returns a tuple with the CoverImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dashboardv2) GetCoverImageOk() (*string, bool) {
	if o == nil || IsNil(o.CoverImage) {
		return nil, false
	}
	return o.CoverImage, true
}

// HasCoverImage returns a boolean if a field has been set.
func (o *Dashboardv2) HasCoverImage() bool {
	if o != nil && !IsNil(o.CoverImage) {
		return true
	}

	return false
}

// SetCoverImage gets a reference to the given string and assigns it to the CoverImage field.
func (o *Dashboardv2) SetCoverImage(v string) {
	o.CoverImage = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Dashboardv2) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dashboardv2) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Dashboardv2) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Dashboardv2) SetName(v string) {
	o.Name = &v
}

// GetSoftDeleted returns the SoftDeleted field value if set, zero value otherwise.
func (o *Dashboardv2) GetSoftDeleted() bool {
	if o == nil || IsNil(o.SoftDeleted) {
		var ret bool
		return ret
	}
	return *o.SoftDeleted
}

// GetSoftDeletedOk returns a tuple with the SoftDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dashboardv2) GetSoftDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.SoftDeleted) {
		return nil, false
	}
	return o.SoftDeleted, true
}

// HasSoftDeleted returns a boolean if a field has been set.
func (o *Dashboardv2) HasSoftDeleted() bool {
	if o != nil && !IsNil(o.SoftDeleted) {
		return true
	}

	return false
}

// SetSoftDeleted gets a reference to the given bool and assigns it to the SoftDeleted field.
func (o *Dashboardv2) SetSoftDeleted(v bool) {
	o.SoftDeleted = &v
}

// GetWidgets returns the Widgets field value if set, zero value otherwise.
func (o *Dashboardv2) GetWidgets() []Widget {
	if o == nil || IsNil(o.Widgets) {
		var ret []Widget
		return ret
	}
	return o.Widgets
}

// GetWidgetsOk returns a tuple with the Widgets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dashboardv2) GetWidgetsOk() ([]Widget, bool) {
	if o == nil || IsNil(o.Widgets) {
		return nil, false
	}
	return o.Widgets, true
}

// HasWidgets returns a boolean if a field has been set.
func (o *Dashboardv2) HasWidgets() bool {
	if o != nil && !IsNil(o.Widgets) {
		return true
	}

	return false
}

// SetWidgets gets a reference to the given []Widget and assigns it to the Widgets field.
func (o *Dashboardv2) SetWidgets(v []Widget) {
	o.Widgets = v
}

func (o Dashboardv2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dashboardv2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CoverImage) {
		toSerialize["cover_image"] = o.CoverImage
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SoftDeleted) {
		toSerialize["soft_deleted"] = o.SoftDeleted
	}
	if !IsNil(o.Widgets) {
		toSerialize["widgets"] = o.Widgets
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Dashboardv2) UnmarshalJSON(data []byte) (err error) {
	varDashboardv2 := _Dashboardv2{}

	err = json.Unmarshal(data, &varDashboardv2)

	if err != nil {
		return err
	}

	*o = Dashboardv2(varDashboardv2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cover_image")
		delete(additionalProperties, "name")
		delete(additionalProperties, "soft_deleted")
		delete(additionalProperties, "widgets")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDashboardv2 struct {
	value *Dashboardv2
	isSet bool
}

func (v NullableDashboardv2) Get() *Dashboardv2 {
	return v.value
}

func (v *NullableDashboardv2) Set(val *Dashboardv2) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardv2) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardv2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardv2(val *Dashboardv2) *NullableDashboardv2 {
	return &NullableDashboardv2{value: val, isSet: true}
}

func (v NullableDashboardv2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardv2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


