package models

import (
	"github.com/tattsun/coopy/helpers"
)

func before_each() {
	m := NewModel("localhost:3306", "coopyuser", "hogehoge", "coopytest")
	m.Open()
	m.DropAll()
	m.Migrate()
}

func specRandomUser() (*User, string) {
	userid := helpers.RandomStr(20)
	email := helpers.RandomStr(20)
	name := helpers.RandomStr(20)
	password := helpers.RandomStr(20)
	return &User{UserID: userid, Email: email, Name: name}, password
}

/*
func specCreateUser() (user *User, password string) {
	_, _ := CreateUser(user, password)
	u, _ := Find
	return
}
*/
