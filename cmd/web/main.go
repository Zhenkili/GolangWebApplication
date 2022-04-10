package main

//update test
import (
	"fmt"
	"log"
	"net/http"

	"github.com/Zhenkili/udemyproject/pkg/config"
	"github.com/Zhenkili/udemyproject/pkg/handler"
	"github.com/Zhenkili/udemyproject/pkg/render"
)

const portNumber = ":8080"

// main is the main function
func main() {

	var app config.Appconfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("error in create template cache")
	}

	app.TemplateCache = tc
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)
	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
