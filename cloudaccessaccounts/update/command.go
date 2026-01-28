package update

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
    var errRes api.ManagementError
    client := resty.New(g)
    res, err := client.R().
        SetBody(map[string]string{
            "nickname": o.Nickname,
        }).
        SetDebug(g.Debug).
        SetError(&errRes).
        SetPathParams(map[string]string{
            "AccountID": o.ID,
            "ProviderShortName": o.ProviderShortName,
        }).
        Put(g.ApiUrl+"/management/v1/cloud_access_providers/{ProviderShortName}/accounts/{AccountID}")
    if (err != nil) {
        return err
    }
    if (res.IsError()) {
        return fmt.Errorf("error: %s", strings.ToLower(errRes.Error.Message))
    }
    tty.Println("updated "+o.ID)
    return nil
}
