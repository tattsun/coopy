package models

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	UserID    string    `gorm:"primary_key"`
	CreatedAt time.Time `sql:"DEFAULT:current_timestamp"`
	Email     string
	Name      string
	Password  string
	Token     string
	IsPending bool
}

type Article struct {
	ID        string `gorm:"primary_key"`
	UserID    string
	CreatedAt time.Time `sql:"DEFAULT:current_timestamp"`
}

type Revision struct {
	ID           string `gorm:"primary_key"`
	ArticleID    string
	UserID       string
	CreatedAt    time.Time `sql:"DEFAULT:current_timestamp"`
	Title        string
	Content      string
	ContentBuilt string
}

type ArticleTag struct {
	ID        string `gorm:"primary_key"`
	ArticleID string
	TagID     string
}

type Tag struct {
	ID string
}

type Model struct {
	user     string
	password string
	host     string
	database string
}

func NewModel(host string, user string, password string, database string) *Model {
	return &Model{user: user, password: password, host: host, database: database}
}

func (m *Model) Migrate() error {
	conInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", m.user, m.password, m.host, m.database)

	db, err := gorm.Open("mysql", conInfo)
	defer db.Close()
	if err != nil {
		return err
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Article{})
	db.AutoMigrate(&Revision{})
	db.AutoMigrate(&ArticleTag{})
	db.AutoMigrate(&Tag{})
	return nil
}
