# ArduinoWidgetv3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasPermissionIncompatibility** | Pointer to **bool** | True if the linked variables permissions are incompatible with the widget | [optional] 
**HasTypeIncompatibility** | Pointer to **bool** | True if the linked variables types are incompatible with the widget | [optional] 
**HasUnlinkedVariable** | Pointer to **bool** | If it&#39;s true the widget is linked to a soft-deleted variable | [optional] 
**Height** | **int64** | Widget current height for desktop | 
**HeightMobile** | Pointer to **int64** | Widget current height for mobile | [optional] 
**Id** | **string** | The UUID of the widget, set by client | 
**Name** | Pointer to **string** | The name of the widget | [optional] 
**Options** | **map[string]interface{}** | Widget options | 
**PageId** | **string** | The ID of the page the widget belongs to, \&quot;0\&quot; if it&#39;s in the main page | 
**Type** | **string** | The type of the widget | 
**Variables** | Pointer to [**[]ArduinoLinkedvariable**](ArduinoLinkedvariable.md) | ArduinoLinkedvariableCollection is the media type for an array of ArduinoLinkedvariable (default view) | [optional] 
**Width** | **int64** | Widget current width for desktop | 
**WidthMobile** | Pointer to **int64** | Widget current width for mobile | [optional] 
**X** | **int64** | Widget x position for desktop | 
**XMobile** | Pointer to **int64** | Widget x position for mobile | [optional] 
**Y** | **int64** | Widget y position for desktop | 
**YMobile** | Pointer to **int64** | Widget y position for mobile | [optional] 

## Methods

### NewArduinoWidgetv3

`func NewArduinoWidgetv3(height int64, id string, options map[string]interface{}, pageId string, type_ string, width int64, x int64, y int64, ) *ArduinoWidgetv3`

NewArduinoWidgetv3 instantiates a new ArduinoWidgetv3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoWidgetv3WithDefaults

`func NewArduinoWidgetv3WithDefaults() *ArduinoWidgetv3`

NewArduinoWidgetv3WithDefaults instantiates a new ArduinoWidgetv3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasPermissionIncompatibility

`func (o *ArduinoWidgetv3) GetHasPermissionIncompatibility() bool`

GetHasPermissionIncompatibility returns the HasPermissionIncompatibility field if non-nil, zero value otherwise.

### GetHasPermissionIncompatibilityOk

`func (o *ArduinoWidgetv3) GetHasPermissionIncompatibilityOk() (*bool, bool)`

GetHasPermissionIncompatibilityOk returns a tuple with the HasPermissionIncompatibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPermissionIncompatibility

`func (o *ArduinoWidgetv3) SetHasPermissionIncompatibility(v bool)`

SetHasPermissionIncompatibility sets HasPermissionIncompatibility field to given value.

### HasHasPermissionIncompatibility

`func (o *ArduinoWidgetv3) HasHasPermissionIncompatibility() bool`

HasHasPermissionIncompatibility returns a boolean if a field has been set.

### GetHasTypeIncompatibility

`func (o *ArduinoWidgetv3) GetHasTypeIncompatibility() bool`

GetHasTypeIncompatibility returns the HasTypeIncompatibility field if non-nil, zero value otherwise.

### GetHasTypeIncompatibilityOk

`func (o *ArduinoWidgetv3) GetHasTypeIncompatibilityOk() (*bool, bool)`

GetHasTypeIncompatibilityOk returns a tuple with the HasTypeIncompatibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTypeIncompatibility

`func (o *ArduinoWidgetv3) SetHasTypeIncompatibility(v bool)`

SetHasTypeIncompatibility sets HasTypeIncompatibility field to given value.

### HasHasTypeIncompatibility

`func (o *ArduinoWidgetv3) HasHasTypeIncompatibility() bool`

HasHasTypeIncompatibility returns a boolean if a field has been set.

### GetHasUnlinkedVariable

`func (o *ArduinoWidgetv3) GetHasUnlinkedVariable() bool`

GetHasUnlinkedVariable returns the HasUnlinkedVariable field if non-nil, zero value otherwise.

### GetHasUnlinkedVariableOk

`func (o *ArduinoWidgetv3) GetHasUnlinkedVariableOk() (*bool, bool)`

GetHasUnlinkedVariableOk returns a tuple with the HasUnlinkedVariable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUnlinkedVariable

`func (o *ArduinoWidgetv3) SetHasUnlinkedVariable(v bool)`

SetHasUnlinkedVariable sets HasUnlinkedVariable field to given value.

### HasHasUnlinkedVariable

