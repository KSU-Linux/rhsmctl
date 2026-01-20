package download

import (
)

// Options is the customization options for the images download command.
type Options struct {
    Checksum string `arg:"" required:"" help:"SHA256 checksum"`
}
