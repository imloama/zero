package models

type User struct {
	ID       string `json:"id" form:"-"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
