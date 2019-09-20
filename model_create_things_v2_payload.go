/*
 * Iot API
 *
 * Collection of all public API endpoints.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iot
// CreateThingsV2Payload ThingPayload describes a thing
type CreateThingsV2Payload struct {
	// The arn of the associated device
	DeviceId string `json:"device_id,omitempty"`
	// The id of the thing
	Id string `json:"id,omitempty"`
	// The friendly name of the thing
	Name string `json:"name"`
	// Webhook uri
	WebhookActive bool `json:"webhook_active,omitempty"`
	// Webhook uri
	WebhookUri string `json:"webhook_uri,omitempty"`
}
