package handler

import (
	"net/http"

	"github.com/Zhenkili/udemyproject/pkg/config"
	"github.com/Zhenkili/udemyproject/pkg/models"
	"github.com/Zhenkili/udemyproject/pkg/render"
)

//repository for the handler
var Repo *Repository

//claim the repository type
type Repository struct {
	App *config.Appconfig
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

// Home handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "I am the homepage")

	//store the ip in session
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// About handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "I am the about page")
	// res := addNum(1, 4)
	// fmt.Fprintf(w, "I do the culculate, 1 + 4 = %d", res)

	//doing some logic over here
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	intMap := make(map[string]int)
	intMap["Birthday"] = 19961024

	// get the ip out from session
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	//send the data to the template
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
		IntMap:    intMap,
	})
}
