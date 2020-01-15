package iot

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	cc "golang.org/x/oauth2/clientcredentials"
)

var (
	ctx          context.Context
	client       *APIClient
	clientID     = flag.String("client-id", "", "Client ID obtained using client credentials")
	clientSecret = flag.String("client-secret", "", "Client Secreted obtained using client credentials")
)

func testCreateDevice(t *testing.T) ArduinoDevicev2 {
	devicePayload := CreateDevicesV2Payload{
		Name: "TestDevice",
		Type: "mkr1000",
	}

	device, _, err := client.DevicesV2Api.DevicesV2Create(ctx, devicePayload)
	assert.NoError(t, err, "No errors creating device")
	assert.Equal(t, devicePayload.Name, device.Name, "Device name was correctly set")
	assert.Equal(t, devicePayload.Type, device.Type, "Device type was correctly set")
	assert.NotNil(t, device.Id, "Device ID was correctly generated")
	return device
}

func testCreateThing(t *testing.T, name string) ArduinoThing {
	thingPayload := CreateThingsV2Payload{
		Name: name,
	}
	thing, _, err := client.ThingsV2Api.ThingsV2Create(ctx, thingPayload, nil)
	assert.NoError(t, err, "No errors creating thing")
	assert.Equal(t, thingPayload.Name, thing.Name, "Thing name was correctly set")
	return thing
}

func testAttachDeviceThing(t *testing.T, thingID, deviceID string) ArduinoThing {
	thing, _, err := client.ThingsV2Api.ThingsV2Update(ctx, thingID, Thing{
		DeviceId: deviceID,
	}, nil)
	assert.NoError(t, err, "No errors updating thing")
	assert.Equal(t, deviceID, thing.DeviceId, "Device was correctly attached")
	return thing
}

func TestMain(m *testing.M) {
	// Check credentials
	flag.Parse()
	*clientID = strings.TrimSpace(*clientID)
	if *clientID == "" {
		*clientID = os.Getenv("CLIENT_ID")
	}

	*clientSecret = strings.TrimSpace(*clientSecret)
	if *clientSecret == "" {
		*clientSecret = os.Getenv("CLIENT_SECRET")
	}

	if *clientID == "" || *clientSecret == "" {
		log.Fatalf("Invalid credentials, use -client-id -client-secret")
	}

	// We need to pass the additional "audience" var to request an access token
	additionalValues := url.Values{}
	additionalValues.Add("audience", "https://api2.arduino.cc/iot")
	// Set up OAuth2 configuration
	config := cc.Config{
		ClientID:       *clientID,
		ClientSecret:   *clientSecret,
		TokenURL:       "https://api2.arduino.cc/iot/v1/clients/token",
		EndpointParams: additionalValues,
	}
	// Get the access token in exchange of client_id and client_secret
	tok, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("Error retrieving access token, %v", err)
	}
	// Confirm we got the token and print expiration time
	log.Printf("Got an access token, will expire on %s", tok.Expiry)

	// We use the token to create a context that will be passed to any API call
	ctx = context.WithValue(context.Background(), ContextAccessToken, tok.AccessToken)

	// Create an instance of the iot-api Go client, we pass an empty config
	// because defaults are ok
	client = NewAPIClient(NewConfiguration())

	code := m.Run()
	cleanup()
	// call flag.Parse() here if TestMain uses flags
	os.Exit(code)
}

func cleanup() {
	log.Printf("Cleaning devices...")
	// Delete devices
	devices, _, err := client.DevicesV2Api.DevicesV2List(ctx, nil)
	if err != nil {
		panic(err)
	}

	for _, d := range devices {
		_, err = client.DevicesV2Api.DevicesV2Delete(ctx, d.Id)
		if err != nil {
			panic(err)
		}
	}

	// Delete things
	log.Printf("Cleaning things...")
	things, _, err := client.ThingsV2Api.ThingsV2List(ctx, nil)
	if err != nil {
		panic(err)
	}

	for _, t := range things {
		_, err = client.ThingsV2Api.ThingsV2Delete(ctx, t.Id, nil)
		if err != nil {
			panic(err)
		}
	}
}

