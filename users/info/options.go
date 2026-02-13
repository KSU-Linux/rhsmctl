package info

import (
)

// Options is the customization options for the users info command.
type Options struct {
    AccountID int `help:"Red Hat account ID."`
    UserID int `arg:"" help:"User account ID."`
}
