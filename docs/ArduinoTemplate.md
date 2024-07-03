# ArduinoTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dashboards** | Pointer to **[]string** |  | [optional] 
**Things** | [**[]ArduinoThingresult**](ArduinoThingresult.md) | ArduinoThingresultCollection is the media type for an array of ArduinoThingresult (default view) | 
**Triggers** | Pointer to **[]string** |  | [optional] 

## Methods

### NewArduinoTemplate

`func NewArduinoTemplate(things []ArduinoThingresult, ) *ArduinoTemplate`

NewArduinoTemplate instantiates a new ArduinoTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoTemplateWithDefaults

`func NewArduinoTemplateWithDefaults() *ArduinoTemplate`

NewArduinoTemplateWithDefaults instantiates a new ArduinoTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDashboards

`func (o *ArduinoTemplate) GetDashboards() []string`

GetDashboards returns the Dashboards field if non-nil, zero value otherwise.

### GetDashboardsOk

`func (o *ArduinoTemplate) GetDashboardsOk() (*[]string, bool)`

GetDashboardsOk returns a tuple with the Dashboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboards

`func (o *ArduinoTemplate) SetDashboards(v []string)`

SetDashboards sets Dashboards field to given value.

### HasDashboards

`func (o *ArduinoTemplate) HasDashboards() bool`

HasDashboards returns a boolean if a field has been set.

### GetThings

`func (o *ArduinoTemplate) GetThings() []ArduinoThingresult`

GetThings returns the Things field if non-nil, zero value otherwise.

### GetThingsOk

`func (o *ArduinoTemplate) GetThingsOk() (*[]ArduinoThingresult, bool)`

GetThingsOk returns a tuple with the Things field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThings

`func (o *ArduinoTemplate) SetThings(v []ArduinoThingresult)`

SetThings sets Things field to given value.


### GetTriggers

`func (o *ArduinoTemplate) GetTriggers() []string`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *ArduinoTemplate) GetTriggersOk() (*[]string, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *ArduinoTemplate) SetTriggers(v []string)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *ArduinoTemplate) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


