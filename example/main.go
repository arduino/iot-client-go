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
	additionalValues := url.Values{}
	additionalValues.Add("audience", "https://api.arduino.cc")
	config := cc.Config{
		ClientID:       clientID,
		ClientSecret:   clientSecret,
		TokenURL:       "https://login.oniudra.cc/oauth/token",
		EndpointParams: additionalValues,
	}

	tok, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("Error retrieving access token, %v", err)
	}

	log.Printf("Got an access token, will expire on %s", tok.Expiry)

	auth := context.WithValue(context.Background(), sw.ContextAccessToken, tok.Token)
	iotCfg := iot.NewConfiguration()
	client := iot.NewAPIClient(iotCfg)
}
