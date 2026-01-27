package main

import (
    "fmt"
    "os"
    "runtime/debug"

    "rhsmctl/account"
    "rhsmctl/allocation"
    "rhsmctl/allocationentitlements"
    "rhsmctl/cloudaccess"
    "rhsmctl/cloudaccessaccounts"
    "rhsmctl/errata"
    "rhsmctl/images"
    "rhsmctl/organization"
    "rhsmctl/packages"
    "rhsmctl/subscription"
    "rhsmctl/system"
    "rhsmctl/systementitlements"
    "rhsmctl/user"
    "rhsmctl/users"
    "rhsmctl/internal/cli"
    "rhsmctl/internal/util"
    "github.com/alecthomas/kong"
    "github.com/alecthomas/kong-yaml"
)

const shaLen = 7

var (
    // Version contains the application version number. It's set via ldflags
    // when building.
    Version = ""

    // CommitSHA contains the SHA of the commit that this application was built
    // against. It's set via ldflags when building.
    CommitSHA = ""
)

// Options is the command-line interface options for rhsmctl.
type Options struct {
    cli.Globals
    Account account.Options `cmd:"" help:"Show the details of your account."`
    Allocation allocation.Options `cmd:"" help:"Manage allocations."`
    AllocationEntitlements allocationentitlements.Options `cmd:"" help:"Manage allocation entitlements."`
    CloudAccess  cloudaccess.Options `cmd:"" help:"Manage cloud access."`
    CloudAccessAccounts  cloudaccessaccounts.Options `cmd:"" help:"Manage cloud access accounts."`
    Errata errata.Options `cmd:"" help:"Manage errata."`
    Images images.Options `cmd:"" help:"Manage images."`
    Organization organization.Options `cmd:"" help:"Show the details of your organization."`
    Packages packages.Options `cmd:"" help:"Manage packages."`
    Subscription subscription.Options `cmd:"" help:"Manage subscriptions."`
    System system.Options `cmd:"" help:"Manage systems."`
    SystemEntitlements systementitlements.Options `cmd:"" help:"Manage entitlements of a system."`
    User user.Options `cmd:"" help:"Show the details of the user."`
    Users users.Options `cmd:"" help:"Manage users."`
    Version kong.VersionFlag `short:"v" help:"Print version information and quit."`
}

func main() {
    if Version == "" {
        if info, ok := debug.ReadBuildInfo(); ok && info.Main.Sum != "" {
            Version = info.Main.Version
        } else {
            Version = "unknown (built from source)"
        }
    }
    version := fmt.Sprintf("rhsmctl version %s", Version)
    if len(CommitSHA) >= shaLen {
        version += " (" + CommitSHA[:shaLen] + ")"
    }

    options := &Options{}
    ctx := kong.Parse(
        options,
        kong.Bind(&options.Globals),
        kong.Configuration(kongyaml.Loader, util.UserConfigFile()),
        kong.ConfigureHelp(kong.HelpOptions{
            Compact:             true,
            Summary:             false,
            NoAppSummary:        true,
            NoExpandSubcommands: true,
        }),
        kong.Description("A tool for controlling RHSM API."),
        kong.UsageOnError(),
        kong.Vars{
            "version":                 version,
            "versionNumber":           Version,
        },
    )
    if err := ctx.Run(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
