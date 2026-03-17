package status

import (
    "fmt"
    "strconv"
    "strings"
    "rhsmctl/internal/api"
    "rhsmctl/internal/cli"
    "rhsmctl/internal/client"
    "rhsmctl/internal/tty"
    "github.com/alecthomas/kong"
)

func (o *Options) Run(ctx *kong.Context, g *cli.Globals) error {
    var (
        acctId string
        errRes api.AccountError
        userId string
    )

    client := client.New(g)

    // Use account id if provided, otherwise attempt to fetch it.
    if (o.AccountID != nil) {
        acctId = strconv.Itoa(*o.AccountID)
    } else {
        id, err := client.AccountID(); if err != nil {
            return err
        }
        acctId = id
    }

    // Use user id if provided, otherwise attempt to fetch it.
    if (o.UserID != 0) {
        userId = strconv.Itoa(o.UserID)
    } else {
        id, err := client.UserID(); if err != nil {
            return err
        }
        userId = id
    }

    res, err := client.R().
        SetDebug(g.Debug).
        SetError(&errRes).
        SetPathParams(map[string]string{
            "accountId": acctId,
            "id": userId,
        }).
        Get(g.ApiUrl+"/account/v1/accounts/{accountId}/users/{id}/status")
    if (err != nil) {
        return err
    }
    if (res.IsError()) {
        return fmt.Errorf("error: %s", strings.ToLower(errRes.Detail))
    }
    tty.Printjson(res.Bytes())
    return nil
}
