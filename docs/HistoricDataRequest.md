# HistoricDataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **time.Time** | Get data starting from this date | 
**Properties** | **[]string** | IDs of properties | 
**To** | **time.Time** | Get data up to this date | 

## Methods

### NewHistoricDataRequest

`func NewHistoricDataRequest(from time.Time, properties []string, to time.Time, ) *HistoricDataRequest`

NewHistoricDataRequest instantiates a new HistoricDataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricDataRequestWithDefaults

`func NewHistoricDataRequestWithDefaults() *HistoricDataRequest`

NewHistoricDataRequestWithDefaults instantiates a new HistoricDataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *HistoricDataRequest) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *HistoricDataRequest) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *HistoricDataRequest) SetFrom(v time.Time)`

SetFrom sets From field to given value.


### GetProperties

`func (o *HistoricDataRequest) GetProperties() []string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *HistoricDataRequest) GetPropertiesOk() (*[]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *HistoricDataRequest) SetProperties(v []string)`

SetProperties sets Properties field to given value.


### GetTo

`func (o *HistoricDataRequest) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *HistoricDataRequest) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *HistoricDataRequest) SetTo(v time.Time)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


