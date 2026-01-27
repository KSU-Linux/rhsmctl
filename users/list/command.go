package list

import (
    "fmt"
    "strconv"
    "strings"
    "rhsmctl/internal/api"
    "rhsmctl/internal/cli"
    "rhsmctl/internal/resty"
    "rhsmctl/internal/tty"
    "github.com/alecthomas/kong"
)

func (o *Options) Run(ctx *kong.Context, g *cli.Globals) error {
    var errRes api.AccountErrorRoot
    client := resty.New(g)
    res, err := client.R().
        SetDebug(g.Debug).
        SetError(&errRes).
        SetPathParams(map[string]string{
            "accountId": strconv.Itoa(o.AccountID),
        }).
        SetQueryParams(map[string]string{
            "maxResults": strconv.Itoa(o.Limit),
            "firstResultIndex": strconv.Itoa(o.Offset),
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
