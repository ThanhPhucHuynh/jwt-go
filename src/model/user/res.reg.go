package model

type RegisterResponse struct {
	User   UserGet
	Token  string `json:"token"`
	Status int    `json:"status"`
}
