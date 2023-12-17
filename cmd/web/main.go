package main

import (
	"log"
	"net/http"
	"time"

	"github.com/KamigamiNoGigan/booking/pkg/config"
	"github.com/KamigamiNoGigan/booking/pkg/handlers"
	"github.com/KamigamiNoGigan/booking/pkg/render"

	"github.com/alexedwards/scs/v2"
)

var portNum = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.InProduction = false

	session = scs.New()
	session.Lifetime = time.Hour * 3
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateCache()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc
	app.UseCache = true
	render.NewTemplate(&app)
	repo := handlers.NewHandler(&app)
	handlers.NewRepo(repo)

	log.Println("Программа запущена на порте", portNum)

	srv := &http.Server{
		Addr:    portNum,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
