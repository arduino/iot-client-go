package v3

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
	"golang.org/x/oauth2"
	cc "golang.org/x/oauth2/clientcredentials"
)

var (
	ctx          context.Context
	client       *APIClient
	clientID     = flag.String("client-id", "", "Client ID obtained using client credentials")
	clientSecret = flag.String("client-secret", "", "Client Secreted obtained using client credentials")
)

func testCreateDevice(t *testing.T) ArduinoDevicev2 {
	deviceName := "TestDevice"
	devicePayload := CreateDevicesV2Payload{
		Name: &deviceName,
		Type: "mkr1000",
	}

	device, _, err := client.DevicesV2API.DevicesV2Create(ctx).CreateDevicesV2Payload(devicePayload).Execute()
	assert.NoError(t, err, "No errors creating device")
	assert.Equal(t, *devicePayload.Name, device.Name, "Device name was correctly set")
	assert.Equal(t, devicePayload.Type, device.Type, "Device type was correctly set")
	assert.NotNil(t, device.Id, "Device ID was correctly generated")
	return *device
}

func testCreateThing(t *testing.T, name string) ArduinoThing {
	time.Sleep(5 * time.Second)
	thingPayload := ThingCreate{
		Name: &name,
	}
	thing, _, err := client.ThingsV2API.ThingsV2Create(ctx).ThingCreate(thingPayload).Execute()
	assert.NoError(t, err, "No errors creating thing")
	assert.Equal(t, *thingPayload.Name, thing.Name, "Thing name was correctly set")
	assert.NotNil(t, thing.Id, "Thing ID was correctly generated")
	t.Logf("Created thing: %s userId: %s", thing.Id, thing.UserId)
	return *thing
}

func testAttachDeviceThing(t *testing.T, thingID, deviceID string) ArduinoThing {
	thing, _, err := client.ThingsV2API.ThingsV2Update(ctx, thingID).ThingUpdate(ThingUpdate{
		DeviceId: &deviceID,
	}).Execute()
	assert.NoError(t, err, "No errors updating thing")
	assert.Equal(t, deviceID, *thing.DeviceId, "Device was correctly attached")
	assert.Equal(t, thingID, thing.Id, "Thing id was correctly set")
	assert.NotNil(t, thing.Id, "Thing ID was correctly generated")
	return *thing
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
	ctx = context.WithValue(context.Background(), ContextOAuth2, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: tok.AccessToken},
	))

	// Create an instance of the iot-api Go client, we pass an empty config
	// because defaults are ok
	client = NewAPIClient(NewConfiguration())
	cleanup()
	code := m.Run()
	cleanup()
	// call flag.Parse() here if TestMain uses flags
	os.Exit(code)
}

func cleanup() {
	log.Printf("Cleaning devices...")
	// Delete devices
	devices, _, err := client.DevicesV2API.DevicesV2List(ctx).Execute()
	if err != nil {
		panic(err)
	}

	for _, d := range devices {
		_, err = client.DevicesV2API.DevicesV2Delete(ctx, d.Id).Execute()
		if err != nil {
			panic(err)
		}
	}

	// Delete things
	log.Printf("Cleaning things...")
	things, _, err := client.ThingsV2API.ThingsV2List(ctx).Execute()
	if err != nil {
		panic(err)
	}

	for _, t := range things {
		_, err = client.ThingsV2API.ThingsV2Delete(ctx, t.Id).Execute()
		if err != nil {
			panic(err)
		}
	}
}

