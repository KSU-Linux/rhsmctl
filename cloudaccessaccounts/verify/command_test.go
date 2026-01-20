package verify

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
        Put("/cloud_access_providers/MSAZ/accounts/b99d9264-8db1-4909-860b-af8f1376ad5d").
        Reply(204)

    // Run our test.
    opts := &Options{
        AccountID: "b99d9264-8db1-4909-860b-af8f1376ad5d",
        ProviderShortName: "MSAZ",
        Identity: "mock_identity",
        Signature: "mock_signature",
    }
    err := opts.Run(&kong.Context{}, mock.Globals)

    // Restore stdout so we can print error messages.
    os.Stdout = stdout

    st.Assert(t, err, nil)
}

//func TestError(t *testing.T) {
//    // Temporarily disable stdout while we mock API calls.
//    stdout := os.Stdout
//    os.Stdout = nil
//
//    // Register mock API endpoints with simulated responses.
//    defer gock.Disable()
//    gock.New(mock.Globals.ApiUrl).
//        Post("/token").
//        Reply(200).
//        JSON(mock.SuccessJSON("/token"))
//    gock.New(mock.Globals.ApiUrl).
//        Delete("/systems/12345678-1234-5678-1234-567812345678").
//        Reply(500).
//        JSON(mock.ErrorJSON("/systems/delete"))
//
//    // Run our test.
//    opts := &Options{}
//    err := opts.Run(&kong.Context{}, mock.Globals)
//
//    // Restore stdout so we can print error messages.
//    os.Stdout = stdout
//
//    st.Refute(t, err, nil)
//}
