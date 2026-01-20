package api

type AccountErrorRoot struct {
	Body []AccountErrorBody `json:"body"`
}

type AccountErrorBody struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type ManagementErrorRoot struct {
    Error ManagementError `json:"error"`
}

type ManagementError struct {
    Code int `json:"code"`
    Message string `json:"message"`
}
