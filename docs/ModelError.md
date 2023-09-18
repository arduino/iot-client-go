# ModelError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | an application-specific error code, expressed as a string value. | [optional] 
**Detail** | Pointer to **string** | a human-readable explanation specific to this occurrence of the problem. | [optional] 
**Id** | Pointer to **string** | a unique identifier for this particular occurrence of the problem. | [optional] 
**Meta** | Pointer to **map[string]interface{}** | a meta object containing non-standard meta-information about the error. | [optional] 
**Status** | Pointer to **int64** | the HTTP status code applicable to this problem | [optional] 

## Methods

### NewModelError

`func NewModelError() *ModelError`

NewModelError instantiates a new ModelError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelErrorWithDefaults

`func NewModelErrorWithDefaults() *ModelError`

NewModelErrorWithDefaults instantiates a new ModelError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ModelError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ModelError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ModelError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ModelError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDetail

`func (o *ModelError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ModelError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ModelError) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ModelError) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetId

`func (o *ModelError) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelError) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelError) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelError) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMeta

`func (o *ModelError) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ModelError) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ModelError) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ModelError) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetStatus

`func (o *ModelError) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelError) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelError) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


