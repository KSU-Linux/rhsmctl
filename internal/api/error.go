package api

type AccountErrorRoot struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}

type ManagementErrorRoot struct {
    Error ManagementError `json:"error"`
}

type ManagementError struct {
    Code int `json:"code"`
    Message string `json:"message"`
}
