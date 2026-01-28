package remove

import (
    "fmt"
    "log"
    "rhsmctl/internal/api"
    "rhsmctl/internal/cli"
    "rhsmctl/internal/resty"
    "rhsmctl/internal/tty"
    "github.com/alecthomas/kong"
    "github.com/charmbracelet/huh"
)

func (o *Options) Run(ctx *kong.Context, g *cli.Globals) error {
    confirm := true
    if !o.Force {
        form := huh.NewConfirm().
            Title("Are you sure?").
            Affirmative("Yes").
            Negative("No").
            Value(&confirm)
        if err := form.Run(); err != nil {
            log.Fatal(err)
        }
    }
    if !confirm{
        return nil
    }
    var errRes api.ManagementError
    client := resty.New(g)
    res, err := client.R().
        SetDebug(g.Debug).
        SetError(&errRes).
        SetPathParams(map[string]string{
            "uuid": o.UUID,
            "EntitlementID": o.EntitlementID,
        }).
        Delete(g.ApiUrl+"/management/v1/allocations/{uuid}/entitlements/{EntitlementID}")
    if (err != nil) {
        return err
    }
    if (res.IsError()) {
        return fmt.Errorf("Error: %s", errRes.Error.Message)
    }
    tty.Println(o.EntitlementID+" removed from "+o.UUID)
    return nil

}
