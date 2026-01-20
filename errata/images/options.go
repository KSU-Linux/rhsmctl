package images

import (
)

// Options is the customization options for the errata images command.
type Options struct {
    AdvisoryID string `arg:"" help:"Unique indentifier for a Red Hat advisory"`
}
