package invite

import (
)

// Options is the customization options for the users invite command.
type Options struct {
    AccountID *int `hidden:"" help:"Red Hat account ID."`
    EMails *[]string `arg:"" help:"E-mail addresses to send invitation to."`
    Locale *string `default:"en_US" enum:"de,en_US,es,fr,it,ja,ko,pt_BR,zh_CN,zh_TW" help:"Locale code (language) to use for the invitation."`
    Permissions *[]string `enum:"portal_download,portal_system_management,portal_manage_subscriptions,portal_manage_cases" placeholder:"PERMS" help:"Customer portal access permissions."`
    Roles *[]string `enum:"organization_administrator" help:"Account roles."`
}