`func (o *ArduinoWidgetv3) HasHasUnlinkedVariable() bool`

HasHasUnlinkedVariable returns a boolean if a field has been set.

### GetHeight

`func (o *ArduinoWidgetv3) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ArduinoWidgetv3) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ArduinoWidgetv3) SetHeight(v int64)`

SetHeight sets Height field to given value.


### GetHeightMobile

`func (o *ArduinoWidgetv3) GetHeightMobile() int64`

GetHeightMobile returns the HeightMobile field if non-nil, zero value otherwise.

### GetHeightMobileOk

`func (o *ArduinoWidgetv3) GetHeightMobileOk() (*int64, bool)`

GetHeightMobileOk returns a tuple with the HeightMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeightMobile

`func (o *ArduinoWidgetv3) SetHeightMobile(v int64)`

SetHeightMobile sets HeightMobile field to given value.

### HasHeightMobile

`func (o *ArduinoWidgetv3) HasHeightMobile() bool`

HasHeightMobile returns a boolean if a field has been set.

### GetId

`func (o *ArduinoWidgetv3) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArduinoWidgetv3) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArduinoWidgetv3) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ArduinoWidgetv3) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArduinoWidgetv3) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArduinoWidgetv3) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ArduinoWidgetv3) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *ArduinoWidgetv3) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ArduinoWidgetv3) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ArduinoWidgetv3) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.


### GetPageId

`func (o *ArduinoWidgetv3) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *ArduinoWidgetv3) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *ArduinoWidgetv3) SetPageId(v string)`

SetPageId sets PageId field to given value.


### GetType

`func (o *ArduinoWidgetv3) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArduinoWidgetv3) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArduinoWidgetv3) SetType(v string)`

SetType sets Type field to given value.


### GetVariables

`func (o *ArduinoWidgetv3) GetVariables() []ArduinoLinkedvariable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ArduinoWidgetv3) GetVariablesOk() (*[]ArduinoLinkedvariable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ArduinoWidgetv3) SetVariables(v []ArduinoLinkedvariable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *ArduinoWidgetv3) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetWidth

`func (o *ArduinoWidgetv3) GetWidth() int64`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ArduinoWidgetv3) GetWidthOk() (*int64, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ArduinoWidgetv3) SetWidth(v int64)`

SetWidth sets Width field to given value.


### GetWidthMobile

`func (o *ArduinoWidgetv3) GetWidthMobile() int64`

GetWidthMobile returns the WidthMobile field if non-nil, zero value otherwise.

### GetWidthMobileOk

`func (o *ArduinoWidgetv3) GetWidthMobileOk() (*int64, bool)`

GetWidthMobileOk returns a tuple with the WidthMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidthMobile

`func (o *ArduinoWidgetv3) SetWidthMobile(v int64)`

SetWidthMobile sets WidthMobile field to given value.

### HasWidthMobile

`func (o *ArduinoWidgetv3) HasWidthMobile() bool`

HasWidthMobile returns a boolean if a field has been set.

### GetX

`func (o *ArduinoWidgetv3) GetX() int64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *ArduinoWidgetv3) GetXOk() (*int64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *ArduinoWidgetv3) SetX(v int64)`

SetX sets X field to given value.


### GetXMobile

`func (o *ArduinoWidgetv3) GetXMobile() int64`

GetXMobile returns the XMobile field if non-nil, zero value otherwise.

### GetXMobileOk

`func (o *ArduinoWidgetv3) GetXMobileOk() (*int64, bool)`

GetXMobileOk returns a tuple with the XMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXMobile

`func (o *ArduinoWidgetv3) SetXMobile(v int64)`

SetXMobile sets XMobile field to given value.

### HasXMobile

`func (o *ArduinoWidgetv3) HasXMobile() bool`

HasXMobile returns a boolean if a field has been set.

### GetY

`func (o *ArduinoWidgetv3) GetY() int64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *ArduinoWidgetv3) GetYOk() (*int64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *ArduinoWidgetv3) SetY(v int64)`

SetY sets Y field to given value.


### GetYMobile

`func (o *ArduinoWidgetv3) GetYMobile() int64`

GetYMobile returns the YMobile field if non-nil, zero value otherwise.

### GetYMobileOk

`func (o *ArduinoWidgetv3) GetYMobileOk() (*int64, bool)`

GetYMobileOk returns a tuple with the YMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYMobile

`func (o *ArduinoWidgetv3) SetYMobile(v int64)`

SetYMobile sets YMobile field to given value.

### HasYMobile

`func (o *ArduinoWidgetv3) HasYMobile() bool`

HasYMobile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


