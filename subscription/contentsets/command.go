package contentsets

import (
    "fmt"
    "strconv"
    "rhsmctl/internal/api"
    "rhsmctl/internal/cli"
    "rhsmctl/internal/resty"
    "rhsmctl/internal/tty"
    "github.com/alecthomas/kong"
)

func (o *Options) Run(ctx *kong.Context, g *cli.Globals) error {
    var errRes api.ManagementErrorRoot
    client := resty.New(g)
    res, err := client.R().
        SetDebug(g.Debug).
        SetError(&errRes).
        SetPathParams(map[string]string{
            "Subscription": o.Subscription,
        }).
        SetQueryParams(map[string]string{
            "limit": strconv.Itoa(o.Limit),
            "offset": strconv.Itoa(o.Offset),
        }).
        Get(g.ApiUrl+"/management/v1/subscriptions/{Subscription}/contentSets")
    if (err != nil) {
        return err
    }
    if (res.IsError()) {
        return fmt.Errorf("Error: %s", errRes.Error.Message)
    }
    tty.Printjson(res.Bytes())
    return nil
}
