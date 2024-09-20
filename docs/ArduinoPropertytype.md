# ArduinoPropertytype

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assistants** | Pointer to **[]string** | The voice assistants available for this type | [optional] 
**Declaration** | **string** | The c++ type we are using for this variable type | 
**Deprecated** | **bool** | Tell if this type is deprecated | 
**Example** | Pointer to **string** | Example of use | [optional] 
**Name** | **string** | The friendly name of the property type | 
**Rw** | **bool** | Tell if the type allow a R/W permission | 
**SupersededBy** | Pointer to **string** | The type of property to use if it&#39;s deprecated | [optional] 
**Tags** | Pointer to **[]string** | The tags related to the type | [optional] 
**Type** | **string** | The api reference of this type | 
**Units** | Pointer to **[]string** | The measure units available for this type | [optional] 

## Methods

### NewArduinoPropertytype

`func NewArduinoPropertytype(declaration string, deprecated bool, name string, rw bool, type_ string, ) *ArduinoPropertytype`

NewArduinoPropertytype instantiates a new ArduinoPropertytype object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoPropertytypeWithDefaults

`func NewArduinoPropertytypeWithDefaults() *ArduinoPropertytype`

NewArduinoPropertytypeWithDefaults instantiates a new ArduinoPropertytype object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssistants

`func (o *ArduinoPropertytype) GetAssistants() []string`

GetAssistants returns the Assistants field if non-nil, zero value otherwise.

### GetAssistantsOk

`func (o *ArduinoPropertytype) GetAssistantsOk() (*[]string, bool)`

GetAssistantsOk returns a tuple with the Assistants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistants

`func (o *ArduinoPropertytype) SetAssistants(v []string)`

SetAssistants sets Assistants field to given value.

### HasAssistants

`func (o *ArduinoPropertytype) HasAssistants() bool`

HasAssistants returns a boolean if a field has been set.

### GetDeclaration

`func (o *ArduinoPropertytype) GetDeclaration() string`

GetDeclaration returns the Declaration field if non-nil, zero value otherwise.

### GetDeclarationOk

`func (o *ArduinoPropertytype) GetDeclarationOk() (*string, bool)`

GetDeclarationOk returns a tuple with the Declaration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaration

`func (o *ArduinoPropertytype) SetDeclaration(v string)`

SetDeclaration sets Declaration field to given value.


### GetDeprecated

`func (o *ArduinoPropertytype) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *ArduinoPropertytype) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *ArduinoPropertytype) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.


### GetExample

`func (o *ArduinoPropertytype) GetExample() string`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *ArduinoPropertytype) GetExampleOk() (*string, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *ArduinoPropertytype) SetExample(v string)`

SetExample sets Example field to given value.

### HasExample

`func (o *ArduinoPropertytype) HasExample() bool`

HasExample returns a boolean if a field has been set.

### GetName

`func (o *ArduinoPropertytype) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArduinoPropertytype) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArduinoPropertytype) SetName(v string)`

SetName sets Name field to given value.


### GetRw

`func (o *ArduinoPropertytype) GetRw() bool`

GetRw returns the Rw field if non-nil, zero value otherwise.

### GetRwOk

`func (o *ArduinoPropertytype) GetRwOk() (*bool, bool)`

GetRwOk returns a tuple with the Rw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRw

`func (o *ArduinoPropertytype) SetRw(v bool)`

SetRw sets Rw field to given value.


### GetSupersededBy

`func (o *ArduinoPropertytype) GetSupersededBy() string`

GetSupersededBy returns the SupersededBy field if non-nil, zero value otherwise.

### GetSupersededByOk

`func (o *ArduinoPropertytype) GetSupersededByOk() (*string, bool)`

GetSupersededByOk returns a tuple with the SupersededBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupersededBy

`func (o *ArduinoPropertytype) SetSupersededBy(v string)`

SetSupersededBy sets SupersededBy field to given value.

### HasSupersededBy

`func (o *ArduinoPropertytype) HasSupersededBy() bool`

HasSupersededBy returns a boolean if a field has been set.

### GetTags

`func (o *ArduinoPropertytype) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ArduinoPropertytype) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ArduinoPropertytype) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ArduinoPropertytype) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *ArduinoPropertytype) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArduinoPropertytype) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArduinoPropertytype) SetType(v string)`

SetType sets Type field to given value.


### GetUnits

`func (o *ArduinoPropertytype) GetUnits() []string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *ArduinoPropertytype) GetUnitsOk() (*[]string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *ArduinoPropertytype) SetUnits(v []string)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *ArduinoPropertytype) HasUnits() bool`

HasUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


