package main

import (
	"net/http"

	"github.com/Zhenkili/udemyproject/pkg/handler"
)

const portNum = ":9090"

func main() {
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)

	_ = http.ListenAndServe(portNum, nil)
}
