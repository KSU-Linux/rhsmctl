package allocation

import (
    "rhsmctl/allocation/list"
    "rhsmctl/allocation/remove"
    "rhsmctl/allocation/versions"
)

type Options struct {
    List list.Options `cmd:"" help:"List allocations."`
    ListVersions versions.Options `cmd:"" help:"List Satellite versions."`
    Remove remove.Options `cmd:"" help:"Remove allocation profile."`
}
