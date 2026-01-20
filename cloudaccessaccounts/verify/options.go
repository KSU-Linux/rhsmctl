package verify

import (
)

// Options is the customization options for the cloud-access-account update command.
type Options struct {
    ProviderShortName string `default:"MSAZ" enum:"AWS,AGOV,ACN,MSAZ,GCE" name:"provider" help:"Short name of the provider (AWS,AGOV,ACN,MSAZ,GCE)."`
    AccountID string `arg:"" help:"ID of the provider account to verify."`
    Identity string `help:"Identity from the cloud provider's metadata service."`
    Signature string `help:"Signature from the cloud provider's metadata service."`
}
