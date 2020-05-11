# ArduinoProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | [**time.Time**](time.Time.md) | Creation date of the property | [optional] 
**DeletedAt** | [**time.Time**](time.Time.md) | Delete date of the property | [optional] 
**Href** | **string** | The api reference of this property | 
**Id** | **string** | The id of the property | 
**LastValue** | [**interface{}**](.md) | Last value of this property | [optional] 
**MaxValue** | **float64** | Maximum value of this property | [optional] 
**MinValue** | **float64** | Minimum value of this property | [optional] 
**Name** | **string** | The friendly name of the property | 
**Permission** | **string** | The permission of the property | 
**Persist** | **bool** | If true, data will persist into a timeseries database | [optional] 
**Tag** | **float64** | The integer id of the property | [optional] 
**ThingId** | **string** | The id of the thing | 
**Type** | **string** | The type of the property | 
**UpdateParameter** | **float64** | The update frequency in seconds, or the amount of the property has to change in order to trigger an update | [optional] 
**UpdateStrategy** | **string** | The update strategy for the property value | 
**UpdatedAt** | [**time.Time**](time.Time.md) | Update date of the property | [optional] 
**ValueUpdatedAt** | [**time.Time**](time.Time.md) | Last update timestamp of this property | [optional] 
**VariableName** | **string** | The sketch variable name of the property | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


