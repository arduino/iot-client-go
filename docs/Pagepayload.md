# Pagepayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Icon** | Pointer to **string** | The icon of the Page | [optional] 
**Id** | **string** | The id of the Page | 
**Name** | **string** | The name of the Page | 
**Position** | **int64** | The position of the Page | 

## Methods

### NewPagepayload

`func NewPagepayload(id string, name string, position int64, ) *Pagepayload`

NewPagepayload instantiates a new Pagepayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagepayloadWithDefaults

`func NewPagepayloadWithDefaults() *Pagepayload`

NewPagepayloadWithDefaults instantiates a new Pagepayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIcon

`func (o *Pagepayload) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *Pagepayload) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *Pagepayload) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *Pagepayload) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetId

`func (o *Pagepayload) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Pagepayload) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Pagepayload) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Pagepayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Pagepayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Pagepayload) SetName(v string)`

SetName sets Name field to given value.


### GetPosition

`func (o *Pagepayload) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *Pagepayload) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *Pagepayload) SetPosition(v int64)`

SetPosition sets Position field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


