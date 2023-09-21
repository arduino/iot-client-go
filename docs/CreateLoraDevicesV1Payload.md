# CreateLoraDevicesV1Payload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | **string** | The app name | 
**AppEui** | Pointer to **string** | The app eui of the lora device | [optional] 
**AppKey** | Pointer to **string** | The app key of the lora device | [optional] 
**Eui** | **string** | The eui of the lora device | 
**FrequencyPlan** | **string** | The frequency plan required by your country  | 
**Name** | **string** | A common name for the device | 
**Serial** | Pointer to **string** | The optional serial number | [optional] 
**Type** | **string** | The type of device | 
**UserId** | **string** | The id of the user. Can be the special string &#39;me&#39; | 

## Methods

### NewCreateLoraDevicesV1Payload

`func NewCreateLoraDevicesV1Payload(app string, eui string, frequencyPlan string, name string, type_ string, userId string, ) *CreateLoraDevicesV1Payload`

NewCreateLoraDevicesV1Payload instantiates a new CreateLoraDevicesV1Payload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLoraDevicesV1PayloadWithDefaults

`func NewCreateLoraDevicesV1PayloadWithDefaults() *CreateLoraDevicesV1Payload`

NewCreateLoraDevicesV1PayloadWithDefaults instantiates a new CreateLoraDevicesV1Payload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *CreateLoraDevicesV1Payload) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *CreateLoraDevicesV1Payload) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *CreateLoraDevicesV1Payload) SetApp(v string)`

SetApp sets App field to given value.


### GetAppEui

`func (o *CreateLoraDevicesV1Payload) GetAppEui() string`

GetAppEui returns the AppEui field if non-nil, zero value otherwise.

### GetAppEuiOk

`func (o *CreateLoraDevicesV1Payload) GetAppEuiOk() (*string, bool)`

GetAppEuiOk returns a tuple with the AppEui field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppEui

`func (o *CreateLoraDevicesV1Payload) SetAppEui(v string)`

SetAppEui sets AppEui field to given value.

### HasAppEui

`func (o *CreateLoraDevicesV1Payload) HasAppEui() bool`

HasAppEui returns a boolean if a field has been set.

### GetAppKey

`func (o *CreateLoraDevicesV1Payload) GetAppKey() string`

GetAppKey returns the AppKey field if non-nil, zero value otherwise.

### GetAppKeyOk

`func (o *CreateLoraDevicesV1Payload) GetAppKeyOk() (*string, bool)`

GetAppKeyOk returns a tuple with the AppKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppKey

`func (o *CreateLoraDevicesV1Payload) SetAppKey(v string)`

SetAppKey sets AppKey field to given value.

### HasAppKey

`func (o *CreateLoraDevicesV1Payload) HasAppKey() bool`

HasAppKey returns a boolean if a field has been set.

### GetEui

`func (o *CreateLoraDevicesV1Payload) GetEui() string`

GetEui returns the Eui field if non-nil, zero value otherwise.

### GetEuiOk

`func (o *CreateLoraDevicesV1Payload) GetEuiOk() (*string, bool)`

GetEuiOk returns a tuple with the Eui field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEui

`func (o *CreateLoraDevicesV1Payload) SetEui(v string)`

SetEui sets Eui field to given value.


### GetFrequencyPlan

`func (o *CreateLoraDevicesV1Payload) GetFrequencyPlan() string`

GetFrequencyPlan returns the FrequencyPlan field if non-nil, zero value otherwise.

### GetFrequencyPlanOk

`func (o *CreateLoraDevicesV1Payload) GetFrequencyPlanOk() (*string, bool)`

GetFrequencyPlanOk returns a tuple with the FrequencyPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyPlan

`func (o *CreateLoraDevicesV1Payload) SetFrequencyPlan(v string)`

SetFrequencyPlan sets FrequencyPlan field to given value.


### GetName

`func (o *CreateLoraDevicesV1Payload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLoraDevicesV1Payload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLoraDevicesV1Payload) SetName(v string)`

SetName sets Name field to given value.


### GetSerial

`func (o *CreateLoraDevicesV1Payload) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *CreateLoraDevicesV1Payload) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *CreateLoraDevicesV1Payload) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *CreateLoraDevicesV1Payload) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetType

`func (o *CreateLoraDevicesV1Payload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateLoraDevicesV1Payload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateLoraDevicesV1Payload) SetType(v string)`

SetType sets Type field to given value.


### GetUserId

`func (o *CreateLoraDevicesV1Payload) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateLoraDevicesV1Payload) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateLoraDevicesV1Payload) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


