package main

import (
	"fmt"
	"net/http"

	"github.com/Zhenkili/udemyproject/pkg/config"
)

func routes(app *config.Appconfig) http.Handler {
	// mux := pat.New()
	return http.HandlerFunc()
	fmt.Println("here")
}
