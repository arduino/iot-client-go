/*
Arduino IoT Cloud API

 Provides a set of endpoints to manage Arduino IoT Cloud **Devices**, **Things**, **Properties** and **Timeseries**. This API can be called just with any HTTP Client, or using one of these clients:  * [Javascript NPM package](https://www.npmjs.com/package/@arduino/arduino-iot-client)  * [Python PYPI Package](https://pypi.org/project/arduino-iot-client/)  * [Golang Module](https://github.com/arduino/iot-client-go)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iot

import (
	"encoding/json"
	"time"
)

// checks if the ArduinoSeriesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArduinoSeriesResponse{}

// ArduinoSeriesResponse ArduinoSeriesResponse media type (default view)
type ArduinoSeriesResponse struct {
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
	// Max of values
	SeriesLimit *int64 `json:"series_limit,omitempty"`
	// Status of the response
	Status string `json:"status"`
	// Timestamp in RFC3339
	Times []time.Time `json:"times"`
	// To date
	ToDate time.Time `json:"to_date"`
	// Values in Float
	Values []float64 `json:"values"`
}

// NewArduinoSeriesResponse instantiates a new ArduinoSeriesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArduinoSeriesResponse(countValues int64, fromDate time.Time, interval int64, query string, respVersion int64, status string, times []time.Time, toDate time.Time, values []float64) *ArduinoSeriesResponse {
	this := ArduinoSeriesResponse{}
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

// NewArduinoSeriesResponseWithDefaults instantiates a new ArduinoSeriesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArduinoSeriesResponseWithDefaults() *ArduinoSeriesResponse {
	this := ArduinoSeriesResponse{}
	var message string = ""
	this.Message = &message
	return &this
}

// GetCountValues returns the CountValues field value
func (o *ArduinoSeriesResponse) GetCountValues() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CountValues
}

// GetCountValuesOk returns a tuple with the CountValues field value
// and a boolean to check if the value has been set.
func (o *ArduinoSeriesResponse) GetCountValuesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountValues, true
}

// SetCountValues sets field value
func (o *ArduinoSeriesResponse) SetCountValues(v int64) {
	o.CountValues = v
}

// GetFromDate returns the FromDate field value
func (o *ArduinoSeriesResponse) GetFromDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.FromDate
}

// GetFromDateOk returns a tuple with the FromDate field value
// and a boolean to check if the value has been set.
func (o *ArduinoSeriesResponse) GetFromDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromDate, true
}

// SetFromDate sets field value
func (o *ArduinoSeriesResponse) SetFromDate(v time.Time) {
	o.FromDate = v
}

// GetInterval returns the Interval field value
func (o *ArduinoSeriesResponse) GetInterval() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *ArduinoSeriesResponse) GetIntervalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *ArduinoSeriesResponse) SetInterval(v int64) {
	o.Interval = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ArduinoSeriesResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoSeriesResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ArduinoSeriesResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ArduinoSeriesResponse) SetMessage(v string) {
	o.Message = &v
}

// GetQuery returns the Query field value
func (o *ArduinoSeriesResponse) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *ArduinoSeriesResponse) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *ArduinoSeriesResponse) SetQuery(v string) {
	o.Query = v
}

// GetRespVersion returns the RespVersion field value
func (o *ArduinoSeriesResponse) GetRespVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RespVersion
}

// GetRespVersionOk returns a tuple with the RespVersion field value
// and a boolean to check if the value has been set.
func (o *ArduinoSeriesResponse) GetRespVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RespVersion, true
}

// SetRespVersion sets field value
func (o *ArduinoSeriesResponse) SetRespVersion(v int64) {
	o.RespVersion = v
}

// GetSeriesLimit returns the SeriesLimit field value if set, zero value otherwise.
func (o *ArduinoSeriesResponse) GetSeriesLimit() int64 {
	if o == nil || IsNil(o.SeriesLimit) {
		var ret int64
		return ret
	}
	return *o.SeriesLimit
}

// GetSeriesLimitOk returns a tuple with the SeriesLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArduinoSeriesResponse) GetSeriesLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.SeriesLimit) {
		return nil, false
	}
	return o.SeriesLimit, true
}

// HasSeriesLimit returns a boolean if a field has been set.
func (o *ArduinoSeriesResponse) HasSeriesLimit() bool {
	if o != nil && !IsNil(o.SeriesLimit) {
		return true
	}

	return false
}

// SetSeriesLimit gets a reference to the given int64 and assigns it to the SeriesLimit field.
func (o *ArduinoSeriesResponse) SetSeriesLimit(v int64) {
	o.SeriesLimit = &v
}

// GetStatus returns the Status field value
func (o *ArduinoSeriesResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ArduinoSeriesResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ArduinoSeriesResponse) SetStatus(v string) {
	o.Status = v
}

// GetTimes returns the Times field value
func (o *ArduinoSeriesResponse) GetTimes() []time.Time {
	if o == nil {
		var ret []time.Time
		return ret
	}

	return o.Times
}

// GetTimesOk returns a tuple with the Times field value
// and a boolean to check if the value has been set.
func (o *ArduinoSeriesResponse) GetTimesOk() ([]time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Times, true
}

// SetTimes sets field value
func (o *ArduinoSeriesResponse) SetTimes(v []time.Time) {
	o.Times = v
}

// GetToDate returns the ToDate field value
func (o *ArduinoSeriesResponse) GetToDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ToDate
}

// GetToDateOk returns a tuple with the ToDate field value
// and a boolean to check if the value has been set.
func (o *ArduinoSeriesResponse) GetToDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToDate, true
}

// SetToDate sets field value
func (o *ArduinoSeriesResponse) SetToDate(v time.Time) {
	o.ToDate = v
}

// GetValues returns the Values field value
func (o *ArduinoSeriesResponse) GetValues() []float64 {
	if o == nil {
		var ret []float64
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *ArduinoSeriesResponse) GetValuesOk() ([]float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *ArduinoSeriesResponse) SetValues(v []float64) {
	o.Values = v
}

func (o ArduinoSeriesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArduinoSeriesResponse) ToMap() (map[string]interface{}, error) {
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

type NullableArduinoSeriesResponse struct {
	value *ArduinoSeriesResponse
	isSet bool
}

func (v NullableArduinoSeriesResponse) Get() *ArduinoSeriesResponse {
	return v.value
}

func (v *NullableArduinoSeriesResponse) Set(val *ArduinoSeriesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableArduinoSeriesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableArduinoSeriesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArduinoSeriesResponse(val *ArduinoSeriesResponse) *NullableArduinoSeriesResponse {
	return &NullableArduinoSeriesResponse{value: val, isSet: true}
}

func (v NullableArduinoSeriesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArduinoSeriesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


