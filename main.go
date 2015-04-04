package main

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"

	"github.com/tattsun/coopy/config"
	"github.com/tattsun/coopy/models"
)

var conf = config.GetConfig()
var model = models.NewModel(conf.MysqlHost, conf.MysqlUser, conf.MysqlPassword, conf.MysqlDatabase)

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	model.Migrate()
	fmt.Fprintf(w, "Hello, %s, %s day", c.URLParams["name"], conf.MysqlHost)
}

func main() {
	goji.Get("/hello/:name", hello)
	goji.Serve()
}
