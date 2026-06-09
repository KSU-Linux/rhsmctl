package businesshours

import (
)

// Options is the customization options for the busineshours command.
type Options struct {
    Timezone string `default:"America/New_York" short:"z" help:"Name of the time zone (e.g. US/Central)."`
}
