package rhel

import (
)

// Options is the customization options for the images rhel command.
type Options struct {
    Arch string `default:"x86_64" enum:"aarch64,ppc64le,x86_64" short:"a" help:"Architecture"`
    RHELVersion string `arg:"" required:"" help:"RHEL version"`
}
