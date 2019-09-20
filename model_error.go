/*
 * Iot API
 *
 * Collection of all public API endpoints.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iot
// ModelError Error response media type (default view)
type ModelError struct {
	// an application-specific error code, expressed as a string value.
	Code string `json:"code,omitempty"`
	// a human-readable explanation specific to this occurrence of the problem.
	Detail string `json:"detail,omitempty"`
	// a unique identifier for this particular occurrence of the problem.
	Id string `json:"id,omitempty"`
	// a meta object containing non-standard meta-information about the error.
	Meta map[string]interface{} `json:"meta,omitempty"`
	// the HTTP status code applicable to this problem
	Status int64 `json:"status,omitempty"`
}
