# BatchQueryRequestMediaV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregation** | Pointer to **string** | Aggregation statistic function. For numeric values, AVG statistic is used by default. PCT_X compute the Xth approximate percentile (e.g. PCT_95 is the 95th approximate percentile). For boolean, BOOL_OR statistic is used as default. | [optional] 
**From** | **time.Time** | From timestamp | 
**Interval** | Pointer to **int64** | Resolution in seconds (max allowed: 86400) | [optional] 
**Q** | **string** | Data selection query (e.g. property.2a99729d-2556-4220-a139-023348a1e6b5 or thing.95717675-4786-4ffc-afcc-799777755391) | 
**SeriesLimit** | Pointer to **int64** | Maximum number of values returned after data aggregation, if any (default: 300, limit: 1000 - 10000 in case of thing query) | [optional] 
**To** | **time.Time** | To timestamp | 

## Methods

### NewBatchQueryRequestMediaV1

`func NewBatchQueryRequestMediaV1(from time.Time, q string, to time.Time, ) *BatchQueryRequestMediaV1`

NewBatchQueryRequestMediaV1 instantiates a new BatchQueryRequestMediaV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchQueryRequestMediaV1WithDefaults

`func NewBatchQueryRequestMediaV1WithDefaults() *BatchQueryRequestMediaV1`

NewBatchQueryRequestMediaV1WithDefaults instantiates a new BatchQueryRequestMediaV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *BatchQueryRequestMediaV1) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *BatchQueryRequestMediaV1) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *BatchQueryRequestMediaV1) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *BatchQueryRequestMediaV1) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetFrom

`func (o *BatchQueryRequestMediaV1) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *BatchQueryRequestMediaV1) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *BatchQueryRequestMediaV1) SetFrom(v time.Time)`

SetFrom sets From field to given value.


### GetInterval

`func (o *BatchQueryRequestMediaV1) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *BatchQueryRequestMediaV1) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *BatchQueryRequestMediaV1) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *BatchQueryRequestMediaV1) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetQ

`func (o *BatchQueryRequestMediaV1) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *BatchQueryRequestMediaV1) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *BatchQueryRequestMediaV1) SetQ(v string)`

SetQ sets Q field to given value.


### GetSeriesLimit

`func (o *BatchQueryRequestMediaV1) GetSeriesLimit() int64`

GetSeriesLimit returns the SeriesLimit field if non-nil, zero value otherwise.

### GetSeriesLimitOk

`func (o *BatchQueryRequestMediaV1) GetSeriesLimitOk() (*int64, bool)`

GetSeriesLimitOk returns a tuple with the SeriesLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesLimit

`func (o *BatchQueryRequestMediaV1) SetSeriesLimit(v int64)`

SetSeriesLimit sets SeriesLimit field to given value.

### HasSeriesLimit

`func (o *BatchQueryRequestMediaV1) HasSeriesLimit() bool`

HasSeriesLimit returns a boolean if a field has been set.

### GetTo

`func (o *BatchQueryRequestMediaV1) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *BatchQueryRequestMediaV1) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *BatchQueryRequestMediaV1) SetTo(v time.Time)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