func TestDevicesAPI(t *testing.T) {
	// Get the list of devices for the current user
	devices, _, err := client.DevicesV2Api.DevicesV2List(ctx, nil)
	assert.NoError(t, err, "No errors listing devices")

	// Ensure is empty
	assert.Equal(t, 0, len(devices), "Device list is empty")

	// Add a new device
	device := testCreateDevice(t)

	// Show device
	newDevice, _, err := client.DevicesV2Api.DevicesV2Show(ctx, device.Id)
	assert.NoError(t, err, "No errors showing device")
	assert.Equal(t, device.Name, newDevice.Name, "Device Name is correct")
	assert.Equal(t, device.Type, newDevice.Type, "Device ID is correct")
	assert.Equal(t, device.Id, newDevice.Id, "Device ID is correct")

	// Check if there's only 1 device
	devices, _, err = client.DevicesV2Api.DevicesV2List(ctx, nil)
	assert.NoError(t, err, "No errors listing devices")
	assert.Equal(t, 1, len(devices), "Device list should contain only 1 device")

	// Update device name
	newName := "TestDevice2"
	device, _, err = client.DevicesV2Api.DevicesV2Update(ctx, device.Id, Devicev2{
		Name: newName,
	})
	assert.NoError(t, err, "No error updating device")
	assert.Equal(t, newName, device.Name, "Name was updated correctly")

	// Delete device
	_, err = client.DevicesV2Api.DevicesV2Delete(ctx, device.Id)
	assert.NoError(t, err, "No errors deleting device")

	// Ensure device list is empty
	devices, _, err = client.DevicesV2Api.DevicesV2List(ctx, nil)
	assert.NoError(t, err, "No errors listing devices")
	assert.Equal(t, 0, len(devices), "Device list is empty")

	// Try to get the no more existing device
	device, _, err = client.DevicesV2Api.DevicesV2Show(ctx, device.Id)
	assert.EqualError(t, err, "401 Unauthorized", "Error should be unauthorized")
	assert.Equal(t, ArduinoDevicev2{}, device, "Device should be empty")
}

func TestThingsAPI(t *testing.T) {
	// Add a new device
	device := testCreateDevice(t)

	// Create a thing without a device
	thingName := "TestThing"
	thing := testCreateThing(t, thingName)

	// Attach a device to the thing
	thing = testAttachDeviceThing(t, thing.Id, device.Id)

	// Show thing
	thing, _, err := client.ThingsV2Api.ThingsV2Show(ctx, thing.Id, nil)
	assert.NoError(t, err, "No errors showing thing")
	assert.Equal(t, thingName, thing.Name, "Name is correct")
	assert.Equal(t, device.Id, thing.DeviceId, "Device is correct")

	// Delete thing
	_, err = client.ThingsV2Api.ThingsV2Delete(ctx, thing.Id, nil)
	assert.NoError(t, err, "No errors deleting thing")

	// Try to get the no more existing thing
	thing, _, err = client.ThingsV2Api.ThingsV2Show(ctx, thing.Id, nil)
	assert.EqualError(t, err, "404 Not Found", "Error should be not found")
	assert.Equal(t, ArduinoThing{}, thing, "Thing should be empty")

	// Delete device
	_, err = client.DevicesV2Api.DevicesV2Delete(ctx, device.Id)
	assert.NoError(t, err, "No errors deleting device")
}

func TestThingsLimit(t *testing.T) {
	// Create 5 things
	things := make([]ArduinoThing, 0)
	for k := 0; k < 5; k++ {
		thingPayload := CreateThingsV2Payload{
			Name: fmt.Sprintf("TestThing-%d", k),
		}
		thing, _, err := client.ThingsV2Api.ThingsV2Create(ctx, thingPayload, nil)
		assert.NoError(t, err, "No errors creating thing")
		assert.Equal(t, thingPayload.Name, thing.Name, "Thing name was correctly set")
		things = append(things, thing)
	}

	// Create the 6th thing over limit
	thingPayload := CreateThingsV2Payload{
		Name: "TestThing-6",
	}
	thing, _, err := client.ThingsV2Api.ThingsV2Create(ctx, thingPayload, nil)
	assert.EqualError(t, err, "412 Precondition Failed", "Cannot create over 5 things")
	assert.Equal(t, ArduinoThing{}, thing, "Thing is empty")

	// Delete things
	for _, thing := range things {
		_, err := client.ThingsV2Api.ThingsV2Delete(ctx, thing.Id, nil)
		assert.NoError(t, err, "No errors deleting thing")
	}
}

