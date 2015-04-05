package main

import (
	"net/http"

	"github.com/tattsun/coopy/models"

	"github.com/gorilla/sessions"
	"github.com/zenazn/goji/web"
)

var store *sessions.CookieStore

func initSession(secretKey string) {
	store = sessions.NewCookieStore([]byte(secretKey))
}

func middlewareCors(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func middlewareLogin(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "coopy")
		userid := session.Values["userid"].(string)
		token := session.Values["token"].(string)
		ok := models.AuthorizeToken(userid, token)
		if !ok {

		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
