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
    params := make(map[string]string)
    params["limit"] = strconv.Itoa(o.Limit)
    params["offset"] = strconv.Itoa(o.Offset)
    if o.Filter != "" {
        params["filter"] = o.Filter
    }
    if o.Username != "" {
        params["username"] = o.Username
    }
    var errRes api.ManagementError
    client := resty.New(g)
    res, err := client.R().
        SetDebug(g.Debug).
        SetError(&errRes).
        SetQueryParams(params).
        Get(g.ApiUrl+"/management/v1/systems")
    if (err != nil) {
        return err
    }
    if (res.IsError()) {
        return fmt.Errorf("error: %s", strings.ToLower(errRes.Error.Message))
    }
    tty.Printjson(res.Bytes())
    return nil
}
