# BatchQuerySampledRequestMediaV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **time.Time** | From timestamp (default: now UTC - 24h) | [optional] 
**Interval** | Pointer to **int32** | Resolution in seconds (allowed min:60, max:86400) | [optional] [default to 300]
**Q** | **string** | Data selection query (e.g. property.2a99729d-2556-4220-a139-023348a1e6b5) | 
**SeriesLimit** | Pointer to **int64** | Maximum number of values returned after data aggregation, if any (default: 300, limit: 1000) | [optional] 
**To** | Pointer to **time.Time** | To timestamp (default: now UTC) | [optional] 

## Methods

### NewBatchQuerySampledRequestMediaV1

`func NewBatchQuerySampledRequestMediaV1(q string, ) *BatchQuerySampledRequestMediaV1`

NewBatchQuerySampledRequestMediaV1 instantiates a new BatchQuerySampledRequestMediaV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchQuerySampledRequestMediaV1WithDefaults

`func NewBatchQuerySampledRequestMediaV1WithDefaults() *BatchQuerySampledRequestMediaV1`

NewBatchQuerySampledRequestMediaV1WithDefaults instantiates a new BatchQuerySampledRequestMediaV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *BatchQuerySampledRequestMediaV1) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *BatchQuerySampledRequestMediaV1) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *BatchQuerySampledRequestMediaV1) SetFrom(v time.Time)`

SetFrom sets From field to given value.

### HasFrom

`func (o *BatchQuerySampledRequestMediaV1) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetInterval

`func (o *BatchQuerySampledRequestMediaV1) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *BatchQuerySampledRequestMediaV1) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *BatchQuerySampledRequestMediaV1) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *BatchQuerySampledRequestMediaV1) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetQ

`func (o *BatchQuerySampledRequestMediaV1) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *BatchQuerySampledRequestMediaV1) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *BatchQuerySampledRequestMediaV1) SetQ(v string)`

SetQ sets Q field to given value.


### GetSeriesLimit

`func (o *BatchQuerySampledRequestMediaV1) GetSeriesLimit() int64`

GetSeriesLimit returns the SeriesLimit field if non-nil, zero value otherwise.

### GetSeriesLimitOk

`func (o *BatchQuerySampledRequestMediaV1) GetSeriesLimitOk() (*int64, bool)`

GetSeriesLimitOk returns a tuple with the SeriesLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesLimit

`func (o *BatchQuerySampledRequestMediaV1) SetSeriesLimit(v int64)`

SetSeriesLimit sets SeriesLimit field to given value.

### HasSeriesLimit

`func (o *BatchQuerySampledRequestMediaV1) HasSeriesLimit() bool`

HasSeriesLimit returns a boolean if a field has been set.

### GetTo

`func (o *BatchQuerySampledRequestMediaV1) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *BatchQuerySampledRequestMediaV1) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *BatchQuerySampledRequestMediaV1) SetTo(v time.Time)`

SetTo sets To field to given value.

### HasTo

`func (o *BatchQuerySampledRequestMediaV1) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


