package list

import (
)

// Options is the customization options for the subscription list command.
type Options struct {
    Filter string `short:"f" help:"Filter systems by system name"`
    Limit int `default:"100" short:"l" help:"Max number of results"`
    Offset int `default:"0" short:"o" help:"Index to start results from"`
    Username string `short:"u" help:"Filter systems by username (system owner)"`
}
