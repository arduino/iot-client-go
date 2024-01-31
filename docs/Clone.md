# Clone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Overrides** | Pointer to [**[]Override**](Override.md) | The overrides to apply to the cloned dashboard. An override is a tuple of ids: the id of the thing to override and the id of the new thing to link | [optional] 

## Methods

### NewClone

`func NewClone() *Clone`

NewClone instantiates a new Clone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloneWithDefaults

`func NewCloneWithDefaults() *Clone`

NewCloneWithDefaults instantiates a new Clone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverrides

`func (o *Clone) GetOverrides() []Override`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *Clone) GetOverridesOk() (*[]Override, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *Clone) SetOverrides(v []Override)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *Clone) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


