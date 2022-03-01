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
// ArduinoTimezone ArduinoTimezone media type (default view)
type ArduinoTimezone struct {
	// Name of the time zone.
	Name string `json:"name"`
	// Current UTC DST offset in seconds.
	Offset int64 `json:"offset"`
	// Date until the offset is valid.
	Until time.Time `json:"until,omitempty"`
}