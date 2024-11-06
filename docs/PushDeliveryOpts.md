# PushDeliveryOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | [**[]UserRecipient**](UserRecipient.md) | The recipient of a push notification | 

## Methods

### NewPushDeliveryOpts

`func NewPushDeliveryOpts(to []UserRecipient, ) *PushDeliveryOpts`

NewPushDeliveryOpts instantiates a new PushDeliveryOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushDeliveryOptsWithDefaults

`func NewPushDeliveryOptsWithDefaults() *PushDeliveryOpts`

NewPushDeliveryOptsWithDefaults instantiates a new PushDeliveryOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *PushDeliveryOpts) GetTo() []UserRecipient`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *PushDeliveryOpts) GetToOk() (*[]UserRecipient, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *PushDeliveryOpts) SetTo(v []UserRecipient)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


