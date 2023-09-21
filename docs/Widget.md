# Widget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Height** | **int64** | Widget current height for desktop | 
**HeightMobile** | Pointer to **int64** | Widget current height for mobile | [optional] 
**Id** | **string** | The UUID of the widget, set by client | 
**Name** | Pointer to **string** | The name of the widget | [optional] 
**Options** | **map[string]interface{}** | Widget options | 
**Type** | **string** | The type of the widget | 
**Variables** | Pointer to **[]string** |  | [optional] 
**Width** | **int64** | Widget current width for desktop | 
**WidthMobile** | Pointer to **int64** | Widget current width for mobile | [optional] 
**X** | **int64** | Widget x position for desktop | 
**XMobile** | Pointer to **int64** | Widget x position for mobile | [optional] 
**Y** | **int64** | Widget y position for desktop | 
**YMobile** | Pointer to **int64** | Widget y position for mobile | [optional] 

## Methods

### NewWidget

`func NewWidget(height int64, id string, options map[string]interface{}, type_ string, width int64, x int64, y int64, ) *Widget`

NewWidget instantiates a new Widget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidgetWithDefaults

`func NewWidgetWithDefaults() *Widget`

NewWidgetWithDefaults instantiates a new Widget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeight

`func (o *Widget) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *Widget) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *Widget) SetHeight(v int64)`

SetHeight sets Height field to given value.


### GetHeightMobile

`func (o *Widget) GetHeightMobile() int64`

GetHeightMobile returns the HeightMobile field if non-nil, zero value otherwise.

### GetHeightMobileOk

`func (o *Widget) GetHeightMobileOk() (*int64, bool)`

GetHeightMobileOk returns a tuple with the HeightMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeightMobile

`func (o *Widget) SetHeightMobile(v int64)`

SetHeightMobile sets HeightMobile field to given value.

### HasHeightMobile

`func (o *Widget) HasHeightMobile() bool`

HasHeightMobile returns a boolean if a field has been set.

### GetId

`func (o *Widget) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Widget) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Widget) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Widget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Widget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Widget) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Widget) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *Widget) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Widget) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Widget) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.


### GetType

`func (o *Widget) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Widget) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Widget) SetType(v string)`

SetType sets Type field to given value.


### GetVariables

`func (o *Widget) GetVariables() []string`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *Widget) GetVariablesOk() (*[]string, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *Widget) SetVariables(v []string)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *Widget) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetWidth

`func (o *Widget) GetWidth() int64`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *Widget) GetWidthOk() (*int64, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *Widget) SetWidth(v int64)`

SetWidth sets Width field to given value.


### GetWidthMobile

`func (o *Widget) GetWidthMobile() int64`

GetWidthMobile returns the WidthMobile field if non-nil, zero value otherwise.

### GetWidthMobileOk

`func (o *Widget) GetWidthMobileOk() (*int64, bool)`

GetWidthMobileOk returns a tuple with the WidthMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidthMobile

`func (o *Widget) SetWidthMobile(v int64)`

SetWidthMobile sets WidthMobile field to given value.

### HasWidthMobile

`func (o *Widget) HasWidthMobile() bool`

HasWidthMobile returns a boolean if a field has been set.

### GetX

`func (o *Widget) GetX() int64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *Widget) GetXOk() (*int64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *Widget) SetX(v int64)`

SetX sets X field to given value.


### GetXMobile

`func (o *Widget) GetXMobile() int64`

GetXMobile returns the XMobile field if non-nil, zero value otherwise.

### GetXMobileOk

`func (o *Widget) GetXMobileOk() (*int64, bool)`

GetXMobileOk returns a tuple with the XMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXMobile

`func (o *Widget) SetXMobile(v int64)`

SetXMobile sets XMobile field to given value.

### HasXMobile

`func (o *Widget) HasXMobile() bool`

HasXMobile returns a boolean if a field has been set.

### GetY

`func (o *Widget) GetY() int64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *Widget) GetYOk() (*int64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *Widget) SetY(v int64)`

SetY sets Y field to given value.


### GetYMobile

`func (o *Widget) GetYMobile() int64`

GetYMobile returns the YMobile field if non-nil, zero value otherwise.

### GetYMobileOk

`func (o *Widget) GetYMobileOk() (*int64, bool)`

GetYMobileOk returns a tuple with the YMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYMobile

`func (o *Widget) SetYMobile(v int64)`

SetYMobile sets YMobile field to given value.

### HasYMobile

`func (o *Widget) HasYMobile() bool`

HasYMobile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


