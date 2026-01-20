package goldimage

import (
)

// Options is the customization options for the cloud-access-account gold-image command.
type Options struct {
    ProviderShortName string `default:"MSAZ" enum:"AWS,AGOV,ACN,MSAZ,GCE" name:"provider" help:"Short name of the provider (AWS,AGOV,ACN,MSAZ,GCE)."`
    AccountID string `arg:"" help:"Account ID of the provider account."`
    Image string `default:"rhel-byos" help:"Name of the image to request access to."`
}
