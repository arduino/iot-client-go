# \ManageApi

All URIs are relative to *http://api-dev.arduino.cc*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ManageDeleteUser**](ManageApi.md#ManageDeleteUser) | **Delete** /iot/v1/manage/users/{id} | delete_user manage



## ManageDeleteUser

> ManageDeleteUser(ctx, id)
delete_user manage

Deletes the user, all their sketches and libraries

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the user. Can be &#39;me&#39; to access the currently authenticated user | 

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

