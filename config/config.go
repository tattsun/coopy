package config

import (
	"os"
)

type Config struct {
	MysqlHost     string
	MysqlUser     string
	MysqlPassword string
	MysqlDatabase string
}

func GetConfig() *Config {
	mysqlHost := os.Getenv("COOPY_MYSQL_HOST")
	mysqlUser := os.Getenv("COOPY_MYSQL_USER")
	mysqlPassword := os.Getenv("COOPY_MYSQL_PASSWORD")
	mysqlDatabase := os.Getenv("COOPY_MYSQL_DATABASE")
	return &Config{MysqlHost: mysqlHost, MysqlUser: mysqlUser, MysqlPassword: mysqlPassword, MysqlDatabase: mysqlDatabase}
}
