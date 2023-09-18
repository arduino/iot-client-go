/*
Arduino IoT Cloud API

Testing DevicesV2OtaApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package iot

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/arduino/iot-client-go"
)

func Test_iot_DevicesV2OtaApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DevicesV2OtaApiService DevicesV2OtaSend", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.DevicesV2OtaApi.DevicesV2OtaSend(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DevicesV2OtaApiService DevicesV2OtaUpload", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.DevicesV2OtaApi.DevicesV2OtaUpload(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}