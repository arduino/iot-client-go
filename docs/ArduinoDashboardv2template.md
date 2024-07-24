# ArduinoDashboardv2template

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoverImage** | Pointer to **string** | The cover image of the dashboard | [optional] 
**Id** | Pointer to **string** | The friendly ID of the dashboard | [optional] 
**Name** | **string** | The friendly name of the dashboard | 
**Widgets** | Pointer to [**[]ArduinoWidgetv2template**](ArduinoWidgetv2template.md) | ArduinoWidgetv2templateCollection is the media type for an array of ArduinoWidgetv2template (default view) | [optional] 

## Methods

### NewArduinoDashboardv2template

`func NewArduinoDashboardv2template(name string, ) *ArduinoDashboardv2template`

NewArduinoDashboardv2template instantiates a new ArduinoDashboardv2template object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoDashboardv2templateWithDefaults

`func NewArduinoDashboardv2templateWithDefaults() *ArduinoDashboardv2template`

NewArduinoDashboardv2templateWithDefaults instantiates a new ArduinoDashboardv2template object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoverImage

`func (o *ArduinoDashboardv2template) GetCoverImage() string`

GetCoverImage returns the CoverImage field if non-nil, zero value otherwise.

### GetCoverImageOk

`func (o *ArduinoDashboardv2template) GetCoverImageOk() (*string, bool)`

GetCoverImageOk returns a tuple with the CoverImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverImage

`func (o *ArduinoDashboardv2template) SetCoverImage(v string)`

SetCoverImage sets CoverImage field to given value.

### HasCoverImage

`func (o *ArduinoDashboardv2template) HasCoverImage() bool`

HasCoverImage returns a boolean if a field has been set.

### GetId

`func (o *ArduinoDashboardv2template) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArduinoDashboardv2template) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArduinoDashboardv2template) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ArduinoDashboardv2template) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ArduinoDashboardv2template) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArduinoDashboardv2template) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArduinoDashboardv2template) SetName(v string)`

SetName sets Name field to given value.


### GetWidgets

`func (o *ArduinoDashboardv2template) GetWidgets() []ArduinoWidgetv2template`

GetWidgets returns the Widgets field if non-nil, zero value otherwise.

### GetWidgetsOk

`func (o *ArduinoDashboardv2template) GetWidgetsOk() (*[]ArduinoWidgetv2template, bool)`

GetWidgetsOk returns a tuple with the Widgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgets

`func (o *ArduinoDashboardv2template) SetWidgets(v []ArduinoWidgetv2template)`

SetWidgets sets Widgets field to given value.

### HasWidgets

`func (o *ArduinoDashboardv2template) HasWidgets() bool`

HasWidgets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


