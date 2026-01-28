package client

import (
    "fmt"
    "strings"
    "rhsmctl/internal/api"
)

// AccountID returns the account ID of the current user.
func (client *RestClient) AccountID() (string, error) {
    var errRes api.AccountError
    var okRes api.Account
    res, err := client.R().
        SetDebug(client.globals.Debug).
        SetError(&errRes).
        SetResult(&okRes).
        Get(client.globals.ApiUrl+"/account/v1/accounts")
    if (err != nil) {
        return "", err
    }
    if (res.IsError()) {
        return "", fmt.Errorf("error: %s", strings.ToLower(errRes.Detail))
    }
    return okRes.Body[0].ID, nil
}
