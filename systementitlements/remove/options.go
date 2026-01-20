package remove

import (
)

// Options is the customization options for the system entitlement remove command.
type Options struct {
    SystemUUID string `arg:"" help:"UUID of the system."`
    EntitlementID string `arg:"" help:"ID of the entitlement."`
}
