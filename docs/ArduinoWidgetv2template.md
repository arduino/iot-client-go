# ArduinoWidgetv2template

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Height** | **int64** | Widget current height for desktop | 
**HeightMobile** | Pointer to **int64** | Widget current height for mobile | [optional] 
**Name** | Pointer to **string** | The name of the widget | [optional] 
**Options** | **map[string]interface{}** | Widget options | 
**Type** | **string** | The type of the widget | 
**Variables** | Pointer to [**[]ArduinoTemplatevariable**](ArduinoTemplatevariable.md) | ArduinoTemplatevariableCollection is the media type for an array of ArduinoTemplatevariable (default view) | [optional] 
**Width** | **int64** | Widget current width for desktop | 
**WidthMobile** | Pointer to **int64** | Widget current width for mobile | [optional] 
**X** | **int64** | Widget x position for desktop | 
**XMobile** | Pointer to **int64** | Widget x position for mobile | [optional] 
**Y** | **int64** | Widget y position for desktop | 
**YMobile** | Pointer to **int64** | Widget y position for mobile | [optional] 

## Methods

### NewArduinoWidgetv2template

`func NewArduinoWidgetv2template(height int64, options map[string]interface{}, type_ string, width int64, x int64, y int64, ) *ArduinoWidgetv2template`

NewArduinoWidgetv2template instantiates a new ArduinoWidgetv2template object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoWidgetv2templateWithDefaults

`func NewArduinoWidgetv2templateWithDefaults() *ArduinoWidgetv2template`

NewArduinoWidgetv2templateWithDefaults instantiates a new ArduinoWidgetv2template object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeight

`func (o *ArduinoWidgetv2template) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ArduinoWidgetv2template) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ArduinoWidgetv2template) SetHeight(v int64)`

SetHeight sets Height field to given value.


### GetHeightMobile

`func (o *ArduinoWidgetv2template) GetHeightMobile() int64`

GetHeightMobile returns the HeightMobile field if non-nil, zero value otherwise.

### GetHeightMobileOk

`func (o *ArduinoWidgetv2template) GetHeightMobileOk() (*int64, bool)`

GetHeightMobileOk returns a tuple with the HeightMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeightMobile

`func (o *ArduinoWidgetv2template) SetHeightMobile(v int64)`

SetHeightMobile sets HeightMobile field to given value.

### HasHeightMobile

`func (o *ArduinoWidgetv2template) HasHeightMobile() bool`

HasHeightMobile returns a boolean if a field has been set.

### GetName

`func (o *ArduinoWidgetv2template) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArduinoWidgetv2template) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArduinoWidgetv2template) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ArduinoWidgetv2template) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *ArduinoWidgetv2template) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ArduinoWidgetv2template) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ArduinoWidgetv2template) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.


### GetType

`func (o *ArduinoWidgetv2template) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArduinoWidgetv2template) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArduinoWidgetv2template) SetType(v string)`

SetType sets Type field to given value.


### GetVariables

`func (o *ArduinoWidgetv2template) GetVariables() []ArduinoTemplatevariable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ArduinoWidgetv2template) GetVariablesOk() (*[]ArduinoTemplatevariable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ArduinoWidgetv2template) SetVariables(v []ArduinoTemplatevariable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *ArduinoWidgetv2template) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetWidth

`func (o *ArduinoWidgetv2template) GetWidth() int64`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ArduinoWidgetv2template) GetWidthOk() (*int64, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ArduinoWidgetv2template) SetWidth(v int64)`

SetWidth sets Width field to given value.


### GetWidthMobile

`func (o *ArduinoWidgetv2template) GetWidthMobile() int64`

GetWidthMobile returns the WidthMobile field if non-nil, zero value otherwise.

### GetWidthMobileOk

`func (o *ArduinoWidgetv2template) GetWidthMobileOk() (*int64, bool)`

GetWidthMobileOk returns a tuple with the WidthMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidthMobile

`func (o *ArduinoWidgetv2template) SetWidthMobile(v int64)`

SetWidthMobile sets WidthMobile field to given value.

### HasWidthMobile

`func (o *ArduinoWidgetv2template) HasWidthMobile() bool`

HasWidthMobile returns a boolean if a field has been set.

### GetX

`func (o *ArduinoWidgetv2template) GetX() int64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *ArduinoWidgetv2template) GetXOk() (*int64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *ArduinoWidgetv2template) SetX(v int64)`

SetX sets X field to given value.


### GetXMobile

`func (o *ArduinoWidgetv2template) GetXMobile() int64`

GetXMobile returns the XMobile field if non-nil, zero value otherwise.

### GetXMobileOk

`func (o *ArduinoWidgetv2template) GetXMobileOk() (*int64, bool)`

GetXMobileOk returns a tuple with the XMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXMobile

`func (o *ArduinoWidgetv2template) SetXMobile(v int64)`

SetXMobile sets XMobile field to given value.

### HasXMobile

`func (o *ArduinoWidgetv2template) HasXMobile() bool`

HasXMobile returns a boolean if a field has been set.

### GetY

`func (o *ArduinoWidgetv2template) GetY() int64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *ArduinoWidgetv2template) GetYOk() (*int64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *ArduinoWidgetv2template) SetY(v int64)`

SetY sets Y field to given value.


### GetYMobile

`func (o *ArduinoWidgetv2template) GetYMobile() int64`

GetYMobile returns the YMobile field if non-nil, zero value otherwise.

### GetYMobileOk

`func (o *ArduinoWidgetv2template) GetYMobileOk() (*int64, bool)`

GetYMobileOk returns a tuple with the YMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYMobile

`func (o *ArduinoWidgetv2template) SetYMobile(v int64)`

SetYMobile sets YMobile field to given value.

### HasYMobile

`func (o *ArduinoWidgetv2template) HasYMobile() bool`

HasYMobile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


