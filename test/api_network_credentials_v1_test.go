/*
Arduino IoT Cloud API

Testing NetworkCredentialsV1ApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package iot

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/arduino/iot-client-go/v3"
)

func Test_iot_NetworkCredentialsV1ApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test NetworkCredentialsV1ApiService NetworkCredentialsV1Show", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string

		resp, httpRes, err := apiClient.NetworkCredentialsV1API.NetworkCredentialsV1Show(context.Background(), type_).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NetworkCredentialsV1ApiService NetworkCredentialsV1ShowByDevice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string

		httpRes, err := apiClient.NetworkCredentialsV1API.NetworkCredentialsV1ShowByDevice(context.Background(), type_).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
