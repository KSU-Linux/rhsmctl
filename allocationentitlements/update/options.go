package update

import (
)

// Options is the customization options for the allocation-entitlement update command.
type Options struct {
    UUID string `arg:"" help:"UUID of the allocation."`
    EntitlementID string `arg:"" help:"ID of the entitlement."`
    Quantity int ` default:"1" help:"Quantity you want to attach (must be less than or equal to the maximum number of allowed entitlements in the entitlement pool)."`
}
