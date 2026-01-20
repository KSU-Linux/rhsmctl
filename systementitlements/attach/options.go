package attach

import (
)

// Options is the customization options for the system entitlement attach command.
type Options struct {
    Quantity int `help:"Quantity you want to attach."`
    SystemUUID string `arg:"" help:"UUID of the system."`
    Pool string `arg:"" help:"ID of the pool."`
}
