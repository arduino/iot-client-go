/*
 * Arduino IoT Cloud API
 *
 *  Provides a set of endpoints to manage Arduino IoT Cloud **Devices**, **Things**, **Properties** and **Timeseries**. This API can be called just with any HTTP Client, or using one of these clients:  * [Javascript NPM package](https://www.npmjs.com/package/@arduino/arduino-iot-client)  * [Python PYPI Package](https://pypi.org/project/arduino-iot-client/)  * [Golang Module](https://github.com/arduino/iot-client-go)
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iot
import (
	"time"
)
// ArduinoSeriesRawLastValueResponse ArduinoSeriesRawLastValueResponse media type (default view)
type ArduinoSeriesRawLastValueResponse struct {
	// Total number of values in the array 'values'
	CountValues int64 `json:"count_values"`
	// Property id
	PropertyId string `json:"property_id"`
	// Thing id
	ThingId string `json:"thing_id"`
	// Timestamp in RFC3339
	Times []time.Time `json:"times"`
	// Values can be in Float, String, Bool, Object
	Values []interface{} `json:"values"`
}
