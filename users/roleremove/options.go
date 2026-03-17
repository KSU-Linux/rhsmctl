package roleremove

import (
)

// Options is the customization options for the users role-remove command.
type Options struct {
    AccountID *int `hidden:"" help:"Red Hat account ID."`
    UserID int `arg:"" help:"User account ID."`
    Role string `enum:"organization_administrator" required:"" help:"Account role to remove."`
}