func TestDevicesAPI(t *testing.T) {
	// Get the list of devices for the current user
	devices, _, err := client.DevicesV2API.DevicesV2List(ctx).Execute()
	assert.NoError(t, err, "No errors listing devices")

	// Ensure is empty
	assert.Equal(t, 0, len(devices), "Device list is empty")

	// Add a new device
	device := testCreateDevice(t)

	// Show device
	newDevice, _, err := client.DevicesV2API.DevicesV2Show(ctx, device.Id).Execute()
	assert.NoError(t, err, "No errors showing device")
	assert.Equal(t, device.Name, newDevice.Name, "Device Name is correct")
	assert.Equal(t, device.Type, newDevice.Type, "Device ID is correct")
	assert.Equal(t, device.Id, newDevice.Id, "Device ID is correct")

	// Check if there's only 1 device
	devices, _, err = client.DevicesV2API.DevicesV2List(ctx).Execute()
	assert.NoError(t, err, "No errors listing devices")
	assert.Equal(t, 1, len(devices), "Device list should contain only 1 device")

	// Update device name
	newName := "TestDevice2"
	updatedDevice, _, err := client.DevicesV2API.DevicesV2Update(ctx, device.Id).Devicev2(Devicev2{
		Name: &newName,
	}).Execute()
	assert.NoError(t, err, "No error updating device")
	assert.Equal(t, newName, updatedDevice.Name, "Name was updated correctly")

	// Delete device
	_, err = client.DevicesV2API.DevicesV2Delete(ctx, device.Id).Execute()
	assert.NoError(t, err, "No errors deleting device")

	// Ensure device list is empty
	devices, _, err = client.DevicesV2API.DevicesV2List(ctx).Execute()
	assert.NoError(t, err, "No errors listing devices")
	assert.Equal(t, 0, len(devices), "Device list is empty")

	// Try to get the no more existing device
	storedDevice, _, err := client.DevicesV2API.DevicesV2Show(ctx, device.Id).Execute()
	assert.Contains(t, err.Error(), "404", "Error should not found")
	assert.Nil(t, storedDevice, "Device should be empty")
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
	createdThing, _, err := client.ThingsV2API.ThingsV2Show(ctx, thing.Id).Execute()
	assert.NoError(t, err, "No errors showing thing")
	assert.Equal(t, thingName, createdThing.Name, "Name is correct")
	assert.Equal(t, device.Id, *createdThing.DeviceId, "Device is correct")

	// Delete thing
	_, err = client.ThingsV2API.ThingsV2Delete(ctx, thing.Id).Execute()
	assert.NoError(t, err, "No errors deleting thing")

	// Try to get the no more existing thing
	storedThing, _, err := client.ThingsV2API.ThingsV2Show(ctx, thing.Id).Execute()
	assert.Contains(t, err.Error(), "404", "Error should be not found")
	assert.Nil(t, storedThing, "Thing should be empty")

	// Delete device
	_, err = client.DevicesV2API.DevicesV2Delete(ctx, device.Id).Execute()
	assert.NoError(t, err, "No errors deleting device")
}

