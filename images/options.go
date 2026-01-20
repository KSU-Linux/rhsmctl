package images

import (
    "rhsmctl/images/contentset"
    "rhsmctl/images/download"
    "rhsmctl/images/rhel"
)

type Options struct {
    ContentSet contentset.Options `cmd:"" help:"List available images in a content set"`
    Download download.Options `cmd:"" help:"Download an image by checksum"`
    RHEL rhel.Options `cmd:"" help:"List RHEL image downloads by version and architecture"`
}
