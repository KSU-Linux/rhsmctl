package account

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
    var errRes api.AccountError
    client := resty.New(g)
    res, err := client.R().
        SetDebug(g.Debug).
        SetError(&errRes).
        SetQueryParams(map[string]string{
            "firstResultIndex": strconv.Itoa(o.Offset),
            "maxResults": strconv.Itoa(o.Limit),
        }).
        Get(g.ApiUrl+"/account/v1/accounts")

    if (err != nil) {
        return err
    }
    if (res.IsError()) {
        return fmt.Errorf("Error: %s", res.String())
    }

    tty.Printjson(res.Bytes())
    return nil
}
