package handler

import (
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	model "github.com/tomoyamatsuyama/db-sample/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	if c.Request.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(c.Writer, nil)
	} else {
		// Validation
		user := validationCheck(c.Request.FormValue("login_name"), c.Request.FormValue("user_name"), c.Request.FormValue("password"), c.Request.FormValue("description"))

		dbCreate(user)
	}
}

func Login(c *gin.Context) {
	db, _ := gorm.Open("sqlite3", "./test.db")
	defer db.Close()

	user := model.User{}

	if err := db.Where("login_name = ?", c.Request.FormValue("login_name")).First(&user).Error; err != nil {
		panic(err)
	}

	er := bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(c.Request.FormValue("password")))
	if er == nil {
		// return
		fmt.Print("Success")
		c.JSON(200, user)
	} else {
		// return
		tt := model.User{}
		fmt.Print("Failure")
		c.JSON(400, tt)
	}
}

func validationCheck(ui, un, pass, des string) model.User {
	db, _ := gorm.Open("sqlite3", "./test.db")
	defer db.Close()

	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		panic("Error")
	}

	userVali := model.User{}
	userVali.LoginName = ui

	if db.Select(userVali) != nil {
		fmt.Println("Error")
	}

	user := model.User{
		Id:          0,
		LoginName:   ui,
		UserName:    un,
		PassWord:    hash,
		Description: des,
	}
	return user
}

func dbCreate(user model.User) {
	db, _ := gorm.Open("sqlite3", "./test.db")
	defer db.Close()

	// db.CreateTable(&model.User{})

	db.Create(&user)
}
