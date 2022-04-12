package handler

import (
	"net/http"

	"github.com/Zhenkili/udemyproject/pkg/config"
	"github.com/Zhenkili/udemyproject/pkg/render"
)

//repository for the handler
var Repo *Repository

//claim the repository type
type Repository struct {
	App *config.Appconfig
}

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}

//create a new repository here
func NewRepo(a *config.Appconfig) *Repository {
	return &Repository{
		App: a,
	}
}

//set the repository for the handlers(above)
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "I am the homepage")
	render.RenderTemplate(w, "home.page.html")
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "I am the about page")
	// res := addNum(1, 4)
	// fmt.Fprintf(w, "I do the culculate, 1 + 4 = %d", res)

	//doing some logic over here

	//send the data to the template
	render.RenderTemplate(w, "about.page.html")
}
