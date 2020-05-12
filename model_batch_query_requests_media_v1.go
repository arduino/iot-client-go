/*
 * Arduino IoT Cloud API
 *
 *  Provides a set of endpoints to manage Arduino IoT Cloud **Devices**, **Things**, **Properties** and **Timeseries**. This API can be called just with any HTTP Client, or using one of these clients:  * [Javascript NPM package](https://www.npmjs.com/package/@arduino/arduino-iot-client)  * [Python PYPI Package](https://pypi.org/project/arduino-iot-client/)  * [Golang Module](https://github.com/arduino/iot-client-go)
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iot
// BatchQueryRequestsMediaV1 struct for BatchQueryRequestsMediaV1
type BatchQueryRequestsMediaV1 struct {
	// Requests
	Requests []BatchQueryRequestMediaV1 `json:"requests"`
	// Response version
	RespVersion int64 `json:"resp_version"`
}
