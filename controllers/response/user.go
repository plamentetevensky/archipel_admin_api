package response

type UserLoginSuccessResponse struct {
    Status  string                  `json:"status"`
    Data    UserLoginSuccessData    `json:"data"`
}

type UserLoginSuccessData struct {
    AuthToken   string              `json:"auth_token"`
}