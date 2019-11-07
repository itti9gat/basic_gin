package model

// ListUser struct
type ListUser []User

// User struct
type User struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Status   string `json:"status"`
}
