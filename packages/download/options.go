package download

import (
)

// Options is the customization options for the packages info command.
type Options struct {
    Checksum string `arg:"" help:"Package checksum"`
}
