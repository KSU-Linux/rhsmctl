package mock

import (
    "rhsmctl/internal/cli"
)

// Create a mock *cli.Globals so we don't test API calls
// against production URLs.
var Globals = &cli.Globals{
        ApiToken: "mock_api_token",
        ApiUrl: "http://localhost:9999",
        ClientId: "rhsm-api",
        Format: "json",
        TokenUrl: "http://localhost:9999/token",
}

func ErrorJSON(s string) map[string]interface{} {
    switch s {
        //case "/token":
        //    return map[string]interface{}{
        //        "access_token": "mock_access_token",
        //        "token_type": "bearer",
        //        "not-before-policy": 0,
        //        "session_state": "f0dbb8d4-4e4e-4654-844c-6f3704c84422",
        //        "scope": "offline_access",
        //    }
        //case "/organization":
        //    return map[string]interface{}{
        //        "body": map[string]interface{}{
        //            "id": "0000000",
        //            "simpleContentAccessCapable": true,
        //            "simpleContentAccess": "enabled"}}
        default:
            return map[string]interface{}{
                "error": map[string]interface{}{
                    "code": 500,
                    "message": "internal server error"}}
    }
}

func SuccessJSON(s string) map[string]interface{} {
    switch s {
        case "/token":
            return map[string]interface{}{
                "access_token": "mock_access_token",
                "token_type": "bearer",
                "not-before-policy": 0,
                "session_state": "f0dbb8d4-4e4e-4654-844c-6f3704c84422",
                "scope": "offline_access",
            }
        case "/organization":
            return map[string]interface{}{
                "body": map[string]interface{}{
                    "id": "0000000",
                    "simpleContentAccessCapable": true,
                    "simpleContentAccess": "enabled"}}
        case "/systems/delete":
            return map[string]interface{}{
                "body": map[string]interface{}{
                    "id": "0000000",
                    "simpleContentAccessCapable": true,
                    "simpleContentAccess": "enabled"}}
        default:
            return map[string]interface{}{
                "error": map[string]interface{}{
                    "code": 500,
                    "message": "internal server error"}}
    }
}
