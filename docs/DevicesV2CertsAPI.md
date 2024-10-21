# \DevicesV2CertsAPI

All URIs are relative to *https://api2.arduino.cc/iot*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevicesV2CertsCreate**](DevicesV2CertsAPI.md#DevicesV2CertsCreate) | **Put** /v2/devices/{id}/certs | create devices_v2_certs
[**DevicesV2CertsDelete**](DevicesV2CertsAPI.md#DevicesV2CertsDelete) | **Delete** /v2/devices/{id}/certs/{cid} | delete devices_v2_certs
[**DevicesV2CertsList**](DevicesV2CertsAPI.md#DevicesV2CertsList) | **Get** /v2/devices/{id}/certs | list devices_v2_certs
[**DevicesV2CertsShow**](DevicesV2CertsAPI.md#DevicesV2CertsShow) | **Get** /v2/devices/{id}/certs/{cid} | show devices_v2_certs
[**DevicesV2CertsUpdate**](DevicesV2CertsAPI.md#DevicesV2CertsUpdate) | **Post** /v2/devices/{id}/certs/{cid} | update devices_v2_certs



## DevicesV2CertsCreate

> ArduinoDevicev2Cert DevicesV2CertsCreate(ctx, id).CreateDevicesV2CertsPayload(createDevicesV2CertsPayload).Execute()

create devices_v2_certs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/arduino/iot-client-go/v2"
)

func main() {
	id := "id_example" // string | The id of the device
	createDevicesV2CertsPayload := *openapiclient.NewCreateDevicesV2CertsPayload("-----BEGIN CERTIFICATE-----
			MIIBeDCCAR4CAQAwgY0xCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpTb21lLVN0YXRl
			MQ8wDQYDVQQHEwZNeUNpdHkxFDASBgNVBAoTC0NvbXBhbnkgTHRkMQswCQYDVQQL
			EwJJVDEUMBIGA1UEAxMLZXhhbXBsZS5jb20xHzAdBgkqhkiG9w0BCQEMEHRlc3RA
			ZXhhbXBsZS5jb20wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATf6J9Gk79XGJ2I
			+v6p/r0UmPufUcUwtlx7gx87+DaI8Vpj9R5KN71HsHYw5uq+Lm0cr0CZIdtZU4cP
			upd6jDQToC4wLAYJKoZIhvcNAQkOMR8wHTAbBgNVHREEFDASgRB0ZXN0QGV4YW1w
			bGUuY29tMAoGCCqGSM49BAMCA0gAMEUCIGQqtlGzYdjPwYZYJ41albMBcdrKI7+8
			oiNSNWyDxJSGAiEAqQPPxMdr6vaXCCjr5s1J01WLKHzGoPFCR40rqAPs8eQ=
			-----END CERTIFICATE-----
			", false) // CreateDevicesV2CertsPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesV2CertsAPI.DevicesV2CertsCreate(context.Background(), id).CreateDevicesV2CertsPayload(createDevicesV2CertsPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesV2CertsAPI.DevicesV2CertsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DevicesV2CertsCreate`: ArduinoDevicev2Cert
	fmt.Fprintf(os.Stdout, "Response from `DevicesV2CertsAPI.DevicesV2CertsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesV2CertsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDevicesV2CertsPayload** | [**CreateDevicesV2CertsPayload**](CreateDevicesV2CertsPayload.md) |  | 

### Return type

[**ArduinoDevicev2Cert**](ArduinoDevicev2Cert.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.arduino.devicev2.cert+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2CertsDelete

> DevicesV2CertsDelete(ctx, cid, id).Execute()

delete devices_v2_certs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/arduino/iot-client-go/v2"
)

func main() {
	cid := "cid_example" // string | The id of the cert
	id := "id_example" // string | The id of the device

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevicesV2CertsAPI.DevicesV2CertsDelete(context.Background(), cid, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesV2CertsAPI.DevicesV2CertsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string** | The id of the cert | 
**id** | **string** | The id of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesV2CertsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.goa.error+json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2CertsList

> []ArduinoDevicev2Cert DevicesV2CertsList(ctx, id).Execute()

list devices_v2_certs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/arduino/iot-client-go/v2"
)

func main() {
	id := "id_example" // string | The id of the device

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesV2CertsAPI.DevicesV2CertsList(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesV2CertsAPI.DevicesV2CertsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DevicesV2CertsList`: []ArduinoDevicev2Cert
	fmt.Fprintf(os.Stdout, "Response from `DevicesV2CertsAPI.DevicesV2CertsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesV2CertsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ArduinoDevicev2Cert**](ArduinoDevicev2Cert.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.devicev2.cert+json; type=collection, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2CertsShow

> ArduinoDevicev2Cert DevicesV2CertsShow(ctx, cid, id).Execute()

show devices_v2_certs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/arduino/iot-client-go/v2"
)

func main() {
	cid := "cid_example" // string | The id of the cert
	id := "id_example" // string | The id of the device

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesV2CertsAPI.DevicesV2CertsShow(context.Background(), cid, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesV2CertsAPI.DevicesV2CertsShow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DevicesV2CertsShow`: ArduinoDevicev2Cert
	fmt.Fprintf(os.Stdout, "Response from `DevicesV2CertsAPI.DevicesV2CertsShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string** | The id of the cert | 
**id** | **string** | The id of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesV2CertsShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ArduinoDevicev2Cert**](ArduinoDevicev2Cert.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.arduino.devicev2.cert+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesV2CertsUpdate

> ArduinoDevicev2Cert DevicesV2CertsUpdate(ctx, cid, id).Devicev2Cert(devicev2Cert).Execute()

update devices_v2_certs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/arduino/iot-client-go/v2"
)

func main() {
	cid := "cid_example" // string | The id of the cert
	id := "id_example" // string | The id of the device
	devicev2Cert := *openapiclient.NewDevicev2Cert() // Devicev2Cert | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesV2CertsAPI.DevicesV2CertsUpdate(context.Background(), cid, id).Devicev2Cert(devicev2Cert).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesV2CertsAPI.DevicesV2CertsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DevicesV2CertsUpdate`: ArduinoDevicev2Cert
	fmt.Fprintf(os.Stdout, "Response from `DevicesV2CertsAPI.DevicesV2CertsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string** | The id of the cert | 
**id** | **string** | The id of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesV2CertsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **devicev2Cert** | [**Devicev2Cert**](Devicev2Cert.md) |  | 

### Return type

[**ArduinoDevicev2Cert**](ArduinoDevicev2Cert.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.arduino.devicev2.cert+json, application/vnd.goa.error+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

