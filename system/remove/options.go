package remove

import (
)

// Options is the customization options for the system remove command.
type Options struct {
    SystemUUID string `arg:"" help:"System UUID."`
}
