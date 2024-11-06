# EmailAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | [**BodyExpression**](BodyExpression.md) |  | 
**Delivery** | [**EmailDeliveryOpts**](EmailDeliveryOpts.md) |  | 
**Subject** | [**TitleExpression**](TitleExpression.md) |  | 

## Methods

### NewEmailAction

`func NewEmailAction(body BodyExpression, delivery EmailDeliveryOpts, subject TitleExpression, ) *EmailAction`

NewEmailAction instantiates a new EmailAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailActionWithDefaults

`func NewEmailActionWithDefaults() *EmailAction`

NewEmailActionWithDefaults instantiates a new EmailAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *EmailAction) GetBody() BodyExpression`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *EmailAction) GetBodyOk() (*BodyExpression, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *EmailAction) SetBody(v BodyExpression)`

SetBody sets Body field to given value.


### GetDelivery

`func (o *EmailAction) GetDelivery() EmailDeliveryOpts`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *EmailAction) GetDeliveryOk() (*EmailDeliveryOpts, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *EmailAction) SetDelivery(v EmailDeliveryOpts)`

SetDelivery sets Delivery field to given value.


### GetSubject

`func (o *EmailAction) GetSubject() TitleExpression`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailAction) GetSubjectOk() (*TitleExpression, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailAction) SetSubject(v TitleExpression)`

SetSubject sets Subject field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


