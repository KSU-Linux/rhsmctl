package organization

import (
    "os"
    "testing"
    "rhsmctl/internal/mock"
    "github.com/alecthomas/kong"
    "github.com/h2non/gock"
    "github.com/nbio/st"
)

func TestSuccess(t *testing.T) {
    // Temporarily disable stdout while we mock API calls.
    stdout := os.Stdout
    os.Stdout = nil

    // Register mock API endpoints with simulated responses.
    defer gock.Disable()
    gock.New(mock.Globals.ApiUrl).
        Post("/token").
        Reply(200).
        JSON(mock.SuccessJSON("/token"))
    gock.New(mock.Globals.ApiUrl).
        Get("/organization").
        Reply(200).
        JSON(mock.SuccessJSON("/organization"))

    // Run our test.
    opts := &Options{}
    err := opts.Run(&kong.Context{}, mock.Globals)

    // Restore stdout so we can print error messages.
    os.Stdout = stdout

    st.Assert(t, err, nil)
}

func TestError(t *testing.T) {
    // Temporarily disable stdout while we mock API calls.
    stdout := os.Stdout
    os.Stdout = nil

    // Register mock API endpoints with simulated responses.
    defer gock.Disable()
    gock.New(mock.Globals.ApiUrl).
        Post("/token").
        Reply(200).
        JSON(mock.SuccessJSON("/token"))
    gock.New(mock.Globals.ApiUrl).
        Get("/organization").
        Reply(500).
        JSON(mock.ErrorJSON("/organization"))

    // Run our test.
    opts := &Options{}
    err := opts.Run(&kong.Context{}, mock.Globals)

    // Restore stdout so we can print error messages.
    os.Stdout = stdout

    st.Refute(t, err, nil)
}
