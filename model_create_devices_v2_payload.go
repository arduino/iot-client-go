/*
 * Iot API
 *
 * Collection of all public API endpoints.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iot
// CreateDevicesV2Payload DeviceV2 describes a device.
type CreateDevicesV2Payload struct {
	// The fully qualified board name
	Fqbn string `json:"fqbn,omitempty"`
	// The UUID of the device
	Id string `json:"id,omitempty"`
	// The friendly name of the device
	Name string `json:"name,omitempty"`
	// The serial uuid of the device
	Serial string `json:"serial,omitempty"`
	// The type of the device
	Type string `json:"type"`
	// The user_id associated to the device. If absent it will be inferred from the authentication header
	UserId string `json:"user_id,omitempty"`
}
