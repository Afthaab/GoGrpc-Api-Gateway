package domain

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"  `
	Password string `json:"password" `
	Email    string `json:"email" `
	Phone    string `json:"phone"`
	Otp      string `json:"Otp"`
}
