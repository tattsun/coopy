package main

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"

	"github.com/realistschuckle/gohaml"

	"github.com/tattsun/coopy/config"
	"github.com/tattsun/coopy/models"
)

var conf = config.GetConfig()
var model = models.NewModel(conf.MysqlHost, conf.MysqlUser, conf.MysqlPassword, conf.MysqlDatabase)

func test(c web.C, w http.ResponseWriter, r *http.Request) {
	scope := make(map[string]interface{})
	scope["lang"] = "HAML"
	content := "I love <\n=lang<\n"
	engine, _ := gohaml.NewEngine(content)
	output := engine.Render(scope)
	fmt.Fprint(w, output)
}

func main() {
	err := model.Open()
	if err != nil {
		panic(err)
	}
	goji.Get("/", test)
	goji.Serve()
}
