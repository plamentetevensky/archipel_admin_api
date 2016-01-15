package response

type ErrorResponse struct {
    Status  string                  `json:"status"`
    Message string                  `json:"message"`
    Code    int64                   `json:"code"`
}

