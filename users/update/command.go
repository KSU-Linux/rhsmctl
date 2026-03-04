package update

import (
    "fmt"
    //"reflect"
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
    if (o.AccountID == nil) {
        acctId = strconv.Itoa(*o.AccountID)
    } else {
        id, err := client.AccountID(); if err != nil {
            return err
        }
        acctId = id
    }

    // Use user id if provided, otherwise attempt to fetch it.
    if (o.UserID == nil) {
        userId = strconv.Itoa(*o.UserID)
    } else {
        id, err := client.UserID(); if err != nil {
            return err
        }
        userId = id
    }

    body := make(map[string]interface{})

    // Only add options to the body of the request if they were
    // explicitly set by a command-line option.
    if o.EMail != nil {
        body["email"] = *o.EMail
    }
    if o.Permissions != nil {
        body["permissions"] = *o.Permissions
    }
    if o.Roles != nil {
        body["roles"] = *o.Roles
    }

    res, err := client.R().
        SetBody(body).
        SetDebug(g.Debug).
        SetError(&errRes).
        SetPathParams(map[string]string{
            "accountId": acctId,
            "id": userId,
        }).
        Post(g.ApiUrl+"/account/v1/accounts/{accountId}/users/{id}")
    if (err != nil) {
        return err
    }
    if (res.IsError()) {
        return fmt.Errorf("error: %s", strings.ToLower(errRes.Detail))
    }
    tty.Printjson(res.Bytes())
    return nil
}
