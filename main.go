package main

import (
	"encoding/json"
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

func resetdb(c web.C, w http.ResponseWriter, r *http.Request) {
	model.DropAll()
	model.Migrate()
	engine := views.NewEngine("static/test.haml")
	engine.Add("name", "john")
	engine.Render(w)
}

func test2(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello, world</h1>%s", r.RequestURI)
}

func newUser(c web.C, w http.ResponseWriter, r *http.Request) {
	engine := views.NewEngine("static/newuser.haml")
	engine.Render(w)
}

func staticFile(c web.C, w http.ResponseWriter, r *http.Request) {
	err := views.WriteStatic(r.RequestURI[1:], w)
	if err != nil {
		fmt.Print(err)
		http.NotFound(w, r)
	}
}

func apiNewUser(w http.ResponseWriter, r *http.Request) {
	userid := r.FormValue("userid")
	email := r.FormValue("email")
	password := r.FormValue("password")
	name := r.FormValue("name")
	user, auth, err := models.CreateUser(userid, email, name, password)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	res := make(map[string]interface{})
	res["id"] = user.Id
	res["token"] = auth.Token
	b, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	w.Write(b)
}

func main() {
	err := model.Open()
	if err != nil {
		panic(err)
	}
	// views
	goji.Get("/resetdb", resetdb)
	goji.Get("/test", test2)
	goji.Get("/users/new", newUser)
	goji.Get("/static/assets/*", staticFile)
	// views(login)
	login := web.New()
	goji.Handle("/l/*", login)
	// apis
	api := web.New()
	goji.Handle("/api/*", api)
	api.Use(middlewareCors)
	api.Post("/api/v1/users", apiNewUser)

	goji.Serve()
}
