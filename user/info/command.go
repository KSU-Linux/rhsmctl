package info

import (
    "fmt"
    "strings"
    "rhsmctl/internal/api"
    "rhsmctl/internal/cli"
    "rhsmctl/internal/client"
    "rhsmctl/internal/tty"
    "github.com/alecthomas/kong"
)

func (o *Options) Run(ctx *kong.Context, g *cli.Globals) error {
    var (
        acctId string = ""
        err    error
        errRes api.AccountError
        path   string = "/account/v1/user"
        userId string = ""
    )

    client := client.New(g)

    // Get the current user's account and user ID if extended
    // details are requested.
    if o.Details {
        acctId, err = client.AccountID(); if err != nil {
            return err
        }
        userId, err = client.UserID(); if err != nil {
            return err
        }
        // Update the path string since details are retrieved
        // from another endpoint.
        path = "/account/v1/accounts/{accountId}/users/{id}"
    }

    res, err := client.R().
        SetDebug(g.Debug).
        SetError(&errRes).
        SetPathParams(map[string]string{
            "accountId": acctId,
            "id": userId,
        }).
        Get(g.ApiUrl+path)
    if (err != nil) {
        return err
    }
    if (res.IsError()) {
        return fmt.Errorf("error: %s", strings.ToLower(errRes.Detail))
    }
    tty.Printjson(res.Bytes())
    return nil
}
