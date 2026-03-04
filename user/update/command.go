package update

import (
    "fmt"
    //"strconv"
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
        err    error
        errRes api.AccountError
        userId string
    )

    client := client.New(g)

    // Fetch the account and user IDs for the current user.
    acctId, err = client.AccountID(); if err != nil {
        return err
    }
    userId, err = client.UserID(); if err != nil {
        return err
    }

    body := make(map[string]interface{})

    // Only add options to the body of the request if they were
    // explicitly set by a command-line option.
    if o.Salutation != nil {
        body["salutation"] = *o.Salutation
    }
    if o.FirstName != nil {
        body["firstName"] = *o.FirstName
    }
    if o.LastName != nil {
        body["lastName"] = *o.LastName
    }
    if o.Phone != nil {
        body["phone"] = *o.Phone
    }
    if o.Streets != nil {
        if body["address"] == nil {
            body["address"] = make(map[string]interface{})
        }
        body["address"].(map[string]interface{})["streets"] = *o.Streets
    }
    if o.City != nil {
        if body["address"] == nil {
            body["address"] = make(map[string]interface{})
        }
        body["address"].(map[string]interface{})["city"] = *o.City
    }
    if o.Country != nil {
        if body["address"] == nil {
            body["address"] = make(map[string]interface{})
        }
        body["address"].(map[string]interface{})["country"] = *o.Country
    }
    if o.State != nil {
        if body["address"] == nil {
            body["address"] = make(map[string]interface{})
        }
        body["address"].(map[string]interface{})["state"] = *o.State
    }
    if o.County != nil {
        if body["address"] == nil {
            body["address"] = make(map[string]interface{})
        }
        body["address"].(map[string]interface{})["county"] = *o.County
    }
    if o.Zip != nil {
        if body["address"] == nil {
            body["address"] = make(map[string]string)
        }
        body["address"].(map[string]interface{})["zipCode"] = *o.Zip
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
