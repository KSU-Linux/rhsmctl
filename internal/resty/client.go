package resty

import (
    "context"
    "net/url"
    "resty.dev/v3"
    "rhsmctl/internal/cli"
    "golang.org/x/oauth2/clientcredentials"
)

func New(g *cli.Globals) *resty.Client {
    clientCredCfg := &clientcredentials.Config{
        ClientID: g.ClientId,
        TokenURL: g.TokenUrl,
        EndpointParams: url.Values{
            "grant_type": {"refresh_token"},
            "refresh_token": {g.ApiToken},
        },
    }
    credClient := clientCredCfg.Client(context.Background())
    client := resty.NewWithClient(credClient)
    defer client.Close()
    return client
}
