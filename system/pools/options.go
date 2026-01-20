package pools

import (
)

// Options is the customization options for the system pools command.
type Options struct {
    Limit int `default:"50" short:"l" help:"Max number of results."`
    Offset int `default:"0" short:"o" help:"Index to start results from."`
    SystemUUID string `arg:"" help:"System UUID."`
}
