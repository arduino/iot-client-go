# ArduinoDevicev2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | [**time.Time**](time.Time.md) | Creation date of the device | [optional] 
**Events** | [**[]ArduinoDevicev2SimpleProperties**](ArduinoDevicev2SimpleProperties.md) | ArduinoDevicev2SimplePropertiesCollection is the media type for an array of ArduinoDevicev2SimpleProperties (default view) | [optional] 
**Fqbn** | **string** | The fully qualified board name | [optional] 
**Href** | **string** | The api reference of this device | 
**Id** | **string** | The arn of the device | 
**Label** | **string** | The label of the device | 
**LastActivityAt** | [**time.Time**](time.Time.md) | Last activity date | [optional] 
**LatestWifiFwVersion** | **string** | The latest version of the NINA/WIFI101 firmware available for this device | [optional] 
**Metadata** | **map[string]interface{}** | The metadata of the device | [optional] 
**Name** | **string** | The friendly name of the device | 
**OtaAvailable** | **bool** | True if the device type is ready to receive OTA updated | [optional] 
**OtaCompatible** | **bool** | True if the device type is OTA compatible | [optional] 
**RequiredWifiFwVersion** | **string** | The required version of the NINA/WIFI101 firmware needed by IoT Cloud | [optional] 
**Serial** | **string** | The serial uuid of the device | 
**Thing** | [**ArduinoThing**](ArduinoThing.md) |  | [optional] 
**Type** | **string** | The type of the device | 
**UserId** | **string** | The id of the user | 
**Webhooks** | [**[]ArduinoDevicev2Webhook**](ArduinoDevicev2Webhook.md) | ArduinoDevicev2WebhookCollection is the media type for an array of ArduinoDevicev2Webhook (default view) | [optional] 
**WifiFwVersion** | **string** | The version of the NINA/WIFI101 firmware running on the device | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


