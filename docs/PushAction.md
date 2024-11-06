# PushAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | [**BodyExpression**](BodyExpression.md) |  | 
**Delivery** | [**PushDeliveryOpts**](PushDeliveryOpts.md) |  | 
**Title** | [**TitleExpression**](TitleExpression.md) |  | 

## Methods

### NewPushAction

`func NewPushAction(body BodyExpression, delivery PushDeliveryOpts, title TitleExpression, ) *PushAction`

NewPushAction instantiates a new PushAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushActionWithDefaults

`func NewPushActionWithDefaults() *PushAction`

NewPushActionWithDefaults instantiates a new PushAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *PushAction) GetBody() BodyExpression`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PushAction) GetBodyOk() (*BodyExpression, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PushAction) SetBody(v BodyExpression)`

SetBody sets Body field to given value.


### GetDelivery

`func (o *PushAction) GetDelivery() PushDeliveryOpts`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *PushAction) GetDeliveryOk() (*PushDeliveryOpts, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *PushAction) SetDelivery(v PushDeliveryOpts)`

SetDelivery sets Delivery field to given value.


### GetTitle

`func (o *PushAction) GetTitle() TitleExpression`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PushAction) GetTitleOk() (*TitleExpression, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PushAction) SetTitle(v TitleExpression)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


