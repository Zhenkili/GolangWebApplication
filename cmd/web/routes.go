package main

import (
	"net/http"

	"github.com/Zhenkili/udemyproject/pkg/config"
	"github.com/bmizerany/pat"
)

func routes(app *config.Appconfig) http.Handler {
	mux := pat.New()

	
}
