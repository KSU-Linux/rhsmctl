package allocationentitlements

import (
    "rhsmctl/allocationentitlements/attach"
    "rhsmctl/allocationentitlements/remove"
    "rhsmctl/allocationentitlements/update"
)

type Options struct {
    Attach attach.Options `cmd:"" help:"Attach an entitlement to an allocation."`
    Remove remove.Options `cmd:"" help:"Remove entitlement from an allocation."`
    Update update.Options `cmd:"" help:"Update attached entitlement to an allocation."`
}
