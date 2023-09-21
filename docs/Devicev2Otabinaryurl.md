# Devicev2Otabinaryurl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Async** | Pointer to **bool** | If false, wait for the full OTA process, until it gets a result from the device | [optional] [default to true]
**BinaryKey** | **string** | The object key of the binary | 
**ExpireInMins** | Pointer to **int64** | Binary expire time in minutes, default 10 mins | [optional] [default to 10]

## Methods

### NewDevicev2Otabinaryurl

`func NewDevicev2Otabinaryurl(binaryKey string, ) *Devicev2Otabinaryurl`

NewDevicev2Otabinaryurl instantiates a new Devicev2Otabinaryurl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicev2OtabinaryurlWithDefaults

`func NewDevicev2OtabinaryurlWithDefaults() *Devicev2Otabinaryurl`

NewDevicev2OtabinaryurlWithDefaults instantiates a new Devicev2Otabinaryurl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsync

`func (o *Devicev2Otabinaryurl) GetAsync() bool`

GetAsync returns the Async field if non-nil, zero value otherwise.

### GetAsyncOk

`func (o *Devicev2Otabinaryurl) GetAsyncOk() (*bool, bool)`

GetAsyncOk returns a tuple with the Async field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsync

`func (o *Devicev2Otabinaryurl) SetAsync(v bool)`

SetAsync sets Async field to given value.

### HasAsync

`func (o *Devicev2Otabinaryurl) HasAsync() bool`

HasAsync returns a boolean if a field has been set.

### GetBinaryKey

`func (o *Devicev2Otabinaryurl) GetBinaryKey() string`

GetBinaryKey returns the BinaryKey field if non-nil, zero value otherwise.

### GetBinaryKeyOk

`func (o *Devicev2Otabinaryurl) GetBinaryKeyOk() (*string, bool)`

GetBinaryKeyOk returns a tuple with the BinaryKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryKey

`func (o *Devicev2Otabinaryurl) SetBinaryKey(v string)`

SetBinaryKey sets BinaryKey field to given value.


### GetExpireInMins

`func (o *Devicev2Otabinaryurl) GetExpireInMins() int64`

GetExpireInMins returns the ExpireInMins field if non-nil, zero value otherwise.

### GetExpireInMinsOk

`func (o *Devicev2Otabinaryurl) GetExpireInMinsOk() (*int64, bool)`

GetExpireInMinsOk returns a tuple with the ExpireInMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireInMins

`func (o *Devicev2Otabinaryurl) SetExpireInMins(v int64)`

SetExpireInMins sets ExpireInMins field to given value.

### HasExpireInMins

`func (o *Devicev2Otabinaryurl) HasExpireInMins() bool`

HasExpireInMins returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


