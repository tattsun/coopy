package models

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	UserID    int `sql:"index"`
	Name      string
	CreatedAt time.Time
	Email     string
}

type AuthInfo struct {
	UserID   int `sql:"index"`
	Password string
}

type Article struct {
	ID     int `sql:"index"`
	Title  string
	Body   string
	UserID string `sql:"index"`
}

func Migrate(userId string, password string, dbname string) error {
	conInfo := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", userId, password, dbname)

	db, err := gorm.Open("mysql", conInfo)
	defer db.Close()
	if err != nil {
		return err
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&AuthInfo{})
	db.AutoMigrate(&Article{})
	return nil
}
