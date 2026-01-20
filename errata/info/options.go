package info

import (
)

// Options is the customization options for the errata info command.
type Options struct {
    AdvisoryID string `arg:"" help:"Unique indentifier for a Red Hat advisory"`
}
