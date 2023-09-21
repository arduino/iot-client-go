# PropertiesValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | Pointer to **bool** | If true, send property values to device&#39;s input topic. | [optional] [default to false]
**Properties** | [**[]PropertiesValue**](PropertiesValue.md) |  | 

## Methods

### NewPropertiesValues

`func NewPropertiesValues(properties []PropertiesValue, ) *PropertiesValues`

NewPropertiesValues instantiates a new PropertiesValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertiesValuesWithDefaults

`func NewPropertiesValuesWithDefaults() *PropertiesValues`

NewPropertiesValuesWithDefaults instantiates a new PropertiesValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *PropertiesValues) GetInput() bool`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *PropertiesValues) GetInputOk() (*bool, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *PropertiesValues) SetInput(v bool)`

SetInput sets Input field to given value.

### HasInput

`func (o *PropertiesValues) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetProperties

`func (o *PropertiesValues) GetProperties() []PropertiesValue`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PropertiesValues) GetPropertiesOk() (*[]PropertiesValue, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PropertiesValues) SetProperties(v []PropertiesValue)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


