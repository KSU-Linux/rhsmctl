package contentset

import (
)

// Options is the customization options for the images content-set command.
type Options struct {
    ContentSet string `arg required help:"Content set"`
    Limit int `default:"100" short:"l" help:"Max number of results"`
    Offset int `default:"0" short:"o" help:"Index to start results from"`
}
