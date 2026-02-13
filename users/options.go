package users

import (
	"rhsmctl/users/info"
	"rhsmctl/users/list"
)

type Options struct {
    Info info.Options `cmd:"" help:"Get details of a user under the account."`
    List list.Options `cmd:"" help:"List all the users under the account."`
}
