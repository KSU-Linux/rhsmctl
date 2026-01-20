package packages

import (
)

// Options is the customization options for the system packages command.
type Options struct {
    ErrataDetail bool `default:"false" help:"Show errata details for packages."`
    Limit int `default:"50" short:"l" help:"Max number of results."`
    Offset int `default:"0" short:"o" help:"Index to start results from."`
    SystemUUID string `arg:"" help:"System UUID."`
    Upgradeable bool `default:"false" help:"Show upgradable packages only."`
}
