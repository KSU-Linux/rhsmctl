package client

import (
    "context"
    "net/url"
    "resty.dev/v3"
    "rhsmctl/internal/cli"
    "golang.org/x/oauth2/clientcredentials"
)

func New(g *cli.Globals) *resty.Client {
    ccc := &clientcredentials.Config{
        ClientID: g.ClientId,
        TokenURL: g.TokenUrl,
        EndpointParams: url.Values{
            "grant_type": {"refresh_token"},
            "refresh_token": {g.ApiKey},
        },
    }
    cc := ccc.Client(context.Background())
    c := resty.NewWithClient(cc)
    defer c.Close()
    return c
}
