# EmailDeliveryOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bcc** | Pointer to [**[]UserRecipient**](UserRecipient.md) | The \&quot;bcc:\&quot; field of an e-mail | [optional] 
**Cc** | Pointer to [**[]UserRecipient**](UserRecipient.md) | The \&quot;cc:\&quot; field of an e-mail | [optional] 
**To** | [**[]UserRecipient**](UserRecipient.md) | The \&quot;to:\&quot; field of an e-mail | 

## Methods

### NewEmailDeliveryOpts

`func NewEmailDeliveryOpts(to []UserRecipient, ) *EmailDeliveryOpts`

NewEmailDeliveryOpts instantiates a new EmailDeliveryOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailDeliveryOptsWithDefaults

`func NewEmailDeliveryOptsWithDefaults() *EmailDeliveryOpts`

NewEmailDeliveryOptsWithDefaults instantiates a new EmailDeliveryOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBcc

`func (o *EmailDeliveryOpts) GetBcc() []UserRecipient`

GetBcc returns the Bcc field if non-nil, zero value otherwise.

### GetBccOk

`func (o *EmailDeliveryOpts) GetBccOk() (*[]UserRecipient, bool)`

GetBccOk returns a tuple with the Bcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcc

`func (o *EmailDeliveryOpts) SetBcc(v []UserRecipient)`

SetBcc sets Bcc field to given value.

### HasBcc

`func (o *EmailDeliveryOpts) HasBcc() bool`

HasBcc returns a boolean if a field has been set.

### GetCc

`func (o *EmailDeliveryOpts) GetCc() []UserRecipient`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *EmailDeliveryOpts) GetCcOk() (*[]UserRecipient, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *EmailDeliveryOpts) SetCc(v []UserRecipient)`

SetCc sets Cc field to given value.

### HasCc

`func (o *EmailDeliveryOpts) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetTo

`func (o *EmailDeliveryOpts) GetTo() []UserRecipient`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EmailDeliveryOpts) GetToOk() (*[]UserRecipient, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EmailDeliveryOpts) SetTo(v []UserRecipient)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


