package remove

import (
)

// Options is the customization options for the cloud-access-account remove command.
type Options struct {
    AccountID string `arg:"" help:"ID of the provider account to remove."`
    ProviderShortName string `default:"MSAZ" enum:"AWS,AGOV,ACN,MSAZ,GCE" name:"provider" help:"Short name of the provider (AWS,AGOV,ACN,MSAZ,GCE)."`
}
