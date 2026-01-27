package users

import (
	"rhsmctl/users/list"
)

type Options struct {
    List list.Options `cmd:"" help:"List all the users under the account."`
}
