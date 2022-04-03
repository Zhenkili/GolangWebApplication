package main

import (
	"net/http"

	"github.com/Zhenkili/udemyproject/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "I am the homepage")
	render.RenderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "I am the about page")
	// res := addNum(1, 4)
	// fmt.Fprintf(w, "I do the culculate, 1 + 4 = %d", res)
	render.RenderTemplate(w, "about.page.html")
}
