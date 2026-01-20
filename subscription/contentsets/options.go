package contentsets

import (
)

// Options is the customization options for the subscription content-sets command.
type Options struct {
    Limit int `default:"1000" short:"l" help:"Max number of results"`
    Offset int `default:"0" short:"o" help:"Index to start results from"`
    Subscription string `arg:"" help:"Subscription number"`
}
