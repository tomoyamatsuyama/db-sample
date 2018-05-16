package models

type User struct {
	Id          int    `json:id gorm:"primary_key"`
	LoginName   string `json:login_id`
	UserName    string `json:user_name`
	PassWord    []byte
	Description string `json:description`
}
