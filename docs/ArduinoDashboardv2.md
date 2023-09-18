# ArduinoDashboardv2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to [**ArduinoDashboardowner**](ArduinoDashboardowner.md) |  | [optional] 
**Id** | **string** | The friendly name of the dashboard | 
**Name** | **string** | The friendly name of the dashboard | 
**OrganizationId** | Pointer to **string** | Id of the organization the dashboard belongs to | [optional] 
**SharedBy** | Pointer to [**ArduinoDashboardshare**](ArduinoDashboardshare.md) |  | [optional] 
**SharedWith** | Pointer to [**[]ArduinoDashboardshare**](ArduinoDashboardshare.md) | ArduinoDashboardshareCollection is the media type for an array of ArduinoDashboardshare (default view) | [optional] 
**UpdatedAt** | **time.Time** | Last update date | 
**Widgets** | Pointer to [**[]ArduinoWidgetv2**](ArduinoWidgetv2.md) | ArduinoWidgetv2Collection is the media type for an array of ArduinoWidgetv2 (default view) | [optional] 

## Methods

### NewArduinoDashboardv2

`func NewArduinoDashboardv2(id string, name string, updatedAt time.Time, ) *ArduinoDashboardv2`

NewArduinoDashboardv2 instantiates a new ArduinoDashboardv2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoDashboardv2WithDefaults

`func NewArduinoDashboardv2WithDefaults() *ArduinoDashboardv2`

NewArduinoDashboardv2WithDefaults instantiates a new ArduinoDashboardv2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *ArduinoDashboardv2) GetCreatedBy() ArduinoDashboardowner`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ArduinoDashboardv2) GetCreatedByOk() (*ArduinoDashboardowner, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ArduinoDashboardv2) SetCreatedBy(v ArduinoDashboardowner)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ArduinoDashboardv2) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetId

`func (o *ArduinoDashboardv2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArduinoDashboardv2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArduinoDashboardv2) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ArduinoDashboardv2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArduinoDashboardv2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArduinoDashboardv2) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *ArduinoDashboardv2) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ArduinoDashboardv2) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ArduinoDashboardv2) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ArduinoDashboardv2) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetSharedBy

`func (o *ArduinoDashboardv2) GetSharedBy() ArduinoDashboardshare`

GetSharedBy returns the SharedBy field if non-nil, zero value otherwise.

### GetSharedByOk

`func (o *ArduinoDashboardv2) GetSharedByOk() (*ArduinoDashboardshare, bool)`

GetSharedByOk returns a tuple with the SharedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedBy

`func (o *ArduinoDashboardv2) SetSharedBy(v ArduinoDashboardshare)`

SetSharedBy sets SharedBy field to given value.

### HasSharedBy

`func (o *ArduinoDashboardv2) HasSharedBy() bool`

HasSharedBy returns a boolean if a field has been set.

### GetSharedWith

`func (o *ArduinoDashboardv2) GetSharedWith() []ArduinoDashboardshare`

GetSharedWith returns the SharedWith field if non-nil, zero value otherwise.

### GetSharedWithOk

`func (o *ArduinoDashboardv2) GetSharedWithOk() (*[]ArduinoDashboardshare, bool)`

GetSharedWithOk returns a tuple with the SharedWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedWith

`func (o *ArduinoDashboardv2) SetSharedWith(v []ArduinoDashboardshare)`

SetSharedWith sets SharedWith field to given value.

### HasSharedWith

`func (o *ArduinoDashboardv2) HasSharedWith() bool`

HasSharedWith returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ArduinoDashboardv2) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ArduinoDashboardv2) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ArduinoDashboardv2) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetWidgets

`func (o *ArduinoDashboardv2) GetWidgets() []ArduinoWidgetv2`

GetWidgets returns the Widgets field if non-nil, zero value otherwise.

### GetWidgetsOk

`func (o *ArduinoDashboardv2) GetWidgetsOk() (*[]ArduinoWidgetv2, bool)`

GetWidgetsOk returns a tuple with the Widgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgets

`func (o *ArduinoDashboardv2) SetWidgets(v []ArduinoWidgetv2)`

SetWidgets sets Widgets field to given value.

### HasWidgets

`func (o *ArduinoDashboardv2) HasWidgets() bool`

HasWidgets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


