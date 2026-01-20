package remove

import (
)

// Options is the customization options for the allocation-entitlement remove command.
type Options struct {
    Force bool `default:"false" short:"f" help:"Skip interactive confirmation prompt."`
    UUID string `arg:"" help:"UUID of the allocation."`
    EntitlementID string `arg:"" help:"ID of the entitlement to remove."`
}
