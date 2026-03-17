package update

import (
)

// Options is the customization options for the users update command.
type Options struct {
    AccountID *int `hidden:"" help:"Red Hat account ID."`
    UserID int `arg:"" help:"User account ID."`
    EMail *string `help:"E-mail address."`
    Permissions *[]string `default:"portal_download,portal_system_management,portal_manage_subscriptions,portal_manage_cases" enum:"portal_download,portal_system_management,portal_manage_subscriptions,portal_manage_cases" placeholder:"PERMS" help:"Customer portal access permissions."`
    Roles *[]string `enum:"organization_administrator" help:"Account roles."`
}
