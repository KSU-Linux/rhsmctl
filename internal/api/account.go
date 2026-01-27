package api

type Account struct {
    Body []struct {
        ID   string `json:"id"`
        Name string `json:"name"`
        Type string `json:"type"`
    } `json:"body"`
}
