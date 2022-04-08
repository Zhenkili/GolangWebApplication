package main

//update test
import (
	"fmt"
	"net/http"

	"github.com/Zhenkili/udemyproject/pkg/handler"
)

const portNumber = ":8080"

// main is the main function
func main() {
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
