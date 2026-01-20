package organization

import (
    "fmt"
    "strconv"

    "rhsmctl/internal/api"
    "rhsmctl/internal/cli"
    "rhsmctl/internal/resty"
    "rhsmctl/internal/tty"
    "github.com/alecthomas/kong"
)

type Result struct {
    Organization *Organization `json:"body"`
}

type Organization struct {
    Id string `json:"id"`
    SimpleContentAccessCapable bool `json:"simpleContentAccessCapable"`
    SimpleContentAccess string `json:"simpleContentAccess"`
}

func (o *Options) Run(ctx *kong.Context, g *cli.Globals) error {
    //client := resty.New(g)
    //var result Result
    //var errResult ErrorResult
    //res, _ := client.R().
    //    SetResult(&result).
    //    SetError(&errResult).
    //    Get(g.ApiUrl+"/organization")
    //headers := []string{"id", "simpleContentAccessCapable", "simpleContentAccess"}
    //rows := result.Organization.ToRows()
    //table := util.NewTable().
    //    Headers(headers...).
    //    Rows(rows...)
    //fmt.Println(table)

    var errRes api.ManagementErrorRoot

    client := resty.New(g)
    res, err := client.R().
        SetDebug(g.Debug).
        SetError(&errRes).
        Get(g.ApiUrl+"/management/v1/organization")

    if (err != nil) {
        return err
    }
    if (res.IsError()) {
        return fmt.Errorf("Error: %s", errRes.Error.Message)
    }

    tty.Printjson(res.Bytes())
    return nil
}

func (o *Organization) ToRows() [][]string {
    return [][]string{
        {o.Id, strconv.FormatBool(o.SimpleContentAccessCapable), o.SimpleContentAccess},
    }
}
