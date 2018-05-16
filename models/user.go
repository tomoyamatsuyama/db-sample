package models

type User struct {
	Id          int    `json:id gorm:"primary_key"`
	UserId      string `json:user_id`
	UserName    string `json:user_name`
	PassWord    []byte
	Description string `json:description`
}
