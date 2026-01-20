package remove

import (
)

// Options is the customization options for the cloudaccess remove command.
type Options struct {
    SourceID string `arg:"" help:"Source ID of the provider account."`
}
