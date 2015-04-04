package config

import (
	"os"
)

type Config struct {
	MysqlUrl string
}

func GetConfig() *Config {
	mysqlUrl := os.Getenv("MYSQL_URL")
	return &Config{MysqlUrl: mysqlUrl}
}
