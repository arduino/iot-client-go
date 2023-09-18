# BatchQueryRawRequestMediaV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **time.Time** | From timestamp | [optional] 
**Q** | **string** | Query | 
**SeriesLimit** | Pointer to **int64** | Max of values | [optional] 
**Sort** | Pointer to **string** | Sorting | [optional] [default to "DESC"]
**To** | Pointer to **time.Time** | To timestamp | [optional] 

## Methods

### NewBatchQueryRawRequestMediaV1

`func NewBatchQueryRawRequestMediaV1(q string, ) *BatchQueryRawRequestMediaV1`

NewBatchQueryRawRequestMediaV1 instantiates a new BatchQueryRawRequestMediaV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchQueryRawRequestMediaV1WithDefaults

`func NewBatchQueryRawRequestMediaV1WithDefaults() *BatchQueryRawRequestMediaV1`

NewBatchQueryRawRequestMediaV1WithDefaults instantiates a new BatchQueryRawRequestMediaV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *BatchQueryRawRequestMediaV1) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *BatchQueryRawRequestMediaV1) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *BatchQueryRawRequestMediaV1) SetFrom(v time.Time)`

SetFrom sets From field to given value.

### HasFrom

`func (o *BatchQueryRawRequestMediaV1) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetQ

`func (o *BatchQueryRawRequestMediaV1) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *BatchQueryRawRequestMediaV1) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *BatchQueryRawRequestMediaV1) SetQ(v string)`

SetQ sets Q field to given value.


### GetSeriesLimit

`func (o *BatchQueryRawRequestMediaV1) GetSeriesLimit() int64`

GetSeriesLimit returns the SeriesLimit field if non-nil, zero value otherwise.

### GetSeriesLimitOk

`func (o *BatchQueryRawRequestMediaV1) GetSeriesLimitOk() (*int64, bool)`

GetSeriesLimitOk returns a tuple with the SeriesLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesLimit

`func (o *BatchQueryRawRequestMediaV1) SetSeriesLimit(v int64)`

SetSeriesLimit sets SeriesLimit field to given value.

### HasSeriesLimit

`func (o *BatchQueryRawRequestMediaV1) HasSeriesLimit() bool`

HasSeriesLimit returns a boolean if a field has been set.

### GetSort

`func (o *BatchQueryRawRequestMediaV1) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *BatchQueryRawRequestMediaV1) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *BatchQueryRawRequestMediaV1) SetSort(v string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *BatchQueryRawRequestMediaV1) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetTo

`func (o *BatchQueryRawRequestMediaV1) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *BatchQueryRawRequestMediaV1) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *BatchQueryRawRequestMediaV1) SetTo(v time.Time)`

SetTo sets To field to given value.

### HasTo

`func (o *BatchQueryRawRequestMediaV1) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


