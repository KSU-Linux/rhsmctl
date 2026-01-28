package client

import (
    "context"
    "net/url"
    "resty.dev/v3"
    "rhsmctl/internal/cli"
    "golang.org/x/oauth2/clientcredentials"
)

// RestClient wraps resty.Client.
type RestClient struct {
    resty *resty.Client
    globals *cli.Globals
}

func New(globals *cli.Globals) *RestClient {
    // Initialize client credentials.
    clientCredCfg := &clientcredentials.Config{
        ClientID: globals.ClientId,
        TokenURL: globals.TokenUrl,
        EndpointParams: url.Values{
            "grant_type": {"refresh_token"},
            "refresh_token": {globals.ApiToken},
        },
    }

    // Create a credentials client.
    credClient := clientCredCfg.Client(context.Background())

    // Create a new Resty client with the credentials client.
    restyClient := resty.NewWithClient(credClient)
    defer restyClient.Close()

    // Create a new rest client.
    client := &RestClient{resty: restyClient, globals: globals}

    return client
}

// R returns new resty.Request from client.
func (client *RestClient) R() *resty.Request {
    request := client.resty.R()
    return request
}
