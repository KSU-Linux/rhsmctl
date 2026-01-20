package attach

import (
)

// Options is the customization options for the allocation-entitlement attach command.
type Options struct {
    Quantity int `help:"Quantity you want to attach."`
    UUID string `arg:"" help:"UUID of the allocation."`
    Pool string `arg:"" help:"ID of the pool."`
}
