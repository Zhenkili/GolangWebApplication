package main

import (
	"net/http"

	"github.com/Zhenkili/udemyproject/pkg/config"
	"github.com/Zhenkili/udemyproject/pkg/handler"
	"github.com/bmizerany/pat"
)

//use router instead of main
func routes(app *config.Appconfig) http.Handler {
	mux := pat.New()
	///repalce under instead of the handlerFunc in main
	mux.Get("/", http.HandlerFunc(handler.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handler.Repo.About))

	return mux
}
