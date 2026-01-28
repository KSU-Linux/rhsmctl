package api

type AccountError struct {
    Type     string `json:"type"`
    Title    string `json:"title"`
    Status   int    `json:"status"`
    Detail   string `json:"detail"`
    Instance string `json:"instance"`
}

type ManagementError struct {
    Error struct {
        Code    int    `json:"code"`
        Message string `json:"message"`
    } `json:"error"`
}
