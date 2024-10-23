/*
Arduino IoT Cloud API

Testing DevicesV2TagsApiService

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

func Test_iot_DevicesV2TagsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DevicesV2TagsApiService DevicesV2TagsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var key string

		httpRes, err := apiClient.DevicesV2TagsAPI.DevicesV2TagsDelete(context.Background(), id, key).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DevicesV2TagsApiService DevicesV2TagsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.DevicesV2TagsAPI.DevicesV2TagsList(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DevicesV2TagsApiService DevicesV2TagsUpsert", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.DevicesV2TagsAPI.DevicesV2TagsUpsert(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