func TestProperties(t *testing.T) {
	// Create a device
	device := testCreateDevice(t)

	// Create a thing
	thing := testCreateThing(t, "ThingName")

	// Attach the device to the thing
	thing = testAttachDeviceThing(t, thing.Id, device.Id)

	// Create a property
	propertyPayload := Property{
		Name:           "testInt",
		Type:           "INT",
		Permission:     "READ_WRITE",
		UpdateStrategy: "ON_CHANGE",
	}
	property, _, err := client.PropertiesV2Api.PropertiesV2Create(ctx, thing.Id, propertyPayload)
	assert.NoError(t, err, "No errors creating properties")
	assert.Equal(t, propertyPayload.Name, property.Name, "Property name was set correctly")
	assert.Equal(t, propertyPayload.Type, property.Type, "Property type was set correctly")
	assert.Equal(t, propertyPayload.Permission, property.Permission, "Property permission was set correctly")
	assert.Equal(t, propertyPayload.UpdateStrategy, property.UpdateStrategy, "Property update strategy was set correctly")

	// Generate a sketch
	thing, _, err = client.ThingsV2Api.ThingsV2CreateSketch(ctx, thing.Id, ThingSketch{
		Ssid: "ssid",
		Pass: "password",
	})
	assert.NoError(t, err, "No errors creating sketch")
	assert.NotNil(t, thing.SketchId, "Sketch ID is not null")

	// Create another property
	propertyPayload = Property{
		Name:           "testInt2",
		Type:           "INT",
		Permission:     "READ_WRITE",
		UpdateStrategy: "ON_CHANGE",
		Persist:        true,
	}
	property, _, err = client.PropertiesV2Api.PropertiesV2Create(ctx, thing.Id, propertyPayload)
	assert.NoError(t, err, "No errors creating properties")
	assert.Equal(t, propertyPayload.Name, property.Name, "Property name was set correctly")
	assert.Equal(t, propertyPayload.Type, property.Type, "Property type was set correctly")
	assert.Equal(t, propertyPayload.Permission, property.Permission, "Property permission was set correctly")
	assert.Equal(t, propertyPayload.UpdateStrategy, property.UpdateStrategy, "Property update strategy was set correctly")

	// Update sketch
	thing, _, err = client.ThingsV2Api.ThingsV2UpdateSketch(ctx, thing.Id, thing.SketchId, nil)
	assert.NoError(t, err, "No errors updating sketch")
	assert.NotNil(t, thing.SketchId, "Sketch ID is not null")

	// Publish property
	propertyValue := float64(100)
	_, err = client.PropertiesV2Api.PropertiesV2Publish(ctx, thing.Id, property.Id, PropertyValue{
		Value: propertyValue,
	})
	assert.NoError(t, err, "No errors publishing property")

	// Wait for data pipeline ingest the last value
	time.Sleep(5 * time.Second)

	// Get Last value
	property, _, err = client.PropertiesV2Api.PropertiesV2Show(ctx, thing.Id, property.Id, nil)
	assert.NoError(t, err, "No errors showing propery")
	assert.Equal(t, propertyValue, property.LastValue, "Last value is correct")

	// Get value from series batch query
	request := BatchQueryRequestMediaV1{
		From:        time.Now().Add(-60 * time.Second),
		To:          time.Now(),
		Interval:    1,
		SeriesLimit: 1000,
		Q:           "property." + property.Id,
	}
	batch, _, err := client.SeriesV2Api.SeriesV2BatchQuery(ctx, BatchQueryRequestsMediaV1{
		Requests: []BatchQueryRequestMediaV1{
			request,
		},
	})
	assert.NoError(t, err, "No errors in batch query")
	assert.Equal(t, int64(1), batch.Responses[0].CountValues, "Only 1 value should be present")
	assert.Equal(t, propertyValue, batch.Responses[0].Values[0], "Value should be correct")

	// Get value from series batch query raw
	batchRaw, _, err := client.SeriesV2Api.SeriesV2BatchQueryRaw(ctx, BatchQueryRawRequestsMediaV1{
		Requests: []BatchQueryRawRequestMediaV1{
			BatchQueryRawRequestMediaV1{
				From:        time.Now().Add(-60 * time.Second),
				To:          time.Now(),
				Q:           "property." + property.Id,
				SeriesLimit: 1000,
				Sort:        "ASC",
			},
		},
	})

	assert.NoError(t, err, "No errors getting raw series")
	assert.Equal(t, int64(1), batchRaw.Responses[0].CountValues, "Only 1 value should be present")
	assert.Equal(t, propertyValue, batchRaw.Responses[0].Values[0], "Value should be correct")

	batchLastValue, _, err := client.SeriesV2Api.SeriesV2BatchQueryRawLastValue(ctx, BatchLastValueRequestsMediaV1{
		Requests: []BatchQueryRawLastValueRequestMediaV1{
			BatchQueryRawLastValueRequestMediaV1{
				PropertyId: property.Id,
				ThingId:    thing.Id,
			},
		},
	})

	assert.Equal(t, int64(1), batchLastValue.Responses[0].CountValues, "Only 1 value should be present")
	assert.Equal(t, propertyValue, batchLastValue.Responses[0].Values[0], "Value should be correct")
	assert.NoError(t, err, "No errors getting raw series last value")

	// Delete sketch
	thing, _, err = client.ThingsV2Api.ThingsV2DeleteSketch(ctx, thing.Id)
	assert.NoError(t, err, "No errors updating sketch")
	assert.Equal(t, "", thing.SketchId, "Sketch ID is empty")

	// Delete property
	_, err = client.PropertiesV2Api.PropertiesV2Delete(ctx, thing.Id, property.Id, nil)
	assert.NoError(t, err, "No errors deleting property")

	// Delete device and thing
	_, err = client.DevicesV2Api.DevicesV2Delete(ctx, device.Id)
	assert.NoError(t, err, "No errors deleting device")

	_, err = client.ThingsV2Api.ThingsV2Delete(ctx, thing.Id, nil)
	assert.NoError(t, err, "No errors deleting thing")
}
