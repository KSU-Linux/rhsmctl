package api

type Account struct {
    Body []struct {
        ID   string `json:"id,omitempty"`
        Name string `json:"name,omitempty"`
        Type string `json:"type,omitempty"`
    } `json:"body,omitempty"`
}

type User struct {
    Body struct {
        ID         string `json:"id,omitempty"`
        Username   string `json:"username,omitempty"`
        Salutation string `json:"salutation,omitempty"`
        FirstName  string `json:"firstName,omitempty"`
        LastName   string `json:"lastName,omitempty"`
        Status     string `json:"status,omitempty"`
        Email      string `json:"email,omitempty"`
    } `json:"body,omitempty"`
}
