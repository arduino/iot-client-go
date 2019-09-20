/*
 * Iot API
 *
 * Collection of all public API endpoints.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iot
import (
	"time"
)
// ArduinoProperty ArduinoProperty media type (default view)
type ArduinoProperty struct {
	// Creation date of the property
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Delete date of the property
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// The api reference of this property
	Href string `json:"href"`
	// The id of the property
	Id string `json:"id"`
	// Last value of this property
	LastValue interface{} `json:"last_value,omitempty"`
	// Maximum value of this property
	MaxValue float64 `json:"max_value,omitempty"`
	// Minimum value of this property
	MinValue float64 `json:"min_value,omitempty"`
	// The friendly name of the property
	Name string `json:"name"`
	// The permission of the property
	Permission string `json:"permission"`
	// If true, data will persist into a timeseries database
	Persist bool `json:"persist,omitempty"`
	// The id of the thing
	ThingId string `json:"thing_id"`
	// The type of the property
	Type string `json:"type"`
	// The update frequency in seconds, or the amount of the property has to change in order to trigger an update
	UpdateParameter float64 `json:"update_parameter,omitempty"`
	// The update strategy for the property value
	UpdateStrategy string `json:"update_strategy"`
	// Update date of the property
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Last update timestamp of this property
	ValueUpdatedAt time.Time `json:"value_updated_at,omitempty"`
	// The sketch variable name of the property
	VariableName string `json:"variable_name,omitempty"`
}
