package contentset

import (
)

// Options is the customization options for the subscription list command.
type Options struct {
    Arch string `default:"x86_64" enum:"aarch64,ia64,i386,i686,ppc64,ppc64le,x86_64" short:"a" help:"Architecture"`
    ContentSet string `arg:"" help:"Content set"`
    Limit int `default:"100" short:"l" help:"Max number of results"`
    Offset int `default:"0" short:"o" help:"Index to start results from"`
}
