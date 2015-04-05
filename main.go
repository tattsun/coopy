package main

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"

	"github.com/tattsun/coopy/config"
	"github.com/tattsun/coopy/models"
	"github.com/tattsun/coopy/views"
)

var conf = config.GetConfig()
var model = models.NewModel(conf.MysqlHost, conf.MysqlUser, conf.MysqlPassword, conf.MysqlDatabase)

func test(c web.C, w http.ResponseWriter, r *http.Request) {
	engine := views.NewEngine("static/test.haml")
	engine.Add("name", "john")
	engine.Render(w)
}

func test2(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, world</h1>")
}

func main() {
	err := model.Open()
	if err != nil {
		panic(err)
	}
	goji.Get("/", test)
	goji.Get("/test", test2)
	goji.Serve()
}
