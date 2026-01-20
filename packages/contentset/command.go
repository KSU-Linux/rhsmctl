package contentset

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
    var errRes api.ManagementErrorRoot
    client := resty.New(g)
    res, err := client.R().
        SetDebug(g.Debug).
        SetError(&errRes).
        SetPathParams(map[string]string{
            "Arch": o.Arch,
            "ContentSet": o.ContentSet,
        }).
        SetQueryParams(map[string]string{
            "filter": o.Filter,
            "limit": strconv.Itoa(o.Limit),
            "offset": strconv.Itoa(o.Offset),
        }).
        Get(g.ApiUrl+"/management/v1/packages/cset/{ContentSet}/arch/{Arch}")
    if (err != nil) {
        return err
    }
    if (res.IsError()) {
        return fmt.Errorf("error: %s", strings.ToLower(errRes.Error.Message))
    }
    tty.Printjson(res.Bytes())
    return nil
}
