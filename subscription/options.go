package subscription

import (
	"rhsmctl/subscription/contentsets"
	"rhsmctl/subscription/list"
	"rhsmctl/subscription/systems"
)

type Options struct {
    ContentSets contentsets.Options `cmd:"" help:"List all content sets for a subscription"`
    List list.Options `cmd:"" default:"withargs" help:"List all subscriptions"`
    Systems systems.Options `cmd:"" help:"List all content sets for a subscription"`
}
