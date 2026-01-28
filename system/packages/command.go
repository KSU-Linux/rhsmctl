package packages

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
    var errRes api.ManagementError
    client := resty.New(g)
    res, err := client.R().
        SetDebug(g.Debug).
        SetError(&errRes).
        SetPathParams(map[string]string{
            "SystemUUID": o.SystemUUID,
        }).
        SetQueryParams(map[string]string{
            "errata_detail": strconv.FormatBool(o.ErrataDetail),
            "limit": strconv.Itoa(o.Limit),
            "offset": strconv.Itoa(o.Offset),
            "upgradeable": strconv.FormatBool(o.Upgradeable),
        }).
        Get(g.ApiUrl+"/management/v1/systems/{SystemUUID}/packages")
    if (err != nil) {
        return err
    }
    if (res.IsError()) {
        return fmt.Errorf("error: %s", strings.ToLower(errRes.Error.Message))
    }
    tty.Printjson(res.Bytes())
    return nil
}
