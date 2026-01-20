package info

import (
)

// Options is the customization options for the system info command.
type Options struct {
    Include struct {
        Entitlements bool `default:"false" help:"Include entitlements."`
        Facts bool `default:"false" help:"Include facts."`
        Products bool `default:"false" help:"Include installed products."`
    } `embed:"" prefix:"include-"`
    SystemUUID string `arg:"" help:"System UUID."`
}
