# Template

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomTemplateId** | Pointer to **string** | The name of the directory on S3 bucket containing the user&#39;s template | [optional] 
**PrefixName** | Pointer to **string** | The prefix to apply to the names of the generated resources | [optional] 
**TemplateName** | **string** | The name of the directory on S3 bucket containing the template | 
**ThingsOptions** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewTemplate

`func NewTemplate(templateName string, ) *Template`

NewTemplate instantiates a new Template object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateWithDefaults

`func NewTemplateWithDefaults() *Template`

NewTemplateWithDefaults instantiates a new Template object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomTemplateId

`func (o *Template) GetCustomTemplateId() string`

GetCustomTemplateId returns the CustomTemplateId field if non-nil, zero value otherwise.

### GetCustomTemplateIdOk

`func (o *Template) GetCustomTemplateIdOk() (*string, bool)`

GetCustomTemplateIdOk returns a tuple with the CustomTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTemplateId

`func (o *Template) SetCustomTemplateId(v string)`

SetCustomTemplateId sets CustomTemplateId field to given value.

### HasCustomTemplateId

`func (o *Template) HasCustomTemplateId() bool`

HasCustomTemplateId returns a boolean if a field has been set.

### GetPrefixName

`func (o *Template) GetPrefixName() string`

GetPrefixName returns the PrefixName field if non-nil, zero value otherwise.

### GetPrefixNameOk

`func (o *Template) GetPrefixNameOk() (*string, bool)`

GetPrefixNameOk returns a tuple with the PrefixName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixName

`func (o *Template) SetPrefixName(v string)`

SetPrefixName sets PrefixName field to given value.

### HasPrefixName

`func (o *Template) HasPrefixName() bool`

HasPrefixName returns a boolean if a field has been set.

### GetTemplateName

`func (o *Template) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *Template) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *Template) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.


### GetThingsOptions

`func (o *Template) GetThingsOptions() map[string]interface{}`

GetThingsOptions returns the ThingsOptions field if non-nil, zero value otherwise.

### GetThingsOptionsOk

`func (o *Template) GetThingsOptionsOk() (*map[string]interface{}, bool)`

GetThingsOptionsOk returns a tuple with the ThingsOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThingsOptions

`func (o *Template) SetThingsOptions(v map[string]interface{})`

SetThingsOptions sets ThingsOptions field to given value.

### HasThingsOptions

`func (o *Template) HasThingsOptions() bool`

HasThingsOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


