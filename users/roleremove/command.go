package roleremove

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
        userId string = strconv.Itoa(o.UserID)
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

    res, err := client.R().
        SetBody(map[string]string{
            "role": o.Role,
        }).
        SetDebug(g.Debug).
        SetError(&errRes).
        SetPathParams(map[string]string{
            "accountId": acctId,
            "id": userId,
        }).
        Delete(g.ApiUrl+"/account/v1/accounts/{accountId}/users/{id}/roles")
    if (err != nil) {
        return err
    }
    if (res.IsError()) {
        return fmt.Errorf("error: %s", strings.ToLower(errRes.Detail))
    }
    tty.Printjson(res.Bytes())
    return nil
}
