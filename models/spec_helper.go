package models

import (
	"github.com/tattsun/coopy/helpers"
)

func before_each() {
	m := NewModel("localhost:3306", "coopyuser", "hogehoge", "coopytest")
	m.debug = true
	m.Open()
	m.DropAll()
	m.Migrate()
}

func specCreateUser() (*User, string, error) {
	userid := helpers.RandomStr(20)
	email := helpers.RandomStr(20)
	name := helpers.RandomStr(20)
	password := helpers.RandomStr(20)
	u, _, err := CreateUser(userid, email, name, password)
	return u, password, err
}

/*
func specCreateUser() (user *User, password string) {
	_, _ := CreateUser(user, password)
	u, _ := Find
	return
}
*/
