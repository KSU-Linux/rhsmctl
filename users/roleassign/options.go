package roleassign

import (
)

// Options is the customization options for the users role-assign command.
type Options struct {
    AccountID *int `hidden:"" help:"Red Hat account ID."`
    UserID int `arg:"" help:"User account ID."`
    Roles *[]string `enum:"organization_administrator" required:"" help:"Account roles."`
}
