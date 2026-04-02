# ArduinoDashboardv3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoverImage** | Pointer to **string** | The cover image of the dashboard | [optional] 
**CreatedBy** | Pointer to [**ArduinoDashboardowner**](ArduinoDashboardowner.md) |  | [optional] 
**Id** | **string** | The friendly name of the dashboard | 
**Name** | **string** | The friendly name of the dashboard | 
**OrganizationId** | Pointer to **string** | Id of the organization the dashboard belongs to | [optional] 
**Pages** | Pointer to [**[]ArduinoPagevariable**](ArduinoPagevariable.md) | ArduinoPagevariableCollection is the media type for an array of ArduinoPagevariable (default view) | [optional] 
**SharedBy** | Pointer to [**ArduinoDashboardshare**](ArduinoDashboardshare.md) |  | [optional] 
**SharedWith** | Pointer to [**[]ArduinoDashboardshare**](ArduinoDashboardshare.md) | ArduinoDashboardshareCollection is the media type for an array of ArduinoDashboardshare (default view) | [optional] 
**UpdatedAt** | **time.Time** | Last update date | 
**Widgets** | Pointer to [**[]ArduinoWidgetv3**](ArduinoWidgetv3.md) | ArduinoWidgetv3Collection is the media type for an array of ArduinoWidgetv3 (default view) | [optional] 

## Methods

### NewArduinoDashboardv3

`func NewArduinoDashboardv3(id string, name string, updatedAt time.Time, ) *ArduinoDashboardv3`

NewArduinoDashboardv3 instantiates a new ArduinoDashboardv3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoDashboardv3WithDefaults

`func NewArduinoDashboardv3WithDefaults() *ArduinoDashboardv3`

NewArduinoDashboardv3WithDefaults instantiates a new ArduinoDashboardv3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoverImage

`func (o *ArduinoDashboardv3) GetCoverImage() string`

GetCoverImage returns the CoverImage field if non-nil, zero value otherwise.

### GetCoverImageOk

`func (o *ArduinoDashboardv3) GetCoverImageOk() (*string, bool)`

GetCoverImageOk returns a tuple with the CoverImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverImage

`func (o *ArduinoDashboardv3) SetCoverImage(v string)`

SetCoverImage sets CoverImage field to given value.

### HasCoverImage

`func (o *ArduinoDashboardv3) HasCoverImage() bool`

HasCoverImage returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ArduinoDashboardv3) GetCreatedBy() ArduinoDashboardowner`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ArduinoDashboardv3) GetCreatedByOk() (*ArduinoDashboardowner, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ArduinoDashboardv3) SetCreatedBy(v ArduinoDashboardowner)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ArduinoDashboardv3) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetId

`func (o *ArduinoDashboardv3) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArduinoDashboardv3) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArduinoDashboardv3) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ArduinoDashboardv3) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArduinoDashboardv3) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArduinoDashboardv3) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *ArduinoDashboardv3) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ArduinoDashboardv3) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ArduinoDashboardv3) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ArduinoDashboardv3) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetPages

`func (o *ArduinoDashboardv3) GetPages() []ArduinoPagevariable`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *ArduinoDashboardv3) GetPagesOk() (*[]ArduinoPagevariable, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *ArduinoDashboardv3) SetPages(v []ArduinoPagevariable)`

SetPages sets Pages field to given value.

### HasPages

`func (o *ArduinoDashboardv3) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetSharedBy

`func (o *ArduinoDashboardv3) GetSharedBy() ArduinoDashboardshare`

GetSharedBy returns the SharedBy field if non-nil, zero value otherwise.

### GetSharedByOk

`func (o *ArduinoDashboardv3) GetSharedByOk() (*ArduinoDashboardshare, bool)`

GetSharedByOk returns a tuple with the SharedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedBy

`func (o *ArduinoDashboardv3) SetSharedBy(v ArduinoDashboardshare)`

SetSharedBy sets SharedBy field to given value.

### HasSharedBy

`func (o *ArduinoDashboardv3) HasSharedBy() bool`

HasSharedBy returns a boolean if a field has been set.

### GetSharedWith

`func (o *ArduinoDashboardv3) GetSharedWith() []ArduinoDashboardshare`

GetSharedWith returns the SharedWith field if non-nil, zero value otherwise.

### GetSharedWithOk

`func (o *ArduinoDashboardv3) GetSharedWithOk() (*[]ArduinoDashboardshare, bool)`

GetSharedWithOk returns a tuple with the SharedWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedWith

`func (o *ArduinoDashboardv3) SetSharedWith(v []ArduinoDashboardshare)`

SetSharedWith sets SharedWith field to given value.

### HasSharedWith

`func (o *ArduinoDashboardv3) HasSharedWith() bool`

HasSharedWith returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ArduinoDashboardv3) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ArduinoDashboardv3) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ArduinoDashboardv3) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetWidgets

`func (o *ArduinoDashboardv3) GetWidgets() []ArduinoWidgetv3`

GetWidgets returns the Widgets field if non-nil, zero value otherwise.

### GetWidgetsOk

`func (o *ArduinoDashboardv3) GetWidgetsOk() (*[]ArduinoWidgetv3, bool)`

GetWidgetsOk returns a tuple with the Widgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgets

`func (o *ArduinoDashboardv3) SetWidgets(v []ArduinoWidgetv3)`

SetWidgets sets Widgets field to given value.

### HasWidgets

`func (o *ArduinoDashboardv3) HasWidgets() bool`

HasWidgets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


