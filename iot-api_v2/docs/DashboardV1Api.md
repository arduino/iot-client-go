# \DashboardV1Api

All URIs are relative to *http://api-dev.arduino.cc*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DashboardV1Create**](DashboardV1Api.md#DashboardV1Create) | **Post** /iot/v1/dashboards | create dashboard_v1
[**DashboardV1List**](DashboardV1Api.md#DashboardV1List) | **Get** /iot/v1/dashboards | list dashboard_v1
[**DashboardV1Show**](DashboardV1Api.md#DashboardV1Show) | **Get** /iot/v1/dashboards/{id}/types/{type} | show dashboard_v1
[**DashboardV1Update**](DashboardV1Api.md#DashboardV1Update) | **Put** /iot/v1/dashboards/{id} | update dashboard_v1



## DashboardV1Create

> ArduinoDashboard DashboardV1Create(ctx, )
create dashboard_v1

Create a dashboard    Http body example:       {    \"name\": \"Dashboard name\",    \"type\": \"THING|SPACE|PROPERTIES\",    \"configuration\": {     \"description\": \"The configuration key is a general json field where the front end will store all the information that are useful to visualize the dashboard\"    }   }     

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ArduinoDashboard**](ArduinoDashboard.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.dashboard+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardV1List

> ArduinoDashboardlist DashboardV1List(ctx, )
list dashboard_v1

Returns the list of dashboards

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ArduinoDashboardlist**](ArduinoDashboardlist.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.dashboardlist+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardV1Show

> ArduinoDashboard DashboardV1Show(ctx, id, type_)
show dashboard_v1

Returns the dashboard requested by the user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the dashboard | 
**type_** | **string**| The dashboard type | 

### Return type

[**ArduinoDashboard**](ArduinoDashboard.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.dashboard+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardV1Update

> DashboardV1Update(ctx, id)
update dashboard_v1

Creates or update a dashboard.    Http body example:     {     \"id\": \"a860f4ba-f346-4782-8595-748396b36f5d\",     \"name\": \"Dashboard name\",     \"type\": \"THING|SPACE|PROPERTIES\",     \"configuration\": {    \"description\": \"The configuration key is a general json field where the front end will store all the information that are useful to visualize the dashboard\"     }   }

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of the dashboard | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

