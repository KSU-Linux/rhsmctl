package status

import (
)

// Options is the customization options for the users status command.
type Options struct {
    AccountID *int `hidden:"" help:"Red Hat account ID."`
    UserID int `arg:"" help:"User account ID."`
}
