package packages

import (
	"rhsmctl/packages/download"
	"rhsmctl/packages/info"
	"rhsmctl/packages/contentset"
)

type Options struct {
    Download download.Options `cmd:"" help:"Download a package by its checksum"`
    Info info.Options `cmd:"" help:"Show the details of a package"`
    ContentSet contentset.Options `cmd:"" default:"withargs" help:"List all packages for the specified content set"`
}
