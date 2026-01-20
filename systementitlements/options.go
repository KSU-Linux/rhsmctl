package systementitlements

import (
    "rhsmctl/systementitlements/attach"
    "rhsmctl/systementitlements/remove"
)

type Options struct {
    Attach attach.Options `cmd:"" help:"Attach entitlement to a system."`
    Remove remove.Options `cmd:"" help:"Remove entitlement from a system."`
}
