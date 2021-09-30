# \DashboardsV2Api

All URIs are relative to *https://api2.arduino.cc/iot*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DashboardsV2Create**](DashboardsV2Api.md#DashboardsV2Create) | **Post** /v2/dashboards | create dashboards_v2
[**DashboardsV2Delete**](DashboardsV2Api.md#DashboardsV2Delete) | **Delete** /v2/dashboards/{id} | delete dashboards_v2
[**DashboardsV2DeleteShare**](DashboardsV2Api.md#DashboardsV2DeleteShare) | **Delete** /v2/dashboards/{id}/shares/{user_id} | deleteShare dashboards_v2
[**DashboardsV2Link**](DashboardsV2Api.md#DashboardsV2Link) | **Put** /v2/dashboards/{id}/widgets/{widgetId}/variables | link dashboards_v2
[**DashboardsV2List**](DashboardsV2Api.md#DashboardsV2List) | **Get** /v2/dashboards | list dashboards_v2
[**DashboardsV2ListShares**](DashboardsV2Api.md#DashboardsV2ListShares) | **Get** /v2/dashboards/{id}/shares | listShares dashboards_v2
[**DashboardsV2RequestAccess**](DashboardsV2Api.md#DashboardsV2RequestAccess) | **Put** /v2/dashboards/{id}/share_request | requestAccess dashboards_v2
[**DashboardsV2Share**](DashboardsV2Api.md#DashboardsV2Share) | **Put** /v2/dashboards/{id}/shares | share dashboards_v2
[**DashboardsV2Show**](DashboardsV2Api.md#DashboardsV2Show) | **Get** /v2/dashboards/{id} | show dashboards_v2
[**DashboardsV2Update**](DashboardsV2Api.md#DashboardsV2Update) | **Put** /v2/dashboards/{id} | update dashboards_v2



## DashboardsV2Create

> ArduinoDashboardv2 DashboardsV2Create(ctx, dashboardv2)

create dashboards_v2

Create a new dashboard

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardv2** | [**Dashboardv2**](Dashboardv2.md)| DashboardV2Payload describes a dashboard | 

### Return type

[**ArduinoDashboardv2**](ArduinoDashboardv2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardsV2Delete

> DashboardsV2Delete(ctx, id)

delete dashboards_v2

Delete a dashboard

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the dashboard | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardsV2DeleteShare

> DashboardsV2DeleteShare(ctx, id, userId)

deleteShare dashboards_v2

Delete a user the dashboard has been shared with

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the dashboard | 
**userId** | **string**| The id of the user | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardsV2Link

> ArduinoVariableslinks DashboardsV2Link(ctx, id, widgetId, widgetlink)

link dashboards_v2

Link or detach widget variables

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the dashboard | 
**widgetId** | **string**| The id of the widget | 
**widgetlink** | [**Widgetlink**](Widgetlink.md)|  | 

### Return type

[**ArduinoVariableslinks**](ArduinoVariableslinks.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardsV2List

> []ArduinoDashboardv2 DashboardsV2List(ctx, optional)

list dashboards_v2

Returns the list of dashboards

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DashboardsV2ListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DashboardsV2ListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| The name of the dashboard | 
 **userId** | **optional.String**| The user_id of the dashboard&#39;s owner | 

### Return type

[**[]ArduinoDashboardv2**](ArduinoDashboardv2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardsV2ListShares

> []ArduinoDashboardshare DashboardsV2ListShares(ctx, id)

listShares dashboards_v2

List of users the dashboard has been shared with

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the dashboard | 

### Return type

[**[]ArduinoDashboardshare**](ArduinoDashboardshare.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardsV2RequestAccess

> DashboardsV2RequestAccess(ctx, id, sharerequest)

requestAccess dashboards_v2

Request access to a dashboard

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the dashboard | 
**sharerequest** | [**Sharerequest**](Sharerequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardsV2Share

> DashboardsV2Share(ctx, id, dashboardshare)

share dashboards_v2

Share a dashboard

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the dashboard | 
**dashboardshare** | [**Dashboardshare**](Dashboardshare.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardsV2Show

> ArduinoDashboardv2 DashboardsV2Show(ctx, id)

show dashboards_v2

Show a dashboard

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the dashboard | 

### Return type

[**ArduinoDashboardv2**](ArduinoDashboardv2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardsV2Update

> ArduinoDashboardv2 DashboardsV2Update(ctx, id, dashboardv2)

update dashboards_v2

Updates an existing dashboard

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the dashboard | 
**dashboardv2** | [**Dashboardv2**](Dashboardv2.md)| DashboardV2Payload describes a dashboard | 

### Return type

[**ArduinoDashboardv2**](ArduinoDashboardv2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

