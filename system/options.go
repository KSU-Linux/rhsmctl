package system

import (
    "rhsmctl/system/errata"
    "rhsmctl/system/info"
    "rhsmctl/system/list"
    "rhsmctl/system/packages"
    "rhsmctl/system/pools"
    "rhsmctl/system/remove"
)

type Options struct {
    Errata errata.Options `cmd:"" help:"List all applicable errata for a system."`
    Info info.Options `cmd:"" help:"Show the details of a system."`
    List list.Options `cmd:"" default:"withargs" help:"List all systems."`
    Packages packages.Options `cmd:"" help:"List all packages for a system."`
    Pools pools.Options `cmd:"" help:"List all pools for a system"`
    Remove remove.Options `cmd:"" help:"Remove system profile."`
}
