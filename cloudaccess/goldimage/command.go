package goldimage

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
        SetBody(map[string][]string{
            "accounts": []string{
                o.AccountID,
            },
            "images": []string{
                o.Image,
            },
        }).
        SetDebug(g.Debug).
        SetError(&errRes).
        SetPathParams(map[string]string{
            "ProviderShortName": o.ProviderShortName,
        }).
        Post(g.ApiUrl+"/management/v1/cloud_access_providers/{ProviderShortName}/goldimage")
    if (err != nil) {
        return err
    }
    if (res.IsError()) {
        return fmt.Errorf("error: %s", strings.ToLower(errRes.Error.Message))
    }
    tty.Println("requested gold image for "+o.AccountID)
    return nil
}
