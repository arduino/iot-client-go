/*
Arduino IoT Cloud API

 Provides a set of endpoints to manage Arduino IoT Cloud **Devices**, **Things**, **Properties** and **Timeseries**. This API can be called just with any HTTP Client, or using one of these clients:  * [Javascript NPM package](https://www.npmjs.com/package/@arduino/arduino-iot-client)  * [Python PYPI Package](https://pypi.org/project/arduino-iot-client/)  * [Golang Module](https://github.com/arduino/iot-client-go)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"time"
)

// checks if the ArduinoSeriesSampledResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoSeriesSampledResponse{}

// ArduinoSeriesSampledResponse ArduinoSeriesSampledResponse media type (default view)
type ArduinoSeriesSampledResponse struct {
	// Total number of values in the array 'values'
	CountValues int64 `json:"count_values"`
	// From date
	FromDate time.Time `json:"from_date"`
	// Resolution in seconds
	Interval int64 `json:"interval"`
	// If the response is different than 'ok'
	Message *string `json:"message,omitempty"`
	// Query of for the data
	Query string `json:"query"`
	// Response version
	RespVersion int64 `json:"resp_version"`
	// Maximum number of values returned after data aggregation, if any
	SeriesLimit *int64 `json:"series_limit,omitempty"`
	// Status of the response
	Status string `json:"status"`
	// Timestamp in RFC3339
	Times []time.Time `json:"times"`
	// To date
	ToDate time.Time `json:"to_date"`
	// Values in Float
	Values []interface{} `json:"values"`
}

// NewArduinoSeriesSampledResponse instantiates a new ArduinoSeriesSampledResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoSeriesSampledResponse(countValues int64, fromDate time.Time, interval int64, query string, respVersion int64, status string, times []time.Time, toDate time.Time, values []interface{}) *ArduinoSeriesSampledResponse {
	this := ArduinoSeriesSampledResponse{}
	this.CountValues = countValues
	this.FromDate = fromDate
	this.Interval = interval
	var message string = ""
	this.Message = &message
	this.Query = query
	this.RespVersion = respVersion
	this.Status = status
	this.Times = times
	this.ToDate = toDate
	this.Values = values
	return &this
}

// NewArduinoSeriesSampledResponseWithDefaults instantiates a new ArduinoSeriesSampledResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoSeriesSampledResponseWithDefaults() *ArduinoSeriesSampledResponse {
	this := ArduinoSeriesSampledResponse{}
	var message string = ""
	this.Message = &message
	return &this
}

// GetCountValues returns the CountValues field value
func (o *ArduinoSeriesSampledResponse) GetCountValues() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CountValues
}

// GetCountValuesOk returns a tuple with the CountValues field value
// and a boolean to check if the value has been set.
func (o *ArduinoSeriesSampledResponse) GetCountValuesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountValues, true
}

// SetCountValues sets field value
func (o *ArduinoSeriesSampledResponse) SetCountValues(v int64) {
	o.CountValues = v
}

// GetFromDate returns the FromDate field value
func (o *ArduinoSeriesSampledResponse) GetFromDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.FromDate
}

// GetFromDateOk returns a tuple with the FromDate field value
// and a boolean to check if the value has been set.
func (o *ArduinoSeriesSampledResponse) GetFromDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromDate, true
}

// SetFromDate sets field value
func (o *ArduinoSeriesSampledResponse) SetFromDate(v time.Time) {
	o.FromDate = v
}

// GetInterval returns the Interval field value
func (o *ArduinoSeriesSampledResponse) GetInterval() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *ArduinoSeriesSampledResponse) GetIntervalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *ArduinoSeriesSampledResponse) SetInterval(v int64) {
	o.Interval = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ArduinoSeriesSampledResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoSeriesSampledResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ArduinoSeriesSampledResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ArduinoSeriesSampledResponse) SetMessage(v string) {
	o.Message = &v
}

// GetQuery returns the Query field value
func (o *ArduinoSeriesSampledResponse) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *ArduinoSeriesSampledResponse) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *ArduinoSeriesSampledResponse) SetQuery(v string) {
	o.Query = v
}

// GetRespVersion returns the RespVersion field value
func (o *ArduinoSeriesSampledResponse) GetRespVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RespVersion
}

// GetRespVersionOk returns a tuple with the RespVersion field value
// and a boolean to check if the value has been set.
func (o *ArduinoSeriesSampledResponse) GetRespVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RespVersion, true
}

// SetRespVersion sets field value
func (o *ArduinoSeriesSampledResponse) SetRespVersion(v int64) {
	o.RespVersion = v
}

// GetSeriesLimit returns the SeriesLimit field value if set, zero value otherwise.
func (o *ArduinoSeriesSampledResponse) GetSeriesLimit() int64 {
	if o == nil || IsNil(o.SeriesLimit) {
		var ret int64
		return ret
	}
	return *o.SeriesLimit
}

// GetSeriesLimitOk returns a tuple with the SeriesLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoSeriesSampledResponse) GetSeriesLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.SeriesLimit) {
		return nil, false
	}
	return o.SeriesLimit, true
}

// HasSeriesLimit returns a boolean if a field has been set.
func (o *ArduinoSeriesSampledResponse) HasSeriesLimit() bool {
	if o != nil && !IsNil(o.SeriesLimit) {
		return true
	}

	return false
}

// SetSeriesLimit gets a reference to the given int64 and assigns it to the SeriesLimit field.
func (o *ArduinoSeriesSampledResponse) SetSeriesLimit(v int64) {
	o.SeriesLimit = &v
}

// GetStatus returns the Status field value
func (o *ArduinoSeriesSampledResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ArduinoSeriesSampledResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ArduinoSeriesSampledResponse) SetStatus(v string) {
	o.Status = v
}

// GetTimes returns the Times field value
func (o *ArduinoSeriesSampledResponse) GetTimes() []time.Time {
	if o == nil {
		var ret []time.Time
		return ret
	}

	return o.Times
}

// GetTimesOk returns a tuple with the Times field value
// and a boolean to check if the value has been set.
func (o *ArduinoSeriesSampledResponse) GetTimesOk() ([]time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Times, true
}

// SetTimes sets field value
func (o *ArduinoSeriesSampledResponse) SetTimes(v []time.Time) {
	o.Times = v
}

// GetToDate returns the ToDate field value
func (o *ArduinoSeriesSampledResponse) GetToDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ToDate
}

// GetToDateOk returns a tuple with the ToDate field value
// and a boolean to check if the value has been set.
func (o *ArduinoSeriesSampledResponse) GetToDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToDate, true
}

// SetToDate sets field value
func (o *ArduinoSeriesSampledResponse) SetToDate(v time.Time) {
	o.ToDate = v
}

// GetValues returns the Values field value
func (o *ArduinoSeriesSampledResponse) GetValues() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *ArduinoSeriesSampledResponse) GetValuesOk() ([]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *ArduinoSeriesSampledResponse) SetValues(v []interface{}) {
	o.Values = v
}

func (o ArduinoSeriesSampledResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoSeriesSampledResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count_values"] = o.CountValues
	toSerialize["from_date"] = o.FromDate
	toSerialize["interval"] = o.Interval
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	toSerialize["query"] = o.Query
	toSerialize["resp_version"] = o.RespVersion
	if !IsNil(o.SeriesLimit) {
		toSerialize["series_limit"] = o.SeriesLimit
	}
	toSerialize["status"] = o.Status
	toSerialize["times"] = o.Times
	toSerialize["to_date"] = o.ToDate
	toSerialize["values"] = o.Values
	return toSerialize, nil
}

type NullableArduinoSeriesSampledResponse struct {
	value *ArduinoSeriesSampledResponse
	isSet bool
}

func (v NullableArduinoSeriesSampledResponse) Get() *ArduinoSeriesSampledResponse {
	return v.value
}

func (v *NullableArduinoSeriesSampledResponse) Set(val *ArduinoSeriesSampledResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoSeriesSampledResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoSeriesSampledResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoSeriesSampledResponse(val *ArduinoSeriesSampledResponse) *NullableArduinoSeriesSampledResponse {
	return &NullableArduinoSeriesSampledResponse{value: val, isSet: true}
}

func (v NullableArduinoSeriesSampledResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoSeriesSampledResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


