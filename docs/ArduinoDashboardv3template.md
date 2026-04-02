# ArduinoDashboardv3template

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoverImage** | Pointer to **string** | The cover image of the dashboard | [optional] 
**Id** | Pointer to **string** | The friendly ID of the dashboard | [optional] 
**Name** | **string** | The friendly name of the dashboard | 
**Pages** | Pointer to [**[]ArduinoPagevariable**](ArduinoPagevariable.md) | ArduinoPagevariableCollection is the media type for an array of ArduinoPagevariable (default view) | [optional] 
**Widgets** | Pointer to [**[]ArduinoWidgetv3template**](ArduinoWidgetv3template.md) | ArduinoWidgetv3templateCollection is the media type for an array of ArduinoWidgetv3template (default view) | [optional] 

## Methods

### NewArduinoDashboardv3template

`func NewArduinoDashboardv3template(name string, ) *ArduinoDashboardv3template`

NewArduinoDashboardv3template instantiates a new ArduinoDashboardv3template object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoDashboardv3templateWithDefaults

`func NewArduinoDashboardv3templateWithDefaults() *ArduinoDashboardv3template`

NewArduinoDashboardv3templateWithDefaults instantiates a new ArduinoDashboardv3template object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoverImage

`func (o *ArduinoDashboardv3template) GetCoverImage() string`

GetCoverImage returns the CoverImage field if non-nil, zero value otherwise.

### GetCoverImageOk

`func (o *ArduinoDashboardv3template) GetCoverImageOk() (*string, bool)`

GetCoverImageOk returns a tuple with the CoverImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverImage

`func (o *ArduinoDashboardv3template) SetCoverImage(v string)`

SetCoverImage sets CoverImage field to given value.

### HasCoverImage

`func (o *ArduinoDashboardv3template) HasCoverImage() bool`

HasCoverImage returns a boolean if a field has been set.

### GetId

`func (o *ArduinoDashboardv3template) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArduinoDashboardv3template) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArduinoDashboardv3template) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ArduinoDashboardv3template) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ArduinoDashboardv3template) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArduinoDashboardv3template) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArduinoDashboardv3template) SetName(v string)`

SetName sets Name field to given value.


### GetPages

`func (o *ArduinoDashboardv3template) GetPages() []ArduinoPagevariable`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *ArduinoDashboardv3template) GetPagesOk() (*[]ArduinoPagevariable, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *ArduinoDashboardv3template) SetPages(v []ArduinoPagevariable)`

SetPages sets Pages field to given value.

### HasPages

`func (o *ArduinoDashboardv3template) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetWidgets

`func (o *ArduinoDashboardv3template) GetWidgets() []ArduinoWidgetv3template`

GetWidgets returns the Widgets field if non-nil, zero value otherwise.

### GetWidgetsOk

`func (o *ArduinoDashboardv3template) GetWidgetsOk() (*[]ArduinoWidgetv3template, bool)`

GetWidgetsOk returns a tuple with the Widgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgets

`func (o *ArduinoDashboardv3template) SetWidgets(v []ArduinoWidgetv3template)`

SetWidgets sets Widgets field to given value.

### HasWidgets

`func (o *ArduinoDashboardv3template) HasWidgets() bool`

HasWidgets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


