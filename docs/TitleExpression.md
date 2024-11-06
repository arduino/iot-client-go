# TitleExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expression** | **string** | Content of the title (or subject) of a message, variables are allowed | 
**Variables** | Pointer to [**[]Variable**](Variable.md) | Variables used by the expression | [optional] 

## Methods

### NewTitleExpression

`func NewTitleExpression(expression string, ) *TitleExpression`

NewTitleExpression instantiates a new TitleExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTitleExpressionWithDefaults

`func NewTitleExpressionWithDefaults() *TitleExpression`

NewTitleExpressionWithDefaults instantiates a new TitleExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpression

`func (o *TitleExpression) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *TitleExpression) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *TitleExpression) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetVariables

`func (o *TitleExpression) GetVariables() []Variable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *TitleExpression) GetVariablesOk() (*[]Variable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *TitleExpression) SetVariables(v []Variable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *TitleExpression) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


