# UserRecipient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | The email address of the user | [optional] 
**Id** | **string** | The id of the user | 
**Username** | Pointer to **string** | The username of the user | [optional] 

## Methods

### NewUserRecipient

`func NewUserRecipient(id string, ) *UserRecipient`

NewUserRecipient instantiates a new UserRecipient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRecipientWithDefaults

`func NewUserRecipientWithDefaults() *UserRecipient`

NewUserRecipientWithDefaults instantiates a new UserRecipient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserRecipient) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserRecipient) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserRecipient) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserRecipient) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *UserRecipient) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserRecipient) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserRecipient) SetId(v string)`

SetId sets Id field to given value.


### GetUsername

`func (o *UserRecipient) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserRecipient) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserRecipient) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserRecipient) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


