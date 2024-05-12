package domains

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LogoutRequest struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	AccountType []int  `json:"account_type"`
}
