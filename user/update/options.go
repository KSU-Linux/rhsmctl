package update

import (
)

// Options is the customization options for the users update command.
type Options struct {
    Salutation *string `help:"Salutation (e.g. Mr, Mrs)."`
    FirstName *string `help:"Your first name."`
    LastName *string `help:"Your last name."`
    Phone *string `help:"Your phone number."`
    Streets *[]string `default:"" help:"Your street address."`
    City *string `help:"Your city."`
    Country *string `help:"Your country."`
    State *string `help:"Your state."`
    County *string `help:"Your county."`
    Zip *string `help:"Your zip code."`
}
