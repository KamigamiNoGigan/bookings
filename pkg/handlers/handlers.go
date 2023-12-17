package handlers

import (
	"net/http"

	"github.com/KamigamiNoGigan/booking/pkg/config"
	"github.com/KamigamiNoGigan/booking/pkg/models"
	"github.com/KamigamiNoGigan/booking/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewHandler(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewRepo(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplates(w, "home.page.html", &models.DataStruct{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap := make(map[string]string)
	stringMap["test"] = "I love my dear Lisochka!"
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplates(w, "about.page.html", &models.DataStruct{StringMap: stringMap})
}
