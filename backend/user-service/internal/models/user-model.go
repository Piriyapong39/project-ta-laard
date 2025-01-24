package models

type User struct {
	UserId    uint   `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	IsSeller  bool   `json:"is_seller"`
}
