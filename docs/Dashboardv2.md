# Dashboardv2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoverImage** | Pointer to **string** | The cover image of the dashboard | [optional] 
**Name** | Pointer to **string** | The friendly name of the dashboard | [optional] 
**Widgets** | Pointer to [**[]Widget**](Widget.md) | Widgets attached to this dashboard | [optional] 

## Methods

### NewDashboardv2

`func NewDashboardv2() *Dashboardv2`

NewDashboardv2 instantiates a new Dashboardv2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardv2WithDefaults

`func NewDashboardv2WithDefaults() *Dashboardv2`

NewDashboardv2WithDefaults instantiates a new Dashboardv2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoverImage

`func (o *Dashboardv2) GetCoverImage() string`

GetCoverImage returns the CoverImage field if non-nil, zero value otherwise.

### GetCoverImageOk

`func (o *Dashboardv2) GetCoverImageOk() (*string, bool)`

GetCoverImageOk returns a tuple with the CoverImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverImage

`func (o *Dashboardv2) SetCoverImage(v string)`

SetCoverImage sets CoverImage field to given value.

### HasCoverImage

`func (o *Dashboardv2) HasCoverImage() bool`

HasCoverImage returns a boolean if a field has been set.

### GetName

`func (o *Dashboardv2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Dashboardv2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Dashboardv2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Dashboardv2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWidgets

`func (o *Dashboardv2) GetWidgets() []Widget`

GetWidgets returns the Widgets field if non-nil, zero value otherwise.

### GetWidgetsOk

`func (o *Dashboardv2) GetWidgetsOk() (*[]Widget, bool)`

GetWidgetsOk returns a tuple with the Widgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgets

`func (o *Dashboardv2) SetWidgets(v []Widget)`

SetWidgets sets Widgets field to given value.

### HasWidgets

`func (o *Dashboardv2) HasWidgets() bool`

HasWidgets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


