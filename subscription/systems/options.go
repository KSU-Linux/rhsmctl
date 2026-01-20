package systems

import (
)

// Options is the customization options for the subscription list command.
type Options struct {
    Limit int `default:"100" short:"l" help:"Max number of results"`
    Offset int `default:"0" short:"o" help:"Index to start results from"`
    Subscription string `arg:"" help:"Subscription number"`
}
