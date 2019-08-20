# \CleanupApi

All URIs are relative to *http://api-dev.arduino.cc*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CleanupIot**](CleanupApi.md#CleanupIot) | **Post** /iot/v1/cleanup | iot cleanup



## CleanupIot

> CleanupIot(ctx, )
iot cleanup

Iot removes the things created more than an hour ago and never got reclaimed. It then deletes the certificates that are no longer associated with a thing

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

