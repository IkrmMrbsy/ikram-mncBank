package models

type Customer struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	IsLoggedIn bool   `json:"isLoggedIn"`
}
