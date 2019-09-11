/*
 * Iot API
 *
 * Collection of all public API endpoints.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iot

// PropertyStringValuePayload describes a property value
type PropertyStringValue struct {
	// The device who send the property
	DeviceId string `json:"device_id,omitempty"`
	// The property value, as string
	Value string `json:"value"`
}
