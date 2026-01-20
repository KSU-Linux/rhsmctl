package add

import (
)

// Options is the customization options for the cloud-access-account add command.
type Options struct {
    AccountID string `arg:"" help:"Account ID of the provider account to add."`
    Nickname string `arg:"" optional:"" help:"Nickname for the provider account."`
    ProviderShortName string `default:"MSAZ" enum:"AWS,AGOV,ACN,MSAZ,GCE" name:"provider" help:"Short name of the provider (AWS,AGOV,ACN,MSAZ,GCE)."`
}
