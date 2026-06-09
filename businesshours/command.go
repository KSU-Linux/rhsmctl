package businesshours

import (
    "fmt"

    "rhsmctl/internal/api"
    "rhsmctl/internal/cli"
    "rhsmctl/internal/resty"
    "rhsmctl/internal/tty"
    "github.com/alecthomas/kong"
)

func (o *Options) Run(ctx *kong.Context, g *cli.Globals) error {
    var errRes api.SupportError

    client := resty.New(g)
    res, err := client.R().
        SetDebug(g.Debug).
        SetError(&errRes).
        SetQueryParams(map[string]string{
            "timezone": o.Timezone,
        }).
        Get(g.ApiUrl+"/support/v1/businesshours")

    if (err != nil) {
        return err
    }
    if (res.IsError()) {
        return fmt.Errorf("Error: %s", errRes.DetailMessage)
    }

    tty.Printjson(res.Bytes())
    return nil
}
