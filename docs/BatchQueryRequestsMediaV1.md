# BatchQueryRequestsMediaV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requests** | [**[]BatchQueryRequestMediaV1**](BatchQueryRequestMediaV1.md) | Requests | 
**RespVersion** | **int64** | Response version | 

## Methods

### NewBatchQueryRequestsMediaV1

`func NewBatchQueryRequestsMediaV1(requests []BatchQueryRequestMediaV1, respVersion int64, ) *BatchQueryRequestsMediaV1`

NewBatchQueryRequestsMediaV1 instantiates a new BatchQueryRequestsMediaV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchQueryRequestsMediaV1WithDefaults

`func NewBatchQueryRequestsMediaV1WithDefaults() *BatchQueryRequestsMediaV1`

NewBatchQueryRequestsMediaV1WithDefaults instantiates a new BatchQueryRequestsMediaV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequests

`func (o *BatchQueryRequestsMediaV1) GetRequests() []BatchQueryRequestMediaV1`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *BatchQueryRequestsMediaV1) GetRequestsOk() (*[]BatchQueryRequestMediaV1, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *BatchQueryRequestsMediaV1) SetRequests(v []BatchQueryRequestMediaV1)`

SetRequests sets Requests field to given value.


### GetRespVersion

`func (o *BatchQueryRequestsMediaV1) GetRespVersion() int64`

GetRespVersion returns the RespVersion field if non-nil, zero value otherwise.

### GetRespVersionOk

`func (o *BatchQueryRequestsMediaV1) GetRespVersionOk() (*int64, bool)`

GetRespVersionOk returns a tuple with the RespVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespVersion

`func (o *BatchQueryRequestsMediaV1) SetRespVersion(v int64)`

SetRespVersion sets RespVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


