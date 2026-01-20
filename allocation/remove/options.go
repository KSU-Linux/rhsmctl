package remove

import (
)

// Options is the customization options for the allocation remove command.
type Options struct {
    Force bool `default:"false" short:"f" help:"Skip interactive confirmation prompt."`
    UUID string `arg:"" required:"" help:"UUID of the allocation to delete"`
}
