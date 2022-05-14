package handlers

import (
	"github.com/Aalis/bookings/pkg/config"
	"github.com/Aalis/bookings/pkg/models"
	"github.com/Aalis/bookings/pkg/renderer"
	"net/http"
)

//Repo the repository used by the handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

//NewRepo creates the new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	renderer.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

//About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	strMap := make(map[string]string)
	strMap["t"] = "Revolut"
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	strMap["remote_ip"] = remoteIP
	//send the data to the template
	renderer.RenderTemplate(w, "about.page.html", &models.TemplateData{StringMap: strMap})
}
