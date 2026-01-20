package cloudaccessaccounts

import (
    "rhsmctl/cloudaccessaccounts/add"
    "rhsmctl/cloudaccessaccounts/remove"
    "rhsmctl/cloudaccessaccounts/update"
    "rhsmctl/cloudaccessaccounts/verify"
)

// Options is the customization options for the cloud-access-accounts command.
type Options struct {
    Add add.Options `cmd:"" help:"Add a new provider account to a currently-enabled provider for Red Hat Cloud Access."`
    Remove remove.Options `cmd:"" help:"Remove a currently-enabled provider account."`
    Update update.Options `cmd:"" help:"Update the nickname for a currently-enabled provider account."`
    Verify verify.Options `cmd:"" help:"Verify a cloud provider account for use with RHSM auto registration."`
}
