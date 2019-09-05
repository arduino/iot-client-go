package main

import (
	"context"
	"log"
	"net/url"

	iot "github.com/bcmi-labs/iot-api-client-go"
	cc "golang.org/x/oauth2/clientcredentials"
)

var (
	clientID     = "yr9Tb0P4qFOI3cziKizqUFZTrUvaeL47"
	clientSecret = "D-P2atuyJEZpB64_vymSQQYr1MPMuZhhK0e1U3jgPGVLRBSMH5jhKaivyRybu_4I"
)

func main() {
	// We need to pass the additional "audience" var to request an access token
	additionalValues := url.Values{}
	additionalValues.Add("audience", "https://api.arduino.cc")
	// Set up OAuth2 configuration
	config := cc.Config{
		ClientID:       clientID,
		ClientSecret:   clientSecret,
		TokenURL:       "https://login.oniudra.cc/oauth/token",
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
	ctx := context.WithValue(context.Background(), iot.ContextAccessToken, tok.AccessToken)

	// Create an instance of the iot-api Go client, we pass an empty config
	// because defaults are ok
	client := iot.NewAPIClient(iot.NewConfiguration())

	// Get the list of devices for the current user
	devices, _, err := client.DevicesV2Api.DevicesV2List(ctx, nil)
	if err != nil {
		log.Fatalf("Error getting devices, %v", err)
	}

	// Print a meaningful message if the api call succeeded
	if len(devices) == 0 {
		log.Printf("No device found")
	} else {
		for _, d := range devices {
			log.Printf("Device found: %s", d.Name)
		}
	}
}
