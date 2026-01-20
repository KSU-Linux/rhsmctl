package errata

import (
	"rhsmctl/errata/contentset"
	"rhsmctl/errata/images"
	"rhsmctl/errata/info"
	"rhsmctl/errata/list"
	"rhsmctl/errata/packages"
	"rhsmctl/errata/systems"
)

type Options struct {
    Info info.Options `cmd:"" help:"Show the details of an advisory."`
    List list.Options `cmd:"" default:"withargs" help:"List all errata for your systems."`
    ListContentSet contentset.Options `cmd:"" help:"List all errata for the specified content set."`
    ListImages images.Options `cmd:"" help:"List all updated container images for an advisory."`
    ListPackages packages.Options `cmd:"" help:"List all packages for an advisory."`
    ListSystems systems.Options `cmd:"" help:"List all systems for an advisory."`
}
