/*
Arduino IoT Cloud API

 Provides a set of endpoints to manage Arduino IoT Cloud **Devices**, **Things**, **Properties** and **Timeseries**. This API can be called just with any HTTP Client, or using one of these clients:  * [Javascript NPM package](https://www.npmjs.com/package/@arduino/arduino-iot-client)  * [Python PYPI Package](https://pypi.org/project/arduino-iot-client/)  * [Golang Module](https://github.com/arduino/iot-client-go)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iot

import (
	"encoding/json"
)

// checks if the ArduinoWidgetv2template type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoWidgetv2template{}

// ArduinoWidgetv2template ArduinoWidgetv2template media type (default view)
type ArduinoWidgetv2template struct {
	// Widget current height for desktop
	Height int64 `json:"height"`
	// Widget current height for mobile
	HeightMobile *int64 `json:"height_mobile,omitempty"`
	// The name of the widget
	Name *string `json:"name,omitempty"`
	// Widget options
	Options map[string]interface{} `json:"options"`
	// The type of the widget
	Type string `json:"type"`
	// ArduinoTemplatevariableCollection is the media type for an array of ArduinoTemplatevariable (default view)
	Variables []ArduinoTemplatevariable `json:"variables,omitempty"`
	// Widget current width for desktop
	Width int64 `json:"width"`
	// Widget current width for mobile
	WidthMobile *int64 `json:"width_mobile,omitempty"`
	// Widget x position for desktop
	X int64 `json:"x"`
	// Widget x position for mobile
	XMobile *int64 `json:"x_mobile,omitempty"`
	// Widget y position for desktop
	Y int64 `json:"y"`
	// Widget y position for mobile
	YMobile *int64 `json:"y_mobile,omitempty"`
}

// NewArduinoWidgetv2template instantiates a new ArduinoWidgetv2template object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoWidgetv2template(height int64, options map[string]interface{}, type_ string, width int64, x int64, y int64) *ArduinoWidgetv2template {
	this := ArduinoWidgetv2template{}
	this.Height = height
	this.Options = options
	this.Type = type_
	this.Width = width
	this.X = x
	this.Y = y
	return &this
}

// NewArduinoWidgetv2templateWithDefaults instantiates a new ArduinoWidgetv2template object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoWidgetv2templateWithDefaults() *ArduinoWidgetv2template {
	this := ArduinoWidgetv2template{}
	return &this
}

// GetHeight returns the Height field value
func (o *ArduinoWidgetv2template) GetHeight() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *ArduinoWidgetv2template) GetHeightOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *ArduinoWidgetv2template) SetHeight(v int64) {
	o.Height = v
}

// GetHeightMobile returns the HeightMobile field value if set, zero value otherwise.
func (o *ArduinoWidgetv2template) GetHeightMobile() int64 {
	if o == nil || IsNil(o.HeightMobile) {
		var ret int64
		return ret
	}
	return *o.HeightMobile
}

// GetHeightMobileOk returns a tuple with the HeightMobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoWidgetv2template) GetHeightMobileOk() (*int64, bool) {
	if o == nil || IsNil(o.HeightMobile) {
		return nil, false
	}
	return o.HeightMobile, true
}

// HasHeightMobile returns a boolean if a field has been set.
func (o *ArduinoWidgetv2template) HasHeightMobile() bool {
	if o != nil && !IsNil(o.HeightMobile) {
		return true
	}

	return false
}

// SetHeightMobile gets a reference to the given int64 and assigns it to the HeightMobile field.
func (o *ArduinoWidgetv2template) SetHeightMobile(v int64) {
	o.HeightMobile = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ArduinoWidgetv2template) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoWidgetv2template) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ArduinoWidgetv2template) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ArduinoWidgetv2template) SetName(v string) {
	o.Name = &v
}

// GetOptions returns the Options field value
func (o *ArduinoWidgetv2template) GetOptions() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *ArduinoWidgetv2template) GetOptionsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *ArduinoWidgetv2template) SetOptions(v map[string]interface{}) {
	o.Options = v
}

// GetType returns the Type field value
func (o *ArduinoWidgetv2template) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ArduinoWidgetv2template) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ArduinoWidgetv2template) SetType(v string) {
	o.Type = v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *ArduinoWidgetv2template) GetVariables() []ArduinoTemplatevariable {
	if o == nil || IsNil(o.Variables) {
		var ret []ArduinoTemplatevariable
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoWidgetv2template) GetVariablesOk() ([]ArduinoTemplatevariable, bool) {
	if o == nil || IsNil(o.Variables) {
		return nil, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *ArduinoWidgetv2template) HasVariables() bool {
	if o != nil && !IsNil(o.Variables) {
		return true
	}

	return false
}

// SetVariables gets a reference to the given []ArduinoTemplatevariable and assigns it to the Variables field.
func (o *ArduinoWidgetv2template) SetVariables(v []ArduinoTemplatevariable) {
	o.Variables = v
}

// GetWidth returns the Width field value
func (o *ArduinoWidgetv2template) GetWidth() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *ArduinoWidgetv2template) GetWidthOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *ArduinoWidgetv2template) SetWidth(v int64) {
	o.Width = v
}

// GetWidthMobile returns the WidthMobile field value if set, zero value otherwise.
func (o *ArduinoWidgetv2template) GetWidthMobile() int64 {
	if o == nil || IsNil(o.WidthMobile) {
		var ret int64
		return ret
	}
	return *o.WidthMobile
}

// GetWidthMobileOk returns a tuple with the WidthMobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoWidgetv2template) GetWidthMobileOk() (*int64, bool) {
	if o == nil || IsNil(o.WidthMobile) {
		return nil, false
	}
	return o.WidthMobile, true
}

// HasWidthMobile returns a boolean if a field has been set.
func (o *ArduinoWidgetv2template) HasWidthMobile() bool {
	if o != nil && !IsNil(o.WidthMobile) {
		return true
	}

	return false
}

// SetWidthMobile gets a reference to the given int64 and assigns it to the WidthMobile field.
func (o *ArduinoWidgetv2template) SetWidthMobile(v int64) {
	o.WidthMobile = &v
}

// GetX returns the X field value
func (o *ArduinoWidgetv2template) GetX() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *ArduinoWidgetv2template) GetXOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value
func (o *ArduinoWidgetv2template) SetX(v int64) {
	o.X = v
}

// GetXMobile returns the XMobile field value if set, zero value otherwise.
func (o *ArduinoWidgetv2template) GetXMobile() int64 {
	if o == nil || IsNil(o.XMobile) {
		var ret int64
		return ret
	}
	return *o.XMobile
}

// GetXMobileOk returns a tuple with the XMobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoWidgetv2template) GetXMobileOk() (*int64, bool) {
	if o == nil || IsNil(o.XMobile) {
		return nil, false
	}
	return o.XMobile, true
}

// HasXMobile returns a boolean if a field has been set.
func (o *ArduinoWidgetv2template) HasXMobile() bool {
	if o != nil && !IsNil(o.XMobile) {
		return true
	}

	return false
}

// SetXMobile gets a reference to the given int64 and assigns it to the XMobile field.
func (o *ArduinoWidgetv2template) SetXMobile(v int64) {
	o.XMobile = &v
}

// GetY returns the Y field value
func (o *ArduinoWidgetv2template) GetY() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *ArduinoWidgetv2template) GetYOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Y, true
}

// SetY sets field value
func (o *ArduinoWidgetv2template) SetY(v int64) {
	o.Y = v
}

// GetYMobile returns the YMobile field value if set, zero value otherwise.
func (o *ArduinoWidgetv2template) GetYMobile() int64 {
	if o == nil || IsNil(o.YMobile) {
		var ret int64
		return ret
	}
	return *o.YMobile
}

// GetYMobileOk returns a tuple with the YMobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoWidgetv2template) GetYMobileOk() (*int64, bool) {
	if o == nil || IsNil(o.YMobile) {
		return nil, false
	}
	return o.YMobile, true
}

// HasYMobile returns a boolean if a field has been set.
func (o *ArduinoWidgetv2template) HasYMobile() bool {
	if o != nil && !IsNil(o.YMobile) {
		return true
	}

	return false
}

// SetYMobile gets a reference to the given int64 and assigns it to the YMobile field.
func (o *ArduinoWidgetv2template) SetYMobile(v int64) {
	o.YMobile = &v
}

func (o ArduinoWidgetv2template) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoWidgetv2template) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["height"] = o.Height
	if !IsNil(o.HeightMobile) {
		toSerialize["height_mobile"] = o.HeightMobile
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["options"] = o.Options
	toSerialize["type"] = o.Type
	if !IsNil(o.Variables) {
		toSerialize["variables"] = o.Variables
	}
	toSerialize["width"] = o.Width
	if !IsNil(o.WidthMobile) {
		toSerialize["width_mobile"] = o.WidthMobile
	}
	toSerialize["x"] = o.X
	if !IsNil(o.XMobile) {
		toSerialize["x_mobile"] = o.XMobile
	}
	toSerialize["y"] = o.Y
	if !IsNil(o.YMobile) {
		toSerialize["y_mobile"] = o.YMobile
	}
	return toSerialize, nil
}

type NullableArduinoWidgetv2template struct {
	value *ArduinoWidgetv2template
	isSet bool
}

func (v NullableArduinoWidgetv2template) Get() *ArduinoWidgetv2template {
	return v.value
}

func (v *NullableArduinoWidgetv2template) Set(val *ArduinoWidgetv2template) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoWidgetv2template) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoWidgetv2template) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoWidgetv2template(val *ArduinoWidgetv2template) *NullableArduinoWidgetv2template {
	return &NullableArduinoWidgetv2template{value: val, isSet: true}
}

func (v NullableArduinoWidgetv2template) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoWidgetv2template) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

