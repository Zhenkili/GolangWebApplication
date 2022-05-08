package main

//update test
import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Zhenkili/udemyproject/pkg/config"
	"github.com/Zhenkili/udemyproject/pkg/handler"
	"github.com/Zhenkili/udemyproject/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.Appconfig
var session *scs.SessionManager

// main is the main function
func main() {

	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("error in create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	HandlerRepo := handler.NewRepo(&app)
	handler.NewHandlers(HandlerRepo)
	render.NewTemplates(&app)

	// http.HandleFunc("/", handler.Repo.Home)
	// http.HandleFunc("/about", handler.Repo.About)

	fmt.Printf("Staring application on port %s", portNumber)
	// _ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
