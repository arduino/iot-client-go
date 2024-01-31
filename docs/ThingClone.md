# ThingClone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeTags** | Pointer to **bool** | Include tags in clone procedure | [optional] 
**Name** | **string** | The friendly name of the thing | 

## Methods

### NewThingClone

`func NewThingClone(name string, ) *ThingClone`

NewThingClone instantiates a new ThingClone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThingCloneWithDefaults

`func NewThingCloneWithDefaults() *ThingClone`

NewThingCloneWithDefaults instantiates a new ThingClone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeTags

`func (o *ThingClone) GetIncludeTags() bool`

GetIncludeTags returns the IncludeTags field if non-nil, zero value otherwise.

### GetIncludeTagsOk

`func (o *ThingClone) GetIncludeTagsOk() (*bool, bool)`

GetIncludeTagsOk returns a tuple with the IncludeTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTags

`func (o *ThingClone) SetIncludeTags(v bool)`

SetIncludeTags sets IncludeTags field to given value.

### HasIncludeTags

`func (o *ThingClone) HasIncludeTags() bool`

HasIncludeTags returns a boolean if a field has been set.

### GetName

`func (o *ThingClone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ThingClone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ThingClone) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


