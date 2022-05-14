package main

import (
	"fmt"
	"github.com/Aalis/bookings/pkg/config"
	"github.com/Aalis/bookings/pkg/handlers"
	"github.com/Aalis/bookings/pkg/renderer"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const (
	portNumber = ":8080"
)

var app config.AppConfig
var sessions *scs.SessionManager

func main() {

	//change to true when in production
	app.InProduction = false

	sessions = scs.New()
	sessions.Lifetime = 24 * time.Hour
	sessions.Cookie.Persist = true
	sessions.Cookie.SameSite = http.SameSiteLaxMode
	sessions.Cookie.Secure = false

	app.Session = sessions

	tc, err := renderer.CreateTemplateCache()
	if err != nil {
		log.Fatal("can`t create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false //use cache or always create template cache (true)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	renderer.NewTemplates(&app)

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)
	fmt.Println(fmt.Sprintf("starting server at port"), portNumber)
	//_ = http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routers(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
