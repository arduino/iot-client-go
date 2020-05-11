# Arduino iot-api Go client

## Getting Started

The client requires a valid access token to authenticate, you can use the
`golang.org/x/oauth2` to easily get one with the Client Credentials OAuth2 flow:

```Go
import cc "golang.org/x/oauth2/clientcredentials"

// We need to pass the additional "audience" var to request an access token
additionalValues := url.Values{}
additionalValues.Add("audience", "https://api2.arduino.cc/iot")
// Set up OAuth2 configuration
config := cc.Config{
    ClientID:       clientID,
    ClientSecret:   clientSecret,
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
```

For a working example, see [the example folder](./example) in this repo.
