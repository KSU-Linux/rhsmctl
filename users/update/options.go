package update

import (
)

// Options is the customization options for the users update command.
type Options struct {
    AccountID *int `help:"Red Hat account ID."`
    UserID *int `arg:"" help:"User account ID."`
    EMail *string `help:"E-mail address."`
    Permissions *[]string `default:"portal_download,portal_system_management,portal_manage_subscriptions,portal_manage_cases" enum:"portal_download,portal_system_management,portal_manage_subscriptions,portal_manage_cases" help:"Customer portal access permissions (portal_download,portal_system_management,portal_manage_subscriptions,portal_manage_cases)."`
    Roles *[]string `enum:"organization_administrator" help:"Account roles (organization_administrator)."`
}
