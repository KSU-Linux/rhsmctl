package info

import (
)

// Options is the customization options for the user command.
type Options struct {
    Details bool `default:"false" help:"Show extended details of the current user."`
}
