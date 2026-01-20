package info

import (
    "fmt"
    "strings"
    "rhsmctl/internal/api"
    "rhsmctl/internal/cli"
    "rhsmctl/internal/resty"
    "rhsmctl/internal/tty"
    "github.com/alecthomas/kong"
)

func (o *Options) Run(ctx *kong.Context, g *cli.Globals) error {
    include := make([]string, 0)
    if o.Include.Entitlements {
        include = append(include, "entitlements")
    }
    if o.Include.Facts {
        include = append(include, "facts")
    }
    if o.Include.Products {
        include = append(include, "installedProducts")
    }
    var errRes api.ManagementErrorRoot
    client := resty.New(g)
    res, err := client.R().
        SetDebug(g.Debug).
        SetError(&errRes).
        SetPathParams(map[string]string{
            "SystemUUID": o.SystemUUID,
        }).
        SetQueryParams(map[string]string{
            "include": strings.Join(include, ","),
        }).
        Get(g.ApiUrl+"/management/v1/systems/{SystemUUID}")
    if (err != nil) {
        return err
    }
    if (res.IsError()) {
        return fmt.Errorf("error: %s", strings.ToLower(errRes.Error.Message))
    }
    tty.Printjson(res.Bytes())
    return nil
}
