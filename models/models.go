package models

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	Id        int
	UserID    string    `sql:"unique_index"`
	CreatedAt time.Time `sql:"DEFAULT:current_timestamp"`
	Email     string    `sql:"not null"`
	Name      string    `sql:"not null"`
	IsPending bool      `sql:"not null; DEFAULT:true"`
}

type UserAuthInfo struct {
	Id             int
	UserID         string `sql:"unique_index"`
	HashedPassword string `sql:"not null"`
	Salt           string `sql:"not null"`
	Token          string `sql:"not null"`
}

type Article struct {
	Id        int
	UserID    string    `sql:"not null"`
	CreatedAt time.Time `sql:"DEFAULT:current_timestamp"`
	Tags      []Tag
}

type Revision struct {
	Id           int
	ArticleID    string    `sql:"not null"`
	UserID       string    `sql:"not null"`
	CreatedAt    time.Time `sql:"DEFAULT:current_timestamp"`
	Title        string    `sql:"not null"`
	Content      string    `sql:"not null"`
	ContentBuilt string
}

type Tag struct {
	Id      int
	TagName string
}

type Model struct {
	user     string
	password string
	host     string
	database string
}

var db *gorm.DB = nil

func NewModel(host string, user string, password string, database string) *Model {
	return &Model{user: user, password: password, host: host, database: database}
}

func (m *Model) DropAll() {
	db.DropTableIfExists(&User{})
	db.DropTableIfExists(&UserAuthInfo{})
	db.DropTableIfExists(&Article{})
	db.DropTableIfExists(&Revision{})
	db.DropTableIfExists(&Tag{})
}

func (m *Model) Migrate() {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&UserAuthInfo{})
	db.AutoMigrate(&Article{})
	db.AutoMigrate(&Revision{})
	db.AutoMigrate(&Tag{})
}

func (m *Model) Open() error {
	if db != nil {
		err := m.Close()
		if err != nil {
			return err
		}
	}

	conInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", m.user, m.password, m.host, m.database)
	dbn, err := gorm.Open("mysql", conInfo)
	if err != nil {
		return err
	}
	db = &dbn
	db.LogMode(true)
	return nil
}

func (m *Model) Close() error {
	err := db.Close()
	if err != nil {
		return err
	}
	return nil
}
