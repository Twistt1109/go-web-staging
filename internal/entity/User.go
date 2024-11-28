package entity

type User struct {
	ID       int    `json:"id"`
	UserID   int    `json:"userId"`
	Username string `json:"username"`
	Password string `json:"password"`
}
