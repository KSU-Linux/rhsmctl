package cloudaccess

import (
    "rhsmctl/cloudaccess/goldimage"
    "rhsmctl/cloudaccess/list"
    "rhsmctl/cloudaccess/remove"
)

type Options struct {
    GoldImage goldimage.Options `cmd:"" help:"Request access to Red Hat Gold Images, where available, for currently-enabled products and provider accounts."`
    List list.Options `cmd:"" help:"List all enabled cloud access providers."`
    Remove remove.Options `cmd:"" help:"Remove a currently-enabled provider account by sourceID."`
}
