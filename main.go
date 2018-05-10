package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	type User struct {
		Id   int
		Name string
	}

	db, _ := gorm.Open("sqlite3", "./test.db")
	defer db.Close()

	db.AutoMigrate(&User{})

	tomoya := &User{}
	tomoya.Id = 0
	tomoya.Name = "tomoya"

	db.Create(&tomoya)

	var users []User
	db.Find(&users) // SELECT * FROM users;
	fmt.Println(users)
}
