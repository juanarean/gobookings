package handlers

import (
	"errors"
	"net/http"

	"github.com/juanarean/gobookings/pkg/config"
	"github.com/juanarean/gobookings/pkg/models"
	"github.com/juanarean/gobookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository{
	return &Repository{
		App: a,
	}
}

func NewHandlers(r * Repository) {
	Repo = r
}

func (m *Repository)Home(res http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(res, "Hello")
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(res, "home.page.html", &models.TemplateData{})
}

func (m *Repository)About(res http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{"test": "Hello, again."}

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"]=remoteIP

	render.RenderTemplate(res, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

func divideValues(x, y float32) (float32, error) {
	if y == 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	div := x / y
	return div, nil
}
