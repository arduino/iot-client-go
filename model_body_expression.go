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

// checks if the BodyExpression type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BodyExpression{}

// BodyExpression struct for BodyExpression
type BodyExpression struct {
	// Content of the body of a message, variables are allowed
	Expression string `json:"expression"`
	// Variables used by the expression
	Variables []Variable `json:"variables,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BodyExpression BodyExpression

// NewBodyExpression instantiates a new BodyExpression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBodyExpression(expression string) *BodyExpression {
	this := BodyExpression{}
	this.Expression = expression
	return &this
}

// NewBodyExpressionWithDefaults instantiates a new BodyExpression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBodyExpressionWithDefaults() *BodyExpression {
	this := BodyExpression{}
	return &this
}

// GetExpression returns the Expression field value
func (o *BodyExpression) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *BodyExpression) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *BodyExpression) SetExpression(v string) {
	o.Expression = v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *BodyExpression) GetVariables() []Variable {
	if o == nil || IsNil(o.Variables) {
		var ret []Variable
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BodyExpression) GetVariablesOk() ([]Variable, bool) {
	if o == nil || IsNil(o.Variables) {
		return nil, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *BodyExpression) HasVariables() bool {
	if o != nil && !IsNil(o.Variables) {
		return true
	}

	return false
}

// SetVariables gets a reference to the given []Variable and assigns it to the Variables field.
func (o *BodyExpression) SetVariables(v []Variable) {
	o.Variables = v
}

func (o BodyExpression) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BodyExpression) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expression"] = o.Expression
	if !IsNil(o.Variables) {
		toSerialize["variables"] = o.Variables
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BodyExpression) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"expression",
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

	varBodyExpression := _BodyExpression{}

	err = json.Unmarshal(data, &varBodyExpression)

	if err != nil {
		return err
	}

	*o = BodyExpression(varBodyExpression)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "expression")
		delete(additionalProperties, "variables")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBodyExpression struct {
	value *BodyExpression
	isSet bool
}

func (v NullableBodyExpression) Get() *BodyExpression {
	return v.value
}

func (v *NullableBodyExpression) Set(val *BodyExpression) {
	v.value = val
	v.isSet = true
}

func (v NullableBodyExpression) IsSet() bool {
	return v.isSet
}

func (v *NullableBodyExpression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBodyExpression(val *BodyExpression) *NullableBodyExpression {
	return &NullableBodyExpression{value: val, isSet: true}
}

func (v NullableBodyExpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBodyExpression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


