package list

import (
    "fmt"
    r "resty.dev/v3"
    "strconv"
    "strings"
    "rhsmctl/internal/api"
    "rhsmctl/internal/cli"
    "rhsmctl/internal/resty"
    "rhsmctl/internal/tty"
    "github.com/alecthomas/kong"
)

func (o *Options) Run(ctx *kong.Context, g *cli.Globals) error {
    var (
        acctId string
        errRes api.AccountErrorRoot
        okRes api.Account
        res *r.Response
        err error
    )
    client := resty.New(g)

    if (o.AccountID != 0) {
        acctId = strconv.Itoa(o.AccountID)
    } else {
        res, err = client.R().
            SetDebug(g.Debug).
            SetError(&errRes).
            SetResult(&okRes).
            SetQueryParams(map[string]string{
                "firstResultIndex": strconv.Itoa(o.Offset),
                "maxResults": strconv.Itoa(o.Limit),
            }).
            Get(g.ApiUrl+"/account/v1/accounts")
        if (err != nil) {
            return err
        }
        if (res.IsError()) {
            return fmt.Errorf("error: %s", strings.ToLower(errRes.Detail))
        }
        acctId = okRes.Body[0].ID
    }

    res, err = client.R().
        SetDebug(g.Debug).
        SetError(&errRes).
        SetPathParams(map[string]string{
            //"accountId": strconv.Itoa(o.AccountID),
            "accountId": acctId,
        }).
        SetQueryParams(map[string]string{
            "firstResultIndex": strconv.Itoa(o.Offset),
            "maxResults": strconv.Itoa(o.Limit),
            "status": o.Status,
        }).
        Get(g.ApiUrl+"/account/v1/accounts/{accountId}/users")
    if (err != nil) {
        return err
    }
    if (res.IsError()) {
        return fmt.Errorf("error: %s", strings.ToLower(errRes.Detail))
    }
    tty.Printjson(res.Bytes())
    return nil
}
