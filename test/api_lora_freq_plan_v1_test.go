/*
Arduino IoT Cloud API

Testing LoraFreqPlanV1ApiService

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

func Test_iot_LoraFreqPlanV1ApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LoraFreqPlanV1ApiService LoraFreqPlanV1List", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LoraFreqPlanV1API.LoraFreqPlanV1List(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
