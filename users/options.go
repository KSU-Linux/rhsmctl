package users

import (
    "rhsmctl/users/info"
    "rhsmctl/users/invite"
    "rhsmctl/users/list"
    "rhsmctl/users/roleassign"
    "rhsmctl/users/rolelist"
    "rhsmctl/users/roleremove"
    "rhsmctl/users/status"
    "rhsmctl/users/update"
)

type Options struct {
    Info info.Options `cmd:"" help:"Get details of a user under the account."`
    Invite invite.Options `cmd:"" help:"Invite new users to join the account."`
    List list.Options `cmd:"" help:"List all the users under the account."`
    RoleAssign roleassign.Options `cmd:"" help:"Assign a new role to a user."`
    RoleList rolelist.Options `cmd:"" help:"Get all roles associated with a user."`
    RoleRemove roleremove.Options `cmd:"" help:"Remove a role associated with a user."`
    Status status.Options `cmd:"" help:"Get current status of a user under the account."`
    Update update.Options `cmd:"" help:"Update details of a user under the account."`
}
