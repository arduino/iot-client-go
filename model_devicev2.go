/*
 * Arduino IoT Cloud API
 *
 *  Provides a set of endpoints to manage Arduino IoT Cloud **Devices**, **Things**, **Properties** and **Timeseries**. This API can be called just with any HTTP Client, or using one of these clients:  * [Javascript NPM package](https://www.npmjs.com/package/@arduino/arduino-iot-client)  * [Python PYPI Package](https://pypi.org/project/arduino-iot-client/)  * [Golang Module](https://github.com/arduino/iot-client-go)
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iot
// Devicev2 DeviceV2 describes a device.
type Devicev2 struct {
	// The type of the connections selected by the user when multiple connections are available
	ConnectionType string `json:"connection_type,omitempty"`
	// The fully qualified board name
	Fqbn string `json:"fqbn,omitempty"`
	// The friendly name of the device
	Name string `json:"name,omitempty"`
	// The serial uuid of the device
	Serial string `json:"serial,omitempty"`
	// The type of the device
	Type string `json:"type,omitempty"`
	// The user_id associated to the device. If absent it will be inferred from the authentication header
	UserId string `json:"user_id,omitempty"`
	// The version of the NINA/WIFI101 firmware running on the device
	WifiFwVersion string `json:"wifi_fw_version,omitempty"`
}
