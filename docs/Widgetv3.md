# Widgetv3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Height** | **int64** | Widget current height for desktop | 
**HeightMobile** | Pointer to **int64** | Widget current height for mobile | [optional] 
**Id** | **string** | The UUID of the widget, set by client | 
**Name** | Pointer to **string** | The name of the widget | [optional] 
**Options** | **map[string]interface{}** | Widget options | 
**PageId** | Pointer to **string** | The ID of the page the widget belongs to, \&quot;0\&quot; if it&#39;s in the main page | [optional] 
**Type** | **string** | The type of the widget | 
**Variables** | Pointer to **[]string** |  | [optional] 
**Width** | **int64** | Widget current width for desktop | 
**WidthMobile** | Pointer to **int64** | Widget current width for mobile | [optional] 
**X** | **int64** | Widget x position for desktop | 
**XMobile** | Pointer to **int64** | Widget x position for mobile | [optional] 
**Y** | **int64** | Widget y position for desktop | 
**YMobile** | Pointer to **int64** | Widget y position for mobile | [optional] 

## Methods

### NewWidgetv3

`func NewWidgetv3(height int64, id string, options map[string]interface{}, type_ string, width int64, x int64, y int64, ) *Widgetv3`

NewWidgetv3 instantiates a new Widgetv3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidgetv3WithDefaults

`func NewWidgetv3WithDefaults() *Widgetv3`

NewWidgetv3WithDefaults instantiates a new Widgetv3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeight

`func (o *Widgetv3) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *Widgetv3) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *Widgetv3) SetHeight(v int64)`

SetHeight sets Height field to given value.


### GetHeightMobile

`func (o *Widgetv3) GetHeightMobile() int64`

GetHeightMobile returns the HeightMobile field if non-nil, zero value otherwise.

### GetHeightMobileOk

`func (o *Widgetv3) GetHeightMobileOk() (*int64, bool)`

GetHeightMobileOk returns a tuple with the HeightMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeightMobile

`func (o *Widgetv3) SetHeightMobile(v int64)`

SetHeightMobile sets HeightMobile field to given value.

### HasHeightMobile

`func (o *Widgetv3) HasHeightMobile() bool`

HasHeightMobile returns a boolean if a field has been set.

### GetId

`func (o *Widgetv3) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Widgetv3) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Widgetv3) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Widgetv3) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Widgetv3) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Widgetv3) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Widgetv3) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *Widgetv3) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Widgetv3) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Widgetv3) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.


### GetPageId

`func (o *Widgetv3) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *Widgetv3) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *Widgetv3) SetPageId(v string)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *Widgetv3) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetType

`func (o *Widgetv3) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Widgetv3) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Widgetv3) SetType(v string)`

SetType sets Type field to given value.


### GetVariables

`func (o *Widgetv3) GetVariables() []string`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *Widgetv3) GetVariablesOk() (*[]string, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *Widgetv3) SetVariables(v []string)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *Widgetv3) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetWidth

`func (o *Widgetv3) GetWidth() int64`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *Widgetv3) GetWidthOk() (*int64, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *Widgetv3) SetWidth(v int64)`

SetWidth sets Width field to given value.


### GetWidthMobile

`func (o *Widgetv3) GetWidthMobile() int64`

GetWidthMobile returns the WidthMobile field if non-nil, zero value otherwise.

### GetWidthMobileOk

`func (o *Widgetv3) GetWidthMobileOk() (*int64, bool)`

GetWidthMobileOk returns a tuple with the WidthMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidthMobile

`func (o *Widgetv3) SetWidthMobile(v int64)`

SetWidthMobile sets WidthMobile field to given value.

### HasWidthMobile

`func (o *Widgetv3) HasWidthMobile() bool`

HasWidthMobile returns a boolean if a field has been set.

### GetX

`func (o *Widgetv3) GetX() int64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *Widgetv3) GetXOk() (*int64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *Widgetv3) SetX(v int64)`

SetX sets X field to given value.


### GetXMobile

`func (o *Widgetv3) GetXMobile() int64`

GetXMobile returns the XMobile field if non-nil, zero value otherwise.

### GetXMobileOk

`func (o *Widgetv3) GetXMobileOk() (*int64, bool)`

GetXMobileOk returns a tuple with the XMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXMobile

`func (o *Widgetv3) SetXMobile(v int64)`

SetXMobile sets XMobile field to given value.

### HasXMobile

`func (o *Widgetv3) HasXMobile() bool`

HasXMobile returns a boolean if a field has been set.

### GetY

`func (o *Widgetv3) GetY() int64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *Widgetv3) GetYOk() (*int64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *Widgetv3) SetY(v int64)`

SetY sets Y field to given value.


### GetYMobile

`func (o *Widgetv3) GetYMobile() int64`

GetYMobile returns the YMobile field if non-nil, zero value otherwise.

### GetYMobileOk

`func (o *Widgetv3) GetYMobileOk() (*int64, bool)`

GetYMobileOk returns a tuple with the YMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYMobile

`func (o *Widgetv3) SetYMobile(v int64)`

SetYMobile sets YMobile field to given value.

### HasYMobile

`func (o *Widgetv3) HasYMobile() bool`

HasYMobile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


