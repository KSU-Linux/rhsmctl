package add

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
        SetBody([]map[string]string{
            map[string]string{
                "id": o.AccountID,
                "nickname": o.Nickname,
            },
        }).
        SetDebug(g.Debug).
        SetError(&errRes).
        SetPathParams(map[string]string{
            "ProviderShortName": o.ProviderShortName,
        }).
        Post(g.ApiUrl+"/management/v1/cloud_access_providers/{ProviderShortName}/accounts")
    if (err != nil) {
        return err
    }
    if (res.IsError()) {
        return fmt.Errorf("error: %s", strings.ToLower(errRes.Error.Message))
    }
    tty.Println("added "+o.AccountID)
    return nil
}
