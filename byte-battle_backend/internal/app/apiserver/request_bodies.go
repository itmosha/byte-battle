package apiserver

type UserBody struct {
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}
