package list

import (
)

// Options is the customization options for the allocation list command.
type Options struct {
    Limit int `default:"100" short:"l" help:"Max number of results"`
    Offset int `default:"0" short:"o" help:"Index to start results from"`
    Type string `default:"Satellite" short:"t" help:"Allocation type to list (SAM,Satellite)."`
}
