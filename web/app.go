package web

import (
	"log"
	"net/http"

	"html/template"

	chi "github.com/go-chi/chi/v5"
	chi_middleware "github.com/go-chi/chi/v5/middleware"
)

// Global Application State
type App struct{}

func NewApp() *App {
	return &App{}
}

func (app *App) Router() http.Handler {
	router := chi.NewRouter()
	router.Use(chi_middleware.Logger)

	// images

	fs := http.FileServer(http.Dir("web/assets/images"))
	router.Get("/assets/images/*", http.StripPrefix("/assets/images/", fs).ServeHTTP)

	// assets

	router.Get("/favicon.ico", app.serveTabIcon)
	router.Get("/assets/graphics/dach-flat.svg", app.serveFlatLogo)
	router.Get("/assets/graphics/dach.svg", app.serveDetailedLogo)
	router.Get("/assets/styles/bundle.css", app.serveStyleBundle)

	// pages

	router.Get("/", app.serveIndexPage)
	router.NotFound(app.serveNotFoundPage)

	return router
}

func (app *App) serveTabIcon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/assets/favicon.ico")
}

func (app *App) serveDetailedLogo(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/assets/graphics/dach.svg")
}

func (app *App) serveFlatLogo(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/assets/graphics/dach-flat.svg")
}

func (app *App) serveStyleBundle(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/assets/styles/bundle.css")
}

func (app *App) serveIndexPage(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template

	tmpl = template.New("")
	tmpl = template.Must(tmpl.ParseFiles("web/templates/index.html"))
	tmpl = template.Must(tmpl.ParseGlob("web/icons/*.svg"))

	tmpl.ExecuteTemplate(w, "page", map[string]any{})
}

func (app *App) serveNotFoundPage(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("404 not found"))

	if err != nil {
		log.Fatalf("failed to write 404 response: %v", err)
	}
}
