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

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	models.Migrate("test", "hogehoge", "coopy")

	fmt.Fprintf(w, "Hello, %s, %s day", c.URLParams["name"], conf.MysqlUrl)
}

func main() {
	goji.Get("/hello/:name", hello)
	goji.Serve()
}
