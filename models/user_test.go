package models

import (
	"github.com/tattsun/coopy/config"
	"testing"
)

var conf = config.GetConfig()
var model = NewModel(conf.MysqlHost, conf.MysqlUser, conf.MysqlPassword, conf.MysqlDatabase)
var transaction, _ = model.Open()

func TestCreateUser(t *testing.T) {
	val := CreateUser()
	if val != 3 {
		t.Error("err")
	}
}
