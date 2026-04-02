# ArduinoWidgetv3template

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Height** | **int64** | Widget current height for desktop | 
**HeightMobile** | Pointer to **int64** | Widget current height for mobile | [optional] 
**Name** | Pointer to **string** | The name of the widget | [optional] 
**Options** | **map[string]interface{}** | Widget options | 
**PageId** | **string** | The ID of the page the widget belongs to, \&quot;0\&quot; if it&#39;s in the main page | 
**Type** | **string** | The type of the widget | 
**Variables** | Pointer to [**[]ArduinoTemplatevariable**](ArduinoTemplatevariable.md) | ArduinoTemplatevariableCollection is the media type for an array of ArduinoTemplatevariable (default view) | [optional] 
**Width** | **int64** | Widget current width for desktop | 
**WidthMobile** | Pointer to **int64** | Widget current width for mobile | [optional] 
**X** | **int64** | Widget x position for desktop | 
**XMobile** | Pointer to **int64** | Widget x position for mobile | [optional] 
**Y** | **int64** | Widget y position for desktop | 
**YMobile** | Pointer to **int64** | Widget y position for mobile | [optional] 

## Methods

### NewArduinoWidgetv3template

`func NewArduinoWidgetv3template(height int64, options map[string]interface{}, pageId string, type_ string, width int64, x int64, y int64, ) *ArduinoWidgetv3template`

NewArduinoWidgetv3template instantiates a new ArduinoWidgetv3template object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoWidgetv3templateWithDefaults

`func NewArduinoWidgetv3templateWithDefaults() *ArduinoWidgetv3template`

NewArduinoWidgetv3templateWithDefaults instantiates a new ArduinoWidgetv3template object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeight

`func (o *ArduinoWidgetv3template) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ArduinoWidgetv3template) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ArduinoWidgetv3template) SetHeight(v int64)`

SetHeight sets Height field to given value.


### GetHeightMobile

`func (o *ArduinoWidgetv3template) GetHeightMobile() int64`

GetHeightMobile returns the HeightMobile field if non-nil, zero value otherwise.

### GetHeightMobileOk

`func (o *ArduinoWidgetv3template) GetHeightMobileOk() (*int64, bool)`

GetHeightMobileOk returns a tuple with the HeightMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeightMobile

`func (o *ArduinoWidgetv3template) SetHeightMobile(v int64)`

SetHeightMobile sets HeightMobile field to given value.

### HasHeightMobile

`func (o *ArduinoWidgetv3template) HasHeightMobile() bool`

HasHeightMobile returns a boolean if a field has been set.

### GetName

`func (o *ArduinoWidgetv3template) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArduinoWidgetv3template) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArduinoWidgetv3template) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ArduinoWidgetv3template) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *ArduinoWidgetv3template) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ArduinoWidgetv3template) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ArduinoWidgetv3template) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.


### GetPageId

`func (o *ArduinoWidgetv3template) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *ArduinoWidgetv3template) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *ArduinoWidgetv3template) SetPageId(v string)`

SetPageId sets PageId field to given value.


### GetType

`func (o *ArduinoWidgetv3template) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArduinoWidgetv3template) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArduinoWidgetv3template) SetType(v string)`

SetType sets Type field to given value.


### GetVariables

`func (o *ArduinoWidgetv3template) GetVariables() []ArduinoTemplatevariable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ArduinoWidgetv3template) GetVariablesOk() (*[]ArduinoTemplatevariable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ArduinoWidgetv3template) SetVariables(v []ArduinoTemplatevariable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *ArduinoWidgetv3template) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetWidth

`func (o *ArduinoWidgetv3template) GetWidth() int64`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ArduinoWidgetv3template) GetWidthOk() (*int64, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ArduinoWidgetv3template) SetWidth(v int64)`

SetWidth sets Width field to given value.


### GetWidthMobile

`func (o *ArduinoWidgetv3template) GetWidthMobile() int64`

GetWidthMobile returns the WidthMobile field if non-nil, zero value otherwise.

### GetWidthMobileOk

`func (o *ArduinoWidgetv3template) GetWidthMobileOk() (*int64, bool)`

GetWidthMobileOk returns a tuple with the WidthMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidthMobile

`func (o *ArduinoWidgetv3template) SetWidthMobile(v int64)`

SetWidthMobile sets WidthMobile field to given value.

### HasWidthMobile

`func (o *ArduinoWidgetv3template) HasWidthMobile() bool`

HasWidthMobile returns a boolean if a field has been set.

### GetX

`func (o *ArduinoWidgetv3template) GetX() int64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *ArduinoWidgetv3template) GetXOk() (*int64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *ArduinoWidgetv3template) SetX(v int64)`

SetX sets X field to given value.


### GetXMobile

`func (o *ArduinoWidgetv3template) GetXMobile() int64`

GetXMobile returns the XMobile field if non-nil, zero value otherwise.

### GetXMobileOk

`func (o *ArduinoWidgetv3template) GetXMobileOk() (*int64, bool)`

GetXMobileOk returns a tuple with the XMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXMobile

`func (o *ArduinoWidgetv3template) SetXMobile(v int64)`

SetXMobile sets XMobile field to given value.

### HasXMobile

`func (o *ArduinoWidgetv3template) HasXMobile() bool`

HasXMobile returns a boolean if a field has been set.

### GetY

`func (o *ArduinoWidgetv3template) GetY() int64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *ArduinoWidgetv3template) GetYOk() (*int64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *ArduinoWidgetv3template) SetY(v int64)`

SetY sets Y field to given value.


### GetYMobile

`func (o *ArduinoWidgetv3template) GetYMobile() int64`

GetYMobile returns the YMobile field if non-nil, zero value otherwise.

### GetYMobileOk

`func (o *ArduinoWidgetv3template) GetYMobileOk() (*int64, bool)`

GetYMobileOk returns a tuple with the YMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYMobile

`func (o *ArduinoWidgetv3template) SetYMobile(v int64)`

SetYMobile sets YMobile field to given value.

### HasYMobile

`func (o *ArduinoWidgetv3template) HasYMobile() bool`

HasYMobile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


