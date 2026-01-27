package list

import (
)

// Options is the customization options for the users list command.
type Options struct {
    AccountID int `arg:"" help:"Red Hat account ID."`
    Limit int `default:"100" short:"l" help:"Max number of results."`
    Offset int `default:"0" short:"o" help:"Index to start results from."`
    Status string `default:"any" enum:"enabled,disabled,any" help:"Filter by account status (enabled, disabled, any)."`
}
