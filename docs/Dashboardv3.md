# Dashboardv3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoverImage** | Pointer to **string** | The cover image of the dashboard | [optional] 
**Name** | Pointer to **string** | The friendly name of the dashboard | [optional] 
**Pages** | Pointer to [**[]Pagepayload**](Pagepayload.md) | List of sub-pages | [optional] 
**SoftDeleted** | Pointer to **bool** | If false, restore the thing from the soft deletion | [optional] [default to false]
**Widgets** | Pointer to [**[]Widgetv3**](Widgetv3.md) | Widgets attached to this dashboard | [optional] 

## Methods

### NewDashboardv3

`func NewDashboardv3() *Dashboardv3`

NewDashboardv3 instantiates a new Dashboardv3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardv3WithDefaults

`func NewDashboardv3WithDefaults() *Dashboardv3`

NewDashboardv3WithDefaults instantiates a new Dashboardv3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoverImage

`func (o *Dashboardv3) GetCoverImage() string`

GetCoverImage returns the CoverImage field if non-nil, zero value otherwise.

### GetCoverImageOk

`func (o *Dashboardv3) GetCoverImageOk() (*string, bool)`

GetCoverImageOk returns a tuple with the CoverImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverImage

`func (o *Dashboardv3) SetCoverImage(v string)`

SetCoverImage sets CoverImage field to given value.

### HasCoverImage

`func (o *Dashboardv3) HasCoverImage() bool`

HasCoverImage returns a boolean if a field has been set.

### GetName

`func (o *Dashboardv3) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Dashboardv3) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Dashboardv3) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Dashboardv3) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPages

`func (o *Dashboardv3) GetPages() []Pagepayload`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *Dashboardv3) GetPagesOk() (*[]Pagepayload, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *Dashboardv3) SetPages(v []Pagepayload)`

SetPages sets Pages field to given value.

### HasPages

`func (o *Dashboardv3) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetSoftDeleted

`func (o *Dashboardv3) GetSoftDeleted() bool`

GetSoftDeleted returns the SoftDeleted field if non-nil, zero value otherwise.

### GetSoftDeletedOk

`func (o *Dashboardv3) GetSoftDeletedOk() (*bool, bool)`

GetSoftDeletedOk returns a tuple with the SoftDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftDeleted

`func (o *Dashboardv3) SetSoftDeleted(v bool)`

SetSoftDeleted sets SoftDeleted field to given value.

### HasSoftDeleted

`func (o *Dashboardv3) HasSoftDeleted() bool`

HasSoftDeleted returns a boolean if a field has been set.

### GetWidgets

`func (o *Dashboardv3) GetWidgets() []Widgetv3`

GetWidgets returns the Widgets field if non-nil, zero value otherwise.

### GetWidgetsOk

`func (o *Dashboardv3) GetWidgetsOk() (*[]Widgetv3, bool)`

GetWidgetsOk returns a tuple with the Widgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgets

`func (o *Dashboardv3) SetWidgets(v []Widgetv3)`

SetWidgets sets Widgets field to given value.

### HasWidgets

`func (o *Dashboardv3) HasWidgets() bool`

HasWidgets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


