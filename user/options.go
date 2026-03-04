package user

import (
    "rhsmctl/user/info"
    "rhsmctl/user/update"
)

// Options is the customization options for the user command.
type Options struct {
    Info info.Options `cmd:"" help:"Get personal information of the current user."`
    Update update.Options `cmd:"" help:"Update details of the current user."`
}
