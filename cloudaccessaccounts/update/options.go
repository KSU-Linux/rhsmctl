package update

import (
)

// Options is the customization options for the cloud-access-account update command.
type Options struct {
    ProviderShortName string `default:"MSAZ" enum:"AWS,AGOV,ACN,MSAZ,GCE" name:"provider" help:"Short name of the provider (AWS,AGOV,ACN,MSAZ,GCE)."`
    ID string `arg:"" help:"ID of the provider account to remove."`
    Nickname string `arg:"" help:"Nickname for the provider account."`
}
