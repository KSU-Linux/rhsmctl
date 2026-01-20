package cli

type Globals struct {
    ApiKey string `placeholder:"KEY" required:"" help:"API key to use for authentication."`
    ApiUrl string `placeholder:"URL" required:"" help:"URL to the Red Hat API endpoint." default:"https://api.access.redhat.com"`
    ClientId string `required:"" help:"OAuth 2.0 client id to use for authentication." default:"rhsm-api"`
    Debug bool `short:"d" help:"Enable debug mode."`
    Format string `default:"json" enum:"json" hidden:"" help:"Output format to use." `
    TokenUrl string `placeholder:"URL" required:"" help:"OAuth 2.0 token URL to use for authentication." default:"https://sso.redhat.com/auth/realms/redhat-external/protocol/openid-connect/token"`
}
