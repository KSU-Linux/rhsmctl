package users

import (
	"rhsmctl/users/info"
	"rhsmctl/users/invite"
	"rhsmctl/users/list"
	"rhsmctl/users/status"
	"rhsmctl/users/update"
)

type Options struct {
    Info info.Options `cmd:"" help:"Get details of a user under the account."`
    Invite invite.Options `cmd:"" help:"Invite new users to join the account."`
    List list.Options `cmd:"" help:"List all the users under the account."`
    Status status.Options `cmd:"" help:"Get current status of a user under the account."`
    Update update.Options `cmd:"" help:"Update details of a user under the account."`
}
