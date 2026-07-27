package models

type Users struct {
	Users_ID       int    `json:"id"`
	Name     string `json:"name"`	
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
