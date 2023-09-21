# ArduinoDevicev2Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Whether the webhook is active | [optional] [default to true]
**Id** | **string** | The uuid of the webhook | 
**Uri** | **string** | The uri of the webhook | 

## Methods

### NewArduinoDevicev2Webhook

`func NewArduinoDevicev2Webhook(id string, uri string, ) *ArduinoDevicev2Webhook`

NewArduinoDevicev2Webhook instantiates a new ArduinoDevicev2Webhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArduinoDevicev2WebhookWithDefaults

`func NewArduinoDevicev2WebhookWithDefaults() *ArduinoDevicev2Webhook`

NewArduinoDevicev2WebhookWithDefaults instantiates a new ArduinoDevicev2Webhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ArduinoDevicev2Webhook) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ArduinoDevicev2Webhook) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ArduinoDevicev2Webhook) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ArduinoDevicev2Webhook) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetId

`func (o *ArduinoDevicev2Webhook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArduinoDevicev2Webhook) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArduinoDevicev2Webhook) SetId(v string)`

SetId sets Id field to given value.


### GetUri

`func (o *ArduinoDevicev2Webhook) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ArduinoDevicev2Webhook) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ArduinoDevicev2Webhook) SetUri(v string)`

SetUri sets Uri field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