func Disabled_TestProperties(t *testing.T) {
	// Create a device
	device := testCreateDevice(t)

	thingName := fmt.Sprintf("ThingName-%d", time.Now().Unix())
	// Create a thing
	thing := testCreateThing(t, thingName)

	// Attach the device to the thing
	thing = testAttachDeviceThing(t, thing.Id, device.Id)

	// Create a property
	persist := true
	propertyPayload := Property{
		Name:           "testInt",
		Type:           "INT",
		Permission:     "READ_WRITE",
		UpdateStrategy: "ON_CHANGE",
		Persist: 		&persist,
	}
	property, _, err := client.PropertiesV2API.PropertiesV2Create(ctx, thing.Id).Property(propertyPayload).Execute()
	assert.NoError(t, err, "No errors creating properties")
	assert.Equal(t, propertyPayload.Name, property.Name, "Property name was set correctly")
	assert.Equal(t, propertyPayload.Type, property.Type, "Property type was set correctly")
	assert.Equal(t, propertyPayload.Permission, property.Permission, "Property permission was set correctly")
	assert.Equal(t, propertyPayload.UpdateStrategy, property.UpdateStrategy, "Property update strategy was set correctly")

	// Generate a sketch
	storedThing, _, err := client.ThingsV2API.ThingsV2CreateSketch(ctx, thing.Id).ThingSketch(ThingSketch{}).Execute()
	assert.NoError(t, err, "No errors creating sketch")
	assert.NotNil(t, storedThing.SketchId, "Sketch ID is not null")

	// Create another property
	propertyPayload = Property{
		Name:           "testInt2",
		Type:           "INT",
		Permission:     "READ_WRITE",
		UpdateStrategy: "ON_CHANGE",
		Persist:        &persist,
	}
	property, _, err = client.PropertiesV2API.PropertiesV2Create(ctx, thing.Id).Property(propertyPayload).Execute()
	assert.NoError(t, err, "No errors creating properties")
	assert.Equal(t, propertyPayload.Name, property.Name, "Property name was set correctly")
	assert.Equal(t, propertyPayload.Type, property.Type, "Property type was set correctly")
	assert.Equal(t, propertyPayload.Permission, property.Permission, "Property permission was set correctly")
	assert.Equal(t, propertyPayload.UpdateStrategy, property.UpdateStrategy, "Property update strategy was set correctly")
	assert.NotNil(t, property.Id, "Property ID is not null")

	// Update sketch
	storedThing, _, err = client.ThingsV2API.ThingsV2UpdateSketch(ctx, thing.Id, *storedThing.SketchId).Execute()
	assert.NoError(t, err, "No errors updating sketch")
	assert.NotNil(t, storedThing.SketchId, "Sketch ID is not null")

	// Publish property
	propertyValue := float64(100)
	_, err = client.PropertiesV2API.PropertiesV2Publish(ctx, thing.Id, property.Id).PropertyValue(PropertyValue{
		Value: propertyValue,
	}).Execute()
	assert.NoError(t, err, "No errors publishing property")

	// Wait for data pipeline ingest the last value
	time.Sleep(15 * time.Second)

	// Get Last value
	t.Logf("Check property thingid:%s pid:%s", thing.Id, property.Id)
	property, _, err = client.PropertiesV2API.PropertiesV2Show(ctx, thing.Id, property.Id).Execute()
	t.Logf("Error %v", err)
	assert.NoError(t, err, "No errors showing propery")
	assert.Equal(t, propertyValue, property.LastValue, "Last value is correct")

	// Get value from series batch query
	now := time.Now().UTC()
	from := now.Add(-3600 * time.Second)
	limit := int64(1000)
	sort := "ASC"
	request := BatchQueryRawRequestMediaV1{
		From:        &from,
		To:          &now,
		SeriesLimit: &limit,
		Q:           "property." + property.Id,
	}
	t.Logf("Time from %s, to %s", from, now)
	batch, _, err := client.SeriesV2API.SeriesV2BatchQueryRaw(ctx).BatchQueryRawRequestsMediaV1(BatchQueryRawRequestsMediaV1{
		Requests: []BatchQueryRawRequestMediaV1{
			request,
		},
	}).Execute()
	assert.NoError(t, err, "No errors in batch query")
	assert.Equal(t, int64(1), batch.Responses[0].CountValues, "Only 1 value should be present")
	assert.Equal(t, propertyValue, batch.Responses[0].Values[0], "Value should be correct")

	// Get value from series batch query raw
	batchRaw, _, err := client.SeriesV2API.SeriesV2BatchQueryRaw(ctx).BatchQueryRawRequestsMediaV1(BatchQueryRawRequestsMediaV1{
		Requests: []BatchQueryRawRequestMediaV1{
			{
				From:        &from,
				To:          &now,
				Q:           "property." + property.Id,
				SeriesLimit: &limit,
				Sort:        &sort,
			},
		},
	}).Execute()

	assert.NoError(t, err, "No errors getting raw series")
	assert.Equal(t, int64(1), batchRaw.Responses[0].CountValues, "Only 1 value should be present")
	assert.Equal(t, propertyValue, batchRaw.Responses[0].Values[0], "Value should be correct")

	batchLastValue, _, err := client.SeriesV2API.SeriesV2BatchQueryRawLastValue(ctx).BatchLastValueRequestsMediaV1(BatchLastValueRequestsMediaV1{
		Requests: []BatchQueryRawLastValueRequestMediaV1{
			{
				PropertyId: property.Id,
				ThingId:    thing.Id,
			},
		},
	}).Execute()

	assert.Equal(t, int64(1), batchLastValue.Responses[0].CountValues, "Only 1 value should be present")
	assert.Equal(t, propertyValue, batchLastValue.Responses[0].Values[0], "Value should be correct")
	assert.NoError(t, err, "No errors getting raw series last value")

	// Delete sketch
	storedThing, _, err = client.ThingsV2API.ThingsV2DeleteSketch(ctx, thing.Id).Execute()
	assert.NoError(t, err, "No errors updating sketch")
	assert.Nil(t, storedThing.SketchId, "Sketch ID is nil")

	// Delete property
	_, err = client.PropertiesV2API.PropertiesV2Delete(ctx, thing.Id, property.Id).Execute()
	assert.NoError(t, err, "No errors deleting property")

	// Delete device and thing
	_, err = client.DevicesV2API.DevicesV2Delete(ctx, device.Id).Execute()
	assert.NoError(t, err, "No errors deleting device")

	_, err = client.ThingsV2API.ThingsV2Delete(ctx, thing.Id).Execute()
	assert.NoError(t, err, "No errors deleting thing")
}
